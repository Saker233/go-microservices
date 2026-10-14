package main

import "go.mongodb.org/mongo-driver/mongo"

const (
	webPort  = "8000"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	grpcPort = "50001"
)

var client *mongo.Client

type Config struct {

}

func main() {
}
