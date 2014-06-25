package point

import (
	"testing"

	"github.com/chai2010/gopkg/assert"
)

func TestAsserts(t *testing.T) {
	p1 := Point{1, 1}
	p2 := Point{2, 1}

	assert.Equal(t, p1, p2)
}
