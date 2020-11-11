package kafka

import (
	"context"
	"git.iothinking.com/base/gobase/log"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"sync"
	"time"
)

type brokerConfigKey struct{}
type clusterConfigKey struct{}

type RecvMsg struct {
	Topic          string
	Msg            []byte
	Timestamp      time.Time // only set if kafka is version 0.10+, inner message timestamp
	BlockTimestamp time.Time // only set if kafka is version 0.10+, outer (compressed) block timestamp
	Partition      int32
	Offset         int64
}

// 返回 true 则 commit offset 到 kafka 否则,消息不会被消费
type SubscriberHandler func(msg *RecvMsg) bool

type Broker struct {
	addrs []string

	// producer
	c sarama.Client
	p sarama.SyncProducer

	// consumers
	sc      []sarama.Client
	scMutex sync.Mutex

	opts   Options
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewBroker(opts ...Option) *Broker {
	ctx, cancelFunc := context.WithCancel(context.Background())
	options := Options{
		// default to json codec
		Context: ctx,
	}

	for _, o := range opts {
		o(&options)
	}

	var cAddrs []string
	for _, addr := range options.Addrs {
		if len(addr) == 0 {
			continue
		}
		cAddrs = append(cAddrs, addr)
	}
	if len(cAddrs) == 0 {
		cAddrs = []string{"127.0.0.1:9092"}
	}

	return &Broker{
		addrs:  cAddrs,
		opts:   options,
		cancel: cancelFunc,
		wg:     sync.WaitGroup{},
	}
}

func (k *Broker) getSaramaClusterClient(opt SubscribeOptions) (sarama.Client, error) {
	config := k.getClusterConfig(opt)
	cs, err := sarama.NewClient(k.addrs, config)
	if err != nil {
		return nil, err
	}
	k.scMutex.Lock()
	defer k.scMutex.Unlock()
	k.sc = append(k.sc, cs)
	return cs, nil
}

func (k *Broker) getClusterConfig(opt SubscribeOptions) *sarama.Config {
	if c, ok := k.opts.Context.Value(clusterConfigKey{}).(*sarama.Config); ok {
		return c
	}

	clusterConfig := sarama.NewConfig()

	// the oldest supported version is V0_10_2_0
	if !clusterConfig.Version.IsAtLeast(sarama.V0_10_2_0) {
		clusterConfig.Version = sarama.V0_10_2_0
	}
	clusterConfig.Consumer.Return.Errors = true
	clusterConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	clusterConfig.Consumer.Offsets.AutoCommit.Enable = opt.AutoAck
	if !opt.AutoAck {
		clusterConfig.Consumer.Offsets.AutoCommit.Interval = opt.AutoAckTime
	}

	return clusterConfig
}

func (k *Broker) getBrokerConfig() *sarama.Config {
	if c, ok := k.opts.Context.Value(brokerConfigKey{}).(*sarama.Config); ok {
		return c
	}
	return sarama.NewConfig()
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	handler SubscriberHandler
}

func NewConsumer(handler SubscriberHandler) *Consumer {
	return &Consumer{handler: handler}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the c as ready
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		msg := &RecvMsg{
			Topic:          message.Topic,
			Msg:            message.Value,
			Timestamp:      message.Timestamp,
			BlockTimestamp: message.BlockTimestamp,
			Partition:      message.Partition,
			Offset:         message.Offset,
		}
		commit := c.handler(msg)

		// 是否提交
		if commit {
			session.MarkMessage(message, "")
		} else {
			log.Warn("subscribe handler return not commit, offset not commit")
		}
	}

	return nil
}

func (k *Broker) Init(opts ...Option) error {
	for _, o := range opts {
		o(&k.opts)
	}
	var cAddrs []string
	for _, addr := range k.opts.Addrs {
		if len(addr) == 0 {
			continue
		}
		cAddrs = append(cAddrs, addr)
	}
	if len(cAddrs) == 0 {
		cAddrs = []string{"127.0.0.1:9092"}
	}
	k.addrs = cAddrs
	return nil
}

func (k *Broker) Connect() error {
	if k.c != nil {
		return nil
	}

	pconfig := k.getBrokerConfig()
	// For implementation reasons, the SyncProducer requires
	// `Producer.Return.Errors` and `Producer.Return.Successes`
	// to be set to true in its configuration.
	pconfig.Producer.Return.Successes = true
	pconfig.Producer.Return.Errors = true

	c, err := sarama.NewClient(k.addrs, pconfig)
	if err != nil {
		return err
	}

	k.c = c

	p, err := sarama.NewSyncProducerFromClient(c)
	if err != nil {
		return err
	}

	k.p = p

	k.scMutex.Lock()
	defer k.scMutex.Unlock()
	k.sc = make([]sarama.Client, 0)

	log.Info("kafka connect successful")
	return nil
}

func (k *Broker) Disconnect() error {
	k.cancel()

	k.scMutex.Lock()
	defer k.scMutex.Unlock()
	for _, client := range k.sc {
		_ = client.Close()
	}
	k.sc = nil
	_ = k.p.Close()
	_ = k.c.Close()
	k.wg.Wait()

	return nil
}

func (k *Broker) Publish(topic string, msg []byte, opts ...PublishOption) error {
	_, _, err := k.p.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	})
	return err
}

func (k *Broker) Subscribe(topic string, handler SubscriberHandler, opts ...SubscribeOption) error {
	opt := SubscribeOptions{
		AutoAck:     true,
		AutoAckTime: 5 * time.Second,
		Queue:       uuid.New().String(),
	}
	for _, o := range opts {
		o(&opt)
	}
	// we need to create a new client per consumer
	c, err := k.getSaramaClusterClient(opt)
	if err != nil {
		return err
	}
	cg, err := sarama.NewConsumerGroupFromClient(opt.Queue, c)
	if err != nil {
		return err
	}

	consumer := NewConsumer(handler)

	topics := []string{topic}

	k.wg.Add(1)
	go func() {
		defer func() {
			k.wg.Done()
		}()
		for {
			select {
			case err := <-cg.Errors():
				if err != nil {
					log.Error("consumer error", zap.Error(err))
				}
			default:
				err := cg.Consume(k.opts.Context, topics, consumer)
				if err != nil {
					if err == sarama.ErrClosedConsumerGroup {
						log.Info("kafka subscribe graceful exit", zap.String("topic", topic))
						return
					} else {
						log.Error("consumer error", zap.Error(err))
					}
				}
				// 查看是是主动退出的
				err = k.opts.Context.Err()
				if err != nil {
					log.Info("kafka subscribe graceful exit", zap.String("topic", topic))
					return
				}

			}
		}
	}()
	return nil
}
