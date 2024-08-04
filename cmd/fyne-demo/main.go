package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/rbgayoivoye09/fyne-demo/internal/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Demo")
	ui.SetupMainWindow(w)
	w.ShowAndRun()
}
