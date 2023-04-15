package repository

import (
	"github.com/tusfendi/festival-movie-be/config"
	"github.com/tusfendi/festival-movie-be/entity"
)

type UserRepository interface {
	TrxSupportRepo
	CreateUser(user *entity.User) (result *entity.User, err error)
	GetUserByJTI(jti string) (result *entity.User, err error)
}

type UserRepo struct {
	GormTrxSupport
}

func NewUserRepository(mysql *config.Mysql) *UserRepo {
	return &UserRepo{GormTrxSupport{db: mysql.DB}}
}

// create a User
func (r *UserRepo) CreateUser(user *entity.User) (result *entity.User, err error) {
	err = r.db.Create(user).Error
	return user, err
}

func (r *UserRepo) GetUserByJTI(jti string) (result *entity.User, err error) {
	err = r.db.Select("user.*").
		Joins("join user_token on user_token.user_id = user.id").
		Where("user_token.guid = ?", jti).
		First(&result).Error
	return result, err
}
