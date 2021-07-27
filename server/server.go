package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ash3798/url-shortner/config"
)

//Start starts the http server
func Start() {
	addr := fmt.Sprintf(":%d", config.Manager.ApplicationPort)

	mux := http.NewServeMux()
	mux.HandleFunc("/url", HandleRequest)

	server := http.Server{Addr: addr, Handler: mux}

	shutdownChan := make(chan string)

	log.Printf("Starting HTTP server at %s:%d", "localhost", config.Manager.ApplicationPort)
	go func(shChan chan string) {
		err := server.ListenAndServe()
		if err != nil {
			shChan <- err.Error()
		}
	}(shutdownChan)

	msg := <-shutdownChan
	log.Println("Error in server , Error: ", msg)
}
