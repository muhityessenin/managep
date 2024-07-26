package handler

import (
	"github.com/gin-gonic/gin"
	"managep/pkg/model"
	"net/http"
)

func (h *Handler) getProject(c *gin.Context) {}
func (h *Handler) createProject(c *gin.Context) {
	var input model.Project
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.services.ProjectService.CreateProject(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) getProjectById(c *gin.Context)     {}
func (h *Handler) updateProject(c *gin.Context)      {}
func (h *Handler) deleteProject(c *gin.Context)      {}
func (h *Handler) getTasksForProject(c *gin.Context) {}
func (h *Handler) searchProject(c *gin.Context)      {}
