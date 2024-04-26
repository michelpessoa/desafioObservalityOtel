package config

import (
	"github.com/michelpessoa/desafioObservalityOtel/internal/input/orchestration"

	"github.com/spf13/viper"
)

type Config struct {
	OrchestrationURL     string `mapstructure:"ORCHESTRATION_URL"`
	ServiceName          string `mapstructure:"OTEL_SERVICE_INPUT_NAME"`
	OTELExporterEndpoint string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
}

func (c *Config) OrchClientConfig() orchestration.APIConfig {
	return orchestration.APIConfig{OrchestrationURL: c.OrchestrationURL}
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
