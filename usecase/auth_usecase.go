package usecase

import (
	"context"
	"fmt"
	"lc2/domain"
	"lc2/model/users"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const jwtSecretKey = "SECRET_KEY_DONG"

type authUseCase struct {
	authRepository domain.AuthRepository
}

func AuthUseCase(repo domain.AuthRepository) domain.AuthUseCase {
	return &authUseCase{authRepository: repo}
}

// Register validates, hashes password, and creates user + customer records
func (u *authUseCase) Register(ctx context.Context, req users.Register) (users.ResponseRegister, error) {
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return users.ResponseRegister{}, fmt.Errorf("name, email, and password are required")
	}

	// Check if email already registered
	_, err := u.authRepository.GetUserByEmail(req.Email)
	if err == nil {
		return users.ResponseRegister{}, fmt.Errorf("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return users.ResponseRegister{}, fmt.Errorf("error hashing password: %w", err)
	}

	user := users.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	customer := users.Customer{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}

	userID, err := u.authRepository.CreateUserAndCustomer(user, customer)
	if err != nil {
		return users.ResponseRegister{}, fmt.Errorf("error creating user: %w", err)
	}

	return users.ResponseRegister{
		UserID: userID,
		Email:  req.Email,
	}, nil
}

// Login authenticates the user and returns a signed JWT token
func (u *authUseCase) Login(ctx context.Context, req users.Login) (users.ResponseLogin, error) {
	if req.Email == "" || req.Password == "" {
		return users.ResponseLogin{}, fmt.Errorf("email and password are required")
	}

	// Fetch user
	user, err := u.authRepository.GetUserByEmail(req.Email)
	if err != nil {
		return users.ResponseLogin{}, fmt.Errorf("user not found")
	}

	// Verify password
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return users.ResponseLogin{}, fmt.Errorf("invalid password")
	}

	// Generate JWT
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return users.ResponseLogin{}, fmt.Errorf("error generating token: %w", err)
	}

	// Persist token and update login date
	_ = u.authRepository.UpdateUserJWTAndLogin(user.UserID, tokenString)

	return users.ResponseLogin{Token: tokenString}, nil
}
