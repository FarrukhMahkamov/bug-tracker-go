package handler

import (
	"net/http"
	"strconv"

	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllBug(c *gin.Context) {
	Bugs, err := h.services.Bug.GetAllBug()
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Bugs)
}

func (h *Handler) StoreBug(c *gin.Context) {
	var input dto.Bug

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	StoredBug, err := h.services.Bug.StoreBug(input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StoredBug)
}

func (h *Handler) CloseIssue(c *gin.Context) {
	BugId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Bug.CloseIssue(BugId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Issu Closed Successfully",
		Status:  http.StatusNoContent,
	})
}

func (h *Handler) UpdateBug(c *gin.Context) {
	BugId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var UpdateBug dto.BugUpdate
	if err := c.BindJSON(&UpdateBug); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Bug.UpdateBug(BugId, UpdateBug)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Issu Closed Successfully",
		Status:  http.StatusNoContent,
	})
}

func (h *Handler) AddTag(c *gin.Context) {
	var Tags dto.BugTag

	BugId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&Tags); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Bug.AddTag(Tags, BugId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Tags Added Successfully",
		Status:  http.StatusOK,
	})
}

func (h *Handler) GetTagsByBugId(c *gin.Context) {
	BugId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	Tags, err := h.services.Bug.GetTagsByBugId(BugId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Tags) //
}

func (h *Handler) AttachUserToBug(c *gin.Context) {
	BugId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var Users dto.AttachUsers
	if err := c.BindJSON(&Users); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Bug.AttachUserToBug(BugId, Users)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "User(s) attached successfully",
		Status:  http.StatusOK,
	})
}

func (h *Handler) DeattachUserFromBug(c *gin.Context) {
	BugId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var Users dto.AttachUsers
	if err := c.BindJSON(&Users); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Bug.DeattachUserFromBug(BugId, Users)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "User(s) attached successfully",
		Status:  http.StatusOK,
	})
}

func (h *Handler) AttachTeamToBug(c *gin.Context) {

	BugId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	TeamId, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Bug.AttachTeamToBug(BugId, TeamId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Team attached successfully",
		Status:  http.StatusOK,
	})
}

func (h *Handler) DetachTeamFromBug(c *gin.Context) {

	BugId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	TeamId, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Bug.DetachTeamFromBug(BugId, TeamId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Team detached successfully",
		Status:  http.StatusOK,
	})
}
