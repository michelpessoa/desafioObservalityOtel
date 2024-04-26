package input

import (
	"context"

	"github.com/michelpessoa/desafioObservalityOtel/internal/input/orchestration"
	"github.com/michelpessoa/desafioObservalityOtel/internal/platform/log"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Service interface {
	GetTemperatureByCEP(ctx context.Context, request Request) (Response, error)
}

type service struct {
	orchClient orchestration.Client
}

func (s service) GetTemperatureByCEP(ctx context.Context, request Request) (Response, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("cep", request.CEP))

	if err := request.Validate(); err != nil {
		log.Error(ctx, "error validating request", err)
		return Response{}, err
	}

	resp, err := s.orchClient.GetTemperatureByCEP(ctx, request.CEP)
	if err != nil {
		log.Error(ctx, "error calling orchestration service", err)
		return Response{}, err
	}

	span.SetAttributes(attribute.String("city", resp.City))
	span.SetAttributes(attribute.Float64("celsius", resp.TempCelsius))
	return NewResponse(resp), nil
}

func NewService(orchClient orchestration.Client) Service {
	return &service{
		orchClient: orchClient,
	}
}
