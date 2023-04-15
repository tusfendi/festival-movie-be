package entity

import "time"

type User struct {
	ID        int64  `gorm:"column:id; PRIMARY KEY" json:"id"`
	Username  string `gorm:"column:username" json:"username" binding:"required"`
	FullName  string `gorm:"column:full_name" json:"full_name" binding:"required"`
	Password  string `gorm:"column:password" json:"password" binding:"required"`
	Level     string `gorm:"column:level" json:"level" binding:"required"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (User) TableName() string {
	return "user"
}
