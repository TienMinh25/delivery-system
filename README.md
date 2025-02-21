# Delivery System: Microservices-based Application

## Overview

**Delivery System** is a scalable, microservices-based project built with **Go**. It follows **Clean Code Architecture** principles and supports **gRPC** and **message brokers** for service-to-service communication. The project includes **authentication & authorization**, **email and push notifications**, **SQL & NoSQL databases**, **caching**, **file storage**, and **distributed tracing**.

The project is about Delivery System which consists of four microservices:
- **API Gateway**
- **Orders Service**
- **Partners Service**
- **Notifications Service**

Each service is containerized using **Docker** and orchestrated via **Docker Compose** and **Kubernetes**. The project also features **automatic database schema migration**, **mock tests**, and **Swagger documentation** for APIs.

---

## Project Structure

```
delivery-system/
│── cmd/                        # Main service binaries
│   ├── api/                    # API Gateway, authentication & authorization service
│   ├── notifications/          # Push Notifications service
│   ├── orders/                 # Orders service
│   ├── partners/               # Partners and products service
│── configs/                    # Configuration files
│── design/                     # System design diagrams
│── docs/                       # API documentation (Swagger)
│── internal/                   # Core business logic (private)
│── migrations/                 # Database migrations
│── pkg/                        # Shared utility packages
│── postgres_data/              # Folder contains data postgres (mounted from docker)
│── minio_data/                 # Folder contains data S3 (mounted from docker)
│── thirty_party/               # Thirty party used in project (implement interface in package pkg, e.g. minio, keycloak,...)
│── scripts/                    # All scripts are related to database, etc...
│── docker-compose.yml          # Docker Compose configuration
│── .dockerignore               # Files and directories to exclude when building Docker images
│── Makefile                    # Automation commands
│── go.mod, go.sum              # Go dependencies
│── README.md                   # Project documentation
```
---

## Features

- **Clean Architecture**: Well-structured and maintainable.
- **Microservices Architecture**: Independent, scalable services.
- **gRPC & Message Broker**: Efficient inter-service communication.
- **Authentication & Authorization**: Secure access management.
- **Redis Caching**: Performance improvements with in-memory caching.
- **SQL & NoSQL Databases**: PostgreSQL and a NoSQL solution for optimized and flexible data storage.
- **Push Notifications**: Asynchronous messaging for push notifications.
- **File Storage**: Persistent storage solution.
- **Distributed Tracing**: OpenTelemetry for performance monitoring.
- **Swagger API Documentation**: Easy API consumption.
- **Comprehensive Testing**: Unit tests and integration tests.
- **Docker & Kubernetes Ready**: Containerized for scalability.
- **Custom HTTP Router**: Optimized request processing.

---
## Tech stack:
- **Language programming**: Golang
- **Storage**: Postgres, Redis, S3 (minio instead)
- **Interconnection**: gRPC, HTTP, websocket
- **Authention system**: Keycloak
- **Message broker**: Kafka
- **Distributed tracing**: Jeager (handle metric and provide UI), Opentelemetry (collect all information like metric, service name, etc.. and export data to Jeager) 
---

This project was inspired by open-source resources. Thanks to the projects for helping shape the idea and source code to reference and development:
- Github Repository: https://github.com/shahzodshafizod/gocloud