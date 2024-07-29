package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-api-rest/internal/config/database"
	"github.com/go-api-rest/internal/controllers"
	"github.com/go-api-rest/internal/repositories"
	"github.com/go-api-rest/internal/services"
	"github.com/go-api-rest/pkg/router"
)

var PORT = ":3000"

func main() {

	r := gin.Default()

	userRepo := repositories.NewUserRepo(database.NewConnection())
	userServices := services.NewUserService(userRepo)
	userController := controllers.NerUserController(userServices)

	router.InitRouter(r, userController)

	defer func() {
		if reco := recover(); reco != nil {
			fmt.Println("Eneabling custom Port")
			r.Run()
		}
	}()

	err := r.Run(PORT)
	if err != nil {
		panic("[ERROR] failed to start Gin server: " + err.Error())
	} // if is PORT missing, server use the default port
}
