package handlers

import (
	"cake-store/modules/cakes/models/web"
	"cake-store/utils/wrapper"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (handler *HandlerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	l := handler.Logger.LogWithContext(context, "FindById")

	cakeId := params.ByName("cakeId")
	id, err := strconv.Atoi(cakeId)
	if err != nil {
		l.Error(err)
		wrapper.PanicIfError(err)
	}
	req := web.CakeGetRequest{
		CakeId: id,
	}

	err = handler.Validate.Struct(req)
	if err != nil {
		l.Error(err.Error())
		wrapper.PanicIfError(err)
	}

	categoryResponse := handler.Usecase.FindById(request.Context(), &req)
	webResponse := wrapper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}
