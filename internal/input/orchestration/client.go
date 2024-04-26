package orchestration

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	internalErrors "github.com/michelpessoa/desafioObservalityOtel/internal/platform/errors"
	httpPlatform "github.com/michelpessoa/desafioObservalityOtel/internal/platform/http"
	"github.com/michelpessoa/desafioObservalityOtel/internal/platform/log"

	"go.opentelemetry.io/otel/trace"
)

const (
	FailedOrchestration = "ERR_EXECUTING_ORCHESTRATION"
	FailedUnmarshall    = "ERR_FAILED_UNMARSHALL_ORCHESTRATION"
	NotFoundCEP         = "ERR_CANNOT_FIND_ZIPCODE"
)

type Client interface {
	GetTemperatureByCEP(ctx context.Context, cep string) (Response, error)
}

type clientHandler struct {
	tracer    trace.Tracer
	client    *http.Client
	apiConfig APIConfig
}

func (c clientHandler) GetTemperatureByCEP(ctx context.Context, cep string) (Response, error) {
	ctx, span := c.tracer.Start(ctx, "orchestration-api")
	defer span.End()

	url := fmt.Sprintf("%s/%s", c.apiConfig.OrchestrationURL, cep)

	log.Info(ctx, "calling orchestration api", log.Tag("url", url))

	req, err := httpPlatform.NewReq(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedOrchestration, err)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedOrchestration, err)
	}
	if res.StatusCode == http.StatusNotFound {
		return Response{}, internalErrors.NewNotFoundError(NotFoundCEP,
			errors.New(fmt.Sprintf("status_code:%d", res.StatusCode)))
	}
	if res.StatusCode != http.StatusOK {
		return Response{}, internalErrors.NewApplicationError(FailedOrchestration,
			errors.New(fmt.Sprintf("status_code:%d", res.StatusCode)))
	}

	var resp Response
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedUnmarshall, err)
	}

	return resp, nil
}

func NewClient(tracer trace.Tracer, client *http.Client, apiConfig APIConfig) Client {
	return &clientHandler{tracer: tracer, client: client, apiConfig: apiConfig}
}
