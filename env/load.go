package env

import (
	"log"

	"github.com/Viva-con-Agua/vcago"
	"github.com/joho/godotenv"
)

//NatsHost is the ip of the nats service.
var NatsHost string
//NatsPort is the port ot the nats service.
var NatsPort string
//AllowOrigins used in cors header
var AllowOrigins []string
//MailSMTPHost is the ip of the smtp relay
var MailSMTPHost string
//MailSMTPPort is the port of the smpt relay
var MailSMTPPort int

//LoadConfig loads the environment variables form .env file and handle errors
func LoadConfig() {
    var loadEnv vcago.LoadEnv
    godotenv.Load()
    NatsHost, loadEnv = loadEnv.GetEnvString("NATS_HOST", "w", "localhost")
    NatsPort, loadEnv = loadEnv.GetEnvString("NATS_PORT", "w", "4222") 
    AllowOrigins, loadEnv = loadEnv.GetEnvStringList("ALLOW_ORIGINS", "w", []string{"localhost:8080"})
    MailSMTPHost, loadEnv = loadEnv.GetEnvString("MAIL_SMTP_HOST", "w", "localhost")
    MailSMTPPort, loadEnv = loadEnv.GetEnvInt("MAIL_SMTP_HOST", "w", 25)
    log.Print(loadEnv)
    loadEnv.Validate()

}

//LoadConfig loads the environment variables form .env file and handle errors
