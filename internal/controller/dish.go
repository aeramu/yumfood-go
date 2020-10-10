package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aeramu/yumfood-go/internal/domain"
	"github.com/gorilla/mux"
)

func (c controller) getDishes(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["vendorID"]
	dishes := c.dish.GetListByVendorID(id)
	json, _ := json.Marshal(dishes)
	fmt.Fprintf(w, string(json))
}

func (c controller) postDish(w http.ResponseWriter, r *http.Request) {
	body := domain.Dish{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	vendorID := mux.Vars(r)["vendorID"]
	c.dish.Create(vendorID, body.Name, body.Price)
}

func (c controller) putDish(w http.ResponseWriter, r *http.Request) {
	body := domain.Dish{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	id := mux.Vars(r)["id"]
	c.dish.Update(id, body.Name, body.Price)
}

func (c controller) deleteDish(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	c.dish.Delete(id)
}
