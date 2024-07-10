package handlers

import (
	"cake-store/modules/cakes/models/web"
	"cake-store/utils/wrapper"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (handler *HandlerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	l := handler.Logger.LogWithContext(context, "Create")
	categoryCreateRequest := &web.CakeCreateRequest{}
	wrapper.ReadJsonFromRequest(request, &categoryCreateRequest)

	err := handler.Validate.Struct(categoryCreateRequest)
	if err != nil {
		l.Error(err.Error())
		panic(err)
	}

	cakeResponse := handler.Usecase.Create(request.Context(), categoryCreateRequest)
	webResponse := wrapper.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   cakeResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}

func (handler *HandlerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	l := handler.Logger.LogWithContext(context, "Update")

	cakeUpdateRequest := web.CakeUpdateRequest{}
	wrapper.ReadJsonFromRequest(request, &cakeUpdateRequest)

	cakeId := params.ByName("cakeId")
	id, err := strconv.Atoi(cakeId)
	if err != nil {
		l.Error(err)
		panic(err)
	}
	cakeUpdateRequest.Id = id

	err = handler.Validate.Struct(cakeUpdateRequest)
	if err != nil {
		l.Error(err)
		panic(err)
	}

	cakeResponse := handler.Usecase.Update(request.Context(), &cakeUpdateRequest)
	webResponse := wrapper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cakeResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}

func (handler *HandlerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	l := handler.Logger.LogWithContext(context, "Delete")

	cakeId := params.ByName("cakeId")
	id, err := strconv.Atoi(cakeId)
	if err != nil {
		l.Error(err)
		panic(err)
	}
	res := handler.Usecase.Delete(request.Context(), id)
	webResponse := wrapper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}
