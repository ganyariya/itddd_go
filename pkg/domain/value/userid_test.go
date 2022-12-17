package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserId(t *testing.T) {
	tests := []struct {
		Id      string
		IsError bool
	}{
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1", true},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa;", true},
	}

	for _, tt := range tests {
		_, err := NewUserId(tt.Id)
		if tt.IsError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
