package main

import "time"

type todo struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	DueDate string `json:"due"`
}

var todoList = []todo{
	{ID: 1, Title: "Putzen", Content: "Frühjahrsputz",
		DueDate: time.Date(2018, 05, 20, 0, 0, 0, 0, time.UTC).Format("02-01-2006")},
	{ID: 2, Title: "Soziale Netzwerke Hausaufgabe", Content: "Vorstellen einer todo-Webapp",
		DueDate: time.Date(2018, 05, 14, 0, 0, 0, 0, time.UTC).Format("02-01-2006")},
}

func getAllTodos() []todo  {
	return todoList
}