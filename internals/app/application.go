package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Application struct {
	model   *Model
	fApp    *fyne.App
	fWindow *fyne.Window
	fCanvas *fyne.Canvas
}

func NewApp() Application {
	fApp := app.New()
	fWindow := fApp.NewWindow("Launcher")

	app := Application{
		model: &Model{},
		fApp:  &fApp,
		fWindow: &fWindow,
	}

	return app
}

func (a *Application) Run() error {
	(*a.fWindow).ShowAndRun()
	return nil
}
