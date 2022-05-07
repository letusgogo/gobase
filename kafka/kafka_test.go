package kafka

import (
	"fmt"
	"github.com/letusgogo/gobase/log"
	"go.uber.org/zap/zapcore"
	"sync/atomic"
	"testing"
	"time"
)

func TestBroker_Subscribe(t *testing.T) {
	log.InitLog("kafka", zapcore.DebugLevel)

	consumerMsg := int32(0)

	kafkaBroker := NewBroker()
	_ = kafkaBroker.Init(Addrs("kafka1.middleware.com:9092"))
	_ = kafkaBroker.Connect()

	err := kafkaBroker.Subscribe("baldr.baldr110.student", func(msg *RecvMsg) bool {
		fmt.Printf("Offset:%d\n", msg.Offset)
		//time.Sleep(50 * time.Millisecond)
		atomic.AddInt32(&consumerMsg, 1)
		return true
	}, Queue("go.micro.api.bigdata"))
	if err != nil {
		t.Fatal("subscribe error", err.Error())
	}

	//err = kafkaBroker.Subscribe("test2", func(msg *RecvMsg) bool {
	//	fmt.Println("Msg:" + string(msg.Msg))
	//	return true
	//}, Queue("test"))
	//if err != nil {
	//	t.Fatal("subscribe error", err.Error())
	//}

	// 发布数据
	//for i := 0; i < 1000; i++ {
	//	_ = kafkaBroker.Publish("test", []byte("helloworld"))
	//}

	time.Sleep(time.Second * 5)

	_ = kafkaBroker.Disconnect()
	fmt.Printf("kafka consumerMsg:%d\n", consumerMsg)
}
