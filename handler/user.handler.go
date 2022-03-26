package handler

import (
	"net/http"
	"triadmoko-be-golang/entity"
	"triadmoko-be-golang/helper"
	"triadmoko-be-golang/mapping"
	"triadmoko-be-golang/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	userService service.Service
	authService auth.Service
}

func NewHandlerUser(userService service.Service) *handler {
	return &handler{userService}
}

func (h *handler) GetUser(c *gin.Context) {
	var user entity.User
	user, err := h.userService.GetUser()
	if err != nil {
		data := mapping.UserFormatter(user)
		response := helper.ResponseApi("Failed", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
	}
	data := mapping.UserFormatter(user)
	response := helper.ResponseApi("Succes", http.StatusOK, "failed", data)
	c.JSON(http.StatusOK, response)

}

func (h *handler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas di passing sebagai parameter service
	var input mapping.FormatUser

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// error to object
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}

		response := helper.ResponseApi("Register Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// validate confirm password
	if input.Password != input.Confirm_Password {
		response := helper.ResponseApi("Register Failed", http.StatusUnprocessableEntity, "error", "Confirm password is wrong!")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.InputRegister(input)

	if err != nil {
		errorMessage := gin.H{"errors ": err.Error()}

		response := helper.ResponseApi("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID, newUser.Level, newUser.Email, newUser.FName, newUser.Phone, newUser.Avatar, newUser.Address)
	if err != nil {
		response := helper.ResponseApi("Register Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseApi("Account Has been registered", http.StatusOK, "success", token)

	c.JSON(http.StatusOK, response)
}
