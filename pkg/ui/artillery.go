package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func Artillery() fyne.CanvasObject {

	return widget.NewVBox(
		widget.NewGroup("Artillery Calculator"),
	)

}
