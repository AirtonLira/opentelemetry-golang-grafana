# 🚀 OpenTelemetry Tracing & Metrics Project

## 📋 Descrição

Este projeto é uma aplicação completa de monitoramento e rastreamento distribuído utilizando **OpenTelemetry**, **Grafana Tempo**, **Prometheus** e **Grafana**. Ele coleta e armazena traces distribuídos e métricas associadas às operações da aplicação, permitindo a visualização e análise tanto das métricas quanto dos traces.

### 🛠 Tecnologias Utilizadas

- **Grafana Tempo**: Coleta e armazena traces distribuídos para análise de telemetria.
- **Prometheus**: Coleta métricas das aplicações e do Grafana Tempo.
- **Grafana**: Visualiza tanto as métricas coletadas pelo Prometheus quanto os traces do Grafana Tempo.
- **Go (Golang)**: Aplicação que emite métricas e traces usando OpenTelemetry.
- **Docker Compose**: Gerencia os contêineres e a infraestrutura do projeto.
- **PostgreSQL**: Banco de dados relacional para armazenar informações de `Account` e `Payment`.

---

## 🚀 Requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 📂 Estrutura do Projeto

```bash
.
├── docker-compose.yml           # Configuração de todos os serviços Docker
├── prometheus.yml               # Configuração do Prometheus para coletar métricas
├── otel-collector-config.yaml   # Configuração do OpenTelemetry Collector
├── tempo.yaml                   # Configuração do Grafana Tempo
├── go-app/                      # Código da aplicação em Go
│   ├── main.go                  # Arquivo principal da aplicação
│   └── internal/                # Handlers e lógica de negócio da aplicação
│       ├── account/             # Lógica relacionada a contas
│       │   ├── handler.go       # Handler para operações de Account
│       ├── payment/             # Lógica relacionada a pagamentos
│       │   ├── handler.go       # Handler para operações de Payment
└── README.md                    # Este arquivo