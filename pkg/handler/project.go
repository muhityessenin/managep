package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"managep/pkg/model"
	"net/http"
)

// getProject retrieves all projects
// @Summary Get Projects
// @Description Get all projects from the system
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {array} model.Project
// @Failure 404 {object} Response
// @Router /projects [get]
func (h *Handler) getProject(c *gin.Context) {
	res, err := h.services.ProjectService.GetProject()
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		response := newResponse(http.StatusNotFound, "no projects found", nil)
		c.JSON(response.Status, response)
		return
	}
	response := newResponse(http.StatusOK, "projects successfully found", res)
	c.JSON(http.StatusOK, response)
}

// createProject creates a new project
// @Summary Create Project
// @Description Create a new project in the system
// @Tags projects
// @Accept json
// @Produce json
// @Param project body model.ProjectInputResponse true "Project info"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 400 {object} Response
// @Failure 400 {object} Response
// @Router /projects [post]
func (h *Handler) createProject(c *gin.Context) {
	var input model.Project
	var res Response
	if err := c.Bind(&input); err != nil {
		res = newResponse(http.StatusBadRequest, "invalid request body", "")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if h.validator.ValidateProjectInput(input) == false {
		res = newResponse(http.StatusBadRequest, "invalid request body", "")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	_, err := h.services.ProjectService.CreateProject(&input)
	if err != nil {
		res = newResponse(http.StatusBadRequest, "invalid request body", "")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = newResponse(http.StatusCreated, "project successfully created", "")
	c.JSON(http.StatusCreated, res)
}

// getProjectById retrieves a project by ID
// @Summary Get Project by ID
// @Description Retrieve a project from the system by its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} model.Project
// @Failure 404 {object} Response
// @Router /projects/{id} [get]
func (h *Handler) getProjectById(c *gin.Context) {
	res, err := h.services.ProjectService.GetProjectById(c.Param("id"))
	var response Response
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		response = newResponse(http.StatusNotFound, "project not found", nil)
		c.JSON(response.Status, response)
		return
	}
	response = newResponse(http.StatusOK, "project successfully found", res)
	c.JSON(http.StatusOK, response)
}

// updateProject updates a project by ID
// @Summary Update Project
// @Description Update the details of an existing project identified by its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body model.ProjectInputResponse true "Project info"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /projects/{id} [put]
func (h *Handler) updateProject(c *gin.Context) {
	var input model.Project
	var res Response
	if err := c.Bind(&input); err != nil {
		res = newResponse(http.StatusBadRequest, "invalid request body", "")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if h.validator.ValidateProjectInput(input) == false {
		res = newResponse(http.StatusBadRequest, "invalid request body", "")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	code, err := h.services.ProjectService.UpdateProject(&input, c.Param("id"))
	if code == http.StatusBadRequest {
		res = newResponse(http.StatusBadRequest, "invalid request body", "")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		res = newResponse(http.StatusNotFound, "project not found", "")
		c.JSON(http.StatusNotFound, res)
		return
	}
	res = newResponse(http.StatusOK, "project successfully updated", "")
	c.JSON(http.StatusOK, res)
}

// deleteProject deletes a project by ID
// @Summary Delete Project
// @Description Delete a project from the system by its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /projects/{id} [delete]
func (h *Handler) deleteProject(c *gin.Context) {
	_, err := h.services.ProjectService.DeleteProject(c.Param("id"))
	var res Response
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res = newResponse(http.StatusNotFound, "project not found", nil)
			c.JSON(http.StatusNotFound, res)
			return
		}
		res = newResponse(http.StatusBadRequest, "error deleting project", "")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = newResponse(http.StatusOK, "project successfully deleted", "")
	c.JSON(http.StatusOK, res)
}

// getTasksForProject retrieves tasks associated with a specific project
// @Summary Get Tasks for Project
// @Description Retrieve all tasks associated with a project identified by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {array} model.Task
// @Failure 404 {object} Response
// @Failure 404 {object} Response
// @Router /projects/{id}/tasks [get]
func (h *Handler) getTasksForProject(c *gin.Context) {
	res, err := h.services.ProjectService.GetTasksForProject(c.Param("id"))
	var response Response
	if err != nil {
		if len(res) == 1 {
			response = newResponse(http.StatusNotFound, "project not found", "")
			c.JSON(http.StatusNotFound, response)
			return
		} else if errors.Is(err, sql.ErrNoRows) {
			response = newResponse(http.StatusNotFound, "no tasks found", "")
			c.JSON(http.StatusNotFound, response)
			return
		}
		response = newResponse(http.StatusBadRequest, "bad request", "")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response = newResponse(http.StatusOK, "tasks successfully found", res)
	c.JSON(http.StatusOK, response)
}

// searchProject searches for projects based on query parameters
// @Summary Search Projects
// @Description Search for projects by title or manager. Returns a list of projects that match the search criteria.
// @Tags projects
// @Accept json
// @Produce json
// @Param title query string false "Project Title"
// @Param manager query string false "Project Manager"
// @Success 200 {array} model.Project
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /projects/search [get]
func (h *Handler) searchProject(c *gin.Context) {
	var res []model.Project
	var err error
	var check bool
	var response Response
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
		response = newResponse(http.StatusBadRequest, "bad request", "")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response = newResponse(http.StatusNotFound, "no projects found", "")
			c.JSON(http.StatusNotFound, response)
		} else {
			response = newResponse(http.StatusBadRequest, "error searching project", "")
			c.JSON(http.StatusBadRequest, response)
		}
		return
	}
	response = newResponse(http.StatusOK, "projects successfully found", res)
	c.JSON(http.StatusOK, response)
}
