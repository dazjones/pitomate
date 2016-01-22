package pitomate

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

func getConf() Configuration {
	file, _ := os.Open("pitomate.conf")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return configuration
}
