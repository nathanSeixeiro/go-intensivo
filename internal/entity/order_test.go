package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIfIdBlank(t *testing.T) {
	//always tests in go start with Test word
	order := Order{}
	if order.Validate() == nil {
		t.Error("ID id required")
	}
}

func TestWithTestifyIdBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is req")
}