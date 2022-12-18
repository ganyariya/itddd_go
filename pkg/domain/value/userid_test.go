package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserId(t *testing.T) {
	for i := 0; i < 10; i++ {
		_, err := NewUserId()
		assert.NoError(t, err)
	}
}
