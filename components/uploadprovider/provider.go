package uploadprovider

import (
	"RESTaurant_v2/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, destination string) (*common.Image, error)
}
