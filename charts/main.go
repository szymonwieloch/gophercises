package main

import (
	"os"
)

var colors []string = []string{
	"red", "lime", "fuchsia", "green", "blue", "teal", "maroon", "navy", "purple", "olive", "aqua", "yellow",
}

func main() {
	args := parseArgs()
	parseColumns := []string{args.X}
	parseColumns = append(parseColumns, args.Y...)
	data, err := parseCSVFile(args.Input, parseColumns)
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(args.Output)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	x := data[0]
	var ySeries []series
	for idx, dt := range data[1:] {
		ySerie := series{
			data:  dt,
			name:  args.Y[idx],
			color: colors[idx%len(colors)],
		}
		ySeries = append(ySeries, ySerie)
	}

	chart(x, ySeries, outFile)
}
