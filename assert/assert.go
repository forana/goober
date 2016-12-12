package assert

import (
	"testing"
)

func Equal(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("Assertion failed: expected %v, got %v", expected, actual)
	}
}

func NotEqual(t *testing.T, actual interface{}, expected interface{}) {
	if actual == expected {
		t.Errorf("Assertion failed: expected anything other than %v", expected)
	}
}

func True(t *testing.T, value bool) {
	if !value {
		t.Errorf("Assertion failed: expected non-false value")
	}
}
