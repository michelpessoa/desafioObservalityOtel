package cep

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	internalErrors "github.com/michelpessoa/desafioObservalityOtel/internal/platform/errors"
	httpPlatform "github.com/michelpessoa/desafioObservalityOtel/internal/platform/http"

	"go.opentelemetry.io/otel/trace"
)

const (
	NotFoundCEP      = "ERR_NOT_FOUND_CEP"
	FailedGetInfo    = "ERR_FAILED_GET_CEP"
	FailedUnmarshall = "ERR_FAILED_UNMARSHALL_CEP"
)

type Service interface {
	GetInfo(ctx context.Context, request Request) (Response, error)
}

type service struct {
	tracer    trace.Tracer
	client    *http.Client
	apiConfig APIConfig
}

func (s service) GetInfo(ctx context.Context, request Request) (Response, error) {
	ctx, span := s.tracer.Start(ctx, "cep-api")
	defer span.End()

	req, err := httpPlatform.NewReq(ctx, http.MethodGet, fmt.Sprintf(s.apiConfig.URL, request.Cep), nil)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedGetInfo, err)
	}

	res, err := s.client.Do(req)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedGetInfo, err)
	}

	if res.StatusCode == http.StatusNotFound {
		return Response{}, internalErrors.NewNotFoundError(NotFoundCEP, err)
	}
	if res.StatusCode != http.StatusOK {
		return Response{}, internalErrors.NewApplicationError(FailedGetInfo,
			errors.New(fmt.Sprintf("status_code:%d", res.StatusCode)))
	}

	var resp Response
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedUnmarshall, err)
	}

	if resp.City == "" {
		return Response{}, internalErrors.NewNotFoundError(NotFoundCEP, err)
	}

	return resp, nil
}

func NewService(tracer trace.Tracer, client *http.Client, config APIConfig) Service {
	return &service{
		tracer:    tracer,
		client:    client,
		apiConfig: config,
	}
}
