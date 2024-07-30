package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"managep/pkg/model"
	"net/http"
)

// getTask retrieves all tasks
// @Summary Get Tasks
// @Description Get all tasks from the system
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} model.Task
// @Failure 404 {object} Response
// @Router /tasks [get]
func (h *Handler) getTask(c *gin.Context) {
	tasks, err := h.services.TaskService.GetTask()
	var res Response
	if err != nil {
		res = newResponse(http.StatusNotFound, "no tasks found", nil)
		c.JSON(http.StatusNotFound, res)
		return
	}
	res = newResponse(http.StatusOK, "tasks successfully found", tasks)
	c.JSON(http.StatusOK, res)
}

// createTask creates a new task
// @Summary Create Task
// @Description Create a new task in the system
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Task info"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 400 {object} Response
// @Failure 400 {object} Response
// @Router /tasks [post]
func (h *Handler) createTask(c *gin.Context) {
	var input model.Task
	var res Response
	if err := c.BindJSON(&input); err != nil {
		res = newResponse(http.StatusBadRequest, "invalid request body", err)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if h.validator.ValidateTaskInput(input) == false {
		res = newResponse(http.StatusBadRequest, "invalid request body", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	_, err := h.services.TaskService.CreateTask(&input)
	if err != nil {
		res = newResponse(http.StatusBadRequest, "task already exists", err)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = newResponse(http.StatusOK, "task successfully created", nil)
	c.JSON(http.StatusOK, res)
}

// getTaskById retrieves a task by ID
// @Summary Get Task by ID
// @Description Get a task from the system by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} model.Task
// @Failure 404 {object} Response
// @Router /tasks/{id} [get]
func (h *Handler) getTaskById(c *gin.Context) {
	res, err := h.services.TaskService.GetTaskById(c.Param("id"))
	var response Response
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		response = newResponse(http.StatusNotFound, "task not found", "")
		c.JSON(http.StatusNotFound, response)
		return
	}
	response = newResponse(http.StatusOK, "task successfully found", res)
	c.JSON(http.StatusOK, response)
}

// updateTask updates a task by ID
// @Summary Update Task
// @Description Update a task's details by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body model.Task true "Task info"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /tasks/{id} [put]
func (h *Handler) updateTask(c *gin.Context) {
	var input model.Task
	var response Response
	if err := c.BindJSON(&input); err != nil {
		response = newResponse(http.StatusBadRequest, "invalid request body", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if h.validator.ValidateTaskInput(input) == false {
		response = newResponse(http.StatusBadRequest, "invalid request body", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	res, err := h.services.TaskService.UpdateTask(&input, c.Param("id"))
	if res == http.StatusBadRequest {
		response = newResponse(http.StatusBadRequest, "update task failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response = newResponse(http.StatusBadRequest, "task not found", "")
		c.JSON(http.StatusNotFound, response)
		return
	}
	response = newResponse(http.StatusOK, "task successfully updated", "")
	c.JSON(http.StatusOK, response)
}

// deleteTask deletes a task by ID
// @Summary Delete Task
// @Description Delete a task from the system by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /tasks/{id} [delete]
func (h *Handler) deleteTask(c *gin.Context) {
	_, err := h.services.TaskService.DeleteTask(c.Param("id"))
	var response Response
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response = newResponse(http.StatusNotFound, "task not found", "")
			c.JSON(http.StatusNotFound, response)
			return
		}
		response = newResponse(http.StatusBadRequest, "delete task failed", "")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response = newResponse(http.StatusOK, "task successfully deleted", "")
	c.JSON(http.StatusOK, response)
}

// searchTask searches for tasks based on various query parameters
// @Summary Search Tasks
// @Description Search for tasks in the system by title, status, priority, assignee, or project
// @Tags tasks
// @Accept json
// @Produce json
// @Param title query string false "Task Title"
// @Param status query string false "Task Status"
// @Param priority query string false "Task Priority"
// @Param assignee query string false "Task Assignee"
// @Param project query string false "Task Project"
// @Success 200 {array} model.Task
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /tasks/search [get]
func (h *Handler) searchTask(c *gin.Context) {
	title := c.Query("title")
	var res []model.Task
	var err error
	var check bool
	var response Response
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
		response = newResponse(http.StatusNotFound, "bad request", "")
		c.JSON(http.StatusNotFound, response)
		return
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response = newResponse(http.StatusNotFound, "task not found", "")
			c.JSON(http.StatusNotFound, response)
		} else {
			response = newResponse(http.StatusBadRequest, "get task failed", "")
			c.JSON(http.StatusBadRequest, response)
		}
		return
	}
	response = newResponse(http.StatusOK, "task successfully found", res)
	c.JSON(http.StatusOK, response)
}
