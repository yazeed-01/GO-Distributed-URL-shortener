package models

import "time"

type URL struct {
	ID        uint   `gorm:"primaryKey"`
	LongURL   string `gorm:"not null"`
	ShortURL  string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
}
