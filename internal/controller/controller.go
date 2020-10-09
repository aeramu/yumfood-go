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
	r.HandleFunc("/vendor/{id}", c.getVendor)

	return r
}

type controller struct {
	vendor vendor.Service
	dish   dish.Service
	order  order.Service
}
