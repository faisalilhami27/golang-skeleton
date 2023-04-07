package usecase

import (
	entity "go-clean-code/domain/entities"
	model "go-clean-code/domain/models"
	repository "go-clean-code/domain/repositories/user"
	util "go-clean-code/utils"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	UserRepository repository.UserRepository
}

type UserUseCase interface {
	FindAll() ([]entity.UserResponse, error)
	FindByID(int) (*entity.UserResponse, error)
	Save(*entity.UserRequest) (*entity.UserResponse, error)
	Update(*entity.UserRequest, int) (*entity.UserResponse, error)
	Delete(int) error
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &service{UserRepository: userRepository}
}

func (s *service) FindAll() ([]entity.UserResponse, error) {
	var userResponse []entity.UserResponse
	users, err := s.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		userResponse = append(userResponse, entity.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		})
	}
	return userResponse, nil
}

func (s *service) FindByID(id int) (*entity.UserResponse, error) {
	user, err := s.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	result := entity.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	return &result, nil
}

func (s *service) Save(request *entity.UserRequest) (*entity.UserResponse, error) {
	indonesianTimezone := util.ConvertToIndonesiaTimezone()
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user := model.User{
		Name:      request.Name,
		Username:  request.Username,
		Password:  string(passwordHash),
		CreatedAt: indonesianTimezone,
		UpdatedAt: indonesianTimezone,
	}
	newUser, err := s.UserRepository.Save(&user)
	if err != nil {
		return nil, err
	}

	result := entity.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Username:  newUser.Username,
		Password:  newUser.Password,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		DeletedAt: newUser.DeletedAt,
	}
	return &result, nil
}

func (s *service) Update(request *entity.UserRequest, id int) (*entity.UserResponse, error) {
	indonesianTimezone := util.ConvertToIndonesiaTimezone()
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user := model.User{
		Name:      request.Name,
		Username:  request.Username,
		Password:  string(passwordHash),
		UpdatedAt: indonesianTimezone,
	}
	newUser, err := s.UserRepository.Update(&user, id)
	if err != nil {
		return nil, err
	}

	result := entity.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Username:  newUser.Username,
		Password:  newUser.Password,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		DeletedAt: newUser.DeletedAt,
	}
	return &result, nil
}

func (s *service) Delete(id int) error {
	err := s.UserRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
