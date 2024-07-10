package handlers

import (
	"cake-store/modules/cakes/models/web"
	"cake-store/utils/wrapper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (handler *HandlerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	l := handler.Logger.LogWithContext(context, "Create")
	categoryCreateRequest := &web.CakeCreateRequest{}
	wrapper.ReadJsonFromRequest(request, &categoryCreateRequest)

	err := handler.Validate.Struct(categoryCreateRequest)
	if err != nil {
		l.Error(err.Error())
		wrapper.PanicIfError(err)
	}

	categoryResponse := handler.Usecase.Create(request.Context(), categoryCreateRequest)
	webResponse := wrapper.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   categoryResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}
