package hw3_test

import (
	"github.com/lawitz2/em-testing-hw/hw3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetForm(t *testing.T) {
	t.Run("TestMapForm", func(t *testing.T) {
		testData := []string{"apple", "banana", "apple", "pear"}
		result := hw3.SetForm(testData...)

		require.Len(t, result, 3)
		require.Contains(t, result, "apple")
		require.Contains(t, result, "pear")
		require.NotContains(t, result, "strawberry")
		require.Equal(t, result["apple"], struct{}{})
	})
}
