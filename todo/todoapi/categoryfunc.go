package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func CreateCat(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var cat Category
	if err := json.NewDecoder(request.Body).Decode(&cat); err != nil {
		panic(err)
	}
	collection := client.Database("tododatabase").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, cat)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func GetAllCat(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var cats []Category
	collection := client.Database("tododatabase").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var cat Category
		cursor.Decode(&cat)
		cats = append(cats, cat)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"`))
		return
	}
	json.NewEncoder(response).Encode(cats)
}

func CatSlice() []string {
	var cats []string
	collection := client.Database("tododatabase").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var cat Category
		cursor.Decode(&cat)
		cats = append(cats, cat.Name)
	}
	return cats
}