package handler

import (
	"fmt"
	"iwogo/auth"
	"iwogo/helper"
	storebranch "iwogo/storeBranch"
	"iwogo/user"
	"net/http"
	"os"

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
	errs := c.ShouldBind(&input)

	file, err := c.FormFile("logo")

	if errs != nil {
		response := helper.APIResponse("Upload Data Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusOK, response)
		return
	}

	if err != nil {
		response := helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusOK, response)
		return
	}

	// Set Folder untuk menyimpan filenya
	path := os.Getenv("IMG_STORE") + "" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
		// response := helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
		// c.JSON(http.StatusOK, response)
		// return
	}
	//os.Remove(path)

	data, err := h.storeService.CreateStore(input, c.MustGet("currentUser").(user.User).ID, file.Filename, os.Getenv("IMG_STORE"))
	if err != nil {
		os.Remove(path)
		helper.APIResponse("Created Store Failed", http.StatusBadRequest, "error", err)
		// response := helper.APIResponse("Created Store Failed", http.StatusBadRequest, "error", err)
		// c.JSON(http.StatusOK, response)
		// return
	}

	response := helper.APIResponse("Store Created", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *storeHandler) UpdateStore(c *gin.Context) {

	var input storebranch.UpdateStore
	errs := c.Bind(&input)

	file, err := c.FormFile("logo")
	fmt.Println("lokasi file", os.Getenv("IMG_STORE"))
	if errs != nil {
		helper.APIResponse("Upload Data Failed", http.StatusBadRequest, "error", err)
		// response := helper.APIResponse("Upload Data Failed", http.StatusBadRequest, "error", err)
		// c.JSON(http.StatusBadRequest, response)
		// return
	}

	if err != nil {
		data, err := h.storeService.UpdateStore(input, c.MustGet("currentUser").(user.User).ID, c.PostForm("logoOld"), os.Getenv("IMG_STORE"))
		if err != nil {
			response := helper.APIResponse("Update Store Failed", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Store Update", http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)
	} else {
		pathOld := os.Getenv("IMG_STORE") + "" + c.PostForm("logoOld")
		// os.Remove(pathOld)
		if err := os.Remove(pathOld); err != nil {
			helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
			// response := helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
			// c.JSON(http.StatusBadRequest, response)
			// return
		}

		path := os.Getenv("IMG_STORE") + "" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			// response := helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
			// c.JSON(http.StatusBadRequest, response)
			// return
			helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
		}
		//os.Remove(path)
		fmt.Println("id", c.MustGet("currentUser").(user.User).ID)
		data, err := h.storeService.UpdateStore(input, c.MustGet("currentUser").(user.User).ID, file.Filename, os.Getenv("IMG_STORE"))
		if err != nil {
			response := helper.APIResponse("Update Store Failed", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Store Update", http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)
	}
}

func (h *storeHandler) SeachAll(c *gin.Context) {

	items, err := h.storeService.SearchAllStore(c.MustGet("currentUser").(user.User).ID)

	if err != nil {
		response := helper.APIResponse("Search All Products failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := storebranch.FormatStores(items)
	response := helper.APIResponse("Store detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
