package main

import (
	"cake-store/utils/config"
	"fmt"
	"net/http"
)

func main() {
	config := config.GetConfig()

	server := http.Server{
		Addr: fmt.Sprintf("%1s:%2s", config.HOST, config.PORT),
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Starting HTTP server on port ", config.PORT)
	select {}
}
