package paypal

import (
	"errors"
	"log"

	external "github.com/fernandoocampo/micro-patterns/designpatterns/structural/adapters/paypal/externalpaypal"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/adapters/sales/payments"
)

// Client paypal client
type Client struct {
	paypal *external.PaypalExternalAPI
}

// NewClient create a new paypal client.
func NewClient() *Client {
	newClient := Client{
		paypal: &external.PaypalExternalAPI{},
	}
	return &newClient
}

func (c *Client) Pay(data payments.PaymentParameters) (*payments.PaymentResponse, error) {
	paypalParams := external.PaypalPayment{
		TransactionID: data.TransactionID,
	}
	paypalResult, err := c.paypal.DoTransaction(paypalParams)
	if err != nil {
		log.Println("unexpected error paying with paypal", err)
		return nil, errors.New("cannot process payment")
	}
	result := payments.PaymentResponse{
		TransactionID: paypalResult.TransactionID,
	}
	return &result, nil
}
