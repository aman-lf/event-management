# Event Management App

## Description

A app to manage event, participant and expenses

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/aman-lf/event-management.git
   ```

### Setup BE

- **Move to server dir:**

  ```bash
  cd event-management/server

  ```

- **Install dependencies:**

  ```bash
  go mod tidy
  ```

- **Start the project:**

  ```bash
  go run server.go
  ```

## Configuration

- change .env.example to .env in both server and app and provide required credentials

## Migration

```bash
make migrate-up
```
