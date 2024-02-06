package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateUser(c *gin.Context) {

	name:= c.Param("name")
	query:= c.Query("search")

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}

	fmt.Printf("Received POST request with body: %s\n", body)

	fmt.Println(query)

	c.JSON(200, gin.H{
		"message": "Hello" + " " + name,
	})
}