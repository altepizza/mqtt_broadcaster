package main

import (
	"log"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var targetPrefixes = []string{}
	targetPrefixes = strings.Split(viper.GetString("targetPrefixes"), ",")

	log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())
	for _, prefix := range targetPrefixes {
		newTopic := strings.Replace(msg.Topic(), viper.GetString("sourcePrefix"), prefix, 1)
		token := client.Publish(newTopic, 0, false, msg.Payload())
		token.Wait()
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
	client.Subscribe(viper.GetString("sourcePrefix")+"/#", 0, nil)
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connection lost: %v", err)
}

func main() {
	viper.SetEnvPrefix("MQTT_DISTRIBUTER")
	viper.AutomaticEnv()
	opts := mqtt.NewClientOptions()
	opts.AddBroker(viper.GetString("brokerHost") + ":" + viper.GetString("brokerPort"))
	opts.SetUsername(viper.GetString("username"))
	opts.SetPassword(viper.GetString("password"))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		time.Sleep(1 * time.Second)
	}
}
