package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_HOST     string `mapstructure:"DATABASE_HOST"`
	DATABASE_PORT     int    `mapstructure:"DATABASE_PORT"`
	DATABASE_NAME     string `mapstructure:"DATABASE_NAME"`
	DATABASE_USER     string `mapstructure:"DATABASE_USER"`
	DATABASE_PASSWORD string `mapstructure:"DATABASE_PASSWORD"`
	API_PORT          string `mapstructure:"API_PORT"`
	JWT_SECRET        string `mapstructure:"JWT_SECRET"`
	SECRET            string `mapstructure:"SECRET"`
}

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func InitConfigs(filename string) Config {

	if _, err := os.Stat(filename); err == nil {
		viper.SetConfigFile(filename)
		viper.ReadInConfig()
	} else {
		viper.SetDefault("DATABASE_HOST", os.Getenv("DATABASE_HOST"))
		viper.SetDefault("DATABASE_PORT", ConvertInt("DATABASE_PORT"))
		viper.SetDefault("DATABASE_NAME", os.Getenv("DATABASE_NAME"))
		viper.SetDefault("DATABASE_USER", os.Getenv("DATABASE_USER"))
		viper.SetDefault("DATABASE_PASSWORD", os.Getenv("DATABASE_PASSWORD"))
		viper.SetDefault("API_PORT", os.Getenv("API_PORT"))
		viper.SetDefault("JWT_SECRET", os.Getenv("JWT_SECRET"))
		viper.SetDefault("SECRET", os.Getenv("SECRET"))
	}

	var config Config
	viper.Unmarshal(&config)

	return config
}
