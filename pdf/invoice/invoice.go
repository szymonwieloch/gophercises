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

func createTop(pdf *gofpdf.Fpdf, invoiceNumber string, billedTo Company, date Date, totalPrice Cents) float64 {
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

	return lineY + 3

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

const rowStart = 0.05
const rowPrice = 0.45
const rowQuantity = 0.60
const rowVAT = 0.70
const rowTotal = 0.80
const rowEnd = 0.95
const rowHeight = 1.5

func showRow(pdf *gofpdf.Fpdf, y float64, name string, price string, quantity string, vat string, total string) float64 {
	w, _ := pdf.GetPageSize()
	fontSize, _ := pdf.GetFontSize()
	pdf.MoveTo(w*rowStart, y)
	pdf.MultiCell(w*(rowPrice-rowStart), fontSize*rowHeight, name, gofpdf.BorderNone, gofpdf.AlignBaseline, false)
	pdf.MoveTo((w * rowPrice), y)
	pdf.CellFormat(w*(rowQuantity-rowPrice), fontSize*rowHeight, price, gofpdf.BorderNone, 0, gofpdf.AlignRight, false, 0, "")
	pdf.MoveTo((w * rowQuantity), y)
	pdf.CellFormat(w*(rowVAT-rowQuantity), fontSize*rowHeight, quantity, gofpdf.BorderNone, 0, gofpdf.AlignRight, false, 0, "")
	pdf.MoveTo((w * rowVAT), y)
	pdf.CellFormat(w*(rowTotal-rowVAT), fontSize*rowHeight, vat, gofpdf.BorderNone, 0, gofpdf.AlignRight, false, 0, "")
	pdf.MoveTo((w * rowTotal), y)
	pdf.CellFormat(w*(rowEnd-rowTotal), fontSize*rowHeight, total, gofpdf.BorderNone, 0, gofpdf.AlignRight, false, 0, "")

	tmp := pdf.SplitLines([]byte(name), w*(rowPrice-rowStart))
	y += rowHeight * fontSize * float64(len(tmp))

	y += 7
	pdf.Rect(rowStart*w, y, w*(rowEnd-rowStart), 2, "F")
	y += 1
	y += 5

	return y
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

func totalItemPrice(item Item) Cents {
	total := float64(item.NettPrice) * float64(item.Quantity) * (1 + float64(item.VAT)/100.0)
	return Cents(total)
}
