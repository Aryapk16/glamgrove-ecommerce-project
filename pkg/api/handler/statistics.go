package handler

import (
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

// type ProductHandler struct {
// 	ProductService service.ProductService
// }

func (pd *ProductHandler) Statistics(c *gin.Context) {
	var sales request.Sales
	sales.Sdate = c.Query("startDate")
	sales.Edate = c.Query("endDate")

	sDate, err := utils.StringToTime(sales.Sdate)
	if err != nil {
		response := response.ErrorResponse(400, "Please add start date as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	eDate, err := utils.StringToTime(sales.Edate)
	if err != nil {
		response := response.ErrorResponse(400, "Please add end date as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	salesData, err := pd.ProductService.SalesData(sDate, eDate)
	if err != nil {
		response := response.ErrorResponse(400, "Can't calulate details of sales ", err.Error(), salesData)
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully displayed sales data", salesData)
	c.JSON(200, response)

}
