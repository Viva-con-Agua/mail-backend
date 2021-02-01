package utils

import (
	"github.com/jinzhu/configor"
)

var (
	Config = struct {
		Mail struct {
			Host     string
			From     string
			User     string
			Password string
		}
		Nats struct {
			Url string
		}
		DB struct {
			Host     string `required:"true"`
			User     string `default:"admin"`
			Password string `default:"admin"`
			Port     int    `default:"3306"`
			Name     string `required:"true"`
		}
		Alloworigins []string
	}{}
)

func LoadConfig() {
	configor.Load(&Config, "config/config.yml")
}
