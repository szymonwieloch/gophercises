package certificate

import (
	"github.com/jung-kurt/gofpdf"
)

const triangleFraction = 0.15
const lightWidthFraction = 0.8

func createHeaderAndFooter(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()

	headerPoints1 := []gofpdf.PointType{
		{X: w * (1 - lightWidthFraction), Y: 0},
		{X: w, Y: 0},
		{X: w, Y: h * triangleFraction},
	}

	pdf.SetFillColor(122, 85, 102)
	pdf.Polygon(headerPoints1, "F")

	headerPoints2 := []gofpdf.PointType{
		{X: w, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: h * triangleFraction},
	}
	pdf.SetFillColor(103, 60, 79)
	pdf.Polygon(headerPoints2, "F")

	footerPoints1 := []gofpdf.PointType{
		{X: w * lightWidthFraction, Y: h},
		{X: 0, Y: h},
		{X: 0, Y: h * (1 - triangleFraction)},
	}

	pdf.SetFillColor(122, 85, 102)
	pdf.Polygon(footerPoints1, "F")

	footerPoints2 := []gofpdf.PointType{
		{X: 0, Y: h},
		{X: w, Y: h},
		{X: w, Y: h * (1 - triangleFraction)},
	}
	pdf.SetFillColor(103, 60, 79)
	pdf.Polygon(footerPoints2, "F")
}

func createCenteredText(pdf *gofpdf.Fpdf, name string) {
	_, h := pdf.GetPageSize()
	pdf.SetTextColor(50, 50, 50)

	pdf.SetFont("times", "B", 50)
	pdf.MoveTo(0, h*0.15)
	pdf.WriteAligned(0, 70, "Certificate Of Completion", gofpdf.AlignCenter)

	pdf.SetFont("arial", "", 28)
	pdf.MoveTo(0, h*0.30)
	pdf.WriteAligned(0, 28, "This certificate is awarded to", gofpdf.AlignCenter)

	pdf.MoveTo(0, h*0.39)
	pdf.SetFont("times", "B", 42)
	pdf.WriteAligned(0, 42, name, gofpdf.AlignCenter)

	pdf.MoveTo(0, h*0.51)
	pdf.SetFont("arial", "", 22)
	pdf.WriteAligned(0, 30, "For successfully completing all twenty programming exercises in the Gophercises Go programming course", gofpdf.AlignCenter)

}

const gopherWidth = 100

func drawGopher(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.ImageOptions("jump.png", (w-gopherWidth)/2, h*0.67, gopherWidth, 0, false, gofpdf.ImageOptions{}, 0, "")
}

func drawDetails(pdf *gofpdf.Fpdf, description string, content string, image string, mid float64) {
	w, h := pdf.GetPageSize()
	pdf.SetFillColor(100, 100, 100)
	pdf.SetTextColor(100, 100, 100)
	recWidth := w * 0.28
	x := mid - recWidth/2
	y := 0.78 * h
	pdf.Rect(x, y, recWidth, 2, "F")
	pdf.SetFont("arial", "", 12)
	_, lineHt := pdf.GetFontSize()
	pdf.MoveTo(x, y+10)
	pdf.CellFormat(recWidth, lineHt, description, gofpdf.BorderNone, 0, gofpdf.AlignCenter, false, 0, "")

	if content != "" {
		pdf.SetTextColor(50, 50, 50)
		pdf.SetFont("times", "", 25)
		_, lineHt = pdf.GetFontSize()
		pdf.MoveTo(x, y-1.5*lineHt)
		pdf.CellFormat(recWidth, lineHt, content, gofpdf.BorderNone, 0, gofpdf.AlignCenter, false, 0, "")
	}
	if image != "" {
		//imageHeight := 100
		pdf.ImageOptions(image, x+10, y-50, recWidth-20, 0, false, gofpdf.ImageOptions{}, 0, "")
	}
}
