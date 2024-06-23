package main

import (
	"time"

	"github.com/szymonwieloch/gophercises/pdf/invoice"
)

func main() {

	billedTo := invoice.Company{
		Name:    "Famous Corp.",
		Address: [3]string{"760 Church Street", "London", "NW50 6BQ"},
	}
	date := invoice.Date{
		Year:  2024,
		Month: time.April,
		Day:   23,
	}
	items := []invoice.Item{
		{Name: "2x6 Lumber 8'", NettPrice: 375, Quantity: 220, VAT: 12},
		{Name: "2x6 Lumber 10'", NettPrice: 555, Quantity: 18, VAT: 12},
		{Name: "2x4 Lumber 8'", NettPrice: 290, Quantity: 80, VAT: 12},
		{Name: "Drywall Sheet", NettPrice: 820, Quantity: 50, VAT: 12},
		{Name: "Paint", NettPrice: 1450, Quantity: 3, VAT: 8},
		{Name: "An item with a surprisingly long name that is not going to fit in just one row", NettPrice: 999, Quantity: 12, VAT: 0},
	}
	invoice.Create("00000000123", date, billedTo, items)
}
