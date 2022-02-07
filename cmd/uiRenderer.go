package cmd

import (
	"fmt"
	"log"
	"sync"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func uIRenderer(wg *sync.WaitGroup, selectedAlgo string, done, changed chan bool, data *[]float64) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	w, h := ui.TerminalDimensions()
	bc := widgets.NewBarChart()
	bc.Data = *data
	bc.Title = fmt.Sprintf("%s - Bar chart", selectedAlgo)
	bc.SetRect(0, 0, w, h)
	bc.BarGap = 1
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorGreen}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	ui.Render(bc)

	for {
		select {
		case <-done:
			for e := range ui.PollEvents() {
				if e.Type == ui.KeyboardEvent {
					break
				}
			}

			ui.Close()
			defer wg.Done()

			return
		case <-changed:
			bc.Data = *data
			ui.Render(bc)
		}
	}
}
