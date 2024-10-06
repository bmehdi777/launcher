package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Test")
	myCanvas := myWindow.Canvas()

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)

	content := container.NewWithoutLayout(rect, widget.NewEntry())
	myCanvas.SetContent(content)

	myWindow.Resize(fyne.NewSize(1000, 200))
	myWindow.ShowAndRun()
}
