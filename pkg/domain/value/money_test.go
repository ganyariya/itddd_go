package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoneyAdd(t *testing.T) {
	tests := []struct {
		m1       Money
		m2       Money
		expected Money
	}{
		{m1: NewMoney(100, "JPY"), m2: NewMoney(200, "JPY"), expected: NewMoney(300, "JPY")},
	}
	for _, tt := range tests {
		actual, err := tt.m1.Add(tt.m2)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestMoneyAddError(t *testing.T) {
	tests := []struct {
		m1 Money
		m2 Money
	}{
		{m1: NewMoney(100, "JPY"), m2: NewMoney(200, "DOL")},
	}
	for _, tt := range tests {
		_, err := tt.m1.Add(tt.m2)
		assert.Error(t, err)
	}
}
