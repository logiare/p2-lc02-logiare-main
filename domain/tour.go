package domain

import (
	"context"
	"lc2/model/tours"
)

type TourRepository interface {
	GetTourEarnings() ([]tours.TourEarningsResponse, error)
}

type ReportUseCase interface {
	GetTourEarnings(ctx context.Context) ([]tours.TourEarningsResponse, error)
}
