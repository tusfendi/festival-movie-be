package entity

import "time"

type MoviesGenre struct {
	ID        int64 `gorm:"column:id; PRIMARY KEY" json:"id"`
	GenreID   int64 `gorm:"column:genre_id" json:"genre_id" binding:"required"`
	MovieID   int64 `gorm:"column:movie_id" json:"movie_id" binding:"required"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type MoviesGenreSelect struct {
	ID        int64  `gorm:"column:id; PRIMARY KEY" json:"id"`
	GenreID   int64  `gorm:"column:genre_id" json:"genre_id" binding:"required"`
	Genre     string `gorm:"column:genre" json:"genre" binding:"required"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (MoviesGenre) TableName() string {
	return "movies_genre"
}

func (MoviesGenreSelect) TableName() string {
	return "movies_genre"
}
