package main

import (
	"github.com/ZiplEix/API_template/app"
)

// @title API Template
// @version 0.1
// @description This is a sample API template.
// @contact.name ZiplEix
// host localhost:3000
// @BasePath /
func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
