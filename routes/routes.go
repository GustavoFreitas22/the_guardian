package routes

import (
	"github.com/GustavoFreitas22/the_guardian/controller"
	"github.com/gin-gonic/gin"
)

func HandleRoutes() {

	server := gin.Default()

	server.POST("/secret", controller.PostInfoData)
	server.GET("/:id", controller.GetInfoDataById)
	server.GET("/secrets/:context", controller.GetDataByContext)
	server.PUT("/secret", controller.EditInfoData)
	server.DELETE(("/secret/:id"), controller.DeleteInfoData)

	server.Run(":5000")

}
