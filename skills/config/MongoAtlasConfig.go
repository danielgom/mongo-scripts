package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"time"
)

var (
	Client *mongo.Client
)

func Connect() {

	dbUrl := "mongodb+srv://Daniel_GA:Dank3409@cluster0.flnpc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

	clientOptions := options.Client().ApplyURI(dbUrl).SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	clientOptions.SetRetryWrites(false)
	clientOptions.SetMaxPoolSize(100)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error connecting to MongoDB, %s", err.Error()))
	}

	err = Client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalln(fmt.Sprintf("error pinging MongoDB %s", err.Error()))
	}

	log.Println("Successful connection to MongoDB")
}
 
