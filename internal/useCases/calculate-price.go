package usecases

import "github.com/nathanSeixeiro/go-intensivo/internal/entity"

type OrderInput struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutput struct {
	ID     string
	Price  float64
	Tax    float64
	Finalp float64
}

// SOLID - D - Dependency Inversion Principle
type CalculateFinalPrice struct {
	OrderRepository entity.IOrderRepository
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.Constructor(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalprice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutput{
		ID:     order.ID,
		Price:  order.Price,
		Tax:    order.Tax,
		Finalp: order.Finalp,
	}, nil
}
