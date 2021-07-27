package main

import (
	"log"

	"github.com/ash3798/url-shortner/config"
	"github.com/ash3798/url-shortner/server"
)

func main() {
	log.Println("Starting URL Shortener")

	if !config.InitEnv() {
		return
	}

	server.Start()
}
