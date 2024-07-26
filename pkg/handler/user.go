package handler

import (
	"github.com/gin-gonic/gin"
	"managep/pkg/model"
	"net/http"
)

func (h *Handler) getUser(c *gin.Context) {
	users, err := h.services.UserService.GetUser()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}
func (h *Handler) createUser(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.services.UserService.CreateUser(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) getUserById(c *gin.Context) {
	res, err := h.services.UserService.GetUserById(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) updateUser(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.services.UserService.UpdateUser(&input, c.Param("id"))
	if err != nil {
		newErrorResponse(c, 501, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) deleteUser(c *gin.Context) {
	_, err := h.services.UserService.DeleteUser(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (h *Handler) getTasksForUser(c *gin.Context) {
	res, err := h.services.UserService.GetTasksForUser(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) searchUser(c *gin.Context) {
	var users []model.User
	var err error
	name := c.Query("name")
	email := c.Query("email")
	if email == "" {
		users, err = h.services.UserService.SearchUser(name, "name")
	} else {
		users, err = h.services.UserService.SearchUser(email, "email")
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}
