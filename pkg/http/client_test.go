package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Get(t *testing.T) {
	// Setup mock server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message": "get success"}`))
	}))
	defer testServer.Close()

	// Initialize client with mock server URL
	client := NewClient(testServer.URL, nil)

	// Perform GET request
	resp, err := client.Get("/test")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read response body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message": "get success"}`, string(body))
}

func TestClient_Post(t *testing.T) {
	// Setup mock server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// Read request body
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"key":"value"}`, string(body))

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message": "post success"}`))
	}))
	defer testServer.Close()

	// Initialize client with mock server URL
	client := NewClient(testServer.URL, nil)

	// Perform POST request
	body := map[string]string{"key": "value"}
	resp, err := client.Post("/test", body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message": "post success"}`, string(respBody))
}

func TestClient_Put(t *testing.T) {
	// Setup mock server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// Read request body
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"key":"value"}`, string(body))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message": "put success"}`))
	}))
	defer testServer.Close()

	// Initialize client with mock server URL
	client := NewClient(testServer.URL, nil)

	// Perform PUT request
	body := map[string]string{"key": "value"}
	resp, err := client.Put("/test", body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message": "put success"}`, string(respBody))
}

func TestClient_Delete(t *testing.T) {
	// Setup mock server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// Read request body
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"key":"value"}`, string(body))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message": "delete success"}`))
	}))
	defer testServer.Close()

	// Initialize client with mock server URL
	client := NewClient(testServer.URL, nil)

	// Perform DELETE request
	body := map[string]string{"key": "value"}
	resp, err := client.Delete("/test", body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message": "delete success"}`, string(respBody))
}

func TestClient_Patch(t *testing.T) {
	// Setup mock server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// Read request body
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"key":"value"}`, string(body))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message": "patch success"}`))
	}))
	defer testServer.Close()

	// Initialize client with mock server URL
	client := NewClient(testServer.URL, nil)

	// Perform PATCH request
	body := map[string]string{"key": "value"}
	resp, err := client.Patch("/test", body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message": "patch success"}`, string(respBody))
}
