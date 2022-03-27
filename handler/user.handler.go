package handler

import (
	"net/http"
	"triadmoko-be-golang/auth"
	"triadmoko-be-golang/formatter"
	"triadmoko-be-golang/helper"
	"triadmoko-be-golang/mapping"
	"triadmoko-be-golang/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	userService service.Service
	authService auth.Service
}

func NewHandlerUser(userService service.Service, authService auth.Service) *handler {
	return &handler{userService, authService}
}
func (h *handler) Login(c *gin.Context) {
	var input mapping.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		error := helper.FormatValidationError(input)
		errorMessage := gin.H{"errors": error}

		response := helper.ResponseApi("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseApi("Login Failed", 201, "error", errorMessage)
		c.JSON(201, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.ResponseApi("Login Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatJSON := formatter.UserJsonFormatter(token)
	response := helper.ResponseApi("Login Success", http.StatusOK, "success", formatJSON)

	c.JSON(http.StatusOK, response)

}

func (h *handler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas di passing sebagai parameter service
	var input formatter.FormatUser

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

		response := helper.ResponseApi("Register Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// token, err := h.authService.GenerateToken(newUser.ID)
	// if err != nil {
	// 	response := helper.ResponseApi("Register Failed", http.StatusBadRequest, "error", nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	format := formatter.FormateResponseRegisterUser(newUser)
	response := helper.ResponseApi("Register Succesfully ", http.StatusOK, "success", format)

	c.JSON(http.StatusOK, response)
}
