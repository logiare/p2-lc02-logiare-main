package usecase

import (
	"context"
	"fmt"
	"lc2/domain"
	"lc2/model/tours"
)

type reportUseCase struct {
	reportRepository domain.ReportRepository
}

func ReportUseCase(repo domain.ReportRepository) domain.ReportUseCase {
	return &reportUseCase{reportRepository: repo}
}

// GetTotalCustomers returns the count of all registered customers
func (u *reportUseCase) GetTotalCustomers(ctx context.Context) (int, error) {
	total, err := u.reportRepository.GetTotalCustomers()
	if err != nil {
		return 0, fmt.Errorf("error fetching total customers: %w", err)
	}
	return total, nil
}

// GetBookingsPerTour returns total bookings for each tour
func (u *reportUseCase) GetBookingsPerTour(ctx context.Context) ([]tours.BookingsPerTourResponse, error) {
	result, err := u.reportRepository.GetBookingsPerTour()
	if err != nil {
		return nil, fmt.Errorf("error fetching bookings per tour: %w", err)
	}
	return result, nil
}
