package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

const YMDFormat = "2006-01-02 15:04:05 MST"
var client *mongo.Client

func main() {
	fmt.Println("serving...")
	err := Serve()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
