package models

import (
	"time"
)

type Viewership struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`  // ID pengguna
	MovieID   uint      `gorm:"not null"`  // ID film
	WatchedAt time.Time `gorm:"not null"`  // Waktu mulai menonton
	Duration  int       `gorm:"default:0"` // Durasi menonton dalam detik
}
