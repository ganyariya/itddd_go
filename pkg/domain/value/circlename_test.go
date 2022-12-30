package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleName(t *testing.T) {
	tests := []struct {
		Name  string
		IsErr bool
	}{
		{Name: "aaa", IsErr: false},
		{Name: "aaaa", IsErr: false},
		{Name: "aaaaaaaaaaaaaaaaaaaa", IsErr: false},
		{Name: "aa", IsErr: true},
		{Name: "aaaaaaaaaaaaaaaaaaaaa", IsErr: true},
	}

	for _, tt := range tests {
		_, err := NewCircleName(tt.Name)
		if tt.IsErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
