package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"github.com/tusfendi/festival-movie-be/cmd/api/middleware"
	"github.com/tusfendi/festival-movie-be/config"
)

func init() {
	_ = gotenv.Load()
}

func main() {
	cfg := config.NewConfig()
	_, err := config.NewMysql(cfg.AppEnv, &cfg.MysqlOption)
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

	r.Run(":" + fmt.Sprint(cfg.ApiPort))
}
