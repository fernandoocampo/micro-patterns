package externalpaypal

type PaypalPayment struct {
	TransactionID string
}

type PaypalPaymentResponse struct {
	TransactionID string
}

type PaypalExternalAPI struct {
}

func (p *PaypalExternalAPI) DoTransaction(info PaypalPayment) (PaypalPaymentResponse, error) {
	// do something to process the payment
	return PaypalPaymentResponse(info), nil
}
