package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/steveoc64/re/pkg/re"
)

func Artillery(sit *re.ContactSituation) fyne.CanvasObject {
	unitA := sit.Units[0]
	unitB := sit.Units[1]

	return widget.NewVBox(
		widget.NewGroup("Situation",
			widget.NewForm(
				widget.NewFormItem("Range", widget.NewSlider(0.0, 18.0).
					SetOnChanged(func(f float64) {
						sit.Changed("")
						//unitA.CalcFF(0.0)
						//unitB.CalcFF(0.0)
					}).
					Bind(sit.Range)),
				widget.NewFormItem("Weather", widget.NewSelect(nil, sit.Changed).
					Source(re.Weathers).
					Bind(sit.Weather)),
				widget.NewFormItem("", widget.NewCheck("Return Fire", nil).Bind(sit.ReturnFire)),
				widget.NewFormItem("", widget.NewCheck("Enfilade", nil).Bind(sit.Enfilade)),
				widget.NewFormItem("", widget.NewButtonWithIcon("Clear",
					theme.MailReplyIcon(),
					func() {
						sit.Clear()
					},
				)),
			),
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			widget.NewVBox(
				widget.NewGroup("Attacking Unit",
					widget.NewForm(
						widget.NewFormItem("Class", widget.NewSelect(nil, unitA.Changed).
							Source(re.UnitClasses).
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
						widget.NewFormItem("Formation", widget.NewSelect(nil, unitA.FormationChanged).
							Source(re.Formations).
							Bind(unitA.Formation)),
						widget.NewFormItem("Ammo State", widget.NewSelect(nil, unitA.Changed).
							Source(re.AmmoStates).
							Bind(unitA.AmmoState)),
						widget.NewFormItem("Terrain", widget.NewSelect(nil, unitA.Changed).
							Source(re.Terrains).
							Bind(unitA.Terrain)),
						widget.NewFormItem("Fire Factor", widget.NewLabel("").
							Bind(unitA.FireFactor)),
						widget.NewFormItem("Dice", widget.NewLabel("").
							Bind(unitA.DieModDesc)),
						widget.NewFormItem("Result", widget.NewLabel("").
							Bind(unitA.FireResults)),
						widget.NewFormItem("", widget.NewButtonWithIcon("Fire",
							theme.MailReplyIcon(),
							func() {
								unitA.Fire()
							},
						)),
					),
				),
			),
			widget.NewVBox(
				widget.NewGroup("Defending Unit",
					widget.NewForm(
						widget.NewFormItem("Class", widget.NewSelect(nil, unitB.Changed).
							Source(re.UnitClasses).
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
						widget.NewFormItem("Formation", widget.NewSelect(nil, unitB.FormationChanged).
							Source(re.Formations).
							Bind(unitB.Formation)),
						widget.NewFormItem("Ammo State", widget.NewSelect(nil, unitB.Changed).
							Source(re.AmmoStates).
							Bind(unitB.AmmoState)),
						widget.NewFormItem("Terrain", widget.NewSelect(nil, unitB.Changed).
							Source(re.Terrains).
							Bind(unitB.Terrain)),
						widget.NewFormItem("Fire Factor", widget.NewLabel("").
							Bind(unitB.FireFactor)),
						widget.NewFormItem("Dice", widget.NewLabel("").
							Bind(unitB.DieModDesc)),
						widget.NewFormItem("Result", widget.NewLabel("").
							Bind(unitB.FireResults)),
						widget.NewFormItem("", widget.NewButtonWithIcon("Fire",
							theme.MailReplyIcon(),
							func() {
								unitB.Fire()
							},
						)),
					),
				),
			),
		),
	)

}

