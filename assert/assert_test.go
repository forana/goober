package assert

import (
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, 1, 1)
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, 1, 2)
}

func TestTrue(t *testing.T) {
	True(t, true)
}
