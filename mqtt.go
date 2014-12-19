package main

import (
	"fmt"
	"log"

	"time"

	"os"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

type MQTTConfig struct {
	host     string
	port     int
	user     string
	password string
}

var config = MQTTConfig{
	host:     "lite.mqtt.shiguredo.jp",
	port:     1883,
	user:     "kyokomi@github",
	password: os.Getenv("SANGO_MQTT_PASSWORD"),
}

func mqttExample() {

	// 接続用の設定を作成します
	opts := MQTT.NewClientOptions()
	opts.SetUsername(config.user)
	opts.SetPassword(config.password)

	// golangではURIで接続先を指定します。
	brokerUri := fmt.Sprintf("tcp://%s:%d", config.host, config.port)
	opts.AddBroker(brokerUri)

	// 設定を元にクライアントを作成します
	client := MQTT.NewClient(opts)

	// Connectします
	_, err := client.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatalln(Subscribe(client, "say"))
}

func onMessageReceived(client *MQTT.MqttClient, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\n", message.Topic())
	fmt.Printf("Message: %s\n", message.Payload())

	if string(message.Payload()) == "ぬるぽ" {
		Publish(client, "say", "ガッ")
	}
}

func Subscribe(client *MQTT.MqttClient, topic string) error {
	topicName := config.user + "/" + topic
	// 無料枠はQoS 0のみ
	qos := 0

	// Subscribeするtopicを設定します
	topicFilter, err := MQTT.NewTopicFilter(topicName, byte(qos))
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

func Publish(client *MQTT.MqttClient, topic, message string) error {
	topicName := config.user + "/" + topic
	qos := 0

	receipt := client.Publish(MQTT.QoS(qos), topicName, message)
	<-receipt // Publish成功を待ち受ける

	return nil
}
