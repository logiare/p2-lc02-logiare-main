package http

import (
	"lc2/domain"
	"lc2/helper"
	"lc2/model/users"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authUseCase domain.AuthUseCase
}

func AuthHandler(authUseCase domain.AuthUseCase) domain.AuthHandler {
	return &authHandler{authUseCase: authUseCase}
}

// Register godoc
// POST /users/register
func (h *authHandler) Register(c echo.Context) error {
	var req users.Register
	if err := c.Bind(&req); err != nil {
		return helper.RespondError(c, http.StatusBadRequest, "invalid request body")
	}

	if req.Email == "" || req.Password == "" || req.Name == "" {
		return helper.RespondError(c, http.StatusBadRequest, "name, email, and password are required")
	}

	resp, err := h.authUseCase.Register(c.Request().Context(), req)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "already registered") {
			return helper.RespondError(c, http.StatusBadRequest, msg)
		}
		return helper.RespondError(c, http.StatusInternalServerError, "internal server error")
	}

	return helper.RespondJSON(c, http.StatusCreated, resp)
}

// Login godoc
// POST /users/login
func (h *authHandler) Login(c echo.Context) error {
	var req users.Login
	if err := c.Bind(&req); err != nil {
		return helper.RespondError(c, http.StatusBadRequest, "invalid request body")
	}

	if req.Email == "" || req.Password == "" {
		return helper.RespondError(c, http.StatusBadRequest, "email and password are required")
	}

	resp, err := h.authUseCase.Login(c.Request().Context(), req)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "not found") {
			return helper.RespondError(c, http.StatusNotFound, "user not found")
		}
		if strings.Contains(msg, "invalid password") || strings.Contains(msg, "required") {
			return helper.RespondError(c, http.StatusBadRequest, msg)
		}
		return helper.RespondError(c, http.StatusInternalServerError, "internal server error")
	}

	return helper.RespondJSON(c, http.StatusOK, resp)
}
