package main

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

func Serve() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var conerr error
	client, conerr = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if conerr != nil {
		panic(conerr)
	}

	router := mux.NewRouter()
	router.HandleFunc("/todo", Create).Methods("POST")
	router.HandleFunc("/todo", GetAll).Methods("GET")
	router.HandleFunc("/todo/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", Update).Methods("PUT")
	router.HandleFunc("/category", CreateCat).Methods("POST")
	router.HandleFunc("/category", GetAllCat).Methods("GET")
	return http.ListenAndServe(":12345", router)
}