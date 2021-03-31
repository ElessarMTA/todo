package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var client *mongo.Client

func main() {
	fmt.Println("serving...")
	err := Serve()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
