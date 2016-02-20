package main

import (
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"github.com/dazjones/pitomate/conf"
	"github.com/dazjones/pitomate/subscriptions"
	"os"
	"github.com/dazjones/pitomate/djio"
)

var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	sub := subscriptions.Get(msg.Topic())
	fmt.Println(sub)

	if string(msg.Payload()) == "ON" {
		djio.PinHigh(sub.Pin)
	} else {
		djio.PinLow(sub.Pin)
	}
}

func main() {
	cnf := conf.GetConf()

	opts := MQTT.NewClientOptions().AddBroker(cnf.Broker)
	opts.SetClientID("go-simple")
	opts.SetDefaultPublishHandler(f)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	subs := subscriptions.GetAll()

	for _, sub := range subs {
		if token := c.Subscribe(sub.Name, 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}

	for {

	}
}
