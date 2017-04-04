package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	IsHttpServer   bool   `json:"is_http_server"`
	HttpServerPort string `json:"http_server_port"`
	HttpServerHost string `json:"http_server_host"`
}

var config Configuration

func init() {
	config = Configuration{}
}

func LoadConfig() {
	parseConfig()
}

func GetConfig() Configuration {
	return config
}

func parseConfig() {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Printf("Reading config file error: %s", err)
		os.Exit(1)
	}

	json.Unmarshal(raw, &config)
	return
}
