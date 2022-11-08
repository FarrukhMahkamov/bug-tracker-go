package handler

import (
	"net/http"
	"strconv"

	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllTeam(c *gin.Context) {
	Teams, err := h.services.Team.GetAllTeam()
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Teams)
}

func (h *Handler) ShowTeam(c *gin.Context) {
	TeamId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	Team, err := h.services.Team.ShowTeam(TeamId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Team)
}

func (h *Handler) StoreTeam(c *gin.Context) {
	var input dto.Team

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	StoredTeam, err := h.services.Team.StoreTeam(input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StoredTeam)
}

func (h *Handler) UpdateTeam(c *gin.Context) {
	TeamId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input dto.TeamUpdate

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Team.UpdatedTeam(input, TeamId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Updated SuccessFully",
		Status:  http.StatusOK,
	})

}
func (h *Handler) DeleteTeam(c *gin.Context) {
	TeamId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Team.DeleteTeam(TeamId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Deleted Successfully",
		Status:  http.StatusNoContent,
	})
}

func (h *Handler) AddUsersToTeam(c *gin.Context) {
	var Users dto.TeamUsers
	TeamId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&Users); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Team.AddUsersToTeam(TeamId, Users)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Users added successfully",
		Status:  http.StatusNoContent,
	})

}

func (h *Handler) GetTeamUsers(c *gin.Context) {
	TeamId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	Users, err := h.services.Team.GetTeamUsers(TeamId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Users)
}

func (h *Handler) RemoveUsersFromTeam(c *gin.Context) {
	var Users dto.TeamUsers
	TeamId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&Users); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Team.RemoveUsersFromTeam(TeamId, Users)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse{
		Message: "Users removed successfully",
		Status:  http.StatusNoContent,
	})
}
