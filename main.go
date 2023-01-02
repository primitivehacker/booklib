package main

import (
	"github.com/gin-gonic/gin"
	"github.com/primitivehacker/booklib/api"
)

func main() {
	router := gin.Default()
	router.GET("/books", api.L.GetAllBooks)
	router.GET("/books/:id", api.BookById)
	router.POST("/books", api.CreateBook)
	router.PATCH("/checkout", api.CheckOutBook)
	router.PATCH("/return", api.ReturnBook)
	router.Run("localhost:8080")
}
