package entity

import "time"

type MoviesActor struct {
	ID        int64 `gorm:"column:id; PRIMARY KEY" json:"id"`
	ActorID   int64 `gorm:"column:actor_id" json:"actor_id" binding:"required"`
	MovieID   int64 `gorm:"column:movie_id" json:"movie_id" binding:"required"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (MoviesActor) TableName() string {
	return "movies_actor"
}
