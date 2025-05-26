package hw1_test

import (
	"github.com/lawitz2/em-testing-hw/hw1"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	require.Equal(t, 2, hw1.Add(-3, 5))
}
