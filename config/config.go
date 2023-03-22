package config

type AppConfig struct {
  Environment string `env:"APP_ENV,required"`
  
}
