package main

import (
	"net/url"

	"github.com/steveoc64/re/pkg/ui"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/cmd/fyne_demo/data"
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

	screens := ui.NewInfantryScreens()

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Command", theme.HomeIcon(), welcomeScreen(a)),
		widget.NewTabItemWithIcon("Movement", theme.ContentCopyIcon(), welcomeScreen(a)),
		widget.NewTabItemWithIcon("Skirmishers", theme.ViewRefreshIcon(), ui.Skirmish()),
		widget.NewTabItemWithIcon("Infantry", theme.ViewRefreshIcon(), screens.Canvas()),
		widget.NewTabItemWithIcon("Artillery", theme.ViewRefreshIcon(), ui.Artillery()),
		widget.NewTabItemWithIcon("Cavalry", theme.DocumentCreateIcon(), ui.CavalryVCavalry()),
		widget.NewTabItemWithIcon("Morale and Fatigue", theme.ViewFullScreenIcon(), welcomeScreen(a)),
	)
	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	w.SetContent(tabs)

	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}
