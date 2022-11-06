package file

import (
	"context"
	"io"
)

type UploadFileClient interface {
	UploadToBucket(ctx context.Context, bucket string, fileName string, file io.Reader) (*string, error)
}
