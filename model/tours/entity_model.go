package tours

type Tours struct {
	TourID          int     `gorm,:"column:tour_id;primaryKey;autoIncrement" json:"tour_id"`
	TourName        string  `gorm:"column:tour_name;not null" json:"tour_name"`
	TourDescription string  `gorm:"column:tour_description" json:"tour_description"`
	TourPrice       float64 `gorm:"column:tour_price;not null" json:"tour_price"`
	TourDuration    string  `gorm:"column:tour_duration" json:"tour_duration"`
}

func (Tours) TableName() string { return "tours" }
