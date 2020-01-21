package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type InfantryTabs struct {
	ii *InfantryVInfantry
	ia *InfantryVInfantry
	is *InfantryVInfantry
	ic *InfantryVInfantry
	it *InfantryVInfantry
}

func NewInfantryScreens() *InfantryTabs {
	return &InfantryTabs{
		ii: NewInfantryVInfantry(),
		ia: NewInfantryVInfantry(),
		is: NewInfantryVInfantry(),
		ic: NewInfantryVInfantry(),
		it: NewInfantryVInfantry(),
	}
}

func (i *InfantryTabs) Canvas() fyne.CanvasObject {
	return widget.NewTabContainer(
		widget.NewTabItem("Infantry vs Infantry", i.ii.Canvas()),
		widget.NewTabItem("Infantry vs Artillery", i.ia.Canvas()),
		widget.NewTabItem("Infantry vs Skirmishers", i.is.Canvas()),
		widget.NewTabItem("Infantry vs Cavalry", i.ic.Canvas()),
		widget.NewTabItem("Infantry Town Assault", i.it.Canvas()),
	)
}
