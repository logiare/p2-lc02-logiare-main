package db

import (
	"fmt"
	"lc2/domain"
	"lc2/model/tours"

	"gorm.io/gorm"
)

type reportDBConn struct {
	db *gorm.DB
}

func ReportDBconn(db *gorm.DB) domain.ReportRepository {
	return &reportDBConn{db: db}
}

// GetTotalCustomers returns the total number of registered customers
func (r *reportDBConn) GetTotalCustomers() (int, error) {
	var total int
	query := `SELECT COUNT(*) FROM customers`
	if err := r.db.Raw(query).Scan(&total).Error; err != nil {
		return 0, fmt.Errorf("error counting customers: %w", err)
	}
	return total, nil
}

// GetBookingsPerTour returns the number of bookings for each tour
func (r *reportDBConn) GetBookingsPerTour() ([]tours.BookingsPerTourResponse, error) {
	var result []tours.BookingsPerTourResponse
	query := `
		SELECT
			t.tour_name,
			COUNT(tb.booking_id) AS total_bookings
		FROM tours t
		LEFT JOIN tour_bookings tb ON tb.tour_id = t.tour_id
		GROUP BY t.tour_id, t.tour_name
		ORDER BY t.tour_id
	`
	if err := r.db.Raw(query).Scan(&result).Error; err != nil {
		return nil, fmt.Errorf("error querying bookings per tour: %w", err)
	}
	return result, nil
}
