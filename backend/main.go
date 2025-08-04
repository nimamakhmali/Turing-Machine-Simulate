package main

import (
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


func simulateTuringMachine(c *gin.Context) {}

func getExamples(c *gin.Context) {}

func getExample(c *gin.Context) {}