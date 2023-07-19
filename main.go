package main

import "github.com/nathanSeixeiro/go-intensivo/internal/entity"

func main() {
	order, err := entity.Constructor("1", 9, 2)

	if err != nil {
		println(err.Error())
	} else {
		println(order.ID)
	}
}
