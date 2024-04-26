package cep

type Request struct {
	Cep string `json:"cep"`
}

type Response struct {
	City  string `json:"localidade" validate:"required"`
	State string `json:"uf" validate:"required"`
}

type APIConfig struct {
	URL string
}
