package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aeramu/yumfood-go/internal/domain"

	"github.com/gorilla/mux"
)

func (c controller) getVendor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	vendor := c.vendor.Get(id)
	json, _ := json.Marshal(vendor)
	fmt.Fprintf(w, string(json))
}

func (c controller) postVendor(w http.ResponseWriter, r *http.Request) {
	body := domain.Vendor{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	c.vendor.Create(body.Name, body.Description)
}

func (c controller) putVendor(w http.ResponseWriter, r *http.Request) {
	body := domain.Vendor{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	id := mux.Vars(r)["id"]
	c.vendor.Update(id, body.Name, body.Description)
}

func (c controller) deleteVendor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	c.vendor.Delete(id)
}
