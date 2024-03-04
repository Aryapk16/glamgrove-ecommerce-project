package handler

import (
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

// Statistics godoc
// @Summary Get sales statistics
// @Description Retrieves sales statistics within the specified date range.
// @Tags Products
// @Accept json
// @Produce json
// @Param startDate query string true "Start date (YYYY-MM-DD)"
// @Param endDate query string true "End date (YYYY-MM-DD)"
// @Success 200 {object} response.Response{} "Sales data fetched successfully"
// @Failure 400 {object} response.Response{} "Please add start date as params" or "Please add end date as params"
// @Failure 400 {object} response.Response{} "Can't calculate details of sales"
// @Router  /admin/dashboard/salesdata [get]
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
