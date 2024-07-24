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
	user := router.Group("/user")
	{
		user.GET("/users", h.getUser)
		user.POST("/users", h.createUser)
		user.DELETE("/users/:id", h.deleteUser)
		user.GET("/users/:id", h.getUserById)
		user.PUT("/users/:id", h.updateUser)
		user.GET("/users/:id/tasks", h.getTasksForUser)
		user.GET("/users/search", h.searchUser)
	}
	task := router.Group("/task")
	{
		task.GET("/tasks", h.getTask)
		task.POST("/tasks", h.createTask)
		task.DELETE("/tasks/:id", h.deleteTask)
		task.GET("/tasks/:id", h.getTaskById)
		task.PUT("/tasks/:id", h.updateTask)
		task.GET("/tasks/search", h.searchTask)
	}
	project := router.Group("/project")
	{
		project.GET("/projects", h.getProject)
		project.POST("/projects", h.createProject)
		project.DELETE("/projects/:id", h.deleteProject)
		project.GET("/projects/:id", h.getProjectById)
		project.PUT("/projects/:id", h.updateProject)
		project.GET("/projects/search", h.searchProject)
		project.GET("/projects/:id/tasks", h.getTasksForProject)
	}
	return router
}
