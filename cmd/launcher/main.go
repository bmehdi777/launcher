package main

import (
	"fmt"
	"os"

	"github.com/bmehdi777/launcher/internals/app"
)

func main() {
	appli := app.NewApp()
	if err := appli.Run(); err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
}
