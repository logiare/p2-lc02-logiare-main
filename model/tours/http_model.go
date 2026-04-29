package tours

// TourEarningsResponse is the response shape for GET /tours/earning
type TourEarningsResponse struct {
	TourID        int     `json:"tour_id"`
	TourName      string  `json:"tour_name"`
	TotalEarnings float64 `json:"total_earnings"`
}

// BookingsPerTourResponse is the response shape for GET /reports/bookings-per-tour
type BookingsPerTourResponse struct {
	TourName      string `json:"tour_name"`
	TotalBookings int    `json:"total_bookings"`
}
