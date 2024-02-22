package handler

import (
	"fmt"
	service "glamgrove/pkg/usecase/interfaces"
	"image"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	imageService service.ImageService
}

func NewImageHandler(ImageUseCase service.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: ImageUseCase,
	}
}
func (c *ImageHandler) CropImage(ctx *gin.Context) {

	imageId := ctx.Query("product_image_id")
	if imageId == "" {
		ctx.JSON(400, gin.H{"error": "product_image_id is required"})
		return
	}

	imageID, err := strconv.Atoi(imageId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid product_image_id"})
		return
	}
	fmt.Println("id", imageID)

	imageUrl, err := c.imageService.CropImage(ctx, imageID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to crop image", "details": err.Error()})
		return
	}
	fmt.Println("url", imageUrl)

	imageUrl = filepath.Join("/home/arya-pk/Documents/MyProject/GlamGrove/images", imageUrl)

	inputImage, err := imaging.Open(imageUrl)
	if err != nil {
		// fmt.Println("++++++++++", inputImage)
		ctx.JSON(500, gin.H{"error": "Failed to open image", "details": err.Error()})
		return
	}

	filepath.Join()

	cropRect := image.Rect(100, 100, 400, 400)

	croppedImage := imaging.Crop(inputImage, cropRect)

	err = imaging.Save(croppedImage, imageUrl)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to save image", "details": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Image cropped and saved successfully"})
}
