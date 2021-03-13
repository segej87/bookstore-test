package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/bookstore-test/models"
  "github.com/segej87/bookstore-test/controllers"
)

func main() {
  r := gin.Default()
  
  models.ConnectDatabase()
  
  r.GET("/books", controllers.FindBooks)

  r.Run()
}

