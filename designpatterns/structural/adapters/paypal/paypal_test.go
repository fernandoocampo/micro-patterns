package paypal_test

import (
	"fmt"
	"testing"

	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/adapters/paypal"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/adapters/sales/payments"
	"github.com/stretchr/testify/assert"
)

func TestPay(t *testing.T) {
	t.Parallel()
	// Given
	var expectedError error

	expectedPaymentResult := payments.PaymentResponse{
		TransactionID: "123",
	}
	paymentData := payments.PaymentParameters{
		TransactionID: "123",
	}
	paymentProvider := paypal.NewClient()
	// When
	got, err := doPayment(t, paymentProvider, paymentData)
	// Then
	assert.Equal(t, expectedError, err)
	assert.Equal(t, &expectedPaymentResult, got)
}

func doPayment(t *testing.T, paymentProvider PaymentAdapter, paymentData payments.PaymentParameters) (*payments.PaymentResponse, error) {
	t.Helper()

	result, err := paymentProvider.Pay(paymentData)
	if err != nil {
		err = fmt.Errorf("unexpecte error %w", err)
	}

	return result, err
}

type PaymentAdapter interface {
	Pay(payments.PaymentParameters) (*payments.PaymentResponse, error)
}
