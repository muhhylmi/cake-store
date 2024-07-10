package main

import (
	"cake-store/modules/cakes"
	"cake-store/modules/cakes/handlers"
	repository "cake-store/modules/cakes/repositories"
	usecase "cake-store/modules/cakes/usecases"

	"cake-store/utils/config"
	"cake-store/utils/database"
	"cake-store/utils/logger"
	"cake-store/utils/middleware"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

func main() {
	config := config.GetConfig()
	logger := logger.Newlogger()
	l := logger.LogWithContext("server", "main")
	validate := validator.New()

	dbVar := database.DBServiceVar{
		DbUri:   &config.DB_URI,
		Dialect: &config.DB_DIALECT,
		Logger:  logger,
	}
	db := database.NewDB(&dbVar)

	cakeRepo := repository.NewRepository(logger, db)
	cakeUseCase := usecase.NewUsecase(logger, cakeRepo)
	cakeHandler := handlers.NewCakeHandler(logger, cakeUseCase, validate)
	router := cakes.NewRouter(cakeHandler)

	server := http.Server{
		Addr:    fmt.Sprintf("%1s:%2s", config.HOST, config.PORT),
		Handler: middleware.NewAuthMiddleware(router, config),
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
