package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIfIdBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestIfpriceIsBlank(t *testing.T){
	order := Order{
		ID: "1",
	}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestIfTaxIsBlank(t *testing.T){
	order := Order{
		ID: "1",
		price: 1,
	}
	assert.Error(t, order.CalculateFinalprice(), "price must be greater than zero")
}
	