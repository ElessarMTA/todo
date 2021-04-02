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
	Deadline    string             `json:"deadline,omitempty" bson:"deadline,omitempty"`
	CreatedTime    string             `json:"createdTime,omitempty" bson:"createdTime,omitempty"`
	UpdatedTime    string             `json:"updatedTime,omitempty" bson:"updatedTime,omitempty"`
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
	out.Deadline = input.Deadline
	return out
}

func TimeParser(date string) time.Time {
	a, _ := time.Parse(YMDFormat, date)
	return a
}

func (x *Todo) SetCTime() {
	x.CreatedTime = time.Now().Format(YMDFormat)
}

func (x *Todo) SetUTime() {
	x.UpdatedTime = time.Now().Format(YMDFormat)
}

func (x *Todo) SetStatus() {
	if strings.Contains(strings.ToLower(x.Description), "acil") || strings.Contains(strings.ToLower(x.Description), "acıl") {
		x.Status = "Acil"
	} else {
		x.Status = "Normal"
	}
}