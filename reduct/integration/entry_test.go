package integration

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mhmmdihza/reduct-go/reduct/integration/client/operations"
	"github.com/stretchr/testify/assert"
)

type bytesBufferStringWithClose struct {
	*bytes.Buffer
}

func (*bytesBufferStringWithClose) Close() error {
	return nil
}
func NewBytesBufferStringWithClose(str string) *bytesBufferStringWithClose {
	return &bytesBufferStringWithClose{
		bytes.NewBufferString(str),
	}
}

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
		err = integration.WriteEntry(NewBytesBufferStringWithClose("asd"), "bucket_name", "entry_name", 5000, 3, map[string]string{"1": "a", "2": "b"})
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
		err = integration.WriteEntry(NewBytesBufferStringWithClose("zzz"), "bucket_name", "entry_name", 5000, 3, nil)
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
		err = integration.WriteEntry(NewBytesBufferStringWithClose("zzz"), "bucket_name", "entry_name", 5000, 3, nil)
		assert.NotNil(t, err)
		assert.EqualError(t, &operations.PostAPIV1BBucketNameEntryNameUnprocessableEntity{}, err.Error())
	})
	t.Run("success write entry with large file", func(t *testing.T) {
		var expectedSize int64 = 10485760
		server := mockAPIServerForWriteStream(t, expectedSize)
		defer server.Close()
		integration, err := NewIntegration(server.URL, nil)
		assert.NoError(t, err)
		assert.NotNil(t, integration)
		err = integration.WriteEntry(NewBytesBufferStringWithClose(("A" + strings.Repeat("B", 10*1024*1024-2) + "Z")), "bucket_name", "entry_name", 5000, 10485760, nil)
		assert.NoError(t, err)
	})
}

func mockAPIServerForWriteStream(t *testing.T, expectedSize int64) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		totalBytes := 0
		firstByte := ""
		lastByte := ""
		for {
			buf := make([]byte, 1024)
			n, err := r.Body.Read(buf)
			if err == io.EOF {
				totalBytes += n
				lastByte = string(buf[n-1])
				break
			}
			if err != nil {
				http.Error(w, "Error reading body", http.StatusInternalServerError)
				return
			}
			if firstByte == "" {
				firstByte = string(buf[0])
			}
			totalBytes += n
		}
		assert.Equal(t, expectedSize, int64(totalBytes))
		assert.Equal(t, "A", firstByte)
		assert.Equal(t, "Z", lastByte)

		w.WriteHeader(http.StatusOK)
	}))
	return server
}
