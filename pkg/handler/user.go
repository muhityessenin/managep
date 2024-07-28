package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"managep/pkg/model"
	"net/http"
)

func (h *Handler) getUser(c *gin.Context) {
	users, err := h.services.UserService.GetUser()
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no users found",
			})
			return
		}
	}
	c.JSON(http.StatusOK, users)
}
func (h *Handler) createUser(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}
	if h.validator.ValidateUserInput(input) == false {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}
	_, err := h.services.UserService.CreateUser(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, nil)
}
func (h *Handler) getUserById(c *gin.Context) {
	res, err := h.services.UserService.GetUserById(c.Param("id"))
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) updateUser(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if h.validator.ValidateUserInput(input) == false {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}
	res, err := h.services.UserService.UpdateUser(&input, c.Param("id"))
	if res == http.StatusBadRequest {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
		}
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) deleteUser(c *gin.Context) {
	_, err := h.services.UserService.DeleteUser(c.Param("id"))
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) getTasksForUser(c *gin.Context) {
	res, err := h.services.UserService.GetTasksForUser(c.Param("id"))
	if err != nil {
		if len(res) == 1 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
		} else if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no tasks found",
			})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) searchUser(c *gin.Context) {
	var user model.User
	var err error
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
		c.JSON(http.StatusNotFound, gin.H{
			"message": "bad request",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
