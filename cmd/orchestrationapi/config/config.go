package config

import (
	"github.com/michelpessoa/desafioObservalityOtel/internal/orchestration/cep"
	"github.com/michelpessoa/desafioObservalityOtel/internal/orchestration/weather"

	"github.com/spf13/viper"
)

type Config struct {
	WeatherURL           string `mapstructure:"WEATHER_URL"`
	WeatherKey           string `mapstructure:"WEATHER_API_KEY"`
	CepURL               string `mapstructure:"CEP_URL"`
	ServiceName          string `mapstructure:"OTEL_SERVICE_ORCHESTRATE_NAME"`
	OTELExporterEndpoint string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
}

func (c *Config) WeatherAPIConfig() weather.APIConfig {
	return weather.APIConfig{
		URL:    c.WeatherURL,
		APIKey: c.WeatherKey,
	}
}

func (c *Config) CepAPIConfig() cep.APIConfig {
	return cep.APIConfig{
		URL: c.CepURL,
	}
}

func LoadConfig(path string) (cfg Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
