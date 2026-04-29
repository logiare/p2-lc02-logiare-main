package bookings

// BookingResponse is the response shape for GET /bookings and GET /bookings/unpaid
type BookingResponse struct {
	BookingID     int    `json:"booking_id"`
	TourName      string `json:"tour_name"`
	BookingDate   string `json:"booking_date"`
	BookingStatus string `json:"booking_status"`
}
