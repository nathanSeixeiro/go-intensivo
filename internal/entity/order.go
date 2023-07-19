package entity

import "errors"

// in Go we work without class but with structs
type Order struct {
	ID     string
	Price  float64
	Tax    float64
	Finalp float64
}

func Constructor(id string, price, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

// for the methods we use pointers and references for the struct
// Go don't have try/catch, we catching each error e and process it
func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is required")
	}

	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if o.Tax <= 0 {
		return errors.New("tax must be greater than zero")
	}
	return nil
}

func (o *Order) CalculateFinalprice() error {
	o.Finalp = o.Price + o.Tax

	err := o.Validate()
	if err != nil {
		return err
	}
	return nil
}
