package route

import (
	"github.com/gin-gonic/gin"
	controller "go-clean-code/controllers"
)

type Route struct {
	route      *gin.Engine
	controller controller.Controller
}

type UserRoute interface {
	Start(route *gin.RouterGroup) *gin.Engine
}

func NewUserRoute(controller controller.Controller) UserRoute {
	return &Route{
		route:      gin.Default(),
		controller: controller,
	}
}

func (r *Route) Start(route *gin.RouterGroup) *gin.Engine {
	user := route.Group("/user")
	user.GET("/", r.controller.UserController.FindAll)
	user.GET("/:id", r.controller.UserController.FindByID)
	user.POST("/", r.controller.UserController.Save)
	user.PUT("/:id", r.controller.UserController.Update)
	user.DELETE("/:id", r.controller.UserController.Delete)
	return r.route
}
