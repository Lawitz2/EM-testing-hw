package hw4

import (
	"testing"
	"testing/quick"
)

func TestUselessWork(t *testing.T) {
	property := func(a, b int) bool {
		return UselessWork(a, b) == a
	}

	conf := quick.Config{
		MaxCount:      100,
		MaxCountScale: 0,
		Rand:          nil,
		Values:        nil,
	}

	err := quick.Check(property, &conf)
	if err != nil {
		t.Errorf("property check failed, %v", err)
	}
}
