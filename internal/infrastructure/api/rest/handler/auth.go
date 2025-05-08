package handler

import (
	"client_server/pkg/jwt"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterAuthRoutes(g *echo.Group) {
	g.POST("/auth/login", h.login)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Router /auth/login [post]
// @Param request body LoginRequest true "Login Request"
func (h *Handler) login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	// check username and password
	jwtHandler := jwt.NewJWTHandler(h.log)
	accessToken, err := jwtHandler.GenerateJWT(map[string]interface{}{
		"username": req.Username,
	}, []byte("secret"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate jwt",
		})
	}
	var token = map[string]interface{}{
		"access_token":  accessToken,
		"token_type":    "bearer",
		"expires_in":    600,
		"refresh_token": "",
	}
	h.log.Info(fmt.Sprintf("token: %s", token))
	return c.JSON(
		http.StatusOK,
		token,
	)
}
