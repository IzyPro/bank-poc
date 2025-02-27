package main

import (
	"bank-poc/src/internal/handlers"
	"bank-poc/src/internal/repository"
	"bank-poc/src/internal/service"
	"bank-poc/src/internal/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.New()
	services := service.New(repo)

	app := RunHttpServer(services)
	http.Handle("/", app)

	srv := &http.Server{
		Handler:      app,
		Addr:         fmt.Sprintf(":%s", utils.PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Tiny Bank is listening on port %s\n", utils.PORT)
	log.Fatal(srv.ListenAndServe())
}

func RunHttpServer(services *service.Service) *mux.Router {
	r := mux.NewRouter()
	transactionRouter := r.PathPrefix("/transaction").Subrouter()
	handler := &handlers.TransactionHandler{
		TransactionService: services.TransactionService,
	}
	transactionRouter.HandleFunc("", handler.TransactionHistory).Methods("GET")
	transactionRouter.HandleFunc("/balance", handler.Balance).Methods("GET")
	transactionRouter.HandleFunc("/deposit", handler.Deposit).Methods("POST")
	transactionRouter.HandleFunc("/withdraw", handler.Withdraw).Methods("POST")
	r.HandleFunc("/", PingHandler)
	return r
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
