package test

import (
	"fmt"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestMqtt(t *testing.T) {
	opt := mqtt.NewClientOptions().
		AddBroker("tcp://43.139.116.74:1883").
		SetClientID("go-test")

	// 回调
	opt.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("Message: %s\n", message.Payload())
		fmt.Printf("Topic: %s\n", message.Topic())
	})

	c := mqtt.NewClient(opt)

	// 链接
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	// 订阅主播
	if token := c.Subscribe("/sys/#", 0, nil); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	// 发布
	if token := c.Publish("/sys/1/device_key/ping", 0, false, "Hello Toma"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	time.Sleep(time.Second * 10)

	// 取消订阅
	if token := c.Unsubscribe("/sys/#"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	// 关闭链接
	c.Disconnect(250)
}
