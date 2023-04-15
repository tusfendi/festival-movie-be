package usecase

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tusfendi/festival-movie-be/entity"
	"github.com/tusfendi/festival-movie-be/repository"
)

type GenreProvider struct {
	genreRepo repository.GenreRepository
}

func NewGenreUsecase(genreRepo repository.GenreRepository) *GenreProvider {
	return &GenreProvider{genreRepo}
}

func (u *GenreProvider) CreateGenre(c *gin.Context) {
	var params entity.Genre
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Genre tidak bisa kosong"})
		return
	}
	res, err := u.genreRepo.UpdateOrInsert(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (u *GenreProvider) UpdateGenre(c *gin.Context) {
	var params entity.Genre
	ID, _ := strconv.Atoi(c.Param("id"))
	params.ID = int64(ID)
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Genre tidak bisa kosong"})
		return
	}

	err = u.genreRepo.UpdateGenre(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": params})
}

func (u *GenreProvider) Get(c *gin.Context) {
	var params entity.Genre
	result, err := u.genreRepo.GetGenres(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *GenreProvider) GetDetail(c *gin.Context) {
	var params entity.Genre
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Genre tidak bisa kosong"})
		return
	}
	result, err := u.genreRepo.GetGenres(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *GenreProvider) DeleteGenre(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	err := u.genreRepo.DeleteGenre(&entity.Genre{}, int64(ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Data gagal dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "sukses", "message": "Genre berhasil dihapus"})
}
