package hw10_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lawitz2/em-testing-hw/hw10"
	"github.com/stretchr/testify/require"
)

func TestHTTPServer(t *testing.T) {
	server := httptest.NewServer(hw10.AppConf())
	defer server.Close()

	t.Run("/hello endpoint", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/hello")
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		bodyBytes, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, "Hello!", string(bodyBytes))
	})

	t.Run("/data endpoint", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/data")
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var d hw10.DataStr
		err = json.NewDecoder(resp.Body).Decode(&d)
		require.NoError(t, err)

		expectedData := hw10.DataStr{Data: "very useful data"}
		require.Equal(t, expectedData, d)
	})

	t.Run("/bad endpoint", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/bad")
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var actualError map[string]string
		err = json.NewDecoder(resp.Body).Decode(&actualError)
		require.NoError(t, err)

		expectedError := map[string]string{"error": "bad request"}
		require.Equal(t, expectedError, actualError)
	})

	t.Run("non-existent endpoint returns 404", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/nonexistent")
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
