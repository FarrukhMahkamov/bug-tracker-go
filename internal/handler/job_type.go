package handler

import (
	"net/http"

	"github.com/FarrukhMahkamov/bugtracker/internal/datastruct"
	"github.com/FarrukhMahkamov/bugtracker/internal/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllJobType(c *gin.Context) {

}

func (h *Handler) ShowJobType(c *gin.Context) {

}

func (h *Handler) StoreJobType(c *gin.Context) {
	var input datastruct.JobType

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	StoredJobType, err := h.seriveces.JobType.StoreJobType(input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StoredJobType)
}

func (h *Handler) UpdateJobType(c *gin.Context) {

}
func (h *Handler) DeleteJobType(c *gin.Context) {

}
