package weather

type Request struct {
	Query string `json:"query"`
}

type Response struct {
	Current Current `json:"current"`
}

type Current struct {
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
}

type APIConfig struct {
	URL    string
	APIKey string
}
