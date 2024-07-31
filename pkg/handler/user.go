package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"managep/pkg/model"
	"net/http"
)

// @Summary Get User
// @Description Get all users from the system
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Failure 404 {object} Response
// @Router /users [get]
func (h *Handler) getUser(c *gin.Context) {
	users, err := h.services.UserService.GetUser()
	var res Response
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		res = Response{
			Status:  http.StatusNotFound,
			Message: "no users found",
			Data:    nil,
		}
		c.JSON(http.StatusNotFound, res)
		return
	}
	res = Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    users,
	}
	c.JSON(http.StatusOK, res)
}

// createUser creates a new user
// @Summary Create User
// @Description Create a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserInputResponse true "User info"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 400 {object} Response
// @Failure 400 {object} Response
// @Router /users [post]
func (h *Handler) createUser(c *gin.Context) {
	var input model.User
	var res Response
	if err := c.BindJSON(&input); err != nil {
		res = newResponse(http.StatusBadRequest, "invalid input", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if h.validator.ValidateUserInput(input) == false {
		res = newResponse(http.StatusBadRequest, "invalid input", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	_, err := h.services.UserService.CreateUser(&input)
	if err != nil {
		res = newResponse(http.StatusBadRequest, "invalid input", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = newResponse(http.StatusCreated, "user created", nil)
	c.JSON(http.StatusCreated, res)
}

// getUserById retrieves a user by ID
// @Summary Get User by ID
// @Description Get a user from the system by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 404 {object} Response
// @Router /users/{id} [get]
func (h *Handler) getUserById(c *gin.Context) {
	res, err := h.services.UserService.GetUserById(c.Param("id"))
	var result Response
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		result = newResponse(http.StatusNotFound, "user not found", nil)
		c.JSON(http.StatusNotFound, result)
		return
	}
	result = newResponse(http.StatusOK, "user successfully found", res)
	c.JSON(http.StatusOK, result)
}

// updateUser updates a user by ID
// @Summary Update User
// @Description Update a user's details by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body model.UserInputResponse true "User info"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /users/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	var input model.User
	var response Response
	if err := c.BindJSON(&input); err != nil {
		response = newResponse(http.StatusBadRequest, "invalid input", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if h.validator.ValidateUserInput(input) == false {
		response = newResponse(http.StatusBadRequest, "invalid input", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	res, err := h.services.UserService.UpdateUser(&input, c.Param("id"))
	if res == http.StatusBadRequest {
		response = newResponse(http.StatusBadRequest, "invalid input", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		response = newResponse(http.StatusBadRequest, "user not found", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response = newResponse(http.StatusOK, "user successfully updated", nil)
	c.JSON(http.StatusOK, response)
}

// deleteUser deletes a user by ID
// @Summary Delete User
// @Description Delete a user from the system by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /users/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	_, err := h.services.UserService.DeleteUser(c.Param("id"))
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		res := newResponse(http.StatusNotFound, "user not found", "")
		c.JSON(http.StatusOK, res)
		return
	}
	res := newResponse(http.StatusOK, "user deleted successfully", "")
	c.JSON(http.StatusOK, res)
}

// getTasksForUser retrieves tasks for a specific user by their ID
// @Summary Get Tasks for User
// @Description Get all tasks assigned to a specific user by their ID
// @Tags users, tasks
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} model.Task
// @Failure 404 {object} Response
// @Failure 404 {object} Response
// @Router /users/{id}/tasks [get]
func (h *Handler) getTasksForUser(c *gin.Context) {
	res, err := h.services.UserService.GetTasksForUser(c.Param("id"))
	var result Response
	if err != nil {
		if len(res) == 1 {
			result = newResponse(http.StatusNotFound, "user not found", "")
		} else if errors.Is(err, sql.ErrNoRows) {
			result = newResponse(http.StatusNotFound, "no tasks found", "")
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	result = newResponse(http.StatusOK, "tasks for user successfully found", res)
	c.JSON(http.StatusOK, result)
}

// searchUser searches for a user by name or email
// @Summary Search User
// @Description Search for a user in the system by their name or email
// @Tags users
// @Accept json
// @Produce json
// @Param name query string false "Username"
// @Param email query string false "User Email"
// @Success 200 {object} model.User
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /users/search [get]
func (h *Handler) searchUser(c *gin.Context) {
	var user model.User
	var err error
	var res Response
	check := false
	name := c.Query("name")
	email := c.Query("email")
	if name != "" {
		user, err = h.services.UserService.SearchUser(name, "name")
		check = true
	} else if email != "" {
		user, err = h.services.UserService.SearchUser(email, "email")
		check = true
	}
	if !check {
		res = newResponse(http.StatusBadRequest, "invalid input", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		res = newResponse(http.StatusBadRequest, "user not found", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = newResponse(http.StatusOK, "user successfully found", user)
	c.JSON(http.StatusOK, res)
}
