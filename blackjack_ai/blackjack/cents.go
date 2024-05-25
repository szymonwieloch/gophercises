package blackjack

import "fmt"

// Cash is represented as the number of cents
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
