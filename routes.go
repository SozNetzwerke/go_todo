package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/contrib/static"
)

func initializeRoutes(){

	router.Use(static.Serve("/img", static.LocalFile("./img", true)))
	router.StaticFile("/template/progress.js", "./template/progress.js")



	router.GET("/", showIndexPage)

	router.GET("/index.html", showIndexPage)

	router.GET("/edit.html", getTodo)

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