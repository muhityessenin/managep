package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"managep/pkg/model"
	"net/http"
)

func (h *Handler) getProject(c *gin.Context) {
	res, err := h.services.ProjectService.GetProject()
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no projects found",
			})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) createProject(c *gin.Context) {
	var input model.Project
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if h.validator.ValidateProjectInput(input) == false {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}
	_, err := h.services.ProjectService.CreateProject(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, nil)
}
func (h *Handler) getProjectById(c *gin.Context) {
	res, err := h.services.ProjectService.GetProjectById(c.Param("id"))
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{"message": "project not found"})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) updateProject(c *gin.Context) {
	var input model.Project
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if h.validator.ValidateProjectInput(input) == false {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}
	res, err := h.services.ProjectService.UpdateProject(&input, c.Param("id"))
	if res == http.StatusBadRequest {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request or invalid input"})
		return
	}
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{"message": "project not found"})
		}
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) deleteProject(c *gin.Context) {
	_, err := h.services.ProjectService.DeleteProject(c.Param("id"))
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{"message": "project not found"})
		}
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) getTasksForProject(c *gin.Context) {
	res, err := h.services.ProjectService.GetTasksForProject(c.Param("id"))
	if err != nil {
		if len(res) == 1 {
			c.JSON(http.StatusNotFound, gin.H{"message": "project not found"})
		} else if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"message": "no tasks found"})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) searchProject(c *gin.Context) {
	var res []model.Project
	var err error
	var check bool
	title := c.Query("title")
	if title != "" {
		res, err = h.services.ProjectService.SearchProject(title, "title")
		check = true
	}
	manager := c.Query("manager")
	if manager != "" {
		res, err = h.services.ProjectService.SearchProject(manager, "manager")
		check = true
	}
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no projects found",
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
