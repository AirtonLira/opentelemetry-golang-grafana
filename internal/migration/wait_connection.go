package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var conn *sql.DB
var err error

func WaitPostgresUp() {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))

	log.Printf("Conexão com PostgreSQL: %v", connStr)

	for i := 0; i < 10; i++ {
		conn, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		err = conn.Ping()
		if err == nil {
			break
		}

		log.Printf("PostgreSQL ainda não está pronto, tentando novamente... (%d/10)", i+1)
		time.Sleep(5 * time.Second) // Espera 5 segundos antes da próxima tentativa
	}
}
