package usecase

import (
	repository "go-clean-code/domain/repositories"
	"go-clean-code/services/usecases/user"
)

type UseCase struct {
	UserUseCase usecase.UserUseCase
}

func NewUseCase(repository *repository.Repository) *UseCase {
	userUseCase := usecase.NewUserUseCase(repository.UserRepository)
	return &UseCase{
		UserUseCase: userUseCase,
	}
}
