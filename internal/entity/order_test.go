package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfIdBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestIfpriceIsBlank(t *testing.T) {
	order := Order{
		ID: "1",
	}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestIfTaxIsBlank(t *testing.T) {
	order := Order{
		ID:    "1",
		price: 1,
	}
	assert.Error(t, order.CalculateFinalprice(), "price must be greater than zero")
}

func TestFinalprice(t *testing.T) {
	order := Order{
		ID:    "1",
		price: 1.0,
		tax:   1.0,
	}
	order.CalculateFinalprice()
	assert.Equal(t, 2.0, order.finalp)
}
func TestConstructor(t *testing.T) {
	order, err := Constructor("1", 1, 1)
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, 1.0, order.price)
	assert.Equal(t, 1.0, order.tax)
}
