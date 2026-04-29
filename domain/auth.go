package domain

import (
	"context"
	"lc2/model/users"

	"github.com/labstack/echo/v4"
)

// AuthRepository defines the data-layer contract for authentication
type AuthRepository interface {
	CreateUserAndCustomer(user users.User, customer users.Customer) (int, error)
	GetUserByEmail(email string) (users.User, error)
	UpdateUserJWTAndLogin(userID int, token string) error
}

// AuthUseCase defines the business-logic contract for authentication
type AuthUseCase interface {
	Register(ctx context.Context, req users.Register) (users.ResponseRegister, error)
	Login(ctx context.Context, req users.Login) (users.ResponseLogin, error)
}

// AuthHandler defines the HTTP-layer contract for authentication
type AuthHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}
