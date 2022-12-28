package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThesShouldReeiverAnError(t *testing.T) {
	_, err := NewOrder("", 10.1, 1.12)

	assert.NotNil(t, err)
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThesShouldReeiverAnError(t *testing.T) {
	order := Order{Id: "123212", Tax: 100}
	assert.Error(t, order.IsValid(), "invalid price")
}
func TestGivenAnEmptyTax_WhenCreateANewOrder_ThesShouldReeiverAnError(t *testing.T) {
	order := Order{Id: "123212", Price: 100}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAnEmptyTaxParms_WhenCreateANewOrder_ThesShouldReeiverCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123212", 100, 10)
	assert.Nil(t, err)
	assert.Equal(t, order.Id, "123212")
	assert.Equal(t, order.Price, 100.0)
	assert.Equal(t, order.Tax, 10.0)
}

func TestGivenPriceAndTax_WhenCallCalculateTax_ThesShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123212", 100, 10)
	if err != nil {
		errors.New("Intance is wrong")
	}
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, order.FinalPrice, 110.0)
}
