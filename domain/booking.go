package domain

import (
	"context"
	"lc2/model/bookings"

	"github.com/labstack/echo/v4"
)

type BookingRepository interface {
	GetBookingsByCustomerID(customerID int) ([]bookings.BookingResponse, error)
	GetUnpaidBookingsByCustomerID(customerID int) ([]bookings.BookingResponse, error)
	GetCustomerIDByUserId(userID int) (int, error)
}

type BookingUsecase interface {
	GetAllBookings(ctx context.Context, userId int) ([]bookings.BookingResponse, error)
	GetUnpaidBookings(ctx context.Context, userId int) ([]bookings.BookingResponse, error)
}

type BookingHandler interface {
	GetAllBookings(c echo.Context) error
	GetUnpaidBookings(c echo.Context) error
}
