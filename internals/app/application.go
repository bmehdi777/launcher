package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Application struct {
	model  *Model
	fApp    *fyne.App
	fWindow *fyne.Window
	fCanvas *fyne.Canvas
}

func (a *Application) Run() error {
	app := Application {
		model: &Model{},
		fApp: &app.New(),
	}
	return nil
}
