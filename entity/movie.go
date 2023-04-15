package entity

import "time"

type Movie struct {
	ID          int64  `gorm:"column:id; PRIMARY KEY" json:"id"`
	MovieFileID int64  `gorm:"column:movies_file_id" json:"movies_file_id"`
	Title       string `gorm:"column:title" json:"title" binding:"required"`
	Description string `gorm:"column:description" json:"description" binding:"required"`
	Duration    string `gorm:"column:duration" json:"duration"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (Movie) TableName() string {
	return "movie"
}
