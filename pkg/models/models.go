package models

type Person struct {
	Name    string `json:"name" bson:"name"`
	Age     int    `json:"age" bson:"age"`
	IsAdult bool   `json:"isAdult" bson:"isAdult"`
}
