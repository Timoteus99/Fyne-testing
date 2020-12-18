package main

import (
	"testing"

	"fyne.io/fyne/test"
)

func TestCounterStartsAtZero(t *testing.T) {
	c := newCounter()

	if c.out.Text != "0" {
		t.Error("Incorrect starting value!")
	}
}

func TestCounterIncrementTwice(t *testing.T) {
	c := newCounter()

	test.Tap(c.add)
	test.Tap(c.add)

	if c.out.Text != "2" {
		t.Error("Incorrect incremented value!")
	}
}
