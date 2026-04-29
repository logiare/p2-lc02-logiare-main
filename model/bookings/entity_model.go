package bookings

import "time"

type Booking struct {
	BookingId     int       `gorm:"column:booking_id;primaryKey;autoIncrement"  json:"booking_id"`
	CustomerId    int       `gorm:"column:customer_id" json:"customer_id"`
	BookingDate   time.Time `gorm:"column:booking_date" json:"booking_date" `
	BookingStatus string    `gorm:"column:booking_status" json:"booking_status"`
}

func (Booking) TableName() string { return "bookings" }

type TourBooking struct {
	TourBookingId int       `gorm:"column:tour_booking_id;primaryKey" json:"tour_booking_id"`
	BookingId     int       `gorm:"column:booking_id" json:"booking_id"`
	TourId        int       `gorm:"column:tour_id" json:"tour_id"`
	TourDate      time.Time `gorm:"column:tour_date" json:"tour_date"`
	TourEndDate   time.Time `gorm:"column:tour_end_date" json:"tour_end_date"`
	NumOfPeople   int       `gorm:"column:num_of_people" json:"num_of_people"`
}

func (TourBooking) TableName() string { return "tour_bookings" }

type Payment struct {
	PaymentId     int       `gorm:"column:payment_id" json:"payment_id"`
	BookingId     int       `gorm:"column:booking_id" json:"booking_id"`
	PaymentDate   time.Time `gorm:"column:payment_date" json:"payment_date"`
	PaymentAmount int       `gorm:"column:payment_amount" json:"payment_amount"`
}

func (Payment) TableName() string { return "payments" }
