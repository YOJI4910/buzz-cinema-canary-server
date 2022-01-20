// 実処理を書く場所
// DBとのやり取りを書く
package database

import "Yukit02/app/domain"

// movie repositoryはSqlHandlerを持つことを宣言
type MovieRepository struct {
	SqlHandler
}

func (repo *MovieRepository) Select() []domain.Movie {
	movie := []domain.Movie{}
	repo.FindAll(&movie)
	return movie
}

func (repo *MovieRepository) Delete(id string) {
	movie := []domain.Movie{}
	repo.DeleteById(&movie, id)
}
