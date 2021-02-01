module mail-backend

go 1.15

require (
	github.com/Viva-con-Agua/echo-pool v1.0.4
	github.com/Viva-con-Agua/vcago v0.0.8
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.2
	github.com/jinzhu/configor v1.2.1
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.17
	github.com/labstack/gommon v0.3.0
	github.com/nats-io/nats.go v1.10.0
	go.mongodb.org/mongo-driver v1.4.5
	golang.org/x/net v0.0.0-20201216054612-986b41b23924
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

replace github.com/Viva-con-Agua/vcago => ../github.com/Viva-con-Agua/vcago
