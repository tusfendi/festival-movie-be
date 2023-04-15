package repository

import (
	"github.com/tusfendi/festival-movie-be/config"
	"github.com/tusfendi/festival-movie-be/entity"
)

type ActorRepository interface {
	TrxSupportRepo
	CreateActor(actor *entity.Actor) (result *entity.Actor, err error)
	GetActors(keyword string) (result *[]entity.Actor, err error)
	GetActor(ID int) (result *entity.Actor, err error)
	UpdateActor(actor *entity.Actor) (err error)
	DeleteActor(ID int) (actor *entity.Actor, err error)
	UpdateOrInsert(actor *entity.Actor) (result entity.Actor, err error)
}

type ActorRepo struct {
	GormTrxSupport
}

func NewActorRepository(mysql *config.Mysql) *ActorRepo {
	return &ActorRepo{GormTrxSupport{db: mysql.DB}}
}

// create a Actor
func (r *ActorRepo) CreateActor(actor *entity.Actor) (result *entity.Actor, err error) {
	err = r.db.Create(actor).Error
	return actor, err
}

// get Actors
func (r *ActorRepo) GetActors(keyword string) (actor *[]entity.Actor, err error) {
	if keyword == "" {
		err = r.db.Find(&actor).Error
	} else {
		err = r.db.Where("name LIKE ? ", "%"+keyword+"%").Find(&actor).Error
	}
	return actor, err
}

// get Actor by id
func (r *ActorRepo) GetActor(ID int) (result *entity.Actor, err error) {
	err = r.db.Where("id = ?", ID).First(&result).Error
	return result, err
}

// update Actor
func (r *ActorRepo) UpdateActor(actor *entity.Actor) (err error) {
	r.db.Save(actor)
	return nil
}

func (r *ActorRepo) UpdateOrInsert(actor *entity.Actor) (result entity.Actor, err error) {
	err = r.db.Where("name = ?", actor.Name).Assign(actor).FirstOrCreate(&result).Error
	return result, err
}

// delete Actor
func (r *ActorRepo) DeleteActor(ID int) (actor *entity.Actor, err error) {
	err = r.db.Where("id = ?", ID).Delete(&actor).Error
	return actor, err
}
