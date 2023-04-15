package repository

import (
	"github.com/tusfendi/festival-movie-be/config"
	"github.com/tusfendi/festival-movie-be/entity"
)

type MoviesActorRepository interface {
	TrxSupportRepo
	CreateMoviesActor(MoviesActor *entity.MoviesActor) (result *entity.MoviesActor, err error)
	DeleteMoviesActor(MoviesActor *entity.MoviesActor) (err error)
	GetMoviesActorsByMovieID(ID int64) (MoviesActor *[]entity.MoviesActorSelect, err error)
}

type MoviesActorRepo struct {
	GormTrxSupport
}

func NewMoviesActorRepository(mysql *config.Mysql) *MoviesActorRepo {
	return &MoviesActorRepo{GormTrxSupport{db: mysql.DB}}
}

func (r *MoviesActorRepo) CreateMoviesActor(MoviesActor *entity.MoviesActor) (result *entity.MoviesActor, err error) {
	err = r.db.Where("actor_id = ? AND movie_id =?", MoviesActor.ActorID, MoviesActor.MovieID).Assign(MoviesActor).FirstOrCreate(&result).Error
	return result, err
}

func (r *MoviesActorRepo) DeleteMoviesActor(MoviesActor *entity.MoviesActor) (err error) {
	err = r.db.Where("id = ?", MoviesActor.ID).
		Delete(&MoviesActor).Error
	return err
}

func (r *MoviesActorRepo) GetMoviesActorsByMovieID(ID int64) (MoviesActor *[]entity.MoviesActorSelect, err error) {

	err = r.db.Select("movies_actor.*, actor.name, actor.gender").
		Joins("join actor on actor.id = movies_actor.actor_id").
		Where("movie_id = ?", ID).
		Find(&MoviesActor).
		Error

	return MoviesActor, err
}
