package middleware

import (
	"cake-store/utils/config"
	"cake-store/utils/wrapper"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
	Config  *config.Configurations
}

func NewAuthMiddleware(handler http.Handler, config *config.Configurations) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler, Config: config}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("API-Key") == config.GetConfig().API_KEY {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := wrapper.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		wrapper.WriteToResponseBody(writer, webResponse)
	}
}
