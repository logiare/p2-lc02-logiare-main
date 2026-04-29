package tours

type TourEarningsResponse struct {
	TourID        int    `json:"tour_id"`
	TourName      string `json:"tour_name"`
	TotalEarnings string `json:"total_earnings"`
}

type BookingsPerTourResponse struct {
	TourName      string `json:"tour_name"`
	TotalBookings string `json:"total_bookings"`
}
