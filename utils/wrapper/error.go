package wrapper

import "net/http"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// BadRequest struct
type RespondError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type NotFoundError struct {
	Error string
}

type ConflictError struct {
	Error string
}

// NewBadRequest
func NewBadRequest(message string) RespondError {
	errObj := RespondError{}
	errObj.Message = message
	errObj.Code = http.StatusBadRequest

	return errObj
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func NewConflictError(err string) ConflictError {
	return ConflictError{Error: err}
}
