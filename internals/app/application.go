package app

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Application struct {
	state *State
	fApp  *fyne.App
	view  *ApplicationView
}

type ApplicationView struct {
	fWindow *fyne.Window
	fCanvas *fyne.Canvas

	textInput *widget.Entry
}

func (av *ApplicationView) initializeView() {
	av.textInput = widget.NewEntry()

	(*av.fWindow).SetContent(av.textInput)
}

func (a *Application) initializeApp() {
	(*a.view.fWindow).SetFixedSize(true)
	(*a.view.fWindow).Resize(fyne.NewSize(1000, 100))

	(*a.view.fCanvas).SetOnTypedKey(a.onKey)

	a.view.initializeView()
}

func (a *Application) onKey(key *fyne.KeyEvent) {
	log.Println("Key : ", key.Name)
	switch key.Name {
	case "Escape":
		(*a.view.fWindow).Close()
	}
}

func NewApp() Application {
	fApp := app.New()
	fWindow := fApp.NewWindow("Launcher")
	fCanvas := fWindow.Canvas()

	app := Application{
		state: DefaultState(),
		fApp:  &fApp,
		view: &ApplicationView{
			fWindow: &fWindow,
			fCanvas: &fCanvas,
		},
	}

	return app
}

func (a *Application) Run() error {
	a.initializeApp()

	(*a.view.fWindow).ShowAndRun()

	return nil
}
