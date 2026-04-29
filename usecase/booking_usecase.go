package usecase

import (
	"context"
	"fmt"
	"lc2/domain"
	"lc2/model/bookings"
)

type bookingUseCase struct {
	bookingRepository domain.BookingRepository
}

func BookingUseCase(repo domain.BookingRepository) domain.BookingUseCase {
	return &bookingUseCase{bookingRepository: repo}
}

// GetAllBookings retrieves all bookings for the logged-in user
func (u *bookingUseCase) GetAllBookings(ctx context.Context, userID int) ([]bookings.BookingResponse, error) {
	customerID, err := u.bookingRepository.GetCustomerIDByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	result, err := u.bookingRepository.GetBookingsByCustomerID(customerID)
	if err != nil {
		return nil, fmt.Errorf("error fetching bookings: %w", err)
	}
	return result, nil
}

// GetUnpaidBookings retrieves bookings with no payment for the logged-in user
func (u *bookingUseCase) GetUnpaidBookings(ctx context.Context, userID int) ([]bookings.BookingResponse, error) {
	customerID, err := u.bookingRepository.GetCustomerIDByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	result, err := u.bookingRepository.GetUnpaidBookingsByCustomerID(customerID)
	if err != nil {
		return nil, fmt.Errorf("error fetching unpaid bookings: %w", err)
	}
	return result, nil
}
