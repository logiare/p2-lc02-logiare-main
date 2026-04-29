package db

import (
	"fmt"
	"lc2/domain"
	"lc2/model/tours"

	"gorm.io/gorm"
)

type tourDBConn struct {
	db *gorm.DB
}

func TourDBconn(db *gorm.DB) domain.TourRepository {
	return &tourDBConn{db: db}
}

// GetTourEarnings returns total payment amount grouped by tour
func (r *tourDBConn) GetTourEarnings() ([]tours.TourEarningsResponse, error) {
	var result []tours.TourEarningsResponse
	query := `
		SELECT
			t.tour_id,
			t.tour_name,
			COALESCE(SUM(p.payment_amount), 0) AS total_earnings
		FROM tours t
		LEFT JOIN tour_bookings tb ON tb.tour_id = t.tour_id
		LEFT JOIN payments p       ON p.booking_id = tb.booking_id
		GROUP BY t.tour_id, t.tour_name
		ORDER BY t.tour_id
	`
	if err := r.db.Raw(query).Scan(&result).Error; err != nil {
		return nil, fmt.Errorf("error querying tour earnings: %w", err)
	}
	return result, nil
}
