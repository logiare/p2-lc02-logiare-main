package domain

import (
	"context"
	"lc2/model/bookings"

	"github.com/labstack/echo/v4"
)

// BookingRepository defines the data-layer contract for bookings
type BookingRepository interface {
	GetBookingsByCustomerID(customerID int) ([]bookings.BookingResponse, error)
	GetUnpaidBookingsByCustomerID(customerID int) ([]bookings.BookingResponse, error)
	GetCustomerIDByUserID(userID int) (int, error)
}

// BookingUseCase defines the business-logic contract for bookings
type BookingUseCase interface {
	GetAllBookings(ctx context.Context, userID int) ([]bookings.BookingResponse, error)
	GetUnpaidBookings(ctx context.Context, userID int) ([]bookings.BookingResponse, error)
}

// BookingHandler defines the HTTP-layer contract for bookings
type BookingHandler interface {
	GetAllBookings(c echo.Context) error
	GetUnpaidBookings(c echo.Context) error
}
