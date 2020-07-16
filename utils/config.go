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
	}{}
)

func LoadConfig() {
	configor.Load(&Config, "config/config.yml")
}
