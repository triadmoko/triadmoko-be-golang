package handler

import (
	"net/http"
	"triadmoko-be-golang/formatter"
	"triadmoko-be-golang/helper"
	"triadmoko-be-golang/mapping"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateNakes(c *gin.Context) {
	var input mapping.InputNakes

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// error to object
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}

		response := helper.ResponseApi("Insert Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newNakes, err := h.service.InputNakes(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ResponseApi("Insert Failed", 201, "error", errorMessage)
		c.JSON(201, response)
		return
	}
	formatJSON := formatter.FormatterNakes(newNakes)
	response := helper.ResponseApi("Insert Success", http.StatusOK, "success", formatJSON)

	c.JSON(http.StatusOK, response)
}
