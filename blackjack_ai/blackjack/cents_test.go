package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCentsString(t *testing.T) {
	assert.Equal(t, Cents(100).String(), "1.00$")
	assert.Equal(t, Cents(-50).String(), "-0.50$")
	assert.Equal(t, Cents(-256).String(), "-2.56$")
}
