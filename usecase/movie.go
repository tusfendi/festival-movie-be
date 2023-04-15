package usecase

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tusfendi/festival-movie-be/entity"
	"github.com/tusfendi/festival-movie-be/repository"
)

type MovieProvider struct {
	movieRepo      repository.MovieRepository
	genreRepo      repository.MoviesGenreRepository
	actorRepo      repository.MoviesActorRepository
	userHitoryRepo repository.UserHistoryRepository
}

func NewMovieUsecase(movieRepo repository.MovieRepository,
	genreRepo repository.MoviesGenreRepository,
	actorRepo repository.MoviesActorRepository,
	userHitoryRepo repository.UserHistoryRepository) *MovieProvider {
	return &MovieProvider{movieRepo, genreRepo, actorRepo, userHitoryRepo}
}

func (u *MovieProvider) CreateMovie(c *gin.Context) {
	var params entity.Movie
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Data tidak lengkap"})
		return
	}

	result, httpCode, err := u.ProcessCreate(params)
	if err != nil {
		c.JSON(httpCode, gin.H{"response": "gagal", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *MovieProvider) ProcessCreate(params entity.Movie) (result *entity.Movie, httpCode int, err error) {
	// Check movie file ID
	// if params.MovieFileID > 0 {
	// 	_, err = u.movieRepo.GetMovie(params.MovieFileID)
	// 	if err != nil {
	// 		return &params, http.StatusBadRequest, fmt.Errorf("File movie tidak ditemukan")
	// 	}
	// }
	result, err = u.movieRepo.CreateMovie(&params)
	if err != nil {
		return result, http.StatusInternalServerError, fmt.Errorf("Terjadi kesalahan")
	}
	return result, http.StatusOK, nil
}

func (u *MovieProvider) UpdateMovie(c *gin.Context) {
	var params entity.Movie
	ID, _ := strconv.Atoi(c.Param("id"))
	params.ID = int64(ID)
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Movie tidak bisa kosong"})
		return
	}

	if params.MovieFileID > 0 {
		_, err = u.movieRepo.GetMovie(params.MovieFileID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Sequel tidak ditemukan"})
			return
		}
	}
	err = u.movieRepo.UpdateMovie(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": params})
}

func (u *MovieProvider) Get(c *gin.Context) {
	var params entity.SearchMovies
	err := c.Bind(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}

	// update filter nya banyak
	result, err := u.movieRepo.GetMovies(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Terjadi kesalahan " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *MovieProvider) GetDetail(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	movie, err := u.movieRepo.GetMovie(int64(ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Terjadi kesalahan [1]" + err.Error()})
		return
	}

	genre, err := u.genreRepo.GetMoviesGenresByMovieID(int64(ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Terjadi kesalahan [2]" + err.Error()})
		return
	}

	actor, err := u.actorRepo.GetMoviesActorsByMovieID(int64(ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "gagal", "error": "Terjadi kesalahan [2]" + err.Error()})
		return
	}

	for _, v := range *genre {
		u.userHitoryRepo.CreateUserHistory(&entity.UserHistory{
			UserID:  c.GetString("user_id"),
			MovieID: movie.ID,
			GenreID: v.ID,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": entity.MovieDetail{
		MovieWithFilePath: *movie,
		Genres:            *genre,
		Actors:            *actor,
	}})
}

func (u *MovieProvider) DeleteMovie(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	err := u.movieRepo.DeleteMovie(int64(ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "gagal", "error": "Data gagal dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "sukses", "message": "data berhasil dihapus"})
}
