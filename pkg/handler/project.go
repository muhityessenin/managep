package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) getProject(c *gin.Context)         {}
func (h *Handler) createProject(c *gin.Context)      {}
func (h *Handler) getProjectById(c *gin.Context)     {}
func (h *Handler) updateProject(c *gin.Context)      {}
func (h *Handler) deleteProject(c *gin.Context)      {}
func (h *Handler) getTasksForProject(c *gin.Context) {}
func (h *Handler) searchProject(c *gin.Context)      {}
