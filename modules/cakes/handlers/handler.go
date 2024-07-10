package handlers

import (
	"cake-store/modules/cakes/usecases"
	"cake-store/utils/logger"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

const context = "module.cake.handlers"

type HandlerImpl struct {
	Logger   *logger.Logger
	Usecase  usecases.Usecases
	Validate *validator.Validate
}

type CakeHandler interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

func NewCakeHandler(logger *logger.Logger, usecase usecases.Usecases, validate *validator.Validate) CakeHandler {
	return &HandlerImpl{
		Logger:   logger,
		Usecase:  usecase,
		Validate: validate,
	}
}