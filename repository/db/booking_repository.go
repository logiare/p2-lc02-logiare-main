package db

import (
	"fmt"
	"lc2/domain"
	"lc2/model/bookings"

	"gorm.io/gorm"
)

type bookingDBConn struct {
	db *gorm.DB
}

func BookingDBconn(db *gorm.DB) domain.BookingRepository {
	return &bookingDBConn{db: db}
}

// GetCustomerIDByUserID returns the customer_id linked to a given user_id
func (r *bookingDBConn) GetCustomerIDByUserID(userID int) (int, error) {
	var customerID int
	query := `SELECT customer_id FROM customers WHERE user_id = $1 LIMIT 1`
	result := r.db.Raw(query, userID).Scan(&customerID)
	if result.Error != nil {
		return 0, fmt.Errorf("error querying customer: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("customer not found")
	}
	return customerID, nil
}

// GetBookingsByCustomerID retrieves all bookings for a customer with tour names
func (r *bookingDBConn) GetBookingsByCustomerID(customerID int) ([]bookings.BookingResponse, error) {
	var result []bookings.BookingResponse
	query := `
		SELECT
			b.booking_id,
			t.tour_name,
			TO_CHAR(b.booking_date, 'YYYY-MM-DD') AS booking_date,
			b.booking_status
		FROM bookings b
		JOIN tour_bookings tb ON tb.booking_id = b.booking_id
		JOIN tours t          ON t.tour_id = tb.tour_id
		WHERE b.customer_id = $1
		ORDER BY b.booking_id
	`
	if err := r.db.Raw(query, customerID).Scan(&result).Error; err != nil {
		return nil, fmt.Errorf("error querying bookings: %w", err)
	}
	return result, nil
}

// GetUnpaidBookingsByCustomerID retrieves bookings that have no corresponding payment record
func (r *bookingDBConn) GetUnpaidBookingsByCustomerID(customerID int) ([]bookings.BookingResponse, error) {
	var result []bookings.BookingResponse
	query := `
		SELECT
			b.booking_id,
			t.tour_name,
			TO_CHAR(b.booking_date, 'YYYY-MM-DD') AS booking_date,
			b.booking_status
		FROM bookings b
		JOIN tour_bookings tb ON tb.booking_id = b.booking_id
		JOIN tours t          ON t.tour_id = tb.tour_id
		LEFT JOIN payments p  ON p.booking_id = b.booking_id
		WHERE b.customer_id = $1
		  AND p.payment_id IS NULL
		ORDER BY b.booking_id
	`
	if err := r.db.Raw(query, customerID).Scan(&result).Error; err != nil {
		return nil, fmt.Errorf("error querying unpaid bookings: %w", err)
	}
	return result, nil
}
