package main

import (
	"gobookapi/api"

	"github.com/gin-gonic/gin"
)

func main(){
	api.InitDB()
	r := gin.Default()

	r.POST("/books", api.CreateBook)
	r.GET("/books", api.GetBooks)
	r.GET("/books/:id", api.GetBook)
	r.PUT("/books/:id", api.UpdateBook)
	r.DELETE("/books/:id", api.DeleteBook)

	r.Run(":8080")

}