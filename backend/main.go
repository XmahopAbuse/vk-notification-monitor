package main

import (
	"log"
	config "vk-notification-monitor/config"
	"vk-notification-monitor/server"
)

func main() {

	cfg := config.NewConfig()
	app, err := server.NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
