package controller

import (
	"net/http"

	"github.com/aeramu/yumfood-go/internal/service/dish"
	"github.com/aeramu/yumfood-go/internal/service/order"

	"github.com/aeramu/yumfood-go/internal/service/vendor"

	"github.com/gorilla/mux"
)

// NewRouter return router
func NewRouter(
	vendor vendor.Service,
	dish dish.Service,
	order order.Service,
) http.Handler {
	r := mux.NewRouter()
	c := controller{vendor, dish, order}

	r.HandleFunc("/vendors", c.postVendor).Methods("POST")
	r.HandleFunc("/vendors/{id}", c.getVendor).Methods("GET")
	r.HandleFunc("/vendors/{id}", c.getVendor).Methods("PUT")
	r.HandleFunc("/vendors/{id}", c.getVendor).Methods("DELETE")

	r.HandleFunc("/vendors/{vendorID}/dishes", c.postDish).Methods("POST")
	r.HandleFunc("/vendors/{vendorID}/dishes/{id}", c.putDish).Methods("PUT")
	r.HandleFunc("/vendors/{vendorID}/dishes/{id}", c.deleteDish).Methods("DELETE")
	r.HandleFunc("/vendors/{vendorID}/dishes", c.getDishes).Methods("GET")

	r.HandleFunc("/orders", c.postOrder).Methods("POST")
	r.HandleFunc("/orders", c.getOrders).Methods("GET")

	return r
}

type controller struct {
	vendor vendor.Service
	dish   dish.Service
	order  order.Service
}
