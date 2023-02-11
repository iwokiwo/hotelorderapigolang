package handler

import (
	"iwogo/auth"
	"iwogo/helper"
	"iwogo/payment"
	"iwogo/user"
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

	data, err := h.paymentService.CreatePaymentType(input, c.MustGet("currentUser").(user.User).ID)

	if err != nil {
		response := helper.APIResponse("Create payment type failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Create has been registered", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *paymentHandler) UpdatePayment(c *gin.Context) {
	var input payment.UpdatePaymentTypeInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Update Update failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.paymentService.UpdatePaymentType(input)
	if err != nil {
		response := helper.APIResponse("Update Payment Type failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Update has been updated", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *paymentHandler) DeletePaymentType(c *gin.Context) {
	var input payment.DeletePaymentTypeInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Delete Payment Type failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	del, err := h.paymentService.DeletePaymentType(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Please try again", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Delete payment type", http.StatusOK, "success", del)
	c.JSON(http.StatusOK, response)
}

func (h *paymentHandler) FindAllPaymentType(c *gin.Context) {

	items, err := h.paymentService.FindAllServicePaymentType(c.MustGet("currentUser").(user.User).ID)

	if err != nil {
		response := helper.APIResponse("Data Payment Type failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := payment.FormatPaymentTypes(items)
	response := helper.APIResponse("Payment type detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
