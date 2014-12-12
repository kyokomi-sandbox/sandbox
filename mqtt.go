package main

import (
	"fmt"
	"log"

	"time"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func mqttExample() {
	host := "lite.mqtt.shiguredo.jp"
	port := 1883
	user := "handson"
	password := "handson201412"

	// 接続用の設定を作成します
	opts := MQTT.NewClientOptions()
	opts.SetUsername(user)
	opts.SetPassword(password)

	// golangではURIで接続先を指定します。
	brokerUri := fmt.Sprintf("tcp://%s:%d", host, port)
	opts.AddBroker(brokerUri)

	// 設定を元にクライアントを作成します
	client := MQTT.NewClient(opts)

	// Connectします
	_, err := client.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatalln(Subscribe(client))
}

func onMessageReceived(client *MQTT.MqttClient, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\n", message.Topic())
	fmt.Printf("Message: %s\n", message.Payload())

	if string(message.Payload()) == "ぬるぽ" {
		Publish(client, "ガッ")
	}
}

func Subscribe(client *MQTT.MqttClient) error {
	topic := "handson/say"
	qos := 0

	// Subscribeするtopicを設定します
	topicFilter, err := MQTT.NewTopicFilter(topic, byte(qos))
	if err != nil {
		return err
	}

	// Subscribeします
	// onMessageReceived はメッセージが届いたら呼び出されるコールバックです
	_, err = client.StartSubscription(onMessageReceived, topicFilter)
	if err != nil {
		return err
	}

	// そのままではプロセスが終わってしまいますので、待ち受けます
	for {
		time.Sleep(1 * time.Second)
	}
}

func Publish(client *MQTT.MqttClient, message string) error {
	topic := "handson/say"
	qos := 0

	receipt := client.Publish(MQTT.QoS(qos), topic, message)
	<-receipt // Publish成功を待ち受ける

	return nil
}
