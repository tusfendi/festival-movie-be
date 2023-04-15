package repository

import (
	"github.com/tusfendi/festival-movie-be/config"
	"github.com/tusfendi/festival-movie-be/entity"
)

type MovieRepository interface {
	TrxSupportRepo
	CreateMovie(movie *entity.Movie) (result *entity.Movie, err error)
	GetMovies(params entity.SearchMovies) (result *[]entity.Movie, err error)
	GetMovie(ID int64) (result *entity.MovieWithFilePath, err error)
	UpdateMovie(movie *entity.Movie) (err error)
	DeleteMovie(ID int64) (err error)
}

type MovieRepo struct {
	GormTrxSupport
}

func NewMovieRepository(mysql *config.Mysql) *MovieRepo {
	return &MovieRepo{GormTrxSupport{db: mysql.DB}}
}

// create a Movie
func (r *MovieRepo) CreateMovie(movie *entity.Movie) (result *entity.Movie, err error) {
	err = r.db.Create(movie).Error
	return movie, err
}

// get Movies
func (r *MovieRepo) GetMovies(params entity.SearchMovies) (movie *[]entity.Movie, err error) {
	// TBU
	err = r.db.Find(&movie).Error
	return movie, err
}

// get Movie by id
func (r *MovieRepo) GetMovie(ID int64) (result *entity.MovieWithFilePath, err error) {
	err = r.db.Select("movie.*, movies_file.file_path").
		Joins("left join movies_file on movies_file.id = movie.id").
		Where("movie.id = ?", ID).First(&result).Error
	return result, err
}

// update Movie
func (r *MovieRepo) UpdateMovie(movie *entity.Movie) (err error) {
	r.db.Save(movie)
	return nil
}

// delete Movie
func (r *MovieRepo) DeleteMovie(ID int64) (err error) {
	var movie *entity.Movie
	err = r.db.Where("id = ?", ID).Delete(&movie).Error
	return err
}
