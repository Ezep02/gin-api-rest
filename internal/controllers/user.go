package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-api-rest/internal/models"
	"github.com/go-api-rest/internal/services"
)

type UserController struct {
	service *services.UserService
}

func NerUserController(UserService *services.UserService) *UserController {

	return &UserController{
		service: UserService,
	}
}

func (uc *UserController) GetUserController(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing id parameter",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Invalid id parameter",
		})
		return
	}

	user, err := uc.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) CreateNewUser(c *gin.Context) {
	var user models.User

	// Extraer datos del cuerpo de la solicitud
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Agregar el nuevo usuario
	createdUser, err := uc.service.AddNewUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Responder con el usuario creado
	c.JSON(http.StatusOK, createdUser)
}
