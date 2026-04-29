package bookings

type BookingResponse struct {
	BookingId     string `json:"booking_id"`
	TourName      string `json:"tour_name"`
	BookingDate   string `json:"booking_date"`
	BookingStatus string `json:"booking_status"`
}
