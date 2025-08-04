package main

import (
<<<<<<< HEAD
	"encoding/json"
	"log"
	"net/http"
	"strconv"

=======
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/gin-contrib/cors"
>>>>>>> 9f1fe1fdfb5dca5f61d86c58b4ec8eb020820509
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
<<<<<<< HEAD


func simulateTuringMachine(c *gin.Context) {}

func getExamples(c *gin.Context) {}

func getExample(c *gin.Context) {}
=======
=======
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
=======
>>>>>>> 033d287 (feat: add structs and improve code organization)
=======
>>>>>>> 3bfb323552cbb36413a2a7258edcff776262d231
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
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> 29a7883 (feat: add packages)
=======
>>>>>>> 033d287 (feat: add structs and improve code organization)
=======
>>>>>>> 3bfb323552cbb36413a2a7258edcff776262d231
>>>>>>> 9f1fe1fdfb5dca5f61d86c58b4ec8eb020820509
