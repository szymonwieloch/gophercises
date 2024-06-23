package invoice

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const headerLeft = 0.1
const headerRight = 0.15
const footerLeft = 0.98
const footerRigth = 0.96
const invoiceFontSize = 30.0
const headerFontSize = 12.0

func createHeader(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFillColor(103, 60, 79)

	headerPoints := []gofpdf.PointType{
		{X: 0, Y: 0},
		{X: w, Y: 0},
		{X: w, Y: h * headerRight},
		{X: 0, Y: h * headerLeft},
	}
	pdf.Polygon(headerPoints, "F")
	pdf.SetFont("Arial", "B", invoiceFontSize)
	pdf.SetTextColor(255, 255, 255)
	pdf.Text(w*0.05, h*headerLeft/2+invoiceFontSize/2, "INVOICE")

	banner := "(+48) 123 456 790\nsomeone@example.com\nexample.com"
	address := "55 The Green\nLondon\nNW30 5OR\nUnited Kingdom"
	pdf.SetFont("Arial", "", invoiceFontSize)
	pdf.SetFontSize(headerFontSize)
	pdf.MoveTo(w*0.65, 25)
	pdf.MultiCell(w*0.3, headerFontSize*1.5, address, gofpdf.BorderNone, gofpdf.AlignRight, true)
	pdf.MoveTo(w*0.45, 25)
	pdf.MultiCell(w*0.3, headerFontSize*1.5, banner, gofpdf.BorderNone, gofpdf.AlignRight, true)
	pdf.MoveTo(0, 10)
	pdf.Image("jump.png", w*0.32, 100, 80, 0, true, "", 0, "")

}

func createTop(pdf *gofpdf.Fpdf, invoiceNumber string, billedTo Company, date Date, totalPrice Cents) {
	w, h := pdf.GetPageSize()
	top := h * 0.2
	mid := w * 0.35
	left := w * 0.60

	height := showBox(pdf, w*0.05, top, "Billed To:", billedTo.Name, billedTo.Address[0], billedTo.Address[1], billedTo.Address[2])
	showBox(pdf, mid, top, "Invoice Number:", invoiceNumber)
	showBox(pdf, mid, top+height/2, "Date:", fmt.Sprintf("%d-%02d-%02d", date.Year, int(date.Month), date.Day))
	showBox(pdf, left, top, "Invoice Total:")
	pdf.SetFont("times", "B", 50)
	pdf.Text(left, top+50, totalPrice.String())

	lineY := top + height + 10
	pdf.Rect(0.05*w, lineY, 0.9*w, 3, "F")

}

func showBox(pdf *gofpdf.Fpdf, x float64, y float64, title string, lines ...string) float64 {
	pdf.SetTextColor(150, 150, 150)
	pdf.Text(x, y, title)
	fontSize, _ := pdf.GetFontSize()
	pdf.SetTextColor(0, 0, 0)
	top := y + 1.5*fontSize
	for _, line := range lines {
		pdf.Text(x, top, line)
		top += 1.5 * fontSize
	}
	return top - y
}

func createFooter(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	footerPoints := []gofpdf.PointType{
		{X: 0, Y: h},
		{X: w, Y: h},
		{X: w, Y: h * 0.98},
		{X: 0, Y: h * 0.96},
	}
	pdf.Polygon(footerPoints, "F")

}
