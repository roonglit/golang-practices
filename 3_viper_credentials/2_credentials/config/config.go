package config

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	Port    string `mapstructure:"PORT"`
	DBUri   string `mapstructure:"DB_URI"`
}
