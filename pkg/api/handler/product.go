package handler

// import (
// 	"glamgrove/pkg/api/middleware"
// 	"glamgrove/pkg/domain"
// 	services "glamgrove/pkg/usecase/interfaces"
// 	"glamgrove/pkg/utils"
// 	"glamgrove/pkg/utils/request"
// 	"glamgrove/pkg/utils/response"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type ProductHandler struct {
// 	productUsecase services.ProductUseCase
// }

// func NewProductHandler(usecase services.ProductUseCase) *ProductHandler {
// 	return &ProductHandler{
// 		productUsecase: usecase,
// 	}
// }

// func (ph *ProductHandler) SaveProduct(c *gin.Context) {
// 	var product domain.Product

// 	//get id from getid
// 	id, err := middleware.GetId(c, "Admin_Authorization")
// 	if err != nil {
// 		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), product)
// 		c.JSON(400, response)
// 		return
// 	}

// 	product.AdminId = uint(id)

// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		response := response.ErrorResponse(400, "error entering details", err.Error(), product)
// 		c.JSON(400, response)
// 		return
// 	}

// 	product.ProductCode = utils.GenerateProductCode(6)

// 	productDetails, err := ph.productUsecase.AddProduct(c, product)
// 	if err != nil {
// 		response := response.ErrorResponse(400, "can't add product", err.Error(), product)
// 		c.JSON(400, response)
// 		return
// 	}
// 	response := response.SuccessResponse(200, "successfully added product", productDetails)
// 	c.JSON(200, response)
// }

// // to get all products

// func (ph *ProductHandler) GetAllProducts(c *gin.Context) {
// 	page, err := strconv.Atoi(c.Query("page"))
// 	if err != nil {
// 		response := response.ErrorResponse(400, "Please add page number as params", err.Error(), "")
// 		c.JSON(400, response)
// 	}
// 	pagesize, err := strconv.Atoi(c.Query("pagesize"))
// 	if err != nil {
// 		response := response.ErrorResponse(400, "Please add pages size as params", err.Error(), "")
// 		c.JSON(400, response)
// 	}
// 	pagination := utils.Pagination{
// 		Page:     page,
// 		PageSize: pagesize,
// 	}
// 	product, metadata, err := ph.productUsecase.FindAllProducts(c, pagination)
// 	if err != nil {
// 		response := response.ErrorResponse(400, "error while finding products", err.Error(), product)
// 		c.JSON(400, response)
// 		return
// 	}
// 	response := response.SuccessResponse(200, "successfully displayed all products", product, metadata)
// 	c.JSON(200, response)
// }

// //search product

// func (ph *ProductHandler) SearchProduct(c *gin.Context) {

// 	var code request.CodeRequest
// 	product_code := c.Query("code")
// 	code.Code = product_code

// 	product, err := ph.productUsecase.SearchByCode(c, code.Code)
// 	if err != nil {
// 		response := response.ErrorResponse(400, "error while finding products", err.Error(), product)
// 		c.JSON(400, response)
// 		return
// 	}
// 	response := response.SuccessResponse(200, "successfully displayed product details", product)
// 	c.JSON(200, response)

// }
