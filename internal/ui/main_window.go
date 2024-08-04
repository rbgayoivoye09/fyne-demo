package ui

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func SetupMainWindow(myWindow fyne.Window) {

	textColor := theme.Color(theme.ColorNameForeground)

	// Create labels for hours, minutes, and seconds
	hourLabel := canvas.NewText("", textColor)
	minuteLabel := canvas.NewText("", textColor)
	secondLabel := canvas.NewText("", textColor)
	colonLabel := canvas.NewText(":", textColor)

	// Set the font size for the labels
	hourLabel.TextSize = 144
	minuteLabel.TextSize = 144
	secondLabel.TextSize = 144
	colonLabel.TextSize = 144

	// Create a container for each section
	hourContainer := container.New(layout.NewCenterLayout(), hourLabel)
	minuteContainer := container.New(layout.NewCenterLayout(), minuteLabel)
	secondContainer := container.New(layout.NewCenterLayout(), secondLabel)
	colonContainer := container.New(layout.NewCenterLayout(), colonLabel)
	colonContainer1 := container.New(layout.NewCenterLayout(), colonLabel)

	// Create a background rectangle
	background := canvas.NewRectangle(color.White) // Default color

	updateLayout := func() {
		orientation := fyne.CurrentApp().Driver().Device().Orientation()

		var clockContainer *fyne.Container

		if fyne.IsVertical(orientation) {
			// Arrange containers in a vertical box
			clockContainer = container.New(layout.NewVBoxLayout(), hourContainer, minuteContainer, secondContainer)
		} else {
			// Arrange containers in a horizontal box
			clockContainer = container.New(layout.NewHBoxLayout(), hourContainer, colonContainer, minuteContainer, colonContainer1, secondContainer)
		}

		// Overlay the clock on the background
		mainContainer := container.New(layout.NewStackLayout(), clockContainer)

		// Create buttons to open the color picker dialogs
		bgColorButton := widget.NewButton("Change Background Color", func() {
			colorPicker := dialog.NewColorPicker("Pick Background Color", "Choose a color for the background", func(c color.Color) {
				background.FillColor = c
				background.Refresh()
			}, myWindow)
			colorPicker.Show()
		})

		textColorButton := widget.NewButton("Change Text Color", func() {
			colorPicker := dialog.NewColorPicker("Pick Text Color", "Choose a color for the text", func(c color.Color) {
				textColor = c
				hourLabel.Color = textColor
				minuteLabel.Color = textColor
				secondLabel.Color = textColor
				colonLabel.Color = textColor
				hourLabel.Refresh()
				minuteLabel.Refresh()
				secondLabel.Refresh()
				colonLabel.Refresh()
			}, myWindow)
			colorPicker.Show()
		})

		// Create a vertical box layout to hold the main content and the buttons
		content := container.New(layout.NewVBoxLayout(), mainContainer, bgColorButton, textColorButton)

		// Center the content in the window
		centeredContent := container.New(layout.NewCenterLayout(), content)

		myWindow.SetContent(container.New(layout.NewStackLayout(), background, centeredContent))

	}

	// Update the labels every second
	go func() {
		for range time.Tick(time.Second) {
			now := time.Now()
			hour, minute, second := now.Hour(), now.Minute(), now.Second()

			hourLabel.Text = fmt.Sprintf("%02d", hour)
			minuteLabel.Text = fmt.Sprintf("%02d", minute)
			secondLabel.Text = fmt.Sprintf("%02d", second)

			hourLabel.Refresh()
			minuteLabel.Refresh()
			secondLabel.Refresh()

			// Update layout
			updateLayout()
		}
	}()

}

func colorPicked(c color.Color, w fyne.Window) {
	log.Println("Color picked:", c)
	rectangle := canvas.NewRectangle(c)
	size := 2 * theme.IconInlineSize()
	rectangle.SetMinSize(fyne.NewSize(size, size))
	dialog.ShowCustom("Color Picked", "Ok", rectangle, w)
}

func makeCell() fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{128, 128, 128, 255})
	rect.SetMinSize(fyne.NewSize(30, 30))
	return rect
}
