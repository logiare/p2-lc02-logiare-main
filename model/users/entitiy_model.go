package users

import "time"

type User struct {
	UserId        int       `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	Email         string    `gorm:"column:email;uniqueIndex;not null" json:"email"`
	PasswordHash  string    `gorm:"column:password_hash;not null" json:"-"`
	LastLoginDate time.Time `gorm:"column:last_login_date" json:"last_login_date"`
	JWTToken      string    `gorm:"column:jwt_token" json:"jwt_token"`
}

func (User) TableName() string { return "users" }

type Customer struct {
	CustomerId  int    `gorm:"column:customer_id;primaryKey;autoIncrement" json:"customer_id"`
	UserId      int    `gorm:"column:user_id;" json:"user_id" json:"user_id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Email       string `gorm:"column:email;not null" json:"email"`
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`
	Address     string `json:"address" json:"address"`
}

func (Customer) TableName() string { return "customers" }
