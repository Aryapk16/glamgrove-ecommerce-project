package interfaces

import "context"

type ImageService interface {
	CropImage(ctx context.Context, productImageId int) (string, error)
}
