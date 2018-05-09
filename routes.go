package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(){


	router.POST("/index.html/delete/:todo_id", deleteTodo)
	router.GET("/", showIndexPage)

	router.GET("/index.html", showIndexPage)

	router.GET("/edit.html", getTodo)

	router.POST("/edit.html/update/:todo_id" ,updateTodo)

	router.GET("createTODO.html", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"createTODO.html",
			gin.H{
				"title": "Create TODO",
				"navitem2":"active",
			},
		)
	})
	router.GET("impressum.html", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"impressum.html",
			gin.H{
				"title": "Impressum",
				"navitem4":"active",
			},
		)
	})

	router.POST("/createTODO.html/create", createTodo)
	router.GET("howTo.html", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"howTo.html",
			gin.H{
				"title": "How To",
				"navitem3":"active",
			},
		)
	})
}
