package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/nimamakhmali/turing-machine-go/turing"
)

type SimulateRequest struct {
	Definition turing.TuringMachineDefinition `json:"definition"`
	MaxSteps   int                           `json:"maxSteps"`
}

type SimulateResponse struct {
	Result     string   `json:"result"`
	FinalTape  []string `json:"finalTape"`
	Steps      int      `json:"steps"`
	History    []Step   `json:"history,omitempty"`
}

type Step struct {
	State     string   `json:"state"`
	Tape      []string `json:"tape"`
	Head      int      `json:"head"`
	Step      int      `json:"step"`
	Action    string   `json:"action"`
}

func main() {
	r := gin.Default()    //router
	r.Use(cors.Default())
	api := r.Group("/api")
	{
		api.POST("/simulate", simulateTuringMachine)
		api.GET("/examples", getExamples)
		api.GET("/examples/:name", getExample)
	}
	r.Static("/static", "./frontend/static")
	r.LoadHTMLGlob("frontend/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	log.Println("Server starting on :8080")
	r.Run(":8080")
}

func simulateTuringMachine(c *gin.Context) {}

func getExamples(c *gin.Context) {
	examples := []gin.H{
		{"name": "simple", "title": "Simple Example", "description": "Basic TM that accepts strings with 0"},
		{"name": "palindrome", "title": "Palindrome Check", "description": "TM that checks for palindromes"},
		{"name": "binary", "title": "Binary Counter", "description": "TM that increments binary numbers"},
	}
	c.JSON(http.StatusOK, examples)
}

func getExample(c *gin.Context) {
	name := c.Param("name")
	
	// Load example from file
	// This would load from examples/ directory
	c.JSON(http.StatusOK, gin.H{"message": "Example " + name + " loaded"})
} 