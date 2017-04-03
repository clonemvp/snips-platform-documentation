package mqtt

import (
	"fmt"
	paho_mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type clientWrapper struct {
	client paho_mqtt.Client
	broker string
}

type MessageHandler func(topic string, payload []byte)

func NewClient(host string, port int) *clientWrapper {
	opts := paho_mqtt.NewClientOptions()
	broker := fmt.Sprintf("tcp://%s:%d", host, port)
	opts.AddBroker(broker)
	opts.OnConnect = func(client paho_mqtt.Client) {
		log.Printf("Connection to broker %v succeded", broker)
	}
	opts.OnConnectionLost = func(client paho_mqtt.Client, reason error) {
		log.Printf("Connection lost from broker %v -> %v", broker, reason.Error())
	}
	opts.AutoReconnect = true
	return &clientWrapper{
		client: paho_mqtt.NewClient(opts),
		broker: broker,
	}
}

func (cw *clientWrapper) Connect() error {
	log.Printf("Connecting to broker %v...", cw.broker)
	token := cw.client.Connect()
	token.Wait()
	return token.Error()
}

func (cw *clientWrapper) Publish(topic string, payload []byte) error {
	token := cw.client.Publish(topic, byte(1), false, payload)
	token.Wait()
	return token.Error()
}

func (cw *clientWrapper) Subscribe(topic string, handler MessageHandler) error {
	token := cw.client.Subscribe(topic, byte(1), func(client paho_mqtt.Client, message paho_mqtt.Message) {
		handler(message.Topic(), message.Payload())
	})
	token.Wait()
	return token.Error()
}

func (cw *clientWrapper) Disconnect() {
	cw.client.Disconnect(0)
}
