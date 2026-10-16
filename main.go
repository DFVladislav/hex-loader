package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"

	"github.com/DFVladislav/hex-loader/gui"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())

	w := a.NewWindow("Hex Loader")
	u := gui.NewUI(w)
	w.SetContent(u.Build())
	w.Resize(fyne.NewSize(640, 520))
	w.ShowAndRun()
}
