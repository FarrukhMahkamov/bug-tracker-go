package handler

import (
	"net/http"
	"strings"

	"github.com/FarrukhMahkamov/bugtracker/internal/response"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	UserCtx             = "userId"
)

func (h *Handler) Auth(c *gin.Context) {
	// GET HEADER
	header := c.GetHeader(AuthorizationHeader)
	if header == "" {
		response.NewErrorResponse(c, http.StatusUnauthorized, "Empty auth header!")
		return
	}

	// SPLIT HEADER
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		response.NewErrorResponse(c, http.StatusUnauthorized, "Invalid auth header!")
		return
	}

	// PARSE TOKEN

	UserId, err := h.seriveces.Authortization.ParseToken(headerParts[1])
	if err != nil {
		response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(UserCtx, UserId)
}
