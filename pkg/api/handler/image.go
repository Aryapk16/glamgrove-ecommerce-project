package handler

import (
	service "glamgrove/pkg/usecase/interfaces"
	"image"
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
	// Retrieve product_image_id from query parameter
	imageId := ctx.Query("product_image_id")
	if imageId == "" {
		ctx.JSON(400, gin.H{"error": "product_image_id is required"})
		return
	}

	// Convert imageId to integer
	imageID, err := strconv.Atoi(imageId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid product_image_id"})
		return
	}

	// Crop image
	imageUrl, err := c.imageService.CropImage(ctx, imageID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to crop image", "details": err.Error()})
		return
	}

	// Open input image
	inputImage, err := imaging.Open(imageUrl)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to open image"})
		return
	}

	// Define crop rectangle
	cropRect := image.Rect(100, 100, 400, 400) // (x0, y0, x1, y1)

	// Crop image
	croppedImage := imaging.Crop(inputImage, cropRect)

	// Save cropped image
	err = imaging.Save(croppedImage, imageUrl)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to save image"})
		return
	}

	// Respond with success message
	ctx.JSON(200, gin.H{"message": "Image cropped and saved successfully"})
}

// func (c *ImageHandler) Cropimage(ctx *gin.Context) {
// 	imageId := c.Query("product_image_id")
// 	imageID, err := strconv.Atoi(imageId)
// 	if err != nil {
// 		errRes := response.ErrorResponse(500, "error in string conversion", err.Error(), nil)
// 		c.JSON(500, errRes)
// 		return
// 	}
// 	imageUrl, err := usecase.CropImage(imageID)
// 	if err != nil {
// 		errRes := response.ClientResponse(500, "error in cropping", nil, err)
// 		c.JSON(500, errRes)
// 		return
// 	}

// 	inputImage, err := imaging.Open(imageUrl)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to open image"})
// 		return
// 	}

// 	cropRect := image.Rect(100, 100, 400, 400) // (x0, y0, x1, y1)

// 	croppedImage := imaging.Crop(inputImage, cropRect)

// 	err = imaging.Save(croppedImage, imageUrl)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to save image"})
// 		return
// 	}
// 	c.JSON(200, response.SuccessResponse()(200, "Image cropped and saved successfully", nil, nil))
// }
