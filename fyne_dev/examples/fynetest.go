package main

import (
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func showTime(clock *widget.Label) {
	formatted := time.Now().Format("03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello!")

	clock := widget.NewLabel("")
	showTime(clock)

	w.SetContent(clock)

	go func() {
		t := time.NewTicker(time.Second)

		for range t.C {
			showTime(clock)
		}
	}()

	w.ShowAndRun()
}
