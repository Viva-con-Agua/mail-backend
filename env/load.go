package env

import (
	"log"

	"github.com/Viva-con-Agua/vcago"
	"github.com/joho/godotenv"
)

//AllowOrigins used in cors header
var AllowOrigins []string
//MailSMTPHost is the ip of the smtp relay
var MailSMTPHost string
//MailSMTPPort is the port of the smpt relay
var MailSMTPPort int

//LoadConfig loads the environment variables form .env file and handle errors
func LoadConfig() {
    res := true
    var ok bool
    godotenv.Load()
    if AllowOrigins, ok = vcago.GetEnvStringList("ALLOW_ORIGINS", "w", []string{"localhost:8080"}); !ok {
        res = false
    }
    if MailSMTPHost, ok = vcago.GetEnvString("MAIL_SMTP_HOST", "w", "127.0.0.1"); !ok {
        res = false
    }
    if MailSMTPPort, ok = vcago.GetEnvInt("MAIL_SMTP_HOST", "w", 25); !ok {
        res = false
    }
    if !res {
        log.Fatal("Please set enviroment variables in the .env file. Read logs above.")
    }
}
