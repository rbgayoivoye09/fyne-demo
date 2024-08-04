package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/rbgayoivoye09/fyne-demo/internal/ui"
)

func main() {
	a := app.NewWithID("io.fyne.flip-clock")
	w := a.NewWindow("Flip Clock")
	ui.SetupMainWindow(w)
	w.ShowAndRun()
}
