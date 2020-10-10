package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aeramu/yumfood-go/internal/domain"
)

func (c controller) getOrders(w http.ResponseWriter, r *http.Request) {
	order := c.order.GetList()
	json, _ := json.Marshal(order)
	fmt.Fprintf(w, string(json))
}

func (c controller) postOrder(w http.ResponseWriter, r *http.Request) {
	body := []domain.OrderItem{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	c.order.Create(body)
}
