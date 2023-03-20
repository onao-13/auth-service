package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config Настройки сервера
type Config struct {
	Port      string `mapstructure:"PORT"`
	DbUser    string
	DbPass    string
	DbHost    string
	DbPort    string
	DbName    string
	JwtSecret string `mapstructure:"JWTSECRET"`
}

var log = logrus.New()

// UploadDev Загрузка конфигурации для разработки
func UploadDev() Config {
	return uploadConfig("dev")
}

func UploadProd() Config {
	return uploadConfig("prod")
}

// Загрузка файла с конфигурацией
func uploadConfig(fileName string) Config {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName(fileName)
	v.SetConfigType("env")

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("Error reading config. Error: ", err)
	}

	var config Config

	err = v.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Error parse config. Error: ", err)
	}

	return config
}

// DbUrl путь для подключения к базе данных
func (c *Config) DbUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", c.DbUser, c.DbPass,
		c.DbHost, c.DbPort, c.DbName)
}
