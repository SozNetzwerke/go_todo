package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	aList := getAllTodos()

	if len(aList) != len(todoList) {
		t.Fail()
	}

	for i, v := range aList {
		if v.ID != todoList[i].ID ||
			v.Title != todoList[i].Title ||
			v.Content != todoList[i].Content ||
			v.DueDate != todoList[i].DueDate {
				t.Fail()
				break
		}
	}
}
