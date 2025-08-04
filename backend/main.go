package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //router
	r.Use(cors.Default())
	api := r.Group("/api")
	{
		api.POST("/simulate", simulateTuringMachine)
		api.GET("/example", getExamples)
		api.GET("/example/:name", getExample)
	}
	r.Static("/static", "./frontend/static")
	r.LoadHTMLGlob("frontend/templates/*")

}
