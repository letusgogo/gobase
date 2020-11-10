package kafka

import (
	"fmt"
	"git.iothinking.com/base/gobase/log"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func TestBroker_Subscribe(t *testing.T) {
	log.InitLog("httpserver", zapcore.DebugLevel)

	kafkaBroker := NewBroker()
	_ = kafkaBroker.Init(Addrs("kafka1.middleware.com:9092"))
	_ = kafkaBroker.Connect()

	err := kafkaBroker.Subscribe("test", func(msg *RecvMsg) {
		fmt.Println("Msg:" + string(msg.Msg))
	}, Queue("go.micro.api.bigdata"))
	if err != nil {
		t.Fatal("subscribe error", err.Error())
	}

	// 发布数据
	for i := 0; i < 1000; i++ {
		_ = kafkaBroker.Publish("test", []byte("helloworld"))
	}

	time.Sleep(time.Second * 5)

	_ = kafkaBroker.Disconnect()

	time.Sleep(time.Second * 10)

	fmt.Println("kafka exit ok")
}
