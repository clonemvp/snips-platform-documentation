package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type clientWrapper struct {
	client mqtt.Client
	broker string
}

type MessageHandler func(topic string, payload []byte)

func NewClient(host string, port int) *clientWrapper {
	opts := mqtt.NewClientOptions()
	broker := fmt.Sprintf("tcp://%s:%d", host, port)
	opts.AddBroker(broker)
	opts.OnConnect = func(client mqtt.Client) {
		log.Printf("Connection to broker %v succeded", broker)
	}
	opts.OnConnectionLost = func(client mqtt.Client, reason error) {
		log.Printf("Connection lost from broker %v -> %v", broker, reason.Error())
	}
	opts.AutoReconnect = true
	return &clientWrapper{
		client: mqtt.NewClient(opts),
		broker: broker,
	}
}

func (cw *clientWrapper) Connect() error {
	log.Printf("Connecting to broker %v...", cw.broker)
	token := cw.client.Connect()
	token.Wait()
	return token.Error()
}

func (cw *clientWrapper) Subscribe(topic string, handler MessageHandler) error {
	token := cw.client.Subscribe(topic, byte(1), func(client mqtt.Client, message mqtt.Message) {
		handler(message.Topic(), message.Payload())
	})
	token.Wait()
	return token.Error()
}

func (cw *clientWrapper) Disconnect() {
	cw.client.Disconnect(0)
}
