package main

import (
	// "encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"github.com/nathanSeixeiro/go-intensivo/internal/entity"
)

func main() {
	// r := chi.NewRouter() -> router help only on conversation with http methods
	// r.Use(middleware.Logger)
	// r.Get("/", Order)
	// http.ListenAndServe(":8888", r)

	e := echo.New() // echo -> web framework/lib (like express.js?), if necessary change him, you've change all handlers
	e.GET("/", Order)
	e.Logger.Fatal(e.Start(":8888"))
}

func Order(e echo.Context) error {
	order, err := entity.Constructor("1234", 100, 11)
	if err != nil {
		return e.JSON(http.StatusBadRequest, order)
	}
	erro := order.CalculateFinalprice()
	if erro != nil {
		return e.String(http.StatusInternalServerError, erro.Error())
	}
	return e.JSON(http.StatusOK, order)
}
