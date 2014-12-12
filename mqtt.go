package main

import (
	"fmt"
	"log"

	"time"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func main() {
	host := "hoge.mqtt.shiguredo.jp"
	port := 1883
	user := "example@github"
	password := "<password>"

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
}

func onMessageReceived(client *MQTT.MqttClient, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\n", message.Topic())
	fmt.Printf("Message: %s\n", message.Payload())
}

func Subscribe(client *MQTT.MqttClient) error {
	topic := "example@github/a/b"
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

func Publish(client *MQTT.MqttClient) error {
	topic := "example@github/a/b"
	qos := 0
	message := "MQTT from golang"

	receipt := client.Publish(MQTT.QoS(qos), topic, message)
	<-receipt // Publish成功を待ち受ける

	return nil
}
