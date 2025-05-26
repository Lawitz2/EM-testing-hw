package hw9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpecificError(t *testing.T) {
	targetErr := SpecificError
	assert.ErrorIs(t, someError(1), targetErr) // fail (err == nil)
	assert.ErrorIs(t, someError(3), targetErr) // fail (err != targetErr)
	assert.ErrorIs(t, someError(4), targetErr) // pass (err == targetErr (or wraps it))
}
