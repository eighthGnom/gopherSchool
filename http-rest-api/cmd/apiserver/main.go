package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/eighthGnom/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config_path", "configs/apiserver.toml", "Path to the .toml config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Cant decode config file, used default settings", err)
	}
	server := apiserver.New(config)
	log.Fatal(server.Start())
}
