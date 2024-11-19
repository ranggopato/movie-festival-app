package models

// Movie adalah struktur untuk menyimpan informasi film
type Movie struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null" binding:"required"` // Title wajib diisi
	Description string `binding:"required"`                 // Description wajib diisi
	Duration    int    `gorm:"default:0"`
	Artists     string
	Genres      string
	WatchURL    string `gorm:"not null" binding:"required,url"` // WatchURL wajib diisi
	Views       int    `gorm:"default:0"`
}
