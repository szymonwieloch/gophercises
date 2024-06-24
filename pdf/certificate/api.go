package certificate

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func Create(name string, date Date) {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, gofpdf.UnitPoint, gofpdf.PageSizeA4, "")
	pdf.AddPage()

	createHeaderAndFooter(pdf)

	createCenteredText(pdf, name)
	drawGopher(pdf)
	w, _ := pdf.GetPageSize()

	drawDetails(pdf, "Date", fmt.Sprintf("%d-%02d-%02d", date.Year, date.Month, date.Year), "", w*0.22)
	drawDetails(pdf, "Instructor", "", "signature.png", w*0.78)

	err := pdf.OutputFileAndClose("certificate.pdf")
	if err != nil {
		panic(err)
	}
}
