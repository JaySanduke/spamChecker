# Spam Detection Application API

## Tech Stack

- Language: **Go (Golang)**
- Web Framework: **Gin**
- ORM: **GORM**
- Authentication: **JWT**
- Database: **PostgreSQL**
- Env Management: **godotenv**

---

## Features

- User registration and login with JWT auth
- Contact list ingestion (mocked)
- Mark phone numbers as spam
- Search users by name or number
- Conditional email visibility
- Profile API with spam count
- Seeder to populate dummy data

---

## Project Structure

```plaintext
├── cmd/ # Entry point
├── config/ # Env config
├── controllers/ # Request handlers
├── database/ # DB connection
├── middleware/ # JWT Auth
├── models/ # GORM Models
├── routes/ # Route bindings
├── scripts/seed.go # Seeder for mock data
├── utils/ # JWT utilities
├── .env # Environment variables
└── go.mod / go.sum # Go dependencies
```

---

## 🛠 Setup Instructions

### 1. Clone and Setup

```bash
git clone <repo-url>  # NA
cd spamChecker
go mod tidy
```

### 2. Create .env file

```bash
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgresSpamChecker
DB_NAME=spamChecker
DB_PORT=5432
JWT_SECRET=SpamCheckerSuperSecret
```

### 3. Run the app

Prerequisites:

- **Docker**
- **Docker Compose**
- **Golang 1.24^**

### To run the app in development mode:

Run the following command to start the PostgreSQL database using Docker Compose:
```bash
docker compose -f 'docker-compose.yml' up -d --build 'db' 
```

or you can run any other postgres database instance but just change the UserName, Password and the Host in the env file.

Then, run the following command to start the application:

```
go run cmd/main.go --mode=development
```

### To run the app in production mode:

You can run the following command to start application in production mode using Docker Compose:

```
docker-compose up --build -d
```

or if you want to run the app without docker, you can run the following command:

```
go run cmd/main.go --mode=production
```

### 4. Seed Sample Data (optional)

To populate the database with sample data, run the following command if the application is running locally:

```
go run scripts/seed.go
```

or if you are using docker, run the following command:

```
docker exec -it spamChecker_app go run scripts/seed.go
```

## API Endpoints

Method Endpoint Description
```
POST /api/user/register Register a user
POST /api/user/login Login (returns JWT)
POST /api/spam/mark Mark number as spam (auth)
GET /api/user/search?name=x Search by name (auth)
GET /api/user/search?phone=x Search by phone (auth)
GET /api/user/profile/:phone Get profile with spam stats
```

*Note: All routes except registration/login require a Bearer token.

Postman Collection is available in the folder for testing. with the name `SpamChecker.postman_collection.json`
