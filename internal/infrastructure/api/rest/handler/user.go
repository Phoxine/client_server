package handler

import (
	// database "client_server/database"

	users "client_server/internal/domain/users"
	jwt "client_server/pkg/jwt"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	// "gorm.io/gorm"
)

func (h *Handler) RegisterUserRoutes(g *echo.Group) {
	g.GET("/users/:id", h.getUser)
	g.POST("/users", h.saveUser)
	g.PUT("/users/:id", h.updateUser)
	g.DELETE("/users/:id", h.deleteUser)
	g.GET("/users", h.listUser)
	g.POST("/users/client", h.saveClientUser)
}

// @Summary List users
// @Description List users
// @Tags users
// @Accept json
// @Produce json
// @Router /users [get]
// @Security OAuth2Password
func (h *Handler) listUser(c echo.Context) error {
	var users []users.Users
	users, err := h.userService.ListUser(c.Request().Context())
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to list users"})
	}
	return c.JSON(http.StatusOK, users)
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Router /users/{id} [get]
// @Security OAuth2Password
func (h *Handler) getUser(c echo.Context) error {
	var user users.Users
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	user, err := h.userService.GetUser(c.Request().Context(), idInt)
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

type CreateUserRequest struct {
	Name  string `json:"name" example:"Alice"`
	Email string `json:"email" example:"alice@example.com"`
}

// @Summary Save user
// @Description Save user
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Create User Request"
// @Router /users [post]
// @Security OAuth2Password
func (h *Handler) saveUser(c echo.Context) error {
	var req CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	newUser := users.Users{Name: req.Name, Email: req.Email, CreatedAt: time.Now()}
	id, err := h.userService.CreateUser(c.Request().Context(), newUser)
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "Failed to create user",
		})
	}

	return c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

type UpdateUserRequest struct {
	Name  string `json:"name" example:"Alice"`
	Email string `json:"email" example:"alice@example.com"`
}

// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body UpdateUserRequest true "Update User Request"
// @Router /users/{id} [put]
// @Security OAuth2Password
func (h *Handler) updateUser(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	var user users.Users
	user, err := h.userService.GetUser(c.Request().Context(), idInt)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	var req UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	affectedRowCount, err := h.userService.UpdateUser(c.Request().Context(), user)
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Failed to update user"})
	}

	return c.JSON(http.StatusOK, map[string]int{
		"affectedRowCount": affectedRowCount,
	})
}

// @Summary Delete user
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Router /users/{id} [delete]
// @Security OAuth2Password
func (h *Handler) deleteUser(c echo.Context) error {
	// Get user ID from the body
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	if _, err := h.userService.GetUser(c.Request().Context(), idInt); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	_, err := h.userService.DeleteUser(c.Request().Context(), []int{idInt})
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Failed to delete user"})
	}
	return c.String(http.StatusAccepted, "")
}

type CreateClientUserRequest struct {
	Token   string `json:"token" example:"aaa.bbb.ccc"`
	JwksURL string `json:"jwksURL" example:"https://client.com/.well-known/jwks.json"`
}

// @Summary Save Client user
// @Description Save Client user by token
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateClientUserRequest true "Create Client User Request"
// @Router /users/client [post]
// @Security OAuth2Password
func (h *Handler) saveClientUser(c echo.Context) error {
	var req CreateClientUserRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	client := &http.Client{}
	resp, err := client.Get(req.JwksURL)
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "Failed to get jwks",
		})
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "Failed to read jwks",
		})
	}

	jwtHandler := jwt.NewJWTHandler(h.log)

	token, err := jwtHandler.ParseJWTWithJWKSet(req.Token, string(body))
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse jwt",
		})
	}
	var user users.Users
	name := jwtHandler.GetClaimsValuesByKey(token, "user_name", "test").(string)
	email := fmt.Sprintf("%s@%s", name, jwtHandler.GetClaimsValuesByKey(token, "iss", "client.com"))
	user.Name = name
	user.Email = email
	user.CreatedAt = time.Now()
	id, err := h.userService.CreateUser(c.Request().Context(), user)

	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Failed to create user"})
	}
	return c.JSON(http.StatusAccepted, map[string]int{"id": id})
}
