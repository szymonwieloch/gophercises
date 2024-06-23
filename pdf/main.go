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
	items := []invoice.Item{}
	invoice.Create("00000000123", date, billedTo, items)
}
