package main

import (
	"fmt"
	"lc2/config"
	http2 "lc2/delivery/http"
	"lc2/helper"
	"lc2/repository/db"
	"lc2/usecase"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const jwtSecretKey = "SECRET_KEY_DONG"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Database connection
	dbConn, sqlDbConn := config.ConnectDB()
	defer sqlDbConn.Close()

	// Repositories
	authRepository := db.AuthDBconn(dbConn)
	bookingRepository := db.BookingDBconn(dbConn)
	tourRepository := db.TourDBconn(dbConn)
	reportRepository := db.ReportDBconn(dbConn)

	// Use cases (business logic)
	authUseCase := usecase.AuthUseCase(authRepository)
	bookingUseCase := usecase.BookingUseCase(bookingRepository)
	tourUseCase := usecase.TourUseCase(tourRepository)
	reportUseCase := usecase.ReportUseCase(reportRepository)

	// Handlers
	authHandler := http2.AuthHandler(authUseCase)
	bookingHandler := http2.BookingHandler(bookingUseCase)
	tourHandler := http2.TourHandler(tourUseCase)
	reportHandler := http2.ReportHandler(reportUseCase)

	// Echo setup
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(helper.MiddlewareLogger)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Public routes
	e.POST("/users/register", authHandler.Register)
	e.POST("/users/login", authHandler.Login)

	// Protected routes
	protected := e.Group("")
	protected.Use(validateJWT)

	protected.GET("/bookings", bookingHandler.GetAllBookings)
	protected.GET("/bookings/unpaid", bookingHandler.GetUnpaidBookings)
	protected.GET("/tours/earning", tourHandler.GetTourEarnings)
	protected.GET("/reports/total-customers", reportHandler.GetTotalCustomers)
	protected.GET("/reports/bookings-per-tour", reportHandler.GetBookingsPerTour)

	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("failed to start server:", err)
	}
}

// validateJWT is the JWT middleware that protects routes
func validateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "authorization header is required")
		}

		// Expect: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization format, use: Bearer <token>")
		}

		tokenString := parts[1]

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
		}

		// Extract claims and set to context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
		}

		// user_id stored as float64 in JWT claims
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token: missing user_id")
		}

		c.Set("user_id", int(userIDFloat))
		c.Set("email", claims["email"])

		return next(c)
	}
}
