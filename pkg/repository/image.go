package repository

import (
	"context"
	"glamgrove/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type ImageDatabase struct {
	DB *gorm.DB
}

func NewImageRepository(db *gorm.DB) interfaces.ImageRepository {
	return &ImageDatabase{DB: db}
}
func (c *ImageDatabase) GetImageUrl(ctx context.Context, productImageID int) (string, error) {
	var imageUrl string
	if err := c.DB.Raw("select product_image_url from product_images where id = ?", productImageID).Scan(&imageUrl).Error; err != nil {
		return "", err
	}
	return imageUrl, nil
}
