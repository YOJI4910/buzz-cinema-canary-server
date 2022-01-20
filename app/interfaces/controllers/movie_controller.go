package controllers

// controller: routeringと使用するusecaseを紐づける
import (
	"Yukit02/app/domain"
	"Yukit02/app/interfaces/database"
	"Yukit02/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// routerからのリクエストをInteractorに渡す
type MovieController struct {
	Interactor domain.MovieInteractor
}

// 使用するcontrollerを生成しているが、正直この書き方をする必要はなさそう
func NewMovieController(sqlHandler database.SqlHandler) *MovieController {
	return &MovieController{
		Interactor: &usecase.MovieInteractor{},
	}
}

// interface: MovieInteractorはTestメソッドを持っている
// 内容はmovie_interactorで実装される
func (controller *MovieController) Test(c *gin.Context) {
	res := controller.Interactor.Test()
	c.JSON(http.StatusOK, gin.H{"message": res})
}
