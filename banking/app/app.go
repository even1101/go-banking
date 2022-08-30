package app

import (
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:9527", mux))

}
