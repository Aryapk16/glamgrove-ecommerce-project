package response

import (
	"strings"
	"time"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ErrorResponse(statusCode int, message string, err string, data interface{}) Response {

	spiltedError := strings.Split(err, "\n")
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Errors:     spiltedError,
		Data:       data,
	}
}

func SuccessResponse(statusCode int, message string, data ...interface{}) Response {

	return Response{
		StatusCode: statusCode,
		Message:    message,
		Errors:     nil,
		Data:       data,
	}
}

type SalesReport struct {
	UserID          uint      `json:"user_id"`
	Name            string    `json:"name"`
	Email           string    `json:"email,omitempty"`
	OrderDate       time.Time `json:"order_date,omitempty"`
	OrderTotalPrice float64   `json:"order_total_price,omitempty"`
	OrderStatus     string    `json:"order_status,omitempty"`
	DeliveryStatus  string    `json:"delivery_status,omitempty"`
	PaymentType     string    `json:"payment_type,omitempty"`
	PaymentStatus   string    `json:"payment_status,omitempty"`
}
