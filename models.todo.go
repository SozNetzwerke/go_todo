package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"errors"
)

type todo struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Datum   string `json:"datum"`
	Fortschritt float64 `json:"fortschritt"`
}


func getAllTodos() []todo {

	var (
		id int
		content string
		datum string
		fortschritt float64
		todoList []todo
	)

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/sn")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var t todo
	for rows.Next() {
		err := rows.Scan(&id, &content,&datum,&fortschritt)
		if err != nil {
			log.Fatal(err)
		}
		t=todo{ID:id, Content:content, Datum:datum, Fortschritt:fortschritt}
		todoList = append(todoList, t)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}


	return todoList
}

func getTodoById(id int) (*todo, error) {
	todoList:=getAllTodos()
	for _, a := range todoList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("TODO not found")
}