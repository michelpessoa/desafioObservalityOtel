package orchestration

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/michelpessoa/desafioObservalityOtel/internal/orchestration/cep"
	"github.com/michelpessoa/desafioObservalityOtel/internal/orchestration/weather"
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

func (r *Response) ToJSON() []byte {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return jsonBytes
}

func NewRequest(cep string) Request {
	return Request{
		CEP: cep,
	}
}

func (r *Request) BuildCEPRequest() cep.Request {
	return cep.Request{Cep: r.CEP}
}

func NewWeatherRequest(response cep.Response) weather.Request {
	query := fmt.Sprintf("%s,%s", response.City, response.State)
	return weather.Request{Query: query}
}

func NewResponse(cepResp cep.Response, weatherResp weather.Response) Response {
	return Response{
		TempCelsius:    weatherResp.Current.TempCelsius,
		TempFahrenheit: weatherResp.Current.TempFahrenheit,
		TempKelvin:     math.Round(weatherResp.Current.TempCelsius+273.15*100) / 100,
		City:           cepResp.City,
	}
}
