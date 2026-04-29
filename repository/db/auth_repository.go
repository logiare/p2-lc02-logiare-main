package db

import (
	"fmt"
	"lc2/domain"
	"lc2/model/users"

	"gorm.io/gorm"
)

type authDBConn struct {
	db *gorm.DB
}

func AuthDBconn(db *gorm.DB) domain.AuthRepository {
	return &authDBConn{db: db}
}

// CreateUserAndCustomer inserts a new user and customer in a transaction
func (r *authDBConn) CreateUserAndCustomer(user users.User, customer users.Customer) (int, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("error creating user: %w", err)
		}
		customer.UserID = user.UserID
		if err := tx.Create(&customer).Error; err != nil {
			return fmt.Errorf("error creating customer: %w", err)
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return user.UserID, nil
}

// GetUserByEmail fetches a user record by email
func (r *authDBConn) GetUserByEmail(email string) (users.User, error) {
	var user users.User
	query := `SELECT user_id, email, password_hash FROM users WHERE email = $1 LIMIT 1`
	result := r.db.Raw(query, email).Scan(&user)
	if result.Error != nil {
		return users.User{}, fmt.Errorf("error querying user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return users.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

// UpdateUserJWTAndLogin updates jwt_token and last_login_date after successful login
func (r *authDBConn) UpdateUserJWTAndLogin(userID int, token string) error {
	query := `UPDATE users SET jwt_token = $1, last_login_date = NOW() WHERE user_id = $2`
	if err := r.db.Exec(query, token, userID).Error; err != nil {
		return fmt.Errorf("error updating user token: %w", err)
	}
	return nil
}
