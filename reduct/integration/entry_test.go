package integration

import (
	"testing"
	"time"

	"github.com/mhmmdihza/reduct-go/reduct/integration/client/operations"
	"github.com/stretchr/testify/assert"
)

func TestWriteEntry(t *testing.T) {
	t.Run("success write entry", func(t *testing.T) {
		server := MockAPIServer(t, expectedApiRequest{
			method:     "POST",
			path:       "/api/v1/b/bucket_name/entry_name",
			body:       "asd",
			headers:    map[string]string{"Content-length": "3", "X-Reduct-Label-1": "a", "X-Reduct-Label-2": "b"},
			queryParam: map[string]string{"ts": "5000"},
		}, mockApiResponse{
			statusCode: 200,
		}, time.Millisecond)
		defer server.Close()

		integration, err := NewIntegration(server.URL, nil)
		assert.NoError(t, err)
		assert.NotNil(t, integration)
		err = integration.WriteEntry([]byte("asd"), "bucket_name", "entry_name", 5000, 3, map[string]string{"1": "a", "2": "b"})
		assert.NoError(t, err)
	})
	t.Run("success write entry without x-reduct-label-* headers", func(t *testing.T) {
		server := MockAPIServer(t, expectedApiRequest{
			method:     "POST",
			path:       "/api/v1/b/bucket_name/entry_name",
			body:       "zzz",
			headers:    map[string]string{"Content-length": "3"},
			queryParam: map[string]string{"ts": "5000"},
		}, mockApiResponse{
			statusCode: 200,
		}, time.Millisecond)
		defer server.Close()

		integration, err := NewIntegration(server.URL, nil)
		assert.NoError(t, err)
		assert.NotNil(t, integration)
		err = integration.WriteEntry([]byte("zzz"), "bucket_name", "entry_name", 5000, 3, nil)
		assert.NoError(t, err)
	})
	t.Run("unprocessable entity write entry", func(t *testing.T) {
		server := MockAPIServer(t, expectedApiRequest{
			method:     "POST",
			path:       "/api/v1/b/bucket_name/entry_name",
			body:       "zzz",
			headers:    map[string]string{"Content-length": "3"},
			queryParam: map[string]string{"ts": "5000"},
		}, mockApiResponse{
			statusCode: 422,
		}, time.Millisecond)
		defer server.Close()

		integration, err := NewIntegration(server.URL, nil)
		assert.NoError(t, err)
		assert.NotNil(t, integration)
		err = integration.WriteEntry([]byte("zzz"), "bucket_name", "entry_name", 5000, 3, nil)
		assert.NotNil(t, err)
		assert.EqualError(t, &operations.PostAPIV1BBucketNameEntryNameUnprocessableEntity{}, err.Error())
	})
}
