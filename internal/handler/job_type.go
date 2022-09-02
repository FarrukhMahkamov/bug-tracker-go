package handler

import (
	"net/http"
	"strconv"

	"github.com/FarrukhMahkamov/bugtracker/internal/datastruct"
	"github.com/FarrukhMahkamov/bugtracker/internal/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllJobType(c *gin.Context) {
	JobTypes, err := h.seriveces.JobType.GetAllJobType()
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, JobTypes)
}

func (h *Handler) ShowJobType(c *gin.Context) {
	JobTypeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	JobType, err := h.seriveces.JobType.ShowJobType(JobTypeId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, JobType)
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
