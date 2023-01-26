package handler

import (
	"iwogo/auth"
	"iwogo/helper"
	"iwogo/payment"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentService payment.Service
	authService    auth.Service
}

func NewPaymentHandler(paymentService payment.Service, authService auth.Service) *paymentHandler {
	return &paymentHandler{paymentService, authService}
}

func (h *paymentHandler) CreatePaymentType(c *gin.Context) {

	var input payment.CreatePaymentTypeInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create payment type failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	data, err := h.paymentService.CreatePaymentType(input)

	if err != nil {
		response := helper.APIResponse("Create payment type failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Create has been registered", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}