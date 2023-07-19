package entity

import "errors"

type Order struct {
	ID     string
	price  float64
	tax    float64
	finalp float64
}

func Constructor(id string, price, tax float64) (*Order, error){
	order := &Order{
		ID:    id,
		price: price,
		tax:   tax,
	}

	err := order.Validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is required")
	}

	if o.price <= 0 {
		return errors.New("must be greater than zero")
	}

	if o.tax <= 0 {
		return errors.New("must be greater than zero")
	}
	return nil
}

func (o *Order) CalculateFinalprice() error {
	o.finalp = o.price + o.tax

	err := o.Validate()
	if err != nil{
		return err
	}
	return nil
}
