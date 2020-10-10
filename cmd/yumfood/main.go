package main

import (
	"log"
	"net/http"

	"github.com/aeramu/yumfood-go/internal/service/dish"
	"github.com/aeramu/yumfood-go/internal/service/order"
	"github.com/aeramu/yumfood-go/internal/service/vendor"

	"github.com/aeramu/yumfood-go/internal/controller"
	dishRepo "github.com/aeramu/yumfood-go/internal/implementation/repository/sqlite/dish"
	orderRepo "github.com/aeramu/yumfood-go/internal/implementation/repository/sqlite/order"
	vendorRepo "github.com/aeramu/yumfood-go/internal/implementation/repository/sqlite/vendor"
	uid "github.com/aeramu/yumfood-go/internal/implementation/uid/shortid"
)

func main() {
	uidGen := uid.NewUIDGenerator()
	vendorRepo := vendorRepo.NewRepository()
	dishRepo := dishRepo.NewRepository()
	orderRepo := orderRepo.NewRepository()
	vendorService := vendor.NewService(vendorRepo, uidGen)
	dishService := dish.NewService(dishRepo, uidGen)
	orderService := order.NewService(orderRepo, uidGen)
	router := controller.NewRouter(vendorService, dishService, orderService)
	http.Handle("/", router)
	log.Println("server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
