package handler

import (
	"DDD/app/entity"
	"DDD/app/services"
	"DDD/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userService.GetUserById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	err := h.userService.CreateUser(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "server error",
		})
	}
	return c.JSON(http.StatusCreated, user)
}

func AuthMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := middleware.ExtractTokenUser(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "data not found",
			})
		}
		c.Set("user", user)
		return next(c)
	}
}
