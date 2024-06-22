package main

import (
	"fmt"
	"io"
	"math"
	"slices"

	svg "github.com/ajstarks/svgo/float"
)

const bodyHeight = 100
const bodyWidth = 150
const margin = 10
const dataLabelHeight = 10
const dataLabelWidth = 20
const labelHeight = 10
const totalHeight = bodyHeight + 2*margin + dataLabelHeight
const totalWidth = bodyWidth + 2*margin + dataLabelWidth
const asixColor = "black"
const guideStyle = "stroke:silver;stroke-width:0.5;stroke-dasharray:2,3"

type series struct {
	data  []float64
	name  string
	color string
}

func chart(x []float64, data []series, w io.Writer) error {
	scaleX := newScale(x, margin+dataLabelWidth, margin+dataLabelWidth+bodyWidth)
	var yAll []float64
	for _, serie := range data {
		yAll = append(yAll, serie.data...)
	}
	scaleY := newScale(yAll, margin+bodyHeight, margin)
	normX := scaleX.transformSeries(x)
	guidesX := guides(x)
	guidesY := guides(yAll)
	guidesNormX := scaleX.transformSeries(guidesX)
	guidesNormY := scaleY.transformSeries(guidesY)

	canvas := svg.New(w)
	canvasHeight := float64(totalHeight + labelHeight*len(data))
	canvas.Start(totalWidth, canvasHeight)
	createMarker(canvas)
	canvas.Rect(0, 0, totalWidth, canvasHeight, "fill:white")
	drawXGuides(canvas, guidesNormX, guidesX)
	drawYGuides(canvas, guidesNormY, guidesY)
	drawAxis(canvas)

	for _, serie := range data {
		normY := scaleY.transformSeries(serie.data)
		lineStyle := fmt.Sprintf("stroke:%s;stroke-width:1;fill:none; marker-start:url(#datapoint); marker-end:url(#datapoint); marker-mid:url(#datapoint)", serie.color)
		canvas.Polyline(normX, normY, lineStyle)
	}
	drawLabels(canvas, data)
	canvas.End()
	return nil
}

func createMarker(canvas *svg.SVG) {
	canvas.Def()
	canvas.Marker("datapoint", 1, 1, 2, 2)
	canvas.Circle(1, 1, 1, "fill:black")
	canvas.MarkerEnd()
	canvas.DefEnd()
}

func drawAxis(canvas *svg.SVG) {
	lineStyle := fmt.Sprintf("stroke:%s; stroke-width:0.5", asixColor)
	fillStyle := fmt.Sprintf("fill:%s", asixColor)
	canvas.Line(margin+dataLabelWidth, margin, margin+dataLabelWidth, margin+bodyHeight, lineStyle)
	canvas.Line(margin+dataLabelWidth, margin+bodyHeight, margin+dataLabelWidth+bodyWidth, margin+bodyHeight, lineStyle)
	arrowUpX := [...]float64{margin + dataLabelWidth, margin + dataLabelWidth - 1, margin + dataLabelWidth + 1}
	arrowUpY := [...]float64{margin - 3, margin, margin}
	canvas.Polygon(arrowUpX[:], arrowUpY[:], fillStyle)
	arrowRightX := [...]float64{margin + dataLabelWidth + bodyWidth + 3, margin + dataLabelWidth + bodyWidth, margin + dataLabelWidth + bodyWidth}
	arrowRightY := [...]float64{margin + bodyHeight, margin + bodyHeight + 1, margin + bodyHeight - 1}
	canvas.Polygon(arrowRightX[:], arrowRightY[:], fillStyle)
}

func guides(series []float64) []float64 {
	seriesMin := slices.Min(series)
	seriesMax := slices.Max(series)
	diff := seriesMax - seriesMin
	power := math.Floor(math.Log10(diff))
	// create guideline every 10**power
	guideDistance := math.Pow(10, power)
	curr := math.Ceil(seriesMin/guideDistance) * guideDistance
	var result []float64
	for curr <= seriesMax {
		result = append(result, curr)
		curr += guideDistance
	}
	return result

}

func drawXGuides(canvas *svg.SVG, guidesNorm []float64, labels []float64) {
	for idx := range guidesNorm {
		canvas.Line(guidesNorm[idx], margin, guidesNorm[idx], margin+bodyHeight, guideStyle)
		canvas.Text(guidesNorm[idx], margin+bodyHeight+2, fmt.Sprint(labels[idx]), "font-size: 5; text-anchor:middle; alignment-baseline:before-edge")
	}
}
func drawYGuides(canvas *svg.SVG, guidesNorm []float64, labels []float64) {
	for idx := range guidesNorm {
		canvas.Line(margin+dataLabelWidth, guidesNorm[idx], margin+dataLabelWidth+bodyWidth, guidesNorm[idx], guideStyle)
		canvas.Text(margin+dataLabelWidth-2, guidesNorm[idx], fmt.Sprint(labels[idx]), "font-size: 5; text-anchor:end; alignment-baseline:middle")
	}
}

func drawLabels(canvas *svg.SVG, series []series) {
	for idx, serie := range series {
		labelY := margin + bodyHeight + dataLabelHeight + labelHeight*(float64(idx)+0.5)
		labelStyle := fmt.Sprintf("stroke:%s;stroke-width:1;", serie.color)
		canvas.Line(margin+dataLabelWidth, labelY, margin+dataLabelWidth+30, labelY, labelStyle)
		canvas.Text(margin+dataLabelWidth+35, labelY, serie.name, "font-size: 5; text-anchor:start;alignment-baseline:middle")
	}
}

func newScale(series []float64, from, to float64) scale {
	return scale{
		max:  slices.Max(series),
		min:  slices.Min(series),
		from: from,
		to:   to,
	}
}

type scale struct {
	min, max float64
	from, to float64
}

func (scale *scale) transform(val float64) float64 {
	normalized := (val - scale.min) / (scale.max - scale.min)
	return scale.from + (scale.to-scale.from)*normalized
}

func (scale *scale) transformSeries(series []float64) []float64 {
	result := make([]float64, 0, len(series))
	for _, val := range series {
		result = append(result, scale.transform(val))
	}
	return result
}
