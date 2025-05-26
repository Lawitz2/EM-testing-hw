package hw8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter(t *testing.T) {
	assert.Equal(t, 1000000, SafeCounter())

	// При запуске тестов с -race он будет ругаться в этом месте,
	// т.к. UnsafeCounter() не синхронизирован
	assert.Less(t, UnsafeCounter(), 1000000)
}
