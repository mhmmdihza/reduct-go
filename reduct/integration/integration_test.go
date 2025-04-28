package integration

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mhmmdihza/reduct-go/reduct/integration/client/operations"
	"github.com/stretchr/testify/assert"
)

var scenarioNewIntegration = []struct {
	scenarioName    string
	clientOptions   *ClientOptions
	expectedRequest expectedApiRequest
	mockResponse    mockApiResponse
	delayResponse   time.Duration
}{
	{
		scenarioName:  "success without client options",
		clientOptions: nil,
		expectedRequest: expectedApiRequest{
			path:   "/api/v1/b/bucket",
			method: "GET",
			headers: map[string]string{
				"Authorization": "",
			},
		},
		mockResponse: mockApiResponse{},
	},
	{
		scenarioName: "success with API Token",
		clientOptions: &ClientOptions{
			ApiToken: "token",
		},
		expectedRequest: expectedApiRequest{
			path:   "/api/v1/b/bucket",
			method: "GET",
			headers: map[string]string{
				"Authorization": "Bearer token",
			},
		},
		mockResponse: mockApiResponse{},
	},
	{
		scenarioName: "failed timeout",
		clientOptions: &ClientOptions{
			Timeout: time.Second * 1,
		},
		delayResponse: time.Second * 2,
		expectedRequest: expectedApiRequest{
			path:   "/api/v1/b/bucket",
			method: "GET",
			headers: map[string]string{
				"Authorization": "",
			},
		},
		mockResponse: mockApiResponse{},
	},
	{
		scenarioName: "success with timeout",
		clientOptions: &ClientOptions{
			Timeout: time.Second * 3,
		},
		delayResponse: time.Second * 2,
		expectedRequest: expectedApiRequest{
			path:   "/api/v1/b/bucket",
			method: "GET",
			headers: map[string]string{
				"Authorization": "",
			},
		},
		mockResponse: mockApiResponse{},
	},
}

func TestNewIntegration(t *testing.T) {
	// Create a mock HTTP server
	for _, sc := range scenarioNewIntegration {
		t.Run(sc.scenarioName, func(t *testing.T) {
			server := MockAPIServer(t, sc.expectedRequest, sc.mockResponse, sc.delayResponse)
			defer server.Close()

			integration, err := NewIntegration(server.URL, sc.clientOptions)
			assert.NoError(t, err)
			assert.NotNil(t, integration)
			_, err = integration.clientService.Operations.GetAPIV1BBucketName(operations.NewGetAPIV1BBucketNameParams().WithBucketName("bucket"))
			// currently only negative scenario for timeout
			if err != nil && !errors.Is(err, context.DeadlineExceeded) {
				t.Fatalf("expected context deadline exceeded, got: %v", err)
			}
		})
	}
}

// Both cases will fail because the certificate from httptest is not signed by a trusted authority.
func TestNewIntegrationTLS(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer server.Close()

	t.Run("failed with verifySSL true", func(t *testing.T) {
		options := &ClientOptions{
			VerifySSL: true,
		}
		integration, err := NewIntegration(server.URL, options)
		assert.NoError(t, err)
		assert.NotNil(t, integration)
		_, err = integration.clientService.Operations.GetAPIV1BBucketName(
			operations.NewGetAPIV1BBucketNameParams().WithBucketName("bucket"))
		assert.NotNil(t, err)
	})
	t.Run("failed without verifySSL option", func(t *testing.T) {
		integration, err := NewIntegration(server.URL, nil)
		assert.NoError(t, err)
		assert.NotNil(t, integration)
		_, err = integration.clientService.Operations.GetAPIV1BBucketName(
			operations.NewGetAPIV1BBucketNameParams().WithBucketName("bucket"))
		assert.NotNil(t, err)
	})
}

type expectedApiRequest struct {
	method  string
	path    string
	body    string
	headers map[string]string
}

type mockApiResponse struct {
	body       string
	statusCode int
}

func MockAPIServer(
	t *testing.T,
	expectedRequest expectedApiRequest,
	mockResponse mockApiResponse,
	delayResponse time.Duration,
) *httptest.Server {
	t.Helper()

	// Start mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify method
		<-time.After(delayResponse)
		if r.Method != expectedRequest.method {
			t.Errorf("Expected method %s, got %s", expectedRequest.method, r.Method)
		}

		// Verify path
		if r.URL.Path != expectedRequest.path {
			t.Errorf("Expected path %s, got %s", expectedRequest.path, r.URL.Path)
		}

		// Verify headers
		for key, expectedValue := range expectedRequest.headers {
			if got := r.Header.Get(key); got != expectedValue {
				t.Errorf("Expected header %s=%s, got %s", key, expectedValue, got)
			}
		}

		// Verify body if expected
		if expectedRequest.body != "" {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Failed reading body: %v", err)
			}
			defer r.Body.Close()

			bodyString := strings.TrimSpace(string(bodyBytes))
			if bodyString != expectedRequest.body {
				t.Errorf("Expected body %q, got %q", expectedRequest.body, bodyString)
			}
		}
		if mockResponse.statusCode == 0 {
			return
		}

		// Write mock response
		w.WriteHeader(mockResponse.statusCode)
		_, _ = w.Write([]byte(mockResponse.body))
	}))
	return server
}
