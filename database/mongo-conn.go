package database

import (
	"context"
	"time"

	utility "utility"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

//DbConn to establish connection - define a timeout.
//create context and pass it to the connection
func DbConn() {
	mongoURL := "mongodb://localhost:27017"

	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	utility.CheckError(err)
	client, err = mongo.Connect(ctx, mongoURL)
	utility.CheckError(err)

	return client
}
