package integration

import (
	"testing"
	"time"

	"github.com/mhmmdihza/reduct-go/reduct/integration/client/operations"
	"github.com/stretchr/testify/assert"
)

func TestCreateBucket(t *testing.T) {
	t.Run("success creating bucket", func(t *testing.T) {
		server := MockAPIServer(t, expectedApiRequest{
			method: "POST",
			path:   "/api/v1/b/bucket",
		}, mockApiResponse{
			statusCode: 200,
		}, time.Millisecond)
		defer server.Close()

		integration, err := NewIntegration(server.URL, nil)
		assert.NoError(t, err)
		assert.NotNil(t, integration)
		err = integration.CreateBucket("bucket")
		assert.NoError(t, err)
	})
	t.Run("conflicted bucket", func(t *testing.T) {
		server := MockAPIServer(t, expectedApiRequest{
			method: "POST",
			path:   "/api/v1/b/bucket",
		}, mockApiResponse{
			statusCode: 409,
		}, time.Millisecond)
		defer server.Close()

		integration, err := NewIntegration(server.URL, nil)
		assert.NoError(t, err)
		assert.NotNil(t, integration)
		err = integration.CreateBucket("bucket")
		assert.NotNil(t, err)
		assert.EqualError(t, &operations.PostAPIV1BBucketNameConflict{}, err.Error())
	})
}
