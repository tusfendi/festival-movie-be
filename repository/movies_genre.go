package repository

import (
	"github.com/tusfendi/festival-movie-be/config"
	"github.com/tusfendi/festival-movie-be/entity"
)

type MoviesGenreRepository interface {
	TrxSupportRepo
	CreateMoviesGenre(MoviesGenre *entity.MoviesGenre) (result *entity.MoviesGenre, err error)
	DeleteMoviesGenre(MoviesGenre *entity.MoviesGenre) (err error)
	GetMoviesGenresByMovieID(ID int64) (MoviesGenre *[]entity.MoviesGenreSelect, err error)
}

type MoviesGenreRepo struct {
	GormTrxSupport
}

func NewMoviesGenreRepository(mysql *config.Mysql) *MoviesGenreRepo {
	return &MoviesGenreRepo{GormTrxSupport{db: mysql.DB}}
}

func (r *MoviesGenreRepo) CreateMoviesGenre(MoviesGenre *entity.MoviesGenre) (result *entity.MoviesGenre, err error) {
	err = r.db.Where("genre_id = ? AND movie_id =?", MoviesGenre.GenreID, MoviesGenre.MovieID).Assign(MoviesGenre).FirstOrCreate(&result).Error
	return result, err
}

func (r *MoviesGenreRepo) DeleteMoviesGenre(MoviesGenre *entity.MoviesGenre) (err error) {
	err = r.db.Where("id = ?", MoviesGenre.ID).
		Delete(&MoviesGenre).Error
	return err
}

func (r *MoviesGenreRepo) GetMoviesGenresByMovieID(ID int64) (MoviesGenre *[]entity.MoviesGenreSelect, err error) {

	err = r.db.Select("movies_genre.*, genre.genre").
		Joins("join genre on genre.id = movies_genre.genre_id").
		Where("movie_id = ?", ID).
		Find(&MoviesGenre).
		Error

	return MoviesGenre, err
}
