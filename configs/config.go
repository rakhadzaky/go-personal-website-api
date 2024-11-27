package configs

type MainConfig struct {
	AppPort string `env:"APP_PORT"`
	DatabaseConfig
}
