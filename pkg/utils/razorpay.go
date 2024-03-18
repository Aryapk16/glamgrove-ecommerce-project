package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"glamgrove/pkg/config"

	"github.com/razorpay/razorpay-go"
)

func GenerateRazorpayOrder(razorPayAmount uint, recieptIdOptional string) (razorpayOrderID interface{}, err error) {
	// get razor pay key and secret
	razorpayKey, razorpaySecret := config.GetRazorPayConfig()

	//create a razorpay client
	client := razorpay.NewClient(razorpayKey, razorpaySecret)

	data := map[string]interface{}{
		"amount":   razorPayAmount,
		"currency": "INR",
		"receipt":  recieptIdOptional,
	}

	// create an order on razor pay
	razorpayRes, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Println(err)
		return razorpayOrderID, err
	}

	razorpayOrderID = razorpayRes["id"]
	fmt.Println(razorpayOrderID)
	return razorpayOrderID, nil
}

func VeifyRazorpayPayment(razorpayOrderID, razorpayPaymentID, razorpaySignatur string) error {

	razorpayKey, razorpaySecret := config.GetRazorPayConfig()
	fmt.Println(razorpayKey, razorpaySecret)
	//verify signature
	data := razorpayOrderID + "|" + razorpayPaymentID
	h := hmac.New(sha256.New, []byte(razorpaySecret))
	_, err := h.Write([]byte(data))
	if err != nil {
		return errors.New("faild to veify signature")
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(razorpaySignatur)) != 1 {
		return errors.New("razorpay signature not match")
	}

	// then vefiy payment
	client := razorpay.NewClient(razorpayKey, razorpaySecret)

	// fetch payment and vefify
	payment, err := client.Payment.Fetch(razorpayPaymentID, nil, nil)

	if err != nil {
		return err
	}

	// check payment status
	if payment["status"] != "captured" {
		return fmt.Errorf("faild to varify payment \nrazorpay payment with payment_id %v", razorpayPaymentID)
	}

	return nil
}
