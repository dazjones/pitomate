package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	BrokerUrl  []string
	BrokerPort []string
	Hostname   []string
}

func GetConf() Configuration {
	file, _ := os.Open("/opt/pitomate.conf")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return configuration
}
