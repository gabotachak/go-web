package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id       int       `form:"id" json:"Id"`
	Code     string    `form:"code" json:"Code"`
	Currency string    `form:"currency" json:"Currency"`
	Mount    float64   `form:"mount" json:"Mount"`
	Sender   string    `form:"sender" json:"Sender"`
	Receiver string    `form:"receiver" json:"Receiver"`
	Date     time.Time `form:"date" json:"Date"`
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

func getTransactions() []Transaction {
	content, err := ioutil.ReadFile("transactions.json")
	check(err)
	var transactions []Transaction
	json.Unmarshal([]byte(content), &transactions)
	return transactions
}

func getAll(c *gin.Context) {
	transactions := getTransactions()
	c.JSON(200, gin.H{
		"transactions": transactions,
	})
}

func getByFilter(c *gin.Context) {
	transactions := getTransactions()
	filterId, err := strconv.Atoi(c.Param("id"))
	if err != nil || filterId < 0 {
		c.JSON(400, gin.H{
			"error": "Bad request, only integers",
		})
		return
	}
	var filtered []Transaction
	for _, transaction := range transactions {
		if transaction.Id == filterId {
			filtered = append(filtered, transaction)
		}
	}
	if len(filtered) == 0 {
		c.JSON(404, gin.H{
			"error": "Not found any coincidence",
		})
		return
	}
	c.JSON(200, gin.H{
		"transactions": filtered,
	})
}

func getByQuery(c *gin.Context) {
	var filter Transaction
	if c.Bind(&filter) != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}
	transactions := getTransactions()
	var filtered []Transaction
	for _, t := range transactions {
		filtered = append(filtered, t)
	}
	if len(filtered) == 0 {
		c.JSON(404, gin.H{
			"error": "Not found any coincidence",
		})
		return
	}
	c.JSON(200, gin.H{
		"transactions": filtered,
	})
}

func main() {
	router := gin.Default()
	router.GET("/", hello)
	router.GET("/hello-world", hello)
	router.GET("/transactions", getAll)
	router.GET("/transactions/:id", getByFilter)
	router.GET("/transactions/", getByQuery)
	router.Run(":8080")
}
