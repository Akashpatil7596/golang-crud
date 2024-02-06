package main

import (
	"akash/controllers"
	"akash/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func init(){
	initializers.ConnectToDB()
	
}

func main(){
	
	fmt.Println("Hello")
	
	
	r := gin.Default()

	r.POST("/ping/:name", controllers.CreateUser)
	r.Run("127.0.0.1:8080") 

}


