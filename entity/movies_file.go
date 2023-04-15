package entity

import "time"

type MoviesFile struct {
	ID        int64  `gorm:"column:id; PRIMARY KEY" json:"id"`
	FilePath  string `gorm:"column:file_path" json:"file_path" binding:"required"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (MoviesFile) TableName() string {
	return "movies_file"
}
