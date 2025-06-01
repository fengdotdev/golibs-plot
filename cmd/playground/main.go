package main

import (
	"github.com/fengdotdev/golibs-plot/sandbox/draft1/plot2d"
)

func main() {

	myplot := plot2d.NewPlot2D(plot2d.Size{Width: 800, Height: 600})
	myplot.AddPoint("A", 100.0, 100.0)
	myplot.AddPoint("B", 200.0, 200.0)
	myplot.AddPoint("C", 300.0, 300.0)
	myplot.AddPoint("D", 400.0, 400.0)
	myplot.AddPoint("E", 500.0, 500.0)
	err := myplot.Draw()
	if err != nil {
		panic(err)
	}
}
