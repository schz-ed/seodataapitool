package visualization

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

// RenderChart returns an ASCII graph for the provided data.
func RenderChart(data []float64, caption string) string {
	graph := asciigraph.Plot(data, asciigraph.Caption(caption))
	return graph
}

// PrintChart prints the ASCII graph to the terminal.
func PrintChart(data []float64, caption string) {
	fmt.Println(RenderChart(data, caption))
}
