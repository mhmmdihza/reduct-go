package integration

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/mhmmdihza/reduct-go/reduct/integration/client/operations"
)

type readCloserWithLen struct {
	*bytes.Reader
}

func (r *readCloserWithLen) Close() error {
	return nil
}

func NewReadCloserWithLen(data []byte) (io.ReadSeekCloser, int) {
	reader := bytes.NewReader(data)
	return &readCloserWithLen{reader}, reader.Len()
}

func (i *Integration) WriteEntry(data []byte, bucketName, entryName string, ts, contentLength int64, xReductLabelHeader map[string]string) error {
	_, err := i.clientService.Operations.PostAPIV1BBucketNameEntryName(operations.NewPostAPIV1BBucketNameEntryNameParams().
		WithBody(io.NopCloser(bytes.NewReader(data))).
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
