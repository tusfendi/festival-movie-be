package entity

import "time"

type UserHistory struct {
	ID        string `gorm:"column:guid; PRIMARY KEY" json:"id"`
	UserID    string `gorm:"column:user_id" json:"user_id"`
	MovieID   int64  `gorm:"column:movie_id" json:"movie_id"`
	GenreID   int64  `gorm:"column:genre_id" json:"genre_id"`
	CreatedAt *time.Time
}

func (UserHistory) TableName() string {
	return "user_history"
}
