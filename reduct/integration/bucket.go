package integration

import (
	"github.com/mhmmdihza/reduct-go/reduct/integration/client/operations"
)

func (i *Integration) CreateBucket(bucketName string) error {
	_, err := i.clientService.PostAPIV1BBucketName(operations.NewPostAPIV1BBucketNameParams().
		WithBucketName(bucketName))
	if err != nil {
		return err
	}
	return nil
}
