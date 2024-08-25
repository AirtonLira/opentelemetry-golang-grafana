package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	appdb "github.com/airtonlira/opentelemetry/internal/migration"
	"github.com/airtonlira/opentelemetry/internal/routes"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	serviceName = "my-go-service"
	serverAddr  = ":8080"
)

func initTracer() (*trace.TracerProvider, error) {
	ctx := context.Background()

	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	fmt.Println("sdwad")

	otel.SetTracerProvider(tp)
	return tp, nil
}

func setupDatabase() error {
	if err := appdb.EnsureDatabaseExists(); err != nil {
		return err
	}
	return appdb.EnsureTablesExist()
}

func main() {
	// Inicializa o TracerProvider
	tp, err := initTracer()
	if err != nil {
		log.Fatalf("Falha ao inicializar o tracer: %v", err)
	}
	defer tp.Shutdown(context.Background())

	// Configura o banco de dados
	if err := setupDatabase(); err != nil {
		log.Fatalf("Falha na configuração do banco de dados: %v", err)
	}

	// Cria e configura o roteador
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Configura o servidor HTTP
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	// Inicia o servidor em uma goroutine separada
	go func() {
		log.Printf("Servidor rodando em %s", serverAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
	}()

	// Configura o graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Desligando o servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Erro ao desligar o servidor: %v", err)
	}

	log.Println("Servidor encerrado com sucesso")
}
