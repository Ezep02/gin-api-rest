package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-api-rest/internal/controllers"
)

func InitRouter(router *gin.Engine, userController *controllers.UserController) {

	router.GET("/user/:id", userController.GetUserController)

	router.POST("/user", userController.CreateNewUser)
}
