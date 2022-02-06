package cmd

import (
	"fmt"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func uIRenderer(selectedAlgo string, done, changed chan bool, data *[]float64) {
	fmt.Println("uiRender", &data)
	time.Sleep(time.Second * 5)
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	bc := widgets.NewBarChart()
	bc.Data = *data
	bc.Title = fmt.Sprintf("%s - Bar chart", selectedAlgo)
	bc.SetRect(0, 0, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorGreen}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	ui.Render(bc)

	for changed := range changed {
		if changed {
			fmt.Println("change", &data)
			fmt.Println(*data)
			bc.Data = *data
			ui.Render(bc)
		}
	}
}
