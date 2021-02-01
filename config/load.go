package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

var SmptHost string
var SmptPort string 

func GetEnv(key string) string{
    val, ok := os.LookupEnv(key)
    if !ok {
        log.Print("%s not set\n", key)
        return nil
    }
    return os.Getenv(key)
}
