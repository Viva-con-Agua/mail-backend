package dao

import (
	"context"
	"log"
	"os"
	//"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB represents the database
var DB = new(mongo.Database)

//DBNAME database name
var DBNAME = "mail"

//Client represents the database client
var Client = new(mongo.Client)
//Connect connects the DB-Client with mongo database.
func Connect() {
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	uri := "mongodb://" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	Client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("database connection failed", err)
	}
    err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("database connection failed", err)
	}
	DB = Client.Database(DBNAME)
}
