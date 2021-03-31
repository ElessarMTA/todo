package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type	Todo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Process    string             `json:"process,omitempty" bson:"process,omitempty"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
	Deadline    primitive.DateTime             `json:"deadline,omitempty" bson:"deadline,omitempty"`
	CreatedTime    primitive.DateTime             `json:"createdTime,omitempty" bson:"createdTime,omitempty"`
	UpdatedTime    primitive.DateTime             `json:"updatedTime,omitempty" bson:"updatedTime,omitempty"`
	}

type TodoInput struct {
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Process    string             `json:"process,omitempty" bson:"process,omitempty"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
	Deadline    string             `json:"deadline,omitempty" bson:"deadline,omitempty"`
}

func TodoFromInput(input TodoInput) Todo{
	var out Todo
	out.Title = input.Title
	out.Description = input.Description
	out.Category = input.Category
	out.Process = input.Process
	out.Status = input.Status
	out.Deadline = primitive.NewDateTimeFromTime(TimeParser(input.Deadline))
	return out
}

func TimeParser(date string) time.Time {
	a, _ := time.Parse(time.RFC3339, date)
	return a
}

/*func NewTodo(
	id primitive.ObjectID,
	title string,
	description string,
	category string,
	process string,
	deadline string,
	createdTime string,
	updatedTime string,
	status string,
) *Todo {
	return &Todo{
		ID: id,
		Title:  title,
		Description: description,
		Category: category,
		Process: process,
		Deadline: deadline,
		CreatedTime: createdTime,
		UpdatedTime: updatedTime,
		Status: status,
	}
}*/

func (x *Todo) SetCTime() {
	x.CreatedTime = primitive.NewDateTimeFromTime(time.Now())
}

func (x *Todo) SetUTime() {
	x.UpdatedTime = primitive.NewDateTimeFromTime(time.Now())
}

func (x *Todo) SetStatus() {
	if strings.Contains(strings.ToLower(x.Description), "acil") || strings.Contains(strings.ToLower(x.Description), "acıl") {
		x.Status = "Acil"
	} else {
		x.Status = "Normal"
	}
}