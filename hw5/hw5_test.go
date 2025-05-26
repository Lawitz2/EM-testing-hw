package hw5_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lawitz2/em-testing-hw/hw5" // Changed import path from example.com/hw5
)

func TestPingHandler(t *testing.T) {
	router := hw5.AppConf()

	t.Run("GET /ping returns pong", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/ping", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		require.Equal(t, 200, rec.Code)
		require.Equal(t, "pong", rec.Body.String())
	})

	t.Run("GET /nonexistent returns 404", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/nonexistent", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		require.Equal(t, 404, rec.Code)
	})
}
