package router

import (
	"github.com/michelpessoa/desafioObservalityOtel/internal/input"
	"github.com/michelpessoa/desafioObservalityOtel/internal/platform/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

func SetupRouter(serviceName string, handler input.Handler) *mux.Router {
	r := mux.NewRouter()

	r.Use(otelmux.Middleware(serviceName))
	r.Use(http.RequestIDMiddleware)

	r.HandleFunc("/input", handler.GetTemperatureByCEP)
	return r
}
