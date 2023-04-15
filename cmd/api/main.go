package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"github.com/tusfendi/festival-movie-be/cmd/api/middleware"
	"github.com/tusfendi/festival-movie-be/config"
	"github.com/tusfendi/festival-movie-be/repository"
	"github.com/tusfendi/festival-movie-be/usecase"
)

func init() {
	_ = gotenv.Load()
}

func main() {
	cfg := config.NewConfig()
	mysqlConn, err := config.NewMysql(cfg.AppEnv, &cfg.MysqlOption)
	if err != nil {
		log.Fatal(err)
		println("error mysql")
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(middleware.CORSMiddleware())

	r.GET("/foo", func(c *gin.Context) {
		// strings.Split(c.Request.URL.Path, "/")
		fmt.Println("The URL: ", c.Request.Host+c.Request.URL.Path)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"response": c.Request.Host})
	})

	// repository
	actorRepository := repository.NewActorRepository(mysqlConn)
	genreRepository := repository.NewGenreRepository(mysqlConn)
	userHitoryRepository := repository.NewUserHistoryRepository(mysqlConn)
	movieRepository := repository.NewMovieRepository(mysqlConn)
	movieGenreRepository := repository.NewMoviesGenreRepository(mysqlConn)
	movieActorRepository := repository.NewMoviesActorRepository(mysqlConn)

	// usecase
	actorUsecase := usecase.NewActorUsecase(actorRepository)
	genreUsecase := usecase.NewGenreUsecase(genreRepository)
	movieGenreUsecase := usecase.NewMoviesGenreUsecase(movieGenreRepository)
	movieUsecase := usecase.NewMovieUsecase(movieRepository, movieGenreRepository, movieActorRepository, userHitoryRepository)

	// TBU untuk middleware auth user

	// admin use : /_internal
	admin := r.Group("/_internal")
	// add middleware to check role is admin

	// Handler
	// actor
	admin.GET("/artists", actorUsecase.Get)
	admin.POST("/artists", actorUsecase.CreateActor)
	admin.GET("/artists/:id", actorUsecase.GetDetail)
	admin.PATCH("/artists/:id", actorUsecase.UpdateActor)
	admin.DELETE("/artists/:id", actorUsecase.DeleteActor)

	// genre
	admin.GET("/genres", genreUsecase.Get)
	admin.POST("/genres", genreUsecase.CreateGenre)
	admin.GET("/genres/:id", genreUsecase.GetDetail)
	admin.PATCH("/genres/:id", genreUsecase.UpdateGenre)
	admin.DELETE("/genres/:id", genreUsecase.DeleteGenre)

	r.GET("/movies", movieUsecase.Get)
	r.GET("/movies/:id", movieUsecase.GetDetail)
	admin.POST("/movies", movieUsecase.CreateMovie)
	admin.PATCH("/movies/:id", movieUsecase.UpdateMovie)
	admin.DELETE("/movies/:id", movieUsecase.DeleteMovie)

	admin.POST("/movies-genre", movieGenreUsecase.CreateMoviesGenre)
	admin.DELETE("/movies-genre/:id", movieGenreUsecase.Delete)
	admin.GET("/movies-genre/:id", movieGenreUsecase.GetMoviesGenresByMovieID)

	r.Run(":" + fmt.Sprint(cfg.ApiPort))
}
