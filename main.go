package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nathanSeixeiro/go-intensivo/internal/infra/database"
	usecases "github.com/nathanSeixeiro/go-intensivo/internal/useCases"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	orderRepository := database.NewOrderRepository(db)
	usecase := usecases.NewCalculateFinalPrice(orderRepository)
	input := usecases.OrderInput{
		ID: "1",
		Price: 10.0,
		Tax: 10.0,
	}
	output, err := usecase.Execute(input)
	if err != nil{
		panic(err)
	}
	fmt.Println(output)
}
