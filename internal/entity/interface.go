package entity

type IOrderRepository interface {
	Save(order *Order) error
	GetTotalTransaction() (int, error)
}