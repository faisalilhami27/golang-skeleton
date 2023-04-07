package http

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-clean-code/config"
	controller "go-clean-code/controllers"
	repository "go-clean-code/domain/repositories"
	route "go-clean-code/routes"
	useCase2 "go-clean-code/services/usecases"
)

var Command = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		router := gin.Default()
		db := config.NewDBConnection()
		repository := repository.NewRepository(db)
		useCase := useCase2.NewUseCase(repository)
		controller := controller.NewController(useCase)
		route := route.NewRoute(*controller)
		route.Run(router)
		listen(router)
	},
}

func listen(router *gin.Engine) {
	err := router.Run()
	if err != nil {
		return
	}
}
