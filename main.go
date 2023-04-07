package main

import (
	"go-clean-code/cmd"
	"go-clean-code/config"
	util "go-clean-code/utils"
)

func init() {
	err := util.LoadEnv()
	if err != nil {
		panic("Error loading .env file")
	}
	config.InitApp()
}

func main() {
	cmd.Run()
}
