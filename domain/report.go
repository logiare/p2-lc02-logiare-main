package domain

import (
	"context"
	"lc2/model/tours"

	"github.com/labstack/echo/v4"
)

// ReportRepository defines the data-layer contract for reports
type ReportRepository interface {
	GetTotalCustomers() (int, error)
	GetBookingsPerTour() ([]tours.BookingsPerTourResponse, error)
}

// ReportUseCase defines the business-logic contract for reports
type ReportUseCase interface {
	GetTotalCustomers(ctx context.Context) (int, error)
	GetBookingsPerTour(ctx context.Context) ([]tours.BookingsPerTourResponse, error)
}

// ReportHandler defines the HTTP-layer contract for reports
type ReportHandler interface {
	GetTotalCustomers(c echo.Context) error
	GetBookingsPerTour(c echo.Context) error
}
