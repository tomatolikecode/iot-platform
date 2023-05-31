package mqtt

import (
	"fmt"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iot-platform/models"
)

var topic = "/sys/#"
var MC mqtt.Client

func NewMqttServer(mqttBroker, clientID, passwrod string) {
	opt := mqtt.NewClientOptions().
		AddBroker(mqttBroker).
		SetClientID(clientID).SetUsername("get").SetPassword(passwrod)

	// 回调
	opt.SetDefaultPublishHandler(publishHandler)

	MC = mqtt.NewClient(opt)

	// 链接
	if token := MC.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主播
	if token := MC.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer func() {
		// 取消订阅
		if token := MC.Unsubscribe(topic); token.Wait() && token.Error() != nil {
			log.Println("[ERROR]: ", token.Error())
		}

		// 关闭链接
		MC.Disconnect(250)
	}()

	select {}
}

/*
心跳检测: /sys/产品key/设备key/ping
*/

// 回调函数处理
func publishHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Message: %s\n", message.Payload())
	fmt.Printf("Topic: %s\n", message.Topic())

	topicArray := strings.Split(strings.TrimPrefix(message.Topic(), "/"), "/")
	if len(topicArray) >= 4 {
		if topicArray[3] == "ping" {
			// 更新时间
			err := models.UpdateDeviceOnlineTime(topicArray[1], topicArray[2])
			if err != nil {
				log.Println("[DB ERROR]: ", err.Error())
			}
		}
	}
}
