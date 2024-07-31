package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"managep/pkg/model"
	"managep/pkg/service"
	"managep/pkg/validator"
)

type Handler struct {
	services  *service.Service
	validator *validator.Validator
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

func (h *Handler) ParseUserResponse(user *model.User) model.UserInputResponse {
	return model.UserInputResponse{
		FullName:         user.FullName,
		Email:            user.Email,
		RegistrationDate: user.RegistrationDate[:10],
		Role:             user.Role,
	}
}

func (h *Handler) ParseTaskInputResponse(task *model.Task) model.TaskInputResponse {
	return model.TaskInputResponse{
		Name:              task.Name,
		Description:       task.Description,
		Priority:          task.Priority,
		State:             task.State,
		ResponsiblePerson: task.ResponsiblePerson,
		Project:           task.Project,
		CreatedAt:         task.CreatedAt[:10],
		FinishedAt:        task.FinishedAt[:10],
	}
}

func (h *Handler) ParseProjectInputResponse(project *model.Project) model.ProjectInputResponse {
	return model.ProjectInputResponse{
		Name:        project.Name,
		Description: project.Description,
		StartDate:   project.StartDate[:10],
		FinishDate:  project.FinishDate[:10],
		Manager:     project.Manager,
	}
}
