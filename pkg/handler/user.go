package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUser(c *gin.Context)         {}
func (h *Handler) createUser(c *gin.Context)      {}
func (h *Handler) getUserById(c *gin.Context)     {}
func (h *Handler) updateUser(c *gin.Context)      {}
func (h *Handler) deleteUser(c *gin.Context)      {}
func (h *Handler) getTasksForUser(c *gin.Context) {}
func (h *Handler) searchUser(c *gin.Context)      {}
