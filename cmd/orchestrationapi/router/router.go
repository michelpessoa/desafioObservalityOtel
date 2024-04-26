package router

import (
	"github.com/michelpessoa/desafioObservalityOtel/internal/orchestration"
	"github.com/michelpessoa/desafioObservalityOtel/internal/platform/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

func SetupRouter(serviceName string, handler orchestration.Handler) *mux.Router {
	r := mux.NewRouter()

	r.Use(otelmux.Middleware(serviceName))
	r.Use(http.RequestIDMiddleware)

	r.HandleFunc("/orchestrate/{cep}", handler.GetTemperatureByCEP)
	return r
}
