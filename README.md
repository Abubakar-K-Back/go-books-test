# **Go Books API**

## **Overview**
Go Books API is a RESTful service built using the **Gin framework** in Go. It provides CRUD operations to manage books with **PostgreSQL** as the database, **Redis** for caching, and **Kafka** for event streaming. The API is **containerized** using Docker and can be deployed with Docker Compose.

## **Features**
- **CRUD operations** for managing books
- **Pagination support** for retrieving book lists
- **Redis caching** for optimized book retrieval
- **Kafka integration** for event-driven architecture
- **Swagger documentation** for API reference
- **Docker & Docker Compose support**
- **PostgreSQL database with automatic initialization**

## **Tech Stack**
- **Go** (Gin framework)
- **PostgreSQL** (Database)
- **Redis** (Caching)
- **Kafka** (Event Streaming)
- **Docker & Docker Compose** (Containerization)

---

## **Getting Started**

### **Prerequisites**
Ensure you have the following installed:
- **Go 1.24+**
- **Docker & Docker Compose**

### **1. Clone the Repository**
```sh
git clone https://github.com/Abubakar-K-Back/go-books-test.git
cd go-books-api
```

### **2. Set Up Environment Variables**
Create a `.env` file and add the following:
```
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=booksdb
REDIS_ADDR=redis:6379
KAFKA_BROKER=kafka:9092
```

### **3. Start the Application with Docker Compose**
```sh
docker-compose up -d --build
```

### **4. Verify the Setup**
Check running containers:
```sh
docker ps
```
Check logs for PostgreSQL:
```sh
docker logs postgres_db
```
Check logs for the API:
```sh
docker logs go_books_api
```

---

## **API Endpoints**

### **Books Endpoints**
| Method | Endpoint         | Description |
|--------|----------------|-------------|
| GET    | `/books`        | Retrieve all books (supports pagination) |
| GET    | `/books/{id}`   | Retrieve a single book |
| POST   | `/books`        | Create a new book (publishes to Kafka) |
| PUT    | `/books/{id}`   | Update an existing book (publishes to Kafka) |
| DELETE | `/books/{id}`   | Delete a book (publishes to Kafka) |

### **Swagger Documentation**
Once the API is running, access the Swagger UI at:
```
http://localhost:8080/swagger/index.html
```

---

## **Database Initialization**

PostgreSQL is automatically initialized with the `init.sql` script:
- Ensures `postgres` user exists
- Creates `booksdb` if it doesn’t exist
- Grants all privileges to `postgres`

To verify manually:
```sh
docker exec -it postgres_db psql -U postgres -d booksdb -c "\du"
docker exec -it postgres_db psql -U postgres -d booksdb -c "\l"
```

---

## **Testing the API**

### **Using cURL**
```sh
curl -X POST "http://localhost:8080/books" -H "Content-Type: application/json" -d '{
    "title": "Go Programming",
    "author": "John Doe",
    "year": 2023
}'
```

### **Using Postman**
Import the provided Postman collection from the `docs` folder.

---

## **Stopping and Cleaning Up**
To stop and remove all containers:
```sh
docker-compose down -v
```
To remove unused Docker resources:
```sh
docker system prune -a
```


## **Contact**
For questions or support, reach out at **abkawan6@gmail.com** or create an issue on the GitHub repository.


