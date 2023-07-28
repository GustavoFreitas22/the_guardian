package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRoutes() {

	server := gin.Default()
	server.Use(cors.Default())

	api := server.Group("/api")
	{
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, "Salve salve")
		})
	}

	server.Run(":5000")

}
