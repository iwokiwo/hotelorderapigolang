package handler

import (
	"fmt"
	coupon "iwogo/Coupon"
	"iwogo/auth"
	"iwogo/helper"
	"iwogo/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type couponHandler struct {
	couponService coupon.Service
	authService   auth.Service
}

func NewCouponHandler(couponService coupon.Service, authService auth.Service) *couponHandler {
	return &couponHandler{couponService, authService}
}

func (h *couponHandler) CreateCoupon(c *gin.Context) {

	var input coupon.CreateCouponInput

	err := c.ShouldBindJSON(&input)
	fmt.Println(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create coupon failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.couponService.CreateService(input, c.MustGet("currentUser").(user.User).ID)

	if err != nil {

		response := helper.APIResponse("Create coupon failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Coupon has been registered", http.StatusOK, "success", input)
	c.JSON(http.StatusOK, response)
}

func (h *couponHandler) UpdateCoupon(c *gin.Context) {

	var input coupon.UpdateCouponInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update coupon failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.couponService.UpdateService(input, c.MustGet("currentUser").(user.User).ID)

	if err != nil {

		response := helper.APIResponse("Update coupon failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Coupon has been registered", http.StatusOK, "success", input)
	c.JSON(http.StatusOK, response)
}

func (h *couponHandler) FindAllCoupon(c *gin.Context) {

	items, err := h.couponService.FindAllService(c.MustGet("currentUser").(user.User).ID)

	if err != nil {
		response := helper.APIResponse("Search All Branch failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := coupon.FormatCoupons(items)
	response := helper.APIResponse("Coupon detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *couponHandler) FindAllCouponByBranch(c *gin.Context) {

	var input coupon.SearchInput
	errs := c.ShouldBindUri(&input)

	if errs != nil {
		errors := helper.FormatValidationError(errs)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("URI input Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	items, err := h.couponService.FindAllByBranchService(input)

	if err != nil {
		response := helper.APIResponse("Search All Coupon failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := coupon.FormatCoupons(items)
	response := helper.APIResponse("Coupon detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *couponHandler) DeleteCoupon(c *gin.Context) {
	var input coupon.DeleteCouponInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Delete coupon failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	responseServer, err := h.couponService.DeleteService(input)

	if err != nil {
		response := helper.APIResponse("Delete coupon failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Coupon has been Delete", http.StatusOK, "success", responseServer)
	c.JSON(http.StatusOK, response)
}
