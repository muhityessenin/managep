package handler

import (
	"github.com/gin-gonic/gin"
	"managep/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	user := router.Group("/users")
	{
		user.GET("", h.getUser)
		user.POST("", h.createUser)
		user.DELETE("/:id", h.deleteUser)
		user.GET("/:id", h.getUserById)
		user.PUT("/:id", h.updateUser)
		user.GET("/:id/tasks", h.getTasksForUser)
		user.GET("/search", h.searchUser)
	}
	task := router.Group("/tasks")
	{
		task.GET("", h.getTask)
		task.POST("", h.createTask)
		task.DELETE("/:id", h.deleteTask)
		task.GET("/:id", h.getTaskById)
		task.PUT("/:id", h.updateTask)
		task.GET("/search", h.searchTask)
	}
	project := router.Group("/projects")
	{
		project.GET("", h.getProject)
		project.POST("", h.createProject)
		project.DELETE("/:id", h.deleteProject)
		project.GET("/:id", h.getProjectById)
		project.PUT("/:id", h.updateProject)
		project.GET("/search", h.searchProject)
		project.GET("/:id/tasks", h.getTasksForProject)
	}
	return router
}
