# go-boiler-plate

This is a Go boilerplate project following Domain-Driven Design (DDD) principles.

## Technologies Used

- Go
- PostgreSQL
- Docker
- Docker Compose

## Project Structure

// ...existing code...

## How to Run

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/go-boiler-plate.git
    cd go-boiler-plate
    ```

2. **Set up environment variables:**
    Create a `.env` file in the root directory and add the following:
    ```env
    POSTGRES_USER=yourusername
    POSTGRES_PASSWORD=yourpassword
    POSTGRES_DB=yourdatabase
    ```

3. **Run Docker Compose:**
    ```sh
    docker-compose up -d
    ```

4. **Run the application:**
    ```sh
    go run main.go
    ```

## Domain-Driven Design (DDD)

This project is structured according to DDD principles. The main components are:

- **Domain:** Contains the core business logic.
- **Application:** Contains application-specific logic.
- **Infrastructure:** Contains infrastructure-related code such as database connections.