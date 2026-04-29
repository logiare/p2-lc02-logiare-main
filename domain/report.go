package domain

import (
	"context"
	"lc2/model/tours"

	"github.com/labstack/echo/v4"
)

type ReporRepository interface {
	GetTotalCustomers() (int, error)
	GetBookingsPerTour() ([]tours.BookingsPerTourResponse, error)
}

type ReportUseCase interface {
	GetTotalCustomers(ctx context.Context) (int, error)
	GetBookingsPerTour(ctx context.Context) ([]tours.BookingsPerTourResponse, error)
}

type ReportHandler interface {
	GetTotalCustomers(c echo.Context) (int, error)
	GetBookingsPerTour(c echo.Context) (int, error)
}
