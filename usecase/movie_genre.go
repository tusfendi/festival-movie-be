package usecase

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tusfendi/festival-movie-be/entity"
	"github.com/tusfendi/festival-movie-be/repository"
)

type MoviesGenreProvide struct {
	movieGenreRepo repository.MoviesGenreRepository
}

func NewMoviesGenreUsecase(movieGenreRepo repository.MoviesGenreRepository) *MoviesGenreProvide {
	return &MoviesGenreProvide{movieGenreRepo}
}

func (u *MoviesGenreProvide) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	err := u.movieGenreRepo.DeleteMoviesGenre(&entity.MoviesGenre{
		ID: int64(ID),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Data gagal dihapus ", "error_detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "sukses", "message": "data berhasil dihapus"})
}

func (u *MoviesGenreProvide) CreateMoviesGenre(c *gin.Context) {
	var params entity.MoviesGenre
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Data tidak lengkap"})
		return
	}

	result, err := u.movieGenreRepo.CreateMoviesGenre(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan ", "error_detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *MoviesGenreProvide) GetMoviesGenresByMovieID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	result, err := u.movieGenreRepo.GetMoviesGenresByMovieID(int64(ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Terjadi kesalahan" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}
