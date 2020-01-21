package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/steveoc64/re/pkg/re"
)


type InfantryVInfantry struct {
	sit *re.ContactSituation
	saved []*re.ContactSituation
}

func NewInfantryVInfantry() *InfantryVInfantry {
	return &InfantryVInfantry{
		sit: re.NewSmallArmsSituation(),
	}
}

func (i *InfantryVInfantry) Canvas() fyne.CanvasObject {
	return widget.NewVBox(
		widget.NewGroup("Infantry v Infantry",
			widget.NewForm(
				widget.NewFormItem("Range", widget.NewSlider(0.0, 18.0).
					SetOnChanged(func(f float64) {
						i.sit.Changed("")
						//unitA.CalcFF(0.0)
						//unitB.CalcFF(0.0)
					}).
					Bind(i.sit.Range)),
				widget.NewFormItem("Weather", widget.NewSelect(nil, i.sit.Changed).
					Source(re.Weathers).
					Bind(i.sit.Weather)),
				widget.NewFormItem("", widget.NewCheck("Return Fire", nil).Bind(i.sit.ReturnFire)),
				widget.NewFormItem("", widget.NewCheck("Auto Apply Dice Rolls", nil).Bind(i.sit.AutoDice)),
				widget.NewFormItem("", widget.NewLabel("").Bind(i.sit.Status)),
				widget.NewFormItem("", widget.NewButtonWithIcon("Reset",
					theme.MailReplyIcon(),
					func() {
						i.sit.Clear()
					},
				)),
			),
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			widget.NewVBox(
				widget.NewGroup("Attacking Infantry",
					widget.NewForm(
						widget.NewFormItem("Class", widget.NewSelect(nil, i.sit.Attacker.Changed).
							Source(re.UnitClasses).
							Bind(i.sit.Attacker.Class)),
						widget.NewFormItem("Morale", widget.NewSelect(nil, i.sit.Attacker.Changed).
							Source(re.MoraleStates).
							Bind(i.sit.Attacker.MoraleState)),
						widget.NewFormItem("Current Hits", widget.NewEntry().
							Bind(i.sit.Attacker.Hits)),
						widget.NewFormItem("Status", widget.NewLabel("").
							Bind(i.sit.Attacker.BasesDesc)),
						widget.NewFormItem("Close Order Bases", widget.NewSlider(1.0, 6.0).
							Bind(i.sit.Attacker.CloseOrderBases)),
						widget.NewFormItem("Firing Bases", widget.NewSlider(0.0, 6.0).
							Bind(i.sit.Attacker.FiringBases)),
						widget.NewFormItem("Supporting Bases", widget.NewSlider(0.0, 6.0).
							Bind(i.sit.Attacker.SupportingBases)),
						widget.NewFormItem("Formation", widget.NewSelect(nil, i.sit.Attacker.FormationChanged).
							Source(re.Formations).
							Bind(i.sit.Attacker.Formation)),
						widget.NewFormItem("Ammo State", widget.NewSelect(nil, i.sit.Attacker.Changed).
							Source(re.AmmoStates).
							Bind(i.sit.Attacker.AmmoState)),
						widget.NewFormItem("Terrain", widget.NewSelect(nil, i.sit.Attacker.Changed).
							Source(re.Terrains).
							Bind(i.sit.Attacker.Terrain)),
						widget.NewFormItem("", widget.NewCheck("Enfilade", nil).
							Bind(i.sit.Attacker.Enfilade)),
						widget.NewFormItem("Fire Factor", widget.NewLabel("").
							Bind(i.sit.Attacker.FireFactor)),
						widget.NewFormItem("Mods", widget.NewLabel("").
							Bind(i.sit.Attacker.DieModDesc)),
						widget.NewFormItem("Rolls", widget.NewLabel("").
							Bind(i.sit.Attacker.FireRolls)),
						widget.NewFormItem("Result", widget.NewLabel("").
							Bind(i.sit.Attacker.FireResults)),
						widget.NewFormItem("Damage", widget.NewLabel("").
							Bind(i.sit.Attacker.FireTotal)),
						widget.NewFormItem("Morale Check", widget.NewLabel("").
							Bind(i.sit.Attacker.MoraleCheckResult)),
						widget.NewFormItem("", widget.NewButtonWithIcon("Fire",
							theme.MailReplyIcon(),
							func() {
								i.sit.Attacker.Fire()
							},
						)),
					),
				),
			),
			widget.NewVBox(
				widget.NewGroup("Defending Infantry",
					widget.NewForm(
						widget.NewFormItem("Class", widget.NewSelect(nil, i.sit.Defender.Changed).
							Source(re.UnitClasses).
							Bind(i.sit.Defender.Class)),
						widget.NewFormItem("Morale", widget.NewSelect(nil, i.sit.Defender.Changed).
							Source(re.MoraleStates).
							Bind(i.sit.Defender.MoraleState)),
						widget.NewFormItem("Current Hits", widget.NewEntry().
							Bind(i.sit.Defender.Hits)),
						widget.NewFormItem("Status", widget.NewLabel("").
							Bind(i.sit.Defender.BasesDesc)),
						widget.NewFormItem("Close Order Bases", widget.NewSlider(1.0, 6.0).
							Bind(i.sit.Defender.CloseOrderBases)),
						widget.NewFormItem("Firing Bases", widget.NewSlider(0.0, 6.0).
							Bind(i.sit.Defender.FiringBases)),
						widget.NewFormItem("Supporting Bases", widget.NewSlider(0.0, 6.0).
							Bind(i.sit.Defender.SupportingBases)),
						widget.NewFormItem("Formation", widget.NewSelect(nil, i.sit.Defender.FormationChanged).
							Source(re.Formations).
							Bind(i.sit.Defender.Formation)),
						widget.NewFormItem("Ammo State", widget.NewSelect(nil, i.sit.Defender.Changed).
							Source(re.AmmoStates).
							Bind(i.sit.Defender.AmmoState)),
						widget.NewFormItem("Terrain", widget.NewSelect(nil, i.sit.Defender.Changed).
							Source(re.Terrains).
							Bind(i.sit.Defender.Terrain)),
						widget.NewFormItem("", widget.NewCheck("Enfilade", nil).
							Bind(i.sit.Defender.Enfilade)),
						widget.NewFormItem("Fire Factor", widget.NewLabel("").
							Bind(i.sit.Defender.FireFactor)),
						widget.NewFormItem("Mods", widget.NewLabel("").
							Bind(i.sit.Defender.DieModDesc)),
						widget.NewFormItem("Rolls", widget.NewLabel("").
							Bind(i.sit.Defender.FireRolls)),
						widget.NewFormItem("Result", widget.NewLabel("").
							Bind(i.sit.Defender.FireResults)),
						widget.NewFormItem("Damage", widget.NewLabel("").
							Bind(i.sit.Defender.FireTotal)),
						widget.NewFormItem("Morale Check", widget.NewLabel("").
							Bind(i.sit.Defender.MoraleCheckResult)),
						widget.NewFormItem("", widget.NewButtonWithIcon("Fire",
							theme.MailReplyIcon(),
							func() {
								i.sit.Defender.Fire()
								i.sit.FirefightCheck()
							},
						)),
					),
				),
			),
		),
	)

}
