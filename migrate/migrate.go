package main

import (
	"akash/initializers"
	"akash/models"
)

func init(){
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.User{})
}