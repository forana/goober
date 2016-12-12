package commands

import (
	"testing"

	"github.com/forana/goober/assert"
)

func TestMake16(t *testing.T) {
	assert.Equal(t, uint16(0xAABB), make16(0xAA, 0xBB))
}

func TestMake16LSFirst(t *testing.T) {
	assert.Equal(t, uint16(0xBBAA), make16LSFirst(0xAA, 0xBB))
}

func TestRegistry(t *testing.T) {
	// really just retrieve a single command
	assert.Equal(t, uint(4), Registry()[0x00](nil))
}
