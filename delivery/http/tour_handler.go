package http

import (
	"lc2/domain"
	"lc2/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type tourHandler struct {
	tourUseCase domain.TourUseCase
}

func TourHandler(tourUseCase domain.TourUseCase) domain.TourHandler {
	return &tourHandler{tourUseCase: tourUseCase}
}

// GetTourEarnings godoc
// GET /tours/earning
func (h *tourHandler) GetTourEarnings(c echo.Context) error {
	result, err := h.tourUseCase.GetTourEarnings(c.Request().Context())
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, "internal server error")
	}

	if len(result) == 0 {
		return helper.RespondError(c, http.StatusNotFound, "data not found")
	}

	return helper.RespondJSON(c, http.StatusOK, result)
}
