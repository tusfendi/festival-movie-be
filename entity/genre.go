package entity

import "time"

type Genre struct {
	ID        int64  `gorm:"column:id; PRIMARY KEY" json:"id"`
	Genre     string `gorm:"column:genre" json:"genre" binding:"required"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (Genre) TableName() string {
	return "genre"
}
