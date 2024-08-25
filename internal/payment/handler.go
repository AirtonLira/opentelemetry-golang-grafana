package payment

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	migration "github.com/airtonlira/opentelemetry/internal/migration"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

type Payment struct {
	ID        string  `json:"id"`
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
}

func init() {

	migration.WaitPostgresUp()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	log.Printf("String connection payment: %v", connStr)

	var err error
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	log.Printf("String connection handler: %v", connStr)

	var err error
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO payments (id, account_id, amount, status) VALUES ($1, $2, $3, $4)", payment.ID, payment.AccountID, payment.Amount, payment.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var payment Payment

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	log.Printf("String connection handler: %v", connStr)

	var err error
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.QueryRow("SELECT id, account_id, amount, status FROM payments WHERE id = $1", params["id"]).Scan(&payment.ID, &payment.AccountID, &payment.Amount, &payment.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Payment not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(payment)
}
