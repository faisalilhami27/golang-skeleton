package repository

import (
	model "go-clean-code/domain/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(int) (*model.User, error)
	Save(*model.User) (*model.User, error)
	Update(*model.User, int) (*model.User, error)
	Delete(int) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{db: db}
}

func (u *repository) FindAll() ([]model.User, error) {
	var users []model.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *repository) FindByID(id int) (*model.User, error) {
	var user model.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *repository) Save(user *model.User) (*model.User, error) {
	newUser := model.User{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
	}
	err := u.db.Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (u *repository) Update(user *model.User, id int) (*model.User, error) {
	newUser := model.User{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
	}

	err := u.db.Where("id = ?", id).Updates(&newUser).Error
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (u *repository) Delete(id int) error {
	err := u.db.Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
