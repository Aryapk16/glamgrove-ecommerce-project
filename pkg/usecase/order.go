package usecase

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/config"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
)

type OrderUseCase struct {
	OrderRepository interfaces.OrderRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository) service.OrderService {
	return &OrderUseCase{
		OrderRepository: repo,
	}
}
func (o *OrderUseCase) GetTotalAmount(c context.Context, userid uint) (float64, error) {
	var total_amount float64
	total_amount = 0
	cart, err := o.OrderRepository.GetTotalAmount(c, int(userid))
	if err != nil {
		return 0, err
	}

	for _, c := range cart {
		total_amount = total_amount + float64(c.Total)
	}
	return total_amount, nil
}
func (o *OrderUseCase) CreateOrder(c context.Context, order domain.Order) (response.OrderResponse, error) {
	//Checking whether the payment id exist
	_, err := o.OrderRepository.FindPaymentMethodById(c, uint(order.PaymentMethodID))

	if err != nil {
		return response.OrderResponse{}, errors.New("payment method doesn't exists")
	}
	orderresp, err := o.OrderRepository.CreateOrder(c, order)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderresp, nil
}

func (o *OrderUseCase) UpdateOrderDetails(c context.Context, uporder request.UpdateOrder) (response.OrderResponse, error) {
	//Checking whether the payment id exist
	_, err := o.OrderRepository.FindPaymentMethodById(c, uporder.PaymentMethodID)

	if err != nil {
		return response.OrderResponse{}, errors.New("payment method doesn't exists")
	}
	orderup, err := o.OrderRepository.UpdateOrderDetails(c, uporder)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderup, nil
}
func (o *OrderUseCase) ListAllOrders(c context.Context, page request.ReqPagination, userId uint) (orders []response.OrderResponse, err error) {

	return o.OrderRepository.ListAllOrders(c, page, userId)

}

// List order for admin
func (o *OrderUseCase) GetAllOrders(c context.Context, page request.ReqPagination) (orders []response.OrderResponse, err error) {

	return o.OrderRepository.GetAllOrders(c, page)

}
func (o *OrderUseCase) DeleteOrder(c context.Context, order_id uint) error {
	err := o.OrderRepository.DeleteOrder(c, order_id)
	if err != nil {
		return err
	}
	return nil
}
func (o *OrderUseCase) PlaceOrder(c context.Context, order domain.Order) (response.PaymentResponse, error) {
	err1 := o.OrderRepository.FindOrder(c, order)
	if err1 != nil {
		return response.PaymentResponse{}, err1
	}
	method_Id, err := o.OrderRepository.FindPaymentMethodIdByOrderId(c, order.Order_Id)
	if err != nil {
		return response.PaymentResponse{}, err
	}
	if method_Id == 1 {
		order.Order_Status = "order confirmed"
	} else {
		order.Order_Status = "order confirmed payment pending"
	}
	paymentresp, err := o.OrderRepository.PlaceOrder(c, order)
	if err != nil {
		return response.PaymentResponse{}, err
	}
	return paymentresp, nil
}

func (o *OrderUseCase) FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error) {
	method_id, err := o.OrderRepository.FindPaymentMethodIdByOrderId(c, order_id)
	if err != nil {
		return 0, err
	}
	return method_id, nil
}

func (o *OrderUseCase) UpdateOrderStatus(c context.Context, order_id uint) (response.OrderResponse, error) {
	order_status := "Order confirmed"
	orderResp, err := o.OrderRepository.UpdateOrderStatus(c, order_id, order_status)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderResp, nil
}

func (o *OrderUseCase) FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error) {
	totalAmount, err := o.OrderRepository.FindTotalAmountByOrderId(c, order_id)
	if err != nil {
		return 0, err
	}
	return totalAmount, nil
}

func (o *OrderUseCase) FindPhoneEmailByUserId(c context.Context, usr_id int) (response.PhoneEmailResp, error) {
	phnEmail, err := o.OrderRepository.FindPhoneEmailByUserId(c, usr_id)
	if err != nil {
		return response.PhoneEmailResp{}, err
	}
	return phnEmail, nil
}
func (o *OrderUseCase) ReturnRequest(c context.Context, returnOrder domain.OrderReturn) (response.ReturnResponse, error) {
	returnResp, err := o.OrderRepository.ReturnRequest(c, returnOrder)
	if err != nil {
		return response.ReturnResponse{}, err
	}
	return returnResp, nil
}

func (o *OrderUseCase) GetRazorpayOrder(c context.Context, userID uint, razorPay request.RazorPayReq) (response.ResRazorpayOrder, error) {
	var razorpayOrder response.ResRazorpayOrder
	//fmt.Println(razorPay)
	//razorpay amount is caluculate on pisa for india so make the actual price into paisa
	razorPayAmount := uint(razorPay.Total_Amount * 100)

	razopayOrderId, err := utils.GenerateRazorpayOrder(razorPayAmount, "test reciept")
	if err != nil {
		return razorpayOrder, err
	}
	fmt.Println(razopayOrderId)
	// set all details on razopay order
	razorpayOrder.AmountToPay = uint(razorPay.Total_Amount)

	razorpayOrder.RazorpayKey, _ = config.GetRazorPayConfig()

	razorpayOrder.UserID = userID
	razorpayOrder.RazorpayOrderID = razopayOrderId

	razorpayOrder.Email = razorPay.Email
	razorpayOrder.Phone = razorPay.Phone

	return razorpayOrder, nil
}

func (o *OrderUseCase) UpdateStatusRazorpay(c context.Context, order_id uint) (response.OrderResponse, error) {
	order_status := "Order confirmed"
	payment_status := "Payment Done"
	delivery_status := "Order delivered successfully"
	orderResp, err := o.OrderRepository.UpdateStatusRazorpay(c, order_id, order_status, payment_status, delivery_status)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderResp, nil
}

func (o *OrderUseCase) SalesReport(c context.Context, page request.ReqPagination, salesData request.ReqSalesReport) ([]response.SalesReport, error) {
	salesReport, err := o.OrderRepository.SalesReport(c, page, salesData)
	if err != nil {
		return []response.SalesReport{}, err
	}
	return salesReport, nil
}

func (o *OrderUseCase) VerifyOrderID(c context.Context, id uint, orderid uint) error {
	err := o.OrderRepository.VerifyOrderID(c, id, orderid)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderUseCase) GetAllPendingReturnRequest(c context.Context, page request.ReqPagination) ([]response.ReturnRequests, error) {

	returnRequests, err := o.OrderRepository.GetAllPendingReturnOrder(c, page)
	if err != nil {
		return returnRequests, err
	}

	return returnRequests, nil

}

//.............

// func (or *OrderUseCase) PrintInvoice(orderId int) (*gofpdf.Fpdf, error) {

// 	if orderId < 1 {
// 		return nil, errors.New("enter a valid order id")
// 	}

// 	order, err := or.OrderRepository.GetDetailedOrderThroughId(orderId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	items, err := or.OrderRepository.GetItemsByOrderId(orderId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("order details ", order)
// 	fmt.Println("itemssss", items)
// 	fmt.Println("order status", order.OrderStatus)
// 	if order.OrderStatus != "DELIVERED" {
// 		return nil, errors.New("wait for the invoice until the product is received")
// 	}

// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()

// 	pdf.SetFont("Arial", "B", 30)
// 	pdf.SetTextColor(31, 73, 125)
// 	pdf.Cell(0, 20, "Invoice")
// 	pdf.Ln(20)

// 	pdf.SetFont("Arial", "I", 14)
// 	pdf.SetTextColor(51, 51, 51)
// 	pdf.Cell(0, 10, "Customer Details")
// 	pdf.Ln(10)
// 	customerDetails := []string{
// 		"Name: " + order.Name,
// 		"House Name: " + order.HouseName,
// 		"Street: " + order.Street,
// 		"State: " + order.State,
// 		"City: " + order.City,
// 	}
// 	for _, detail := range customerDetails {
// 		pdf.Cell(0, 10, detail)
// 		pdf.Ln(10)
// 	}
// 	pdf.Ln(10)

// 	pdf.SetFont("Arial", "B", 16)
// 	pdf.SetFillColor(217, 217, 217)
// 	pdf.SetTextColor(0, 0, 0)
// 	pdf.CellFormat(40, 10, "Item", "1", 0, "C", true, 0, "")
// 	pdf.CellFormat(40, 10, "Price", "1", 0, "C", true, 0, "")
// 	pdf.CellFormat(40, 10, "Quantity", "1", 0, "C", true, 0, "")
// 	pdf.CellFormat(40, 10, "Total Price", "1", 0, "C", true, 0, "")
// 	pdf.Ln(10)

// 	pdf.SetFont("Arial", "", 12)
// 	pdf.SetFillColor(255, 255, 255)
// 	for _, item := range items {
// 		pdf.CellFormat(40, 10, item.ProductName, "1", 0, "L", true, 0, "")
// 		pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(item.Price, 'f', 2, 64), "1", 0, "C", true, 0, "")
// 		pdf.CellFormat(40, 10, strconv.Itoa(item.Quantity), "1", 0, "C", true, 0, "")
// 		pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(item.Total, 'f', 2, 64), "1", 0, "C", true, 0, "")
// 		pdf.Ln(10)
// 	}
// 	pdf.Ln(10)

// 	var totalPrice float64
// 	for _, item := range items {
// 		totalPrice += item.Total
// 	}

// 	pdf.SetFont("Arial", "B", 16)
// 	pdf.SetFillColor(217, 217, 217)
// 	pdf.CellFormat(120, 10, "Total Price:", "1", 0, "R", true, 0, "")
// 	pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(totalPrice, 'f', 2, 64), "1", 0, "C", true, 0, "")
// 	pdf.Ln(10)

// 	offerApplied := totalPrice - order.FinalPrice

// 	pdf.SetFont("Arial", "B", 16)
// 	pdf.SetFillColor(217, 217, 217)
// 	pdf.CellFormat(120, 10, "Offer Applied:", "1", 0, "R", true, 0, "")
// 	pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(offerApplied, 'f', 2, 64), "1", 0, "C", true, 0, "")
// 	pdf.Ln(10)

// 	pdf.SetFont("Arial", "B", 16)
// 	pdf.SetFillColor(217, 217, 217)
// 	pdf.CellFormat(120, 10, "Final Amount:", "1", 0, "R", true, 0, "")
// 	pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(order.FinalPrice, 'f', 2, 64), "1", 0, "C", true, 0, "")
// 	pdf.Ln(10)
// 	pdf.SetFont("Arial", "I", 12)
// 	pdf.Cell(0, 10, "Generated by HeadZone India Pvt Ltd. - "+time.Now().Format("2006-01-02 15:04:05"))
// 	pdf.Ln(10)

// 	return pdf, nil
// }
