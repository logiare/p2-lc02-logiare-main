package http

import (
	"lc2/domain"
	"lc2/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bookingHandler struct {
	bookingUseCase domain.BookingUseCase
}

func BookingHandler(bookingUseCase domain.BookingUseCase) domain.BookingHandler {
	return &bookingHandler{bookingUseCase: bookingUseCase}
}

// GetAllBookings godoc
// GET /bookings
func (h *bookingHandler) GetAllBookings(c echo.Context) error {
	userID := c.Get("user_id").(int)

	result, err := h.bookingUseCase.GetAllBookings(c.Request().Context(), userID)
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, "internal server error")
	}

	if len(result) == 0 {
		return helper.RespondError(c, http.StatusNotFound, "data not found")
	}

	return helper.RespondJSON(c, http.StatusOK, result)
}

// GetUnpaidBookings godoc
// GET /bookings/unpaid
func (h *bookingHandler) GetUnpaidBookings(c echo.Context) error {
	userID := c.Get("user_id").(int)

	result, err := h.bookingUseCase.GetUnpaidBookings(c.Request().Context(), userID)
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, "internal server error")
	}

	if len(result) == 0 {
		return helper.RespondError(c, http.StatusNotFound, "data not found")
	}

	return helper.RespondJSON(c, http.StatusOK, result)
}
