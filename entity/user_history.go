package entity

import "time"

type UserHisory struct {
	ID        string `gorm:"column:guid; PRIMARY KEY" json:"id"`
	UserID    int64  `gorm:"column:user_id" json:"user_id"`
	MovieID   int64  `gorm:"column:movie_id" json:"movie_id"`
	CreatedAt *time.Time
}

func (UserHisory) TableName() string {
	return "user_history"
}
