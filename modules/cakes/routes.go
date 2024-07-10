package cakes

import (
	"cake-store/modules/cakes/handlers"
	"cake-store/utils/wrapper"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(cakeHandler handlers.CakeHandler) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/cakes", cakeHandler.Create)
	router.GET("/api/cakes", cakeHandler.List)
	router.GET("/api/cakes/:cakeId", cakeHandler.FindById)
	router.PATCH("/api/cakes/:cakeId", cakeHandler.Update)

	router.PanicHandler = wrapper.ErrorHandler

	return router
}
