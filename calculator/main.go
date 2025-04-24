package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Go Calculator")

	//entry for displaying numbers
	entry := widget.NewEntry()
	entry.Disable()

	var currentInput string

	//helper func updates display
	updateDisplay := func(input string) {
		currentInput += input
		entry.SetText(currentInput)
	}

	//handle "="
	calculate := func() {
		//simple calc parser, handles 2 ops only
		var op string
		var a, b float64

		for _, char := range []string{"+", "-", "*", "/"} {
			if pos := findOperator(currentInput, char); pos != -1 {
				op = char
				aStr := currentInput[:pos]
				bStr := currentInput[pos+1:]

				a, _ = strconv.ParseFloat(aStr, 64)
				b, _ = strconv.ParseFloat(bStr, 64)

				var result float64

				switch op {
				case "+":
					result = a + b
				case "-":
					result = a - b
				case "*":
					result = a * b

				case "/":
					if b != 0 {
						result = a / b
					
					} else {
						entry.SetText("Error, cannot divide by 0")
						currentInput = ""
						return
					}

				}

				currentInput = fmt.Sprintf("%v", result)
				entry.SetText(currentInput)
				return
			}
		}
	}

	clear := func() {
		currentInput = ""
		entry.SetText("")
	}

	//buttons

	buttons := []string {
		"7", "8", "9", "/",
		"4", "5", "6", "*",
		"1", "2", "3", "-",
		"0", "C", "=", "+"
	}

	grid := container.NewGridWithColumns(4)

	for _, btn := range buttons {
		b := btn //capture range variable
		button := widget.NewButton(b, func() {
			switch b {
			case "=":
				cacalculate()

			case "C":
				clear()
			default:
				updateDisplay(b)

			}
		})
		grid.Add(button)
	}

	myWindow.SetContent(container.NewVBox(
		entry,
		grid
	))
	myWindow.ShowAndRun()
}

func findOperator (s string, op string) int {
	for i, c := range s {
		if string(c) == op {
			return i
		}
	}
	return -1
}
