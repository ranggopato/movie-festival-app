package models

type Vote struct {
	UserID  uint `gorm:"primaryKey"`
	MovieID uint `gorm:"primaryKey"`
}

type MovieWithVotes struct {
	Movie
	VoteCount int `json:"votes"`
}
