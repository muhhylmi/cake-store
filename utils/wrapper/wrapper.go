package wrapper

import (
	"encoding/json"
	"net/http"
)

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ReadJsonFromRequest(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		panic(err)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	res := response.(WebResponse)
	if res.Code != 0 {
		writer.WriteHeader(res.Code)
	}

	enconder := json.NewEncoder(writer)
	err := enconder.Encode(response)
	if err != nil {
		panic(err)
	}
}
