package domain

import (
	"context"
	"lc2/model/tours"

	"github.com/labstack/echo/v4"
)

// TourRepository defines the data-layer contract for tours
type TourRepository interface {
	GetTourEarnings() ([]tours.TourEarningsResponse, error)
}

// TourUseCase defines the business-logic contract for tours
type TourUseCase interface {
	GetTourEarnings(ctx context.Context) ([]tours.TourEarningsResponse, error)
}

// TourHandler defines the HTTP-layer contract for tours
type TourHandler interface {
	GetTourEarnings(c echo.Context) error
}
