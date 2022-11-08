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
