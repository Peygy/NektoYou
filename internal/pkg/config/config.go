package config

import (
	"github.com/peygy/nektoyou/internal/pkg/gin"
)

type Config struct {
	Gin 	*gin.GinConfig
}

func NewConfig(filePath string) (*Config, *gin.GinConfig, error) {

}
// создание конфига для сервисов
// типа GinConfig (port и т.п)
// grpcConfig (port и бла бла)

/* Пример
Главный конфиг
type Config struct {
	ServiceName  string                        `mapstructure:"serviceName"`
	Logger       *logger.LoggerConfig          `mapstructure:"logger"`
	Rabbitmq     *rabbitmq.RabbitMQConfig      `mapstructure:"rabbitmq"`
	Echo         *echoserver.EchoConfig        `mapstructure:"echo"`
	Grpc         *grpc.GrpcConfig              `mapstructure:"grpc"`
	GormPostgres *gormpgsql.GormPostgresConfig `mapstructure:"gormPostgres"`
	Jaeger       *otel.JaegerConfig            `mapstructure:"jaeger"`
}

func InitConfig() (*Config, *logger.LoggerConfig, *otel.JaegerConfig, *gormpgsql.GormPostgresConfig,
	*grpc.GrpcConfig, *echoserver.EchoConfig, *rabbitmq.RabbitMQConfig, error) {

	......

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, cfg.Logger, cfg.Jaeger, cfg.GormPostgres, cfg.Grpc, cfg.Echo, cfg.Rabbitmq, nil
}


json
  "echo": {
    "port": ":5002",
    "development": true,
    "timeout": 30,
    "basePath": "/api/v1",
    "host": "http://localhost",
    "debugHeaders": true,
    "httpClientDebug": true,
    "debugErrorsResponse": true,
    "ignoreLogUrls": [
      "metrics"
    ]
  },
  "grpc": {
    "port": ":6600",
    "host": "localhost",
    "development": true
  },
*/