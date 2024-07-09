package main

import (
	"cake-store/utils/config"
	"cake-store/utils/database"
	"cake-store/utils/logger"
	"fmt"
	"net/http"
)

func main() {
	config := config.GetConfig()
	logger := logger.Newlogger()
	l := logger.LogWithContext("server", "main")

	dbVar := database.DBServiceVar{
		DbUri:   &config.DB_URI,
		Dialect: &config.DB_DIALECT,
		Logger:  logger,
	}

	database.NewDB(&dbVar)

	server := http.Server{
		Addr: fmt.Sprintf("%1s:%2s", config.HOST, config.PORT),
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	l.Info("Starting HTTP server on port ", config.PORT)
	select {}
}
