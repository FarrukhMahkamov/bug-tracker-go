package handler

import (
	"github.com/FarrukhMahkamov/bugtracker/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitiRoutes() *gin.Engine {
	router := gin.New()

	//Route group for auth
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	//Api routes
	api := router.Group("/api", h.Auth)
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
			teams.POST("/:id/add-users", h.AddUsersToTeam)
			teams.GET("/:id/team-users", h.GetTeamUsers)
			teams.DELETE("/:id/remove-users", h.RemoveUsersFromTeam)
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
			statuses.GET("/:id/bugs", h.ShowBugsByStatusId)
			statuses.POST("/", h.StoreStatus)
			statuses.PUT("/:id", h.UpdateStatus)
			statuses.DELETE("/:id", h.DeleteStatus)
		}

		bugs := api.Group("/bugs")
		{
			bugs.GET("/", h.GetAllBug)
			bugs.POST("/", h.StoreBug)
			bugs.POST("/:id/tags", h.AddTag)
			bugs.GET("/:id/tags", h.GetTagsByBugId)
			bugs.PUT("/:id/close", h.CloseIssue)
			bugs.PUT("/:id", h.UpdateBug)
			bugs.POST("/:id/attach-users", h.AttachUserToBug)
			bugs.DELETE("/:id/detach-users", h.DeattachUserFromBug)
			bugs.POST("/:id/attach-team/:team_id", h.AttachTeamToBug)
			bugs.DELETE("/:id/detach-team/:team_id", h.DetachTeamFromBug)
		}

		projects := api.Group("/projects")
		{
			projects.GET("/", h.GetAllProject)
			projects.POST("/", h.StoreProject)
			projects.GET("/:id", h.ShowProject)
			projects.POST("/:id/update", h.UpdateProject)
			projects.DELETE("/:id", h.DeleteProject)
			projects.POST("/:id/add-users", h.AddUsersToProject)
			projects.POST("/:id/remove-users", h.RemoveUsersFromProject)
			projects.GET("/:id/bugs/:user_id", h.GetAttachedBugs)
			projects.POST("/:id/add-team/:team_id", h.AddTeamToProject)
			projects.GET("/attached-project/:user_id", h.GetAttachedProjects)
		}

	}

	return router
}
