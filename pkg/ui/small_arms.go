package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/steveoc64/re/pkg/re"
)

func SmallArms(unitA, unitB *re.SmallArmsUnit) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewVBox(
			widget.NewGroup("Attacking Unit",
				widget.NewForm(
					widget.NewFormItem("Class", widget.NewSelect(nil, nil).
						Source(unitA.ClassList).
						Bind(unitA.Class)),
					widget.NewFormItem("Current Hits", widget.NewEntry().
						Bind(unitA.Hits)),
					widget.NewFormItem("Status", widget.NewLabel("").
						Bind(unitA.BasesDesc)),
					widget.NewFormItem("Close Order Bases", widget.NewSlider(1.0, 6.0).
						SetOnChanged(unitA.CalcFF).
						Bind(unitA.CloseOrderBases)),
					widget.NewFormItem("Firing Bases", widget.NewSlider(1.0, 6.0).
						SetOnChanged(unitA.CalcFF).
						Bind(unitA.FiringBases)),
					widget.NewFormItem("Supporting Bases", widget.NewSlider(0.0, 6.0).
						SetOnChanged(unitA.CalcFF).
						Bind(unitA.SupportingBases)),
					widget.NewFormItem("Fire Factor", widget.NewLabel("").
						Bind(unitA.FireFactor)),
				),
			),
		),
		widget.NewVBox(
			widget.NewGroup("Defending Unit",
				widget.NewForm(
					widget.NewFormItem("Class", widget.NewSelect(nil, nil).
						Source(unitB.ClassList).
						Bind(unitB.Class)),
					widget.NewFormItem("Current Hits", widget.NewEntry().
						Bind(unitB.Hits)),
					widget.NewFormItem("Status", widget.NewLabel("").
						Bind(unitB.BasesDesc)),
					widget.NewFormItem("Close Order Bases", widget.NewSlider(1.0, 6.0).
						SetOnChanged(unitB.CalcFF).
						Bind(unitB.CloseOrderBases)),
					widget.NewFormItem("Firing Bases", widget.NewSlider(1.0, 6.0).
						SetOnChanged(unitB.CalcFF).
						Bind(unitB.FiringBases)),
					widget.NewFormItem("Supporting Bases", widget.NewSlider(0.0, 6.0).
						SetOnChanged(unitB.CalcFF).
						Bind(unitB.SupportingBases)),
					widget.NewFormItem("Fire Factor", widget.NewLabel("").
						Bind(unitB.FireFactor)),
				),
			),
		),
	)

}
