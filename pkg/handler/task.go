package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"managep/pkg/model"
	"net/http"
)

func (h *Handler) getTask(c *gin.Context) {
	tasks, err := h.services.TaskService.GetTask()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no tasks found",
		})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
func (h *Handler) createTask(c *gin.Context) {
	var input model.Task
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if h.validator.ValidateTaskInput(input) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}
	_, err := h.services.TaskService.CreateTask(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) getTaskById(c *gin.Context) {
	res, err := h.services.TaskService.GetTaskById(c.Param("id"))
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "task not found",
			})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) updateTask(c *gin.Context) {
	var input model.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}
	if h.validator.ValidateTaskInput(input) == false {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}
	res, err := h.services.TaskService.UpdateTask(&input, c.Param("id"))
	if res == http.StatusBadRequest {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input or bad request",
		})
		return
	}
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "task not found",
			})
		}
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) deleteTask(c *gin.Context) {
	res, err := h.services.TaskService.DeleteTask(c.Param("id"))
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "task not found",
			})
		}
		return
	}
	c.JSON(res, nil)
}
func (h *Handler) searchTask(c *gin.Context) {
	title := c.Query("title")
	var res []model.Task
	var err error
	var check bool
	if title != "" {
		res, err = h.services.TaskService.SearchTask(title, "title")
		check = true
	}
	status := c.Query("status")
	if status != "" {
		res, err = h.services.TaskService.SearchTask(status, "status")
		check = true
	}
	priority := c.Query("priority")
	if priority != "" {
		res, err = h.services.TaskService.SearchTask(priority, "priority")
		check = true
	}
	assignee := c.Query("assignee")
	if assignee != "" {
		res, err = h.services.TaskService.SearchTask(assignee, "assignee")
		check = true
	}
	project := c.Query("project")
	if project != "" {
		res, err = h.services.TaskService.SearchTask(project, "project")
		check = true
	}
	if !check {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "bad request",
		})
		return
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no task found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}
