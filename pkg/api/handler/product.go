package handler

import (
	"fmt"
	"glamgrove/pkg/domain"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"net/http"

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
