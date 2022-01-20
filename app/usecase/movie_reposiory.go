package usecase

import "Yukit02/app/domain"

type MovieRepository interface {
	Select() []domain.Movie
	Delete()
}
