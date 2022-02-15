package config

import "github.com/spf13/viper"

type Config struct {
	DBuser string `mapstructure:"POSTGRES_USER"`
	DBpass string `mapstructure:"POSTGRES_PASSWORD"`
	DBname string `mapstructure:"POSTGRES_DB"`
	DBport string `mapstructure:"DB_PORT"`
	DBhost string `mapstructure:"DB_HOST"`

	ServerPort string `mapstructure:"SERVER_PORT"`

	RedisHost string `mapstructure:"R_HOST"`
	RedisPort string `mapstructure:"R_PORT"`

	CHuser string `mapstructure:"CH_USER"`
	CHpass string `mapstructure:"CH_PASSWORD"`
	CHname string `mapstructure:"CH_DB_NAME"`
	CHport string `mapstructure:"CH_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.AddConfigPath("/app")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return config, nil
}
