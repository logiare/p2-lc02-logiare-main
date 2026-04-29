package tours

// Tour maps to the "tours" table
type Tour struct {
	TourID          int     `gorm:"column:tour_id;primaryKey;autoIncrement" json:"tour_id"`
	TourName        string  `gorm:"column:tour_name;not null"               json:"tour_name"`
	TourDescription string  `gorm:"column:tour_description"                 json:"tour_description"`
	TourPrice       float64 `gorm:"column:tour_price;not null"              json:"tour_price"`
	TourDuration    string  `gorm:"column:tour_duration;not null"           json:"tour_duration"`
}

func (Tour) TableName() string { return "tours" }
