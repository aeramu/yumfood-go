package main

import (
	"log"
	"net/http"

	"github.com/aeramu/yumfood-go/internal/service/dish"
	"github.com/aeramu/yumfood-go/internal/service/order"
	"github.com/aeramu/yumfood-go/internal/service/vendor"

	"github.com/aeramu/yumfood-go/internal/controller"
	id "github.com/aeramu/yumfood-go/internal/implementation/id/xid"
	orderRepo "github.com/aeramu/yumfood-go/internal/implementation/repository/dummy/order"
	dishRepo "github.com/aeramu/yumfood-go/internal/implementation/repository/sqlite/dish"
	vendorRepo "github.com/aeramu/yumfood-go/internal/implementation/repository/sqlite/vendor"
)

func main() {
	idGen := id.NewIDGenerator()
	vendorRepo := vendorRepo.NewRepository()
	dishRepo := dishRepo.NewRepository()
	orderRepo := orderRepo.NewRepository()
	vendorService := vendor.NewService(vendorRepo, idGen)
	dishService := dish.NewService(dishRepo, idGen)
	orderService := order.NewService(orderRepo, idGen)
	router := controller.NewRouter(vendorService, dishService, orderService)
	http.Handle("/", router)
	log.Println("server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
