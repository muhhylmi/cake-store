package cakes

import (
	"cake-store/modules/cakes/handlers"
	"cake-store/utils/wrapper"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(cakeHandler handlers.CakeHandler) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/cakes", cakeHandler.Create)
	router.PanicHandler = wrapper.ErrorHandler

	return router
}
