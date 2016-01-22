package main

import (
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"github.com/dazjones/pitomate/conf"
)

var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	conf := conf.GetConf()
	fmt.Println(conf)

	opts := MQTT.NewClientOptions().AddBroker("")
	fmt.Println(opts)
}
