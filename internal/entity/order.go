package entity

import "errors"

// in Go we work without class but structs 
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

// for the methods we use pointers and references for the struct 
// Go don't have try/catch, we catching each error e and process it
func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is required")
	}

	if o.price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if o.tax <= 0 {
		return errors.New("price must be greater than zero")
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
