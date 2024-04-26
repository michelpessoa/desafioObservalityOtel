package input

import (
	"encoding/json"
	"strings"

	"github.com/michelpessoa/desafioObservalityOtel/internal/input/orchestration"
	"github.com/michelpessoa/desafioObservalityOtel/internal/platform/errors"
)

const (
	ErrRequired = "ERR_CEP_REQUIRED"
	ErrInvalid  = "ERR_CEP_INVALID"
)

type Request struct {
	CEP string `json:"cep"`
}

type Response struct {
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
	City           string  `json:"city"`
}

func (r *Request) Validate() error {
	if r.CEP == "" {
		return errors.NewUnprocessableError(ErrRequired)
	}

	r.CEP = strings.ReplaceAll(r.CEP, "-", "")
	if len(r.CEP) != 8 {
		return errors.NewUnprocessableError(ErrInvalid)
	}

	return nil
}

func NewResponse(resp orchestration.Response) Response {
	return Response{
		TempCelsius:    resp.TempCelsius,
		TempFahrenheit: resp.TempFahrenheit,
		TempKelvin:     resp.TempKelvin,
		City:           resp.City,
	}
}

func (r *Response) ToJSON() []byte {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return jsonBytes
}
