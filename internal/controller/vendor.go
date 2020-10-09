package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func (c controller) getVendor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	vendor := c.vendor.Get(id)
	json, _ := json.Marshal(vendor)
	fmt.Fprintf(w, string(json))
}

func (c controller) postVendor(w http.ResponseWriter, r *http.Request) {
	body := struct {
		name        string
		description string
	}{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	c.vendor.Create(body.name, body.description)
}

func (c controller) putVendor(w http.ResponseWriter, r *http.Request) {
	body := struct {
		name        string
		description string
	}{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	id := mux.Vars(r)["id"]
	c.vendor.Update(id, body.name, body.description)
}

func (c controller) deleteVendor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	c.vendor.Delete(id)
}
