package repository

import (
	repository "go-clean-code/domain/repositories/user"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository repository.UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := repository.NewUserRepository(db)
	return &Repository{
		UserRepository: userRepository,
	}
}
