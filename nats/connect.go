package nats

import (
	"log"
	"mail-backend/env"

	"github.com/Viva-con-Agua/vcago/verr"
	nats "github.com/nats-io/nats.go"
)

//Nats represent the global Nats client
var Nats = new(nats.EncodedConn)


//Connect connects nats client to server. The client is reachable over Nats.
func Connect() {
	natsUrl := "nats://" + env.NatsHost + ":" + env.NatsPort
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(verr.ErrorWithColor, err , " ", "NatsUrl: ", natsUrl)
	}
	Nats, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(verr.ErrorWithColor, err)
	}
	log.Print("nats successfully connected!")
}
