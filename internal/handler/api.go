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

		teams := api.Group("/teams")
		{
			teams.GET("/", h.GetAllTeam)
			teams.GET("/:id", h.ShowTeam)
			teams.POST("/", h.StoreTeam)
			teams.PUT("/:id", h.UpdateTeam)
			teams.DELETE("/:id", h.DeleteTeam)
		}

		tags := api.Group("/tags")
		{
			tags.GET("/", h.GetAllTag)
			tags.GET("/:id", h.ShowTag)
			tags.POST("/", h.StoreTag)
			tags.PUT("/:id", h.UpdateTag)
			tags.DELETE("/:id", h.DeleteTag)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", h.GetAllCategory)
			categories.GET("/:id", h.ShowCategory)
			categories.POST("/", h.StoreCategory)
			categories.PUT("/:id", h.UpdateCategory)
			categories.DELETE("/:id", h.DeleteCategory)
		}

		statuses := api.Group("/statuses")
		{
			statuses.GET("/", h.GetAllStatus)
			statuses.GET("/:id", h.ShowStatus)
			statuses.POST("/", h.StoreStatus)
			statuses.PUT("/:id", h.UpdateStatus)
			statuses.DELETE("/:id", h.DeleteStatus)
		}

		bugs := api.Group("/bugs")
		{
			bugs.GET("/", h.GetAllBug)
			bugs.POST("/", h.StoreBug)
			bugs.POST("/:id/tags", h.AddTag)
			bugs.PUT("/:id", h.CloseIssue)
		}

	}

	return router
}
