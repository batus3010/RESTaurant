package uploadBiz

import (
	"RESTaurant_v2/common"
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png" // important: add this to decode png files
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, destination string) (*common.Image, error)
}

type uploadBiz struct {
	provider UploadProvider
}

func NewUploadBiz(provider UploadProvider) *uploadBiz {
	return &uploadBiz{provider: provider}
}

// Upload return a pointer to Image for transport layer to use
func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, errors.New("file is not image")
	}

	if strings.TrimSpace(fileName) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, errors.New("cannot save image file, " + err.Error())
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

// get width, height, and also check if the bytes provided is an image
func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
