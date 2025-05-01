package integration

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/mhmmdihza/reduct-go/reduct/integration/client/operations"
)

func (i *Integration) WriteEntry(data io.ReadCloser, bucketName, entryName string, ts, contentLength int64, xReductLabelHeader map[string]string) error {
	_, err := i.clientService.Operations.PostAPIV1BBucketNameEntryName(operations.NewPostAPIV1BBucketNameEntryNameParams().
		WithBody(data).
		WithBucketName(bucketName).
		WithEntryName(entryName).
		WithTs(ts).
		WithContentLength(contentLength),
		func() []operations.ClientOption {
			if len(xReductLabelHeader) == 0 {
				return make([]operations.ClientOption, 0)
			}
			return []operations.ClientOption{func(co *runtime.ClientOperation) {
				w := co.Params
				co.Params = runtime.ClientRequestWriterFunc(
					func(cr runtime.ClientRequest, r strfmt.Registry) error {
						if err := w.WriteToRequest(cr, r); err != nil {
							return err
						}
						for k, v := range xReductLabelHeader {
							if err := cr.SetHeaderParam(fmt.Sprintf("x-reduct-label-%s", k), v); err != nil {
								return err
							}
						}
						return nil
					})
			}}
		}()...)
	if err != nil {
		return err
	}
	return nil
}
