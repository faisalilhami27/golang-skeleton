package http

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		//_ = config.NewDBConnection()
		router := gin.Default()
		listen(router)
	},
}

func listen(router *gin.Engine) {
	err := router.Run()
	if err != nil {
		return
	}
}
