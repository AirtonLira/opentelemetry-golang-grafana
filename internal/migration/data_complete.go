package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var dbName = os.Getenv("POSTGRES_DB")

func init() {
	WaitPostgresUp()
}

func EnsureDatabaseExists() error {

	// Carregar variáveis de ambiente
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))

	// Conectar ao banco de dados postgres padrão
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Verificar se o banco de dados já existe
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	err = conn.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking if database exists: %v", err)
	}

	// Se o banco de dados não existir, criar
	if !exists {
		_, err := conn.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return fmt.Errorf("error creating database: %v", err)
		}
		log.Printf("Database '%s' created.", dbName)
	} else {
		log.Printf("Database '%s' already exists.", dbName)
	}

	return nil
}

func EnsureTablesExist() error {

	// Carregar variáveis de ambiente
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	// Conectar ao banco de dados postgres padrão
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// SQL para criar a tabela 'accounts'
	createAccountsTable := `
	CREATE TABLE IF NOT EXISTS accounts (
		id VARCHAR(50) PRIMARY KEY,
		name VARCHAR(100),
		balance NUMERIC
	)`

	// SQL para criar a tabela 'payments'
	createPaymentsTable := `
	CREATE TABLE IF NOT EXISTS payments (
		id VARCHAR(50) PRIMARY KEY,
		account_id VARCHAR(50),
		amount NUMERIC,
		status VARCHAR(50)
	)`

	// Executar a criação das tabelas
	_, err = conn.Exec(createAccountsTable)
	if err != nil {
		return fmt.Errorf("error creating accounts table: %v", err)
	}

	_, err = conn.Exec(createPaymentsTable)
	if err != nil {
		return fmt.Errorf("error creating payments table: %v", err)
	}

	log.Println("Tables 'accounts' and 'payments' ensured to exist.")
	return nil
}
