package infrastructure

import (
	"Yukit02/app/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	movieController := controllers.NewMovieController(NewSqlHandler())

	router.GET("/movie", func(c *gin.Context) { movieController.Test(c) })

	Router = router
}
