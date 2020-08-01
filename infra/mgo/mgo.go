package mgo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Common methods
*/

func getDatabase() *mongo.Database {

	clientOpts := options.Client().ApplyURI(os.Getenv("DATABASE_URL"))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		fmt.Println("Error during connection:  ", err)
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Error during connection check:  ", err)
	}
	fmt.Println("Congratulations, you're already connected to MongoDB!")

	return client.Database("hoiLightningTalk")
}
