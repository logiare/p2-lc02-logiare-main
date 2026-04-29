package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, *sql.DB) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://ud1b5uav6a245k:p8bc84b6ffbc9536c187c3e6720cf2f4a72ad041080d64b6afd757abf49814be6@c1jpc731rp0brl.cluster-czrs8kj4isg7.us-east-1.rds.amazonaws.com:5432/d1uup71jiu8nci?sslmode=require&search_path=yuriz_lc2"
	}

	sqlDbExst, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDbExst,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open GORM:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB:", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("DB not reachable:", err)
	}

	fmt.Println("✅ Database connected!")
	return db, sqlDB
}
