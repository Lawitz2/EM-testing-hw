package hw2_test

import (
	"github.com/lawitz2/em-testing-hw/hw2"
	"github.com/stretchr/testify/require"
	"testing"
)

// Получаем ошибку, если число делится на 5
func TestHw2Err(t *testing.T) {
	require.NoError(t, hw2.Err(1))
	require.NoError(t, hw2.Err(3))
	require.NoError(t, hw2.Err(-4))
	require.Error(t, hw2.Err(5))
}
