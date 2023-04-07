package controller

import (
	controller "go-clean-code/controllers/user"
	useCase "go-clean-code/services/usecases"
)

type Controller struct {
	UserController controller.UserController
}

func NewController(useCase *useCase.UseCase) *Controller {
	userController := controller.NewUserController(useCase.UserUseCase)
	return &Controller{
		UserController: userController,
	}
}
