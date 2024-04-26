package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) ToJSON() []byte {
	jsonBytes, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return jsonBytes
}

func Encode(err error) *AppError {
	var unprocessableError UnprocessableError
	var notFoundError NotFoundError
	var badRequestError BadRequestError
	switch {
	case errors.As(err, &unprocessableError):
		return &AppError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	case errors.As(err, &notFoundError):
		return &AppError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
	case errors.As(err, &badRequestError):
		return &AppError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	default:
		return &AppError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
}
