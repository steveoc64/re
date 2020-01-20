package main

import (
	"github.com/steveoc64/re/pkg/re"
	"github.com/steveoc64/re/pkg/ui"
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/cmd/fyne_demo/data"
	"fyne.io/fyne/cmd/fyne_demo/screens"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const preferenceCurrentTab = "currentTab"

func welcomeScreen(a fyne.App) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(data.FyneScene)
	logo.SetMinSize(fyne.NewSize(228, 167))

	link, err := url.Parse("https://fyne.io/")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		widget.NewHyperlinkWithStyle("fyne.io", link, fyne.TextAlignCenter, fyne.TextStyle{}),
		layout.NewSpacer(),

		widget.NewGroup("Theme",
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				widget.NewButton("Dark", func() {
					a.Settings().SetTheme(theme.DarkTheme())
				}),
				widget.NewButton("Light", func() {
					a.Settings().SetTheme(theme.LightTheme())
				}),
			),
		),
	)
}

func main() {
	a := app.NewWithID("io.wargaming.re")
	a.SetIcon(theme.FyneLogo())

	w := a.NewWindow("Revolution and Empire Calculators")
	w.SetMaster()

	unitA := re.NewSmallArmsUnit()
	unitB := re.NewSmallArmsUnit()
	sit := re.NewSmallArmsSituation([]*re.Unit{unitA, unitB})

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Command", theme.HomeIcon(), welcomeScreen(a)),
		widget.NewTabItemWithIcon("Movement", theme.ContentCopyIcon(), screens.WidgetScreen()),
		widget.NewTabItemWithIcon("Skirmish Attack", theme.ViewRefreshIcon(), ui.SmallArms(sit)),
		widget.NewTabItemWithIcon("Musket Fire", theme.ViewRefreshIcon(), ui.SmallArms(sit)),
		widget.NewTabItemWithIcon("Artillery Bombardment", theme.ViewRefreshIcon(), ui.SmallArms(sit)),
		widget.NewTabItemWithIcon("Bayonet Assault", theme.DocumentCreateIcon(), screens.GraphicsScreen()),
		widget.NewTabItemWithIcon("Cavalry Charge", theme.DocumentCreateIcon(), screens.GraphicsScreen()),
		widget.NewTabItemWithIcon("Morale and Fatigue", theme.ViewFullScreenIcon(), screens.DialogScreen(w)),
	)
	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	w.SetContent(tabs)

	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}
