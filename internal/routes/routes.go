package routes

import (
	"github.com/airtonlira/opentelemetry/internal/account"
	migration "github.com/airtonlira/opentelemetry/internal/migration"
	"github.com/airtonlira/opentelemetry/internal/payment"

	"github.com/gorilla/mux"
)

// RegisterRoutes define as rotas da aplicação
func RegisterRoutes(router *mux.Router) {

	migration.WaitPostgresUp()

	// Rotas para contas
	router.HandleFunc("/account", account.CreateAccount).Methods("POST")
	router.HandleFunc("/account/{id}", account.GetAccount).Methods("GET")

	// Rotas para pagamentos
	router.HandleFunc("/payment", payment.CreatePayment).Methods("POST")
	router.HandleFunc("/payment/{id}", payment.GetPayment).Methods("GET")
}
