package main

import (
	"fmt"
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/untref-ayp2/ordenamiento/slice"
	"github.com/untref-ayp2/ordenamiento/sort"
)

// Parametros de la visualización. Cambiar a gusto
const size = 7
const time = 1000

var items = slice.GenerateNumbersConsecutive(size)

var p *widgets.Paragraph
var bc *widgets.BarChart

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Text
	p = widgets.NewParagraph()
	p.Text = strings.Trim(strings.Join(strings.Split(fmt.Sprint(items), " "), ", "), "[]")
	p.SetRect(0, 0, 100, 5)

	// Barchart
	bc = widgets.NewBarChart()
	bc.Data = items
	bc.SetRect(0, 5, size*2+2, size+5+3)
	bc.BarWidth = 1
	bc.BarColors = []ui.Color{ui.ColorWhite}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	render()

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			return
		case "1":
			p.Text = "Ordenando por inserción"
			sort.InsertionWithSleep(items, time, render, color)
			p.Text = "Esperando..."
			render()
		case "2":
			p.Text = "Ordenando por selección"
			sort.SelectionWithSleep(items, time, render, color)
			p.Text = "Esperando..."
			render()
		case "s":
			sort.Shuffle(items)
			render()
		}

	}
}

func render() {
	ui.Clear()
	ui.Render(p, bc)
}

func color(ccolors ...int) {
	colorList := []ui.Color{ui.ColorRed, ui.ColorBlue, ui.ColorYellow}
	colors := getSameColor()
	for i, value := range ccolors {
		colors[value] = colorList[i%len(colorList)]
	}
	bc.BarColors = colors
}

func getSameColor() []ui.Color {
	colors := make([]ui.Color, size)
	for i := range colors {
		colors[i] = ui.ColorGreen
	}
	return colors
}
