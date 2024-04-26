package orchestration

import (
	"context"

	"github.com/michelpessoa/desafioObservalityOtel/internal/orchestration/cep"
	"github.com/michelpessoa/desafioObservalityOtel/internal/orchestration/weather"
	log "github.com/michelpessoa/desafioObservalityOtel/internal/platform/log"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Service interface {
	GetTemperatureByCEP(ctx context.Context, request Request) (Response, error)
}

type service struct {
	weatherService weather.Service
	cepService     cep.Service
}

func (s service) GetTemperatureByCEP(ctx context.Context, request Request) (Response, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("cep", request.CEP))

	log.Info(ctx, "starting cep2temp", log.Tag("cep", request.CEP))

	cepRequest := request.BuildCEPRequest()
	cepResponse, err := s.cepService.GetInfo(ctx, cepRequest)
	if err != nil {
		log.Error(ctx, "error getting cep info", err)
		return Response{}, err
	}
	span.SetAttributes(attribute.String("city", cepResponse.City))
	span.SetAttributes(attribute.String("state", cepResponse.State))

	weatherRequest := NewWeatherRequest(cepResponse)
	weatherResponse, err := s.weatherService.GetInfo(ctx, weatherRequest)
	if err != nil {
		log.Error(ctx, "error getting weather info", err)
		return Response{}, err
	}

	resp := NewResponse(cepResponse, weatherResponse)
	span.SetAttributes(attribute.Float64("celsius", resp.TempCelsius))
	log.Info(ctx, "finished cep2temp", log.Tag("cep", request.CEP), log.Tag("temp_c", resp.TempCelsius),
		log.Tag("temp_f", resp.TempFahrenheit), log.Tag("temp_k", resp.TempKelvin),
		log.Tag("city", resp.City))

	return resp, nil
}

func NewService(cepService cep.Service, weatherService weather.Service) Service {
	return &service{
		weatherService: weatherService,
		cepService:     cepService,
	}
}
