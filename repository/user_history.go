package repository

import (
	"github.com/tusfendi/festival-movie-be/config"
	"github.com/tusfendi/festival-movie-be/entity"
)

type UserHistoryRepository interface {
	TrxSupportRepo
	CreateUserHistory(userHistory *entity.UserHistory) (result *entity.UserHistory, err error)
	// GetUserHistoriesByMovieID(MovieID int64) (result *[]entity.UserHistory, err error)
	// GetUserHistoriesByGenreID(GenreID int64) (result *[]entity.UserHistory, err error)
}

type UserHistoryRepo struct {
	GormTrxSupport
}

func NewUserHistoryRepository(mysql *config.Mysql) *UserHistoryRepo {
	return &UserHistoryRepo{GormTrxSupport{db: mysql.DB}}
}

// create a UserHistory
func (r *UserHistoryRepo) CreateUserHistory(userHistory *entity.UserHistory) (result *entity.UserHistory, err error) {
	err = r.db.Create(userHistory).Error
	return userHistory, err
}
