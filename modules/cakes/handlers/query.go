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

func (handler *HandlerImpl) List(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	l := handler.Logger.LogWithContext(context, "List")

	size := request.URL.Query().Get("size")
	page := request.URL.Query().Get("page")
	keyword := request.URL.Query().Get("q")

	listRequest := web.CakeListRequest{
		Size: 10,
		Page: 1,
	}

	if size != "" {
		size, err := strconv.Atoi(size)
		if err != nil {
			l.Error(err)
			wrapper.PanicIfError(err)
		}
		listRequest.Size = size
	}

	if page != "" {
		page, err := strconv.Atoi(page)
		if err != nil {
			l.Error(err)
			wrapper.PanicIfError(err)
		}
		listRequest.Page = page
	}

	if keyword != "" {
		listRequest.Keyword = keyword
	}

	err := handler.Validate.Struct(listRequest)
	if err != nil {
		l.Error(err.Error())
		wrapper.PanicIfError(err)
	}

	categoryResponses := handler.Usecase.List(request.Context(), &listRequest)
	webResponse := wrapper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}
