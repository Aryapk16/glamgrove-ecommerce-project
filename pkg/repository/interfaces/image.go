package interfaces

import "context"

type ImageRepository interface {
	GetImageUrl(c context.Context, productImageID int) (string, error)
}
