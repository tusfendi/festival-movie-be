package repository

import (
	"github.com/tusfendi/festival-movie-be/config"
	"github.com/tusfendi/festival-movie-be/entity"
)

type GenreRepository interface {
	TrxSupportRepo
	CreateGenre(genre *entity.Genre) (err error)
	GetGenres(genre *entity.Genre) (result *[]entity.Genre, err error)
	GetGenre(ID int64) (result *entity.Genre, err error)
	UpdateGenre(genre *entity.Genre) (err error)
	UpdateOrInsert(genre *entity.Genre) (result entity.Genre, err error)
	DeleteGenre(genre *entity.Genre, ID int64) (err error)
}

type GenreRepo struct {
	GormTrxSupport
}

func NewGenreRepository(mysql *config.Mysql) *GenreRepo {
	return &GenreRepo{GormTrxSupport{db: mysql.DB}}
}

// create a Genre
func (r *GenreRepo) CreateGenre(genre *entity.Genre) (err error) {
	err = r.db.Create(genre).Error
	if err != nil {
		return err
	}
	return nil
}

// get Genres
func (r *GenreRepo) GetGenres(genre *entity.Genre) (result *[]entity.Genre, err error) {
	if genre.Genre != "" {
		err = r.db.Where("genre = ?", genre.Genre).Find(&result).Error
	} else {
		err = r.db.Find(&result).Error
	}
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *GenreRepo) GetGenre(ID int64) (result *entity.Genre, err error) {
	err = r.db.Where("id = ?", ID).First(&result).Error
	return result, err
}

// update Genre
func (r *GenreRepo) UpdateGenre(genre *entity.Genre) (err error) {
	r.db.Save(genre)
	return nil
}

func (r *GenreRepo) UpdateOrInsert(genre *entity.Genre) (result entity.Genre, err error) {
	err = r.db.Where("genre = ?", genre.Genre).Assign(genre).FirstOrCreate(&result).Error
	return result, err
}

// delete Genre
func (r *GenreRepo) DeleteGenre(genre *entity.Genre, ID int64) (err error) {
	err = r.db.Where("id = ?", ID).Delete(genre).Error
	return err
}
