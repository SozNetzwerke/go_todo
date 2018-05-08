// handlers.todo.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
	"log"
)

func showIndexPage(c *gin.Context) {
	todos := getAllTodos()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home",
			"navitem1":"active",
			"payload": todos,
		},
	)
}

func getTodo(c *gin.Context) {

	if c.Query("id") == "" || c.Request.URL.Query().Get("id")=="0"{
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"error.html",
			// Pass the data that the page uses
			gin.H{
				"title":   "ERROR",
			},
		)
	}else {
		if todoID, err := strconv.Atoi(c.Query("id")); err == nil {

			if todo, err := getTodoById(todoID); err == nil {
				// Call the HTML method of the Context to render a template
				log.Println(todo.Datum)
				c.HTML(
					// Set the HTTP status to 200 (OK)
					http.StatusOK,
					// Use the index.html template
					"edit.html",
					// Pass the data that the page uses
					gin.H{
						"title":   todoID,
						"Content": todo.Content,
						"Fortschritt": todo.Fortschritt,
						"Datum": todo.Datum,
					},
				)

			} else {
				// If the article is not found, abort with an error
				c.AbortWithError(http.StatusNotFound, err)
			}

		} else {
			// If an invalid article ID is specified in the URL, abort with an error
			c.AbortWithStatus(http.StatusNotFound)
		}
		}
}

func updateTodo(c *gin.Context){
	update(c.Param("todo_id"),c.PostForm("aufgabe_input"), c.PostForm("date_input"), c.PostForm("fortschritt_input"))

	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/index.html")
}

func deleteTodo(c *gin.Context){
	rm(c.Param("todo_id"))
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/index.html")
}

func createTodo(c *gin.Context){
	create(c.PostForm("aufgabe_input"), c.PostForm("date_input"), c.PostForm("fortschritt_input"))
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/index.html")
}