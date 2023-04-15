package entity

import "time"

type MoviesLike struct {
	ID        int64 `gorm:"column:id; PRIMARY KEY" json:"id"`
	UserID    int64 `gorm:"column:user_id" json:"user_id" binding:"required"`
	MovieID   int64 `gorm:"column:movie_id" json:"movie_id" binding:"required"`
	CreatedAt *time.Time
}

func (MoviesLike) TableName() string {
	return "movies_like"
}
