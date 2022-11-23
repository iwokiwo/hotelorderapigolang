package handler

import (
	"iwogo/auth"
	"iwogo/helper"
	storebranch "iwogo/storeBranch"
	"iwogo/user"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type branchHandler struct {
	branchService storebranch.Service
	authService   auth.Service
}

func NewBranchHandler(store storebranch.Service, authService auth.Service) *branchHandler {
	return &branchHandler{store, authService}
}

func (h *branchHandler) CreateBranch(c *gin.Context) {
	//fmt.Println(c.PostForm("name"))
	var input storebranch.CreateBranch
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
	path := os.Getenv("IMG_BRANCHES") + "" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		response := helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//fmt.Println(path)
	//os.Remove(path)

	data, err := h.branchService.CreateBranch(input, c.MustGet("currentUser").(user.User).ID, file.Filename, os.Getenv("IMG_BRANCHES"))
	if err != nil {
		os.Remove(path)
		response := helper.APIResponse("Created Branch Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Branch Created", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *branchHandler) UpdateBranch(c *gin.Context) {

	var input storebranch.UpdateBranch
	errs := c.Bind(&input)

	file, err := c.FormFile("logo")

	if errs != nil {
		response := helper.APIResponse("Upload Data Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//os.Remove(path)
	if err != nil {
		data, err := h.branchService.UpdateBranch(input, c.MustGet("currentUser").(user.User).ID, c.PostForm("logoOld"), os.Getenv("IMG_BRANCHES"))
		if err != nil {
			response := helper.APIResponse("Update Branch Failed", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Branch Update", http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)

	} else {
		pathOld := os.Getenv("IMG_BRANCHES") + "" + c.PostForm("logoOld")
		// os.Remove(pathOld)
		if err := os.Remove(pathOld); err != nil {
			response := helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		path := os.Getenv("IMG_BRANCHES") + "" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			response := helper.APIResponse("Upload Logo Failed", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		data, err := h.branchService.UpdateBranch(input, c.MustGet("currentUser").(user.User).ID, file.Filename, os.Getenv("IMG_BRANCHES"))
		if err != nil {
			response := helper.APIResponse("Update Branch Failed", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Branch Update", http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)
	}

}

func (h *branchHandler) SeachAll(c *gin.Context) {

	items, err := h.branchService.SearchAllBranch(c.MustGet("currentUser").(user.User).ID)

	if err != nil {
		response := helper.APIResponse("Search All Branch failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Store detail", http.StatusOK, "success", items)
	c.JSON(http.StatusOK, response)
}
