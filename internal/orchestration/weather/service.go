package weather

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	internalErrors "github.com/michelpessoa/desafioObservalityOtel/internal/platform/errors"
	httpPlatform "github.com/michelpessoa/desafioObservalityOtel/internal/platform/http"

	"go.opentelemetry.io/otel/trace"
)

const (
	FailedGetInfo    = "ERR_FAILED_GET_WEATHER"
	FailedUnmarshall = "ERR_FAILED_UNMARSHALL_WEATHER"
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
	ctx, span := s.tracer.Start(ctx, "weather-api")
	defer span.End()

	weatherURL := fmt.Sprintf(s.apiConfig.URL, s.apiConfig.APIKey, url.QueryEscape(request.Query))

	req, err := httpPlatform.NewReq(ctx, http.MethodGet, weatherURL, nil)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedGetInfo, err)
	}

	res, err := s.client.Do(req)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedGetInfo, err)
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

	return resp, nil
}

func NewService(tracer trace.Tracer, client *http.Client, apiConfig APIConfig) Service {
	return &service{tracer: tracer, client: client, apiConfig: apiConfig}
}
