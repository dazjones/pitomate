package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Broker   string
	Hostname string
	MySQLUser string
	MySQLPass string
	MySQLHost string
	MySQLName string
	MySQLPort string
}

func GetConf() Configuration {
	file, _ := os.Open("pitomate.conf")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return configuration
}
