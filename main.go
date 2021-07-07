package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id       int       `json:"Id"`
	Code     string    `json:"Code"`
	Currency string    `json:"Currency"`
	Mount    float64   `json:"Mount"`
	Sender   string    `json:"Sender"`
	Receiver string    `json:"Receiver"`
	Date     time.Time `json:"Date"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hola Gabriel Andres",
	})
}

func getAll(c *gin.Context) {
	content, err := ioutil.ReadFile("transactions.json")

	check(err)

	var transactions []Transaction
	json.Unmarshal([]byte(content), &transactions)

	c.JSON(200, transactions)
}

func main() {
	router := gin.Default()
	router.GET("/hello-world", hello)
	router.GET("/transactions", getAll)
	router.Run()
}
