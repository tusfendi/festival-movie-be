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
		fmt.Println("The URL: ", c.Request.Host+c.Request.URL.Path)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"response": c.Request.Host})
	})

	// repository
	actorRepository := repository.NewActorRepository(mysqlConn)
	genreRepository := repository.NewGenreRepository(mysqlConn)

	// usecase
	actorUsecase := usecase.NewActorUsecase(actorRepository)
	genreUsecase := usecase.NewGenreUsecase(genreRepository)

	// Handler
	// actor
	r.GET("/artists", actorUsecase.Get)
	r.POST("/artists", actorUsecase.CreateActor)
	r.GET("/artists/:id", actorUsecase.GetDetail)
	r.PATCH("/artists/:id", actorUsecase.UpdateActor)
	r.DELETE("/artists/:id", actorUsecase.DeleteActor)

	// genre
	r.GET("/genres", genreUsecase.Get)
	r.POST("/genres", genreUsecase.CreateGenre)
	r.GET("/genres/:id", genreUsecase.GetDetail)
	r.PATCH("/genres/:id", genreUsecase.UpdateGenre)
	r.DELETE("/genres/:id", genreUsecase.DeleteGenre)

	r.Run(":" + fmt.Sprint(cfg.ApiPort))
}
