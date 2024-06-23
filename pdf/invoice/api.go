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
	createTop(pdf, invoiceNumber, billedTo, date, 123456)

	//pdf.Cell(40, 10, "Hello, world")

	err := pdf.OutputFileAndClose("invoice.pdf")
	if err != nil {
		panic(err)
	}
}
