package router

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func (r RouterShema) AuthorRoutes(rg *gin.RouterGroup) {
	author := rg.Group("/authors")

	author.POST("", controllers.CreateAuthor)
	author.GET("", controllers.GetAllAuthors)
	author.GET("/:id", controllers.GetAuthor)
	author.PUT("/:id", controllers.UpdateAuthor)
	author.DELETE("/:id", controllers.DeleteAuthor)
}
