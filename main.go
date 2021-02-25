package main

import (
	"github.com/JenYata/goEx/controllers"
	"github.com/JenYata/goEx/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()
	r.GET("/books", controllers.FindBooks)

	r.POST("/books", controllers.CreateBook)

	r.GET("/books/:id", controllers.FindBook)

	r.PATCH("/books/:id", controllers.UpdateBook)

	r.DELETE("/book/:id", controllers.DeleteBook)
	r.Run()

}
