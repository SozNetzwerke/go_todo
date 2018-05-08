package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
	"errors"
)

/**
Container Todo
 */
type todo struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Datum   string `json:"datum"`
	Fortschritt float64 `json:"fortschritt"`
}

/**
return a list of Todos
 */
func getAllTodos() []todo {

	var (
		id int
		content string
		datum string
		fortschritt float64
		todoList []todo
	)


	// connection to db
	db, err := sql.Open("sqlite3", "db/SN.db")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	//Query to select
	rows, err := db.Query("select * from todo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var t todo

	//append to list
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

/**
param id int
return a tupel (todo error)
 */
func getTodoById(id int) (*todo, error) {
	todoList:=getAllTodos()
	for _, a := range todoList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("TODO not found")
}

/**
update db by changing values of todos
 */
func update(id string, content string, datum string, fortschritt string,){
	//connection
	db, err := sql.Open("sqlite3", "db/SN.db")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	//important " " for inputs
	sqlStatement := `
		UPDATE todo
		SET content ="`+content +`", datum = "`+datum+`", fortschritt = "`+fortschritt+`"
		WHERE id = "`+id+`";`


	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Println("failed")
		panic(err)
	}else{log.Println("success")}

}

func rm(id string){
	//connection
	db, err := sql.Open("sqlite3", "db/SN.db")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	sqlStatement := `
		DELETE FROM todo
		WHERE id="`+id+`";`


	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Println("failed")
		panic(err)
	}else{log.Println("success")}

}

func create(cont string, dat string, fort string){
	//connection
	db, err := sql.Open("sqlite3", "db/SN.db")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	sqlStatement := `
		INSERT INTO todo ("content", "datum", "fortschritt")
		VALUES ("`+cont+`", "`+dat+`", "`+fort+`") ;`


	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Println("failed")
		panic(err)
	}else{log.Println("success")}
}