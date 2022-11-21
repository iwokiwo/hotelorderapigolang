package handler

import (
	"iwogo/auth"
	"iwogo/helper"
	storebranch "iwogo/storeBranch"
	"iwogo/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type storeHandler struct {
	storeService storebranch.Service
	authService  auth.Service
}

func NewStoreHandler(store storebranch.Service, authService auth.Service) *storeHandler {
	return &storeHandler{store, authService}
}

func (h *storeHandler) CreateStore(c *gin.Context) {
	var input storebranch.CreateStore
	errs := c.Bind(&input)
	file, err := c.FormFile("logo")

	if errs != nil {
		response := helper.APIResponse("Upload Data Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Set Folder untuk menyimpan filenya
	path := "storage/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		response := helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.storeService.CreateStore(input, c.MustGet("currentUser").(user.User).ID, file.Filename)
	if err != nil {
		response := helper.APIResponse("Created Store Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Store Created", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}
