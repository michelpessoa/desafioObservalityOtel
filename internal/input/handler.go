package input

import (
	"encoding/json"
	"net/http"

	"github.com/michelpessoa/desafioObservalityOtel/internal/platform/errors"
)

type Handler interface {
	GetTemperatureByCEP(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

func (h handler) GetTemperatureByCEP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appErr := errors.Encode(errors.NewUnprocessableError(err.Error()))
		w.WriteHeader(appErr.Code)
		w.Write(appErr.ToJSON())
		return
	}

	temp, err := h.service.GetTemperatureByCEP(r.Context(), req)
	if err != nil {
		appErr := errors.Encode(err)
		w.WriteHeader(appErr.Code)
		w.Write(appErr.ToJSON())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(temp.ToJSON())
	return
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}
