package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var todo Todo
	var input TodoInput
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		panic(err)
	}
	todo = TodoFromInput(input)
	todo.SetCTime()
	todo.SetUTime()
	todo.SetStatus()
	collection := client.Database("tododatabase").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, todo)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"`))
		return
	}
	if dl := TimeParser(todo.Deadline); dl.Sub(TimeParser(todo.CreatedTime)).Hours() < 72 {
		response.Write([]byte(`"message": "3 günden az kaldı!"`))
	}
	json.NewEncoder(response).Encode(result)
}

func GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var todos []Todo
	collection := client.Database("tododatabase").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := request.URL.Query()
	fmt.Println(query)
	fmt.Println(query["title"])

	processes := []string{"inprogress", "done", "waiting"}
	statuses := []string{"Normal", "Acil"}
	categories := CatSlice()
	before := "21 Apr 99 00:00 +03"
	after := "01 Jan 00 00:00 +03"
	if len(query) != 0 {
		if prosl, ex := query["process"]; ex {
			for _, process := range prosl {
				processes = append(processes, process)
			}
		}
		if stasl, ex := query["status"]; ex {
			for _, status := range stasl {
				statuses = append(statuses, status)
			}
		}
		if catsl, ex := query["category"]; ex {
			for _, category := range catsl {
				categories = append(categories, category)
			}
		}
		if befsl, ex := query["before"]; ex {
				before = befsl[0]
		}
		if aftsl, ex := query["after"]; ex {
				after = aftsl[0]
		}
		if exp, ex := query["expired"]; ex {
			if exp[0] == "true" {
				before = time.Now().Format(YMDFormat)
			} else if exp[0] == "false" {
				after = time.Now().Format(YMDFormat)
			}
		}
	}

	param := bson.M{"process": bson.M{"$in": processes}, "status": bson.M{"$in": statuses}, "category": bson.M{"$in": categories},
		"deadline": bson.M{"$lt": before, "$gt": after} }

	cursor, err := collection.Find(ctx, param)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Todo
		cursor.Decode(&person)
		todos = append(todos, person)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"`))
		return
	}
	json.NewEncoder(response).Encode(todos)
}

func GetTodo(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"`))
		return
	}
	var todo Todo
	collection := client.Database("tododatabase").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, Todo{ID: id}).Decode(&todo)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"`))
		return
	}
	json.NewEncoder(response).Encode(todo)
}

func Update(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"`))
		return
	}
	var todo Todo
	_ = json.NewDecoder(request.Body).Decode(&todo)
	todo.SetStatus()
	todo.SetUTime()
	collection := client.Database("tododatabase").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	updated := bson.M{"$set": todo,}
	result, upderr := collection.UpdateOne(ctx, Todo{ID: id}, updated)
	if upderr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + upderr.Error() + `"`))
		return
	}
	json.NewEncoder(response).Encode(result)
}


