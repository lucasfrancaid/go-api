package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasfrancaid/go-api/controllers"
)

func (r RouterShema) BookRoutes(rg *gin.RouterGroup) {
	book := rg.Group("/books")

	book.POST("", controllers.CreateBook)
	book.GET("", controllers.GetAllBooks)
	book.GET("/:id", controllers.GetBook)
	book.PUT("/:id", controllers.UpdateBook)
	book.DELETE("/:id", controllers.DeleteBook)
}
