package entity

import "time"

type Actor struct {
	ID        int64  `gorm:"column:id; PRIMARY KEY" json:"id"`
	Name      string `gorm:"column:name" json:"name" binding:"required"`
	Gender    string `gorm:"column:gender" json:"gender" binding:"required"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (Actor) TableName() string {
	return "actor"
}
