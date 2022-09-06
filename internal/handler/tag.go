package handler

import (
	"net/http"
	"strconv"

	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllTag(c *gin.Context) {
	Tags, err := h.seriveces.Tag.GetAllTag()
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Tags)
}

func (h *Handler) ShowTag(c *gin.Context) {
	TagId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	Tag, err := h.seriveces.Tag.ShowTag(TagId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Tag)
}

func (h *Handler) StoreTag(c *gin.Context) {
	var input dto.Tag

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	StoredTag, err := h.seriveces.Tag.StoreTag(input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StoredTag)
}

func (h *Handler) UpdateTag(c *gin.Context) {
	TagId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input dto.TagUpdate

	err = h.seriveces.Tag.UpdatedTag(input, TagId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Updated SuccessFully",
		Status:  http.StatusOK,
	})

}
func (h *Handler) DeleteTag(c *gin.Context) {
	TagId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.seriveces.Tag.DeleteTag(TagId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Deleted Successfully",
		Status:  http.StatusNoContent,
	})
}
