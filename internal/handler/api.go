package handler

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	seriveces *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{seriveces: services}
}

func (h *Handler) InitiRoutes() *gin.Engine {
	router := gin.New()

	//Route group for auth
	auth := router.Group("/auth")
	{
		auth.POST("/sing-in", h.SignIn)
		auth.POST("/sing-up", h.SignUp)
	}

	//Api routes
	api := router.Group("/api")
	{
		//Routes for job types
		jobType := api.Group("/job-types")
		{
			jobType.GET("/", h.GetAllJobType)
			jobType.GET("/:id", h.ShowJobType)
			jobType.POST("/", h.StoreJobType)
			jobType.PUT("/:id", h.UpdateJobType)
			jobType.DELETE("/:id", h.DeleteJobType)
		}
	}

	return router
}
