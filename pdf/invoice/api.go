package invoice

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Cents int64

func (c Cents) String() string {
	var sign string
	if c < 0 {
		sign = "-"
		c = -c
	}
	dollars := c / 100
	actualCents := c % 100
	return fmt.Sprintf("%s%d.%02d$", sign, dollars, actualCents)
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

type Company struct {
	Name    string
	Address [3]string
}

type Item struct {
	Name      string
	NettPrice Cents
	Quantity  uint
	VAT       uint8
}

func Create(invoiceNumber string, date Date, billedTo Company, items []Item) {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeA4, "")
	pdf.AddPage()

	createHeader(pdf)
	createFooter(pdf)
	var totalInvoice Cents
	var totalQuantity uint
	for _, item := range items {
		totalInvoice += totalItemPrice(item)
		totalQuantity += item.Quantity
	}

	yPos := createTop(pdf, invoiceNumber, billedTo, date, totalInvoice)
	pdf.SetFont("times", "", 20)
	pdf.SetTextColor(150, 150, 150)
	yPos += 20
	pdf.SetFillColor(150, 150, 150)
	yPos = showRow(pdf, yPos, "Item", "Net Price", "Quant.", "VAT", "Total")
	pdf.SetTextColor(0, 0, 0)
	for _, item := range items {
		total := totalItemPrice(item)
		yPos = showRow(pdf, yPos, item.Name, item.NettPrice.String(), fmt.Sprint(item.Quantity), fmt.Sprintf("%d%%", item.VAT), total.String())
	}
	pdf.SetTextColor(150, 150, 150)
	showRow(pdf, yPos, "SUM", "", fmt.Sprint(totalQuantity), "", totalInvoice.String())

	//pdf.Cell(40, 10, "Hello, world")

	err := pdf.OutputFileAndClose("invoice.pdf")
	if err != nil {
		panic(err)
	}
}
