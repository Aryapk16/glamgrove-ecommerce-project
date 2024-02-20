package usecase

import (
	"context"
	"fmt"
	interfaces "glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"
)

type imageUseCase struct {
	imageRepository interfaces.ImageRepository
}

func NewImageUseCase(ImageRepo interfaces.ImageRepository) service.ImageService {
	return &imageUseCase{imageRepository: ImageRepo}
}

func (c *imageUseCase) CropImage(ctx context.Context, productImageId int) (string, error) {

	imageUrl, err := c.imageRepository.GetImageUrl(ctx, productImageId)
	fmt.Println("image url is ", imageUrl)
	if err != nil {
		return "", err
	}
	return imageUrl, nil

}
