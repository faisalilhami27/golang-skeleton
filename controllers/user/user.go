package controller

import (
	"github.com/gin-gonic/gin"
	"go-clean-code/constant"
	entity "go-clean-code/domain/entities"
	usecase "go-clean-code/services/usecases/user"
	util "go-clean-code/utils"
	"net/http"
	"strconv"
)

type controller struct {
	useCase usecase.UserUseCase
}

type UserController interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	Save(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewUserController(useCase usecase.UserUseCase) UserController {
	return &controller{useCase: useCase}
}

func (c *controller) FindAll(ctx *gin.Context) {
	users, err := c.useCase.FindAll()
	if err != nil {
		error := util.ResponseError(constant.Error(), err.Error())
		ctx.JSON(http.StatusBadRequest, error)
		return
	}
	ctx.JSON(http.StatusOK, util.ResponseSuccess(constant.Success(), users))
}

func (c *controller) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := c.useCase.FindByID(id)
	if err != nil {
		error := util.ResponseError(constant.Error(), err.Error())
		ctx.JSON(http.StatusBadRequest, error)
		return
	}
	ctx.JSON(http.StatusOK, util.ResponseSuccess(constant.Success(), user))
}

func (c *controller) Save(ctx *gin.Context) {
	var data entity.UserRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		error := util.ResponseError(constant.Error(), err.Error())
		ctx.JSON(http.StatusBadRequest, error)
		return
	}

	user, err := c.useCase.Save(&data)
	if err != nil {
		error := util.ResponseError(constant.Error(), err.Error())
		ctx.JSON(http.StatusBadRequest, error)
		return
	}
	ctx.JSON(http.StatusOK, util.ResponseSuccess(constant.Success(), user))
}

func (c *controller) Update(ctx *gin.Context) {
	var data entity.UserRequest
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&data); err != nil {
		error := util.ResponseError(constant.Error(), err.Error())
		ctx.JSON(http.StatusBadRequest, error)
		return
	}

	user, err := c.useCase.Update(&data, id)
	if err != nil {
		error := util.ResponseError(constant.Error(), err.Error())
		ctx.JSON(http.StatusBadRequest, error)
		return
	}
	ctx.JSON(http.StatusOK, util.ResponseSuccess(constant.Success(), user))
}

func (c *controller) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.useCase.Delete(id)
	if err != nil {
		error := util.ResponseError(constant.Error(), err.Error())
		ctx.JSON(http.StatusBadRequest, error)
		return
	}
	ctx.JSON(http.StatusOK, util.ResponseSuccess(constant.Success(), nil))
}
