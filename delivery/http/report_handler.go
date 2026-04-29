package http

import (
	"lc2/domain"
	"lc2/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type reportHandler struct {
	reportUseCase domain.ReportUseCase
}

func ReportHandler(reportUseCase domain.ReportUseCase) domain.ReportHandler {
	return &reportHandler{reportUseCase: reportUseCase}
}

// GetTotalCustomers godoc
// GET /reports/total-customers
func (h *reportHandler) GetTotalCustomers(c echo.Context) error {
	total, err := h.reportUseCase.GetTotalCustomers(c.Request().Context())
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, "internal server error")
	}

	return helper.RespondJSON(c, http.StatusOK, map[string]int{
		"total_customers": total,
	})
}

// GetBookingsPerTour godoc
// GET /reports/bookings-per-tour
func (h *reportHandler) GetBookingsPerTour(c echo.Context) error {
	result, err := h.reportUseCase.GetBookingsPerTour(c.Request().Context())
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, "internal server error")
	}

	if len(result) == 0 {
		return helper.RespondError(c, http.StatusNotFound, "data not found")
	}

	return helper.RespondJSON(c, http.StatusOK, result)
}
