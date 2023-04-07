package route

import (
	"github.com/gin-gonic/gin"
	controller "go-clean-code/controllers"
	route "go-clean-code/routes/user"
)

type router struct {
	controller controller.Controller
}

type Route interface {
	Run(ctx *gin.Engine)
}

func NewRoute(controller controller.Controller) Route {
	return &router{
		controller: controller,
	}
}

func (r *router) Run(router *gin.Engine) {
	group := router.Group("/api/v1")
	userRoute := route.NewUserRoute(r.controller)
	userRoute.Start(group)
}
