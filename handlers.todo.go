// handlers.todo.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
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

	if c.Request.URL.Query().Get("id") == "" || c.Request.URL.Query().Get("id")=="0"{
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
		if todoID, err := strconv.Atoi(c.Request.URL.Query().Get("id")); err == nil {

			if todo, err := getTodoById(todoID); err == nil {
				// Call the HTML method of the Context to render a template
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