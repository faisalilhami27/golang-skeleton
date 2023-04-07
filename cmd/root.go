package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-clean-code/cmd/http"
	"log"
)

var rootCommand = &cobra.Command{
	Use:   "golang-clean-code",
	Short: "Golang clean code architecture",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use serve to start a server")
		fmt.Println("Use -h to see the list of command")
	},
}

func Run() {
	rootCommand.AddCommand(http.Command)
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
