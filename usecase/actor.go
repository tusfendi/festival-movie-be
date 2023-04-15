package usecase

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tusfendi/festival-movie-be/entity"
	"github.com/tusfendi/festival-movie-be/repository"
)

type ActorProvider struct {
	actorRepo repository.ActorRepository
}

func NewActorUsecase(actorRepo repository.ActorRepository) *ActorProvider {
	return &ActorProvider{actorRepo}
}

func (u *ActorProvider) CreateActor(c *gin.Context) {
	var params entity.Actor
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Data tidak lengkap"})
		return
	}

	if params.Gender != "L" && params.Gender != "P" {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Gender tidak ditermukan, gender harus salah satu dari L/P"})
		return
	}

	result, err := u.actorRepo.CreateActor(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *ActorProvider) UpdateActor(c *gin.Context) {
	var params entity.Actor
	ID, _ := strconv.Atoi(c.Param("id"))
	params.ID = int64(ID)
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Actor tidak bisa kosong"})
		return
	}

	if params.Gender != "L" && params.Gender != "P" {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Gender tidak ditermukan, gender harus salah satu dari L/P"})
		return
	}

	err = u.actorRepo.UpdateActor(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": params})
}

func (u *ActorProvider) Get(c *gin.Context) {
	var params entity.SearchByKeyword
	err := c.Bind(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}
	result, err := u.actorRepo.GetActors(params.Keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *ActorProvider) GetDetail(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	result, err := u.actorRepo.GetActor(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *ActorProvider) DeleteActor(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	_, err := u.actorRepo.DeleteActor(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Data gagal dihapus " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "sukses", "message": "data berhasil dihapus"})
}

// get actor role in movie
// update or insert actor rule in movie
