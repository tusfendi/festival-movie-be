package entity

import "time"

type UserToken struct {
	ID        string `gorm:"column:guid; PRIMARY KEY" json:"id"`
	UserID    int64  `gorm:"column:id" json:"user_id"`
	CreatedAt *time.Time
}

func (UserToken) TableName() string {
	return "UserToken"
}
