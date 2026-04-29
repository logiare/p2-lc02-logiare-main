package usecase

import (
	"context"
	"fmt"
	"lc2/domain"
	"lc2/model/tours"
)

type tourUseCase struct {
	tourRepository domain.TourRepository
}

func TourUseCase(repo domain.TourRepository) domain.TourUseCase {
	return &tourUseCase{tourRepository: repo}
}

// GetTourEarnings returns total earnings for every tour
func (u *tourUseCase) GetTourEarnings(ctx context.Context) ([]tours.TourEarningsResponse, error) {
	result, err := u.tourRepository.GetTourEarnings()
	if err != nil {
		return nil, fmt.Errorf("error fetching tour earnings: %w", err)
	}
	return result, nil
}
