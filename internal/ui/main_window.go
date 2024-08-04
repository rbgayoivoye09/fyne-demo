package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func SetupMainWindow(w fyne.Window) {
    label := widget.NewLabel("Hello, Fyne!")
    button := widget.NewButton("Click Me", func() {
        label.SetText("Button Clicked!")
    })
    w.SetContent(container.NewVBox(
        label,
        button,
    ))
}
