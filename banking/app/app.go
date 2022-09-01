package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/ping", ping).Methods(http.MethodGet)

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:9527", router))

}
