package handler

import (
	"fmt"
	"glamgrove/pkg/domain"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(ProductUseCase service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: ProductUseCase,
	}
}

// AddCategory godoc
// @Summary Add a new category
// @Description Adds a new category to the database.
// @Tags Categories
// @Accept json
// @Produce json
// @Param body body request.Category true "Category object to add"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Success message"
// @Failure 400 {object} response.Response{} "Missing or invalid entry or failed to add category"
// @Router  /admin/brands/add  [post]
func (p *ProductHandler) AddCategory(c *gin.Context) {
	var ProductBrand request.Category
	if err := c.ShouldBindJSON(&ProductBrand); err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "Missing or invalid entry", err.Error(), ProductBrand)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := p.ProductService.AddCategory(c, ProductBrand)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to add category", err.Error(), ProductBrand)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly added a new category in database", ProductBrand)
	c.JSON(200, response)
}

// GetAllCategory godoc
// @Summary Get all categories
// @Description Retrieves all categories from the database.
// @Tags Categories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Success message with all categories"
// @Failure 500 {object} response.Response{} "Failed to get categories"
// @Router /admin/brands/get  [get]
// @Router /products/brands  [get]
func (p *ProductHandler) GetAllCategory(c *gin.Context) {
	allcategories, err := p.ProductService.GetAllBrands(c)
	if err != nil {
		response := response.ErrorResponse(500, "Failed to get category", err.Error(), allcategories)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	fmt.Println(allcategories)
	response := response.SuccessResponse(200, "Successfuly listed all category", allcategories)
	c.JSON(200, response)
}

// AddProduct godoc
// @Summary Add a new product
// @Description Adds a new product to the database.
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body request.ProductReq true "Product details"
// @Success 200 {object} response.Response{} "Success message with product details"
// @Failure 400 {object} response.Response{} "Failed to add product"
// @Router /admin/brands/add  [post]
func (p *ProductHandler) AddProduct(c *gin.Context) {
	var body request.ProductReq
	if err := c.ShouldBindJSON(&body); err != nil {
		responce := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, responce)
		return
	}
	var product domain.Product
	copier.Copy(&product, body)
	if err := p.ProductService.AddProduct(c, product); err != nil {
		response := response.ErrorResponse(400, "failed to add product", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(http.StatusOK, "Product added successful", body)
	c.JSON(http.StatusOK, response)
}

// AddImage godoc
// @Summary Add image to a product
// @Description Adds image(s) to a product.
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param product_id formData int true "Product ID"
// @Param image formData file true "Image file(s)"
// @Success 200 {object} response.Response{} "Success message with added images"
// @Failure 400 {object} response.Response{} "Error message with details"
// @Router  /admin/products/addimage [post]

func (p *ProductHandler) AddImage(c *gin.Context) {
	pid, err := strconv.Atoi(c.PostForm("product_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Error while fetching product_id", err.Error(), pid)
		c.JSON(400, response)
		return
	}
	form, err := c.MultipartForm()

	if err != nil {
		response := response.ErrorResponse(400, "error while fetching image file", err.Error(), form)
		c.JSON(400, response)
		return
	}
	files := form.File["image"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image file found"})
		return
	}
	Images, err := p.ProductService.AddImage(c, pid, files)
	if err != nil {
		response := response.ErrorResponse(400, "can't add images", err.Error(), Images)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully added images", Images)
	c.JSON(200, response)

}

// AddItemImage godoc
// @Summary Add image to a product item
// @Description Adds image(s) to a specific product item.
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param product_item_id formData int true "Product Item ID"
// @Param image formData file true "Image file(s)"
// @Success 200 {object} response.Response{} "Success message with added images"
// @Failure 400 {object} response.Response{} "Error message with details"
// @Router  /admin/products/additemimage [post]
func (p *ProductHandler) AddItemImage(c *gin.Context) {
	pid, err := strconv.Atoi(c.PostForm("product_item_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Error while fetching product_item_id", err.Error(), pid)
		c.JSON(400, response)
		return
	}
	form, err := c.MultipartForm()

	if err != nil {
		response := response.ErrorResponse(400, "error while fetching image file", err.Error(), form)
		c.JSON(400, response)
		return
	}
	files := form.File["image"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image file found"})
		return
	}
	Images, err := p.ProductService.AddImage(c, pid, files)
	if err != nil {
		response := response.ErrorResponse(400, "can't add images", err.Error(), Images)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully added images", Images)
	c.JSON(200, response)

}

// ListProducts godoc
// @Summary List products
// @Description Retrieves a list of products with pagination support.
// @Tags Products
// @Accept json
// @Produce json
// @Param count query int false "Number of products to retrieve per page"
// @Param page_number query int false "Page number"
// @Success 200 {object} response.Response{} "Successful response"
// @Success 200 {object} response.Response{} "No products to show"
// @Failure 400 {object} response.Response{} "Invalid inputs"
// @Failure 500 {object} response.Response{} "Failed to get all products"
// @Router /products/   [get]
// @Router  /admin/products/list [get]
func (p *ProductHandler) ListProducts(c *gin.Context) {
	count, err1 := utils.StringToUint(c.Query("count"))
	if err1 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pageNumber, err2 := utils.StringToUint(c.Query("page_number"))
	if err2 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}
	products, err := p.ProductService.GetProducts(c, pagination)
	if err != nil {
		response := response.ErrorResponse(500, "failed to get all products", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if products == nil {
		response := response.SuccessResponse(200, "Oops ! no products to show", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	respones := response.SuccessResponse(200, "Product listed successfuly", products)
	c.JSON(http.StatusOK, respones)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Updates an existing product with the provided details.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID to update"
// @Param body body request.UpdateProductReq true "Product details to update"
// @Success 200 {object} response.Response{} "Product updated successfully"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
// @Failure 400 {object} response.Response{} "Failed to update product"
// @Router   /admin/products/update [put]
func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	var body request.UpdateProductReq
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var product domain.Product
	copier.Copy(&product, &body)
	err := p.ProductService.UpdateProduct(c, product)
	if err != nil {
		response := response.ErrorResponse(400, "failed to update product", err.Error(), body)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Product updated successful", body)
	c.JSON(200, response)
	c.Abort()
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Deletes the product with the specified ID.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID to delete"
// @Success 200 {object} response.Response{} "Product deleted successfully"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
// @Failure 500 {object} response.Response{} "Failed to delete product"
// @Router  /admin/products/delete [delete]
func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	var body request.DeleteProductReq
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	productID := body.ID
	deletedProduct, err := p.ProductService.DeleteProduct(c, productID)
	if err != nil {

		response := response.ErrorResponse(500, "Failed to delete product", err.Error(), body)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := response.SuccessResponse(http.StatusOK, "Successfuly deleted product", deletedProduct)
	c.JSON(200, response)
}

// AddProductItem godoc
// @Summary Add a product item
// @Description Adds a new product item.
// @Tags Products
// @Accept json
// @Produce json
// @Param body body request.ProductItemReq true "Product item details"
// @Success 200 {object} response.Response{} "Product item added successfully"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
// @Failure 500 {object} response.Response{} "Failed to add product item"
// @Router /admin/products/product-item [post]
func (p *ProductHandler) AddProductItem(c *gin.Context) {
	var body request.ProductItemReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := p.ProductService.AddProductItem(c, body)
	if err != nil {
		response := response.ErrorResponse(500, "failed to add product item", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Product item added successful", body)
	c.JSON(200, response)
	c.Abort()
}

// GetProductItem godoc
// @Summary Get product items by product ID
// @Description Retrieves product items based on the provided product ID.
// @Tags Products
// @Accept json
// @Produce json
// @Param product_id query uint true "Product ID"
// @Success 200 {object} response.Response{} "Product items fetched successfully"
// @Failure 400 {object} response.Response{} "Invalid param input" or "No product items for this product ID"
// @Failure 500 {object} response.Response{} "Failed to get product item for given product ID"
// @Router  /admin/products/product-item/:product_id [get]
func (p *ProductHandler) GetProductItem(c *gin.Context) {
	productID, err := utils.StringToUint(c.Query("product_id"))
	if err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "invalid param input", err.Error(), productID)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	productItems, count, err := p.ProductService.GetProductItem(c, productID)
	// Create a map to combine the count and productItems
	fmt.Println("------------->", productItems)
	data := map[string]interface{}{
		"count":        count,
		"productItems": productItems,
	}
	if err != nil {
		response := response.ErrorResponse(400, "failed to get product item for given product id", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	if count == 0 {
		response := response.ErrorResponse(http.StatusBadRequest, "No product items for this product id", "", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Fetching product item successful and listed below", data)
	c.JSON(200, response)

}
