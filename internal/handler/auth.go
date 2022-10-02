package handler

import (
	"net/http"

	"github.com/FarrukhMahkamov/bugtracker/internal/dto"
	"github.com/FarrukhMahkamov/bugtracker/internal/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var InputFields dto.User

	if err := c.BindJSON(&InputFields); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	RegisteredUser, err := h.seriveces.Authortization.RegisterUser(InputFields)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, RegisteredUser)
}

func (h *Handler) SignIn(c *gin.Context) {
	var InputFields dto.UserSignInFields

	if err := c.BindJSON(&InputFields); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	RegisteredUser, err := h.seriveces.Authortization.SignInUser(InputFields.Email, InputFields.Password)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	GetLoginedUser, err := h.seriveces.Authortization.GetUser(InputFields.Email, InputFields.Password)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	GetLoginedUser.Token = &RegisteredUser

	c.JSON(http.StatusOK, GetLoginedUser)
}
