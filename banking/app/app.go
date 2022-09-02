package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/ping", ping).Methods(http.MethodGet)

	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)

	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ah := AccountHandlers{service.NewAccountService(accountRepositoryDb)}
	router.HandleFunc("/customers/{id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost)

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
	// log.Fatal(http.ListenAndServe("localhost:9527", router))

}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "user:pwd@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
