package usecase

type MovieInteractor struct {
	// repository MovieRepository
}

func (interactor *MovieInteractor) Test() string {
	return "this is test, in usecase"
}
