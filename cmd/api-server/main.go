package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	api_server "github.com/asianParadissseee/http-rest/internal/app/api-server"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/api-server.toml", "path to config file")
}

func main() {
	config := api_server.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := api_server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
