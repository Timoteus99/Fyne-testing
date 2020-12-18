package main

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type counter struct {
	value     int
	out       *widget.Label
	add       *widget.Button
	substract *widget.Button
}

func (c *counter) substraction() {
	c.value--
	c.out.SetText(fmt.Sprintf("%d", c.value))
}

func (c *counter) increment() {
	c.value++
	c.out.SetText(fmt.Sprintf("%d", c.value))
}

func newCounter() *counter {
	c := &counter{}
	c.out = widget.NewLabel("0")
	c.add = widget.NewButton("Add", c.increment)
	c.substract = widget.NewButton("Substract", c.substraction)

	return c
}

func main() {
	a := app.New()
	w := a.NewWindow("Counter")

	c := newCounter()

	ui := widget.NewVBox(c.out, c.add, c.substract)

	w.SetContent(ui)

	w.ShowAndRun()
}
