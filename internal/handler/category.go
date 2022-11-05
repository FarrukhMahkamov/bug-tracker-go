package handler

import (
	"net/http"
	"strconv"

	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllCategory(c *gin.Context) {
	Categorys, err := h.services.Category.GetAllCategory()
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Categorys)
}

func (h *Handler) ShowCategory(c *gin.Context) {
	CategoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	Category, err := h.services.Category.ShowCategory(CategoryId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Category)
}

func (h *Handler) StoreCategory(c *gin.Context) {
	var input dto.Category

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	StoredCategory, err := h.services.Category.StoreCategory(input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StoredCategory)
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	CategoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input dto.CategoryUpdate

	err = h.services.Category.UpdatedCategory(input, CategoryId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Updated SuccessFully",
		Status:  http.StatusOK,
	})

}
func (h *Handler) DeleteCategory(c *gin.Context) {
	CategoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Category.DeleteCategory(CategoryId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Deleted Successfully",
		Status:  http.StatusNoContent,
	})
}
