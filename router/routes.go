package router

import (
	"github.com/gin-gonic/gin"
)

type RouterShema struct {
	router *gin.Engine
}

func Routes() RouterShema {
	r := RouterShema{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")
	r.AuthorRoutes(v1)
	r.BookRoutes(v1)

	r.router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	return r
}

func (r RouterShema) Run() error {
	return r.router.Run()
}
