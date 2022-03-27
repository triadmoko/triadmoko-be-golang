package handler

import (
	"net/http"
	"triadmoko-be-golang/entity"
	"triadmoko-be-golang/formatter"
	"triadmoko-be-golang/helper"
	"triadmoko-be-golang/mapping"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateFaskes(c *gin.Context) {
	var input mapping.InputFaskes

	err := c.ShouldBindJSON(&input)
	if err != nil {
		error := helper.FormatValidationError(input)
		errorMessage := gin.H{"errors": error}

		response := helper.ResponseApi("Insert Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	insert, err := h.service.InputFaskes(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseApi("Insert Failed", 201, "error", errorMessage)
		c.JSON(201, response)
		return
	}
	formatJSON := formatter.FormatterFaskes(insert)
	response := helper.ResponseApi("Insert Success", http.StatusOK, "success", formatJSON)

	c.JSON(http.StatusOK, response)
}
func (h *handler) FindAllFaskes(c *gin.Context) {
	var data []entity.Faskes
	data, err := h.service.FindAllFaskes()
	if err != nil {
		data := gin.H{"error": err}
		response := helper.ResponseApi("Failed Request Faskes", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
	}

	formatJSON := formatter.FormattFaskesAll(data)
	response := helper.ResponseApi("Get Faskes Success", http.StatusOK, "success", formatJSON)

	c.JSON(http.StatusOK, response)
}
