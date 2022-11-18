package handler

import (
	"iwogo/auth"
	"iwogo/helper"
	"iwogo/item"
	"net/http"

	"github.com/gin-gonic/gin"
)

type itemHandler struct {
	itemService item.Service
	authService auth.Service
}

func NewItemHandler(itemService item.Service, authService auth.Service) *itemHandler {
	return &itemHandler{itemService, authService}
}

func (h *itemHandler) SeachAll(c *gin.Context) {

	var input item.SearchInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Pagination Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	items, total, err := h.itemService.SearchAll(input)

	if err != nil {
		response := helper.APIResponse("Search All Products failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := item.FormatItems(items)
	response := helper.APIPagination("Search All Products", http.StatusOK, "success", total, formatter)
	c.JSON(http.StatusOK, response)
}

// func (h *unitHandler) CreateUnit(c *gin.Context) {

// 	var input unit.RegisterUnitInput

// 	err := c.ShouldBindJSON(&input)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}
// 		response := helper.APIResponse("Register category failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)

// 		return
// 	}

// 	_, err = h.unitService.RegisterUnit(input)

// 	if err != nil {
// 		response := helper.APIResponse("Create unit failed", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Unit has been registered", http.StatusOK, "success", nil)
// 	c.JSON(http.StatusOK, response)
// }

// func (h *unitHandler) GetAllUnit(c *gin.Context) {
// 	unit, err := h.unitService.FindAllUnit()

// 	if err != nil {
// 		response := helper.APIResponse("Failed to get detail of unit", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Unit detail", http.StatusOK, "success", unit)
// 	c.JSON(http.StatusOK, response)

// }

// func (h *unitHandler) UpdateUnit(c *gin.Context) {
// 	var input unit.UpdateUnitInput
// 	err := c.ShouldBindJSON(&input)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}
// 		response := helper.APIResponse("Update unit failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)

// 		return
// 	}
// 	_, err = h.unitService.UpdateUnit(input)
// 	if err != nil {
// 		response := helper.APIResponse("Update unit failed", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	response := helper.APIResponse("Update has been updated", http.StatusOK, "success", nil)
// 	c.JSON(http.StatusOK, response)

// }

// func (h *unitHandler) DeleteUnit(c *gin.Context) {
// 	var input unit.DeleteUnitInput
// 	err := c.ShouldBindJSON(&input)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}
// 		response := helper.APIResponse("Delete unit failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 	}

// 	del, err := h.unitService.DeleteUnit(input)
// 	if err != nil {
// 		errorMessage := gin.H{"error": err.Error()}
// 		response := helper.APIResponse("Please try again", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	if del != true {
// 		errorMessage := gin.H{"error": err.Error()}
// 		response := helper.APIResponse("Please try again", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	response := helper.APIResponse("Delete unit", http.StatusOK, "success", nil)
// 	c.JSON(http.StatusOK, response)
// }
