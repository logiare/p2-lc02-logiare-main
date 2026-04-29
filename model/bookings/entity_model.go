package bookings

import "time"

// Booking maps to the "bookings" table
type Booking struct {
	BookingID     int       `gorm:"column:booking_id;primaryKey;autoIncrement" json:"booking_id"`
	CustomerID    int       `gorm:"column:customer_id"                         json:"customer_id"`
	BookingDate   time.Time `gorm:"column:booking_date"                        json:"booking_date"`
	BookingStatus string    `gorm:"column:booking_status"                      json:"booking_status"`
}

func (Booking) TableName() string { return "bookings" }

// TourBooking maps to the "tour_bookings" table
type TourBooking struct {
	TourBookingID int       `gorm:"column:tour_booking_id;primaryKey;autoIncrement" json:"tour_booking_id"`
	BookingID     int       `gorm:"column:booking_id"                               json:"booking_id"`
	TourID        int       `gorm:"column:tour_id"                                  json:"tour_id"`
	TourStartDate time.Time `gorm:"column:tour_start_date"                          json:"tour_start_date"`
	TourEndDate   time.Time `gorm:"column:tour_end_date"                            json:"tour_end_date"`
	NumOfPeople   int       `gorm:"column:num_of_people"                            json:"num_of_people"`
}

func (TourBooking) TableName() string { return "tour_bookings" }

// Payment maps to the "payments" table
type Payment struct {
	PaymentID     int     `gorm:"column:payment_id;primaryKey;autoIncrement" json:"payment_id"`
	BookingID     int     `gorm:"column:booking_id"                          json:"booking_id"`
	PaymentDate   string  `gorm:"column:payment_date"                        json:"payment_date"`
	PaymentAmount float64 `gorm:"column:payment_amount"                      json:"payment_amount"`
}

func (Payment) TableName() string { return "payments" }
