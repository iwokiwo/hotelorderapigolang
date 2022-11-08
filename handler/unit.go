package handler

import (
	"iwogo/auth"
	"iwogo/helper"
	"iwogo/unit"
	"net/http"

	"github.com/gin-gonic/gin"
)

type unitHandler struct {
	unitService unit.Service
	authService auth.Service
}

func NewUnitHandler(unitService unit.Service, authService auth.Service) *unitHandler {
	return &unitHandler{unitService, authService}
}

func (h *unitHandler) CreateUnit(c *gin.Context) {

	var input unit.RegisterUnitInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register category failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.unitService.RegisterUnit(input)

	if err != nil {
		response := helper.APIResponse("Create unit failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Unit has been registered", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *unitHandler) GetAllUnit(c *gin.Context) {
	unit, err := h.unitService.FindAllUnit()

	if err != nil {
		response := helper.APIResponse("Failed to get detail of unit", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Unit detail", http.StatusOK, "success", unit)
	c.JSON(http.StatusOK, response)

}
