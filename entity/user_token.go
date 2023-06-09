package entity

import "time"

type UserToken struct {
	ID        string `gorm:"column:guid; PRIMARY KEY" json:"id"`
	UserID    int64  `gorm:"column:user_id" json:"user_id"`
	CreatedAt *time.Time
}

func (UserToken) TableName() string {
	return "user_token"
}
