package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleId(t *testing.T) {
	for i := 0; i < 10; i++ {
		_, err := NewCircleId()
		assert.NoError(t, err)
	}
}
