package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func Skirmish() fyne.CanvasObject {

	return widget.NewVBox(
		widget.NewGroup("Skirmish Attack"),
	)

}
