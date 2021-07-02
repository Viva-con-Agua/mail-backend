module mail-backend

go 1.15

require (
	github.com/Viva-con-Agua/vcago v0.1.15
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.2.0
	github.com/jinzhu/configor v1.2.1
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo/v4 v4.1.17
	github.com/labstack/gommon v0.3.0
	github.com/nats-io/nats-server/v2 v2.3.1 // indirect
	github.com/nats-io/nats.go v1.11.1-0.20210623165838-4b75fc59ae30
	go.mongodb.org/mongo-driver v1.4.6
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

//replace github.com/Viva-con-Agua/vcago => ../github.com/Viva-con-Agua/vcago
