package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "running")
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

// func createCustomer(w http.ResponseWriter, r *http.Request) {

// }

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["id"])
// }
