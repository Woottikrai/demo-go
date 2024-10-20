package main

//go mod is dependencies (package.json)
//go sum is hash  dependencies (package-lock.json)

//go mod init projectName // update package
//go mod tidy manage dependencies
// echo swagger

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	connectionUrl := "mongodb://localhost:27017/?retryWrites=true&loadBalanced=false&connectTimeoutMS=10000"

	clientOptions := options.Client().ApplyURI(connectionUrl)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Ping error: ", err)
	}

	fmt.Println("Connected to MongoDB!")
	fmt.Println("Hello, World!")

	http.ListenAndServe(":8080", nil)
}
