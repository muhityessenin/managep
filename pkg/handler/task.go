package handler

import (
	"github.com/gin-gonic/gin"
	"managep/pkg/model"
	"net/http"
)

func (h *Handler) getTask(c *gin.Context) {}
func (h *Handler) createTask(c *gin.Context) {
	var input model.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task input")
	}
	_, err := h.services.TaskService.CreateTask(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) getTaskById(c *gin.Context) {}
func (h *Handler) updateTask(c *gin.Context)  {}
func (h *Handler) deleteTask(c *gin.Context)  {}
func (h *Handler) searchTask(c *gin.Context)  {}
