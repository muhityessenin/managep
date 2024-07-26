package handler

import (
	"github.com/gin-gonic/gin"
	"managep/pkg/model"
	"net/http"
)

func (h *Handler) getTask(c *gin.Context) {
	tasks, err := h.services.TaskService.GetTask()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, tasks)
}
func (h *Handler) createTask(c *gin.Context) {
	var input model.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task input")
		return
	}
	res, err := h.services.TaskService.CreateTask(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(res, nil)
}
func (h *Handler) getTaskById(c *gin.Context) {
	res, err := h.services.TaskService.GetTaskById(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) updateTask(c *gin.Context) {
	var input model.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task input")
		return
	}
	_, err := h.services.TaskService.UpdateTask(&input, c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) deleteTask(c *gin.Context) {
	res, err := h.services.TaskService.DeleteTask(c.Param("id"))
	if err != nil {
		newErrorResponse(c, 404, err.Error())
		return
	}
	c.JSON(res, nil)
}
func (h *Handler) searchTask(c *gin.Context) {}
