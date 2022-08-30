package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/ping", ping).Methods(http.MethodGet)

	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:9527", router))

}
