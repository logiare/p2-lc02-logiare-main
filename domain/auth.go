package domain

import (
	"context"
	"lc2/model/users"

	"github.com/labstack/echo/v4"
)

type AuthRepository interface {
	CreateUserAndCustomer(user users.User, customer users.Customer) (int, error)
	GetUserByEmail(email string) (users.User, error)
	UpdateUserJWTAndLogin(userID int, token string) error
}

type AuthUseCase interface {
	Register(ctx context.Context, req users.User) (users.ResponseRegister, error)
	Login(ctx context.Context, req users.Login) (users.ResponseLogin, error)
}

type AuthHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}
