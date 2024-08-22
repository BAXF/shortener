package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type URL struct {
	ID        uint   `gorm:"primaryKey"`
	Original  string `gorm:"not null"`
	ShortURL  string `gorm:"not null;unique"`
	UserID    string `gorm:"not null"`
	Visits    int    `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//Shortened  string
}

func ConnectPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err

	}
	db.AutoMigrate(&URL{})
	return db, nil
}
