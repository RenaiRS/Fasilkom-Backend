# Fasilkom-Backend

This is a backend project built with Go, using the Fiber framework. It serves as the API for a university faculty website, managing content such as news, scholarships, vacancies, and more. The project uses MySQL as its database and GORM as the ORM.

---

## Folder Structure

The project is organized into the following directories:

```
redesign_backend/
├── assets/                # Stores static files like covers
├── cache/                 # Cache management for different entities
├── config/                # Configuration management
├── controllers/           # HTTP request handlers
├── database/              # Database connection and migration
├── middlewares/           # Authentication middleware
├── models/                # Database models/entities
├── routes/                # API route definitions
├── utils/                 # Utility functions
├── .env                   # Environment variables
├── go.mod                 # Go module dependencies
└── main.go                # Application entry point
```

---

## How to Run the Project

Follow these steps to set up and run the project locally.

### Prerequisites

* Go (version 1.24.2 or higher)
* MySQL (or a compatible database service)

### 1. Set up the Database

First, you need to create and populate the database.

1.  Create a new database named `fasilkom_be` in your MySQL server.
2.  Import the provided SQL dump to create the necessary tables and populate them with initial data. The file is located at `database/fasilkom_be.sql`.

### 2. Configure Environment Variables

The project uses a `.env` file for configuration.

1.  Copy the `.env.example` file and rename it to `.env`.
2.  Update the values in the new `.env` file with your database credentials:

    ```ini
    DB_HOST=localhost
    DB_PORT=3306
    DB_USER=root
    DB_PASS=your_db_password
    DB_NAME=fasilkom_be
    SERVER_PORT=8080
    ```
    * **DB_USER**: Your MySQL username (e.g., `root`).
    * **DB_PASS**: Your MySQL password.
    * **DB_NAME**: The name of the database you created (`fasilkom_be`).

### 3. Install Dependencies

Install the required Go modules by running this command in the project root directory:

```sh
go mod tidy
   ```


### 4. Run the application
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Documentation

### Base URL
```
http://localhost:8080/api
```

### Authentication

Most endpoints require admin authentication. After logging in, the system uses HTTP-only cookies for session management.

#### Admin Routes
```
POST   /admin/signup          # Create admin account
POST   /admin/login           # Admin login
POST   /admin/logout          # Admin logout
POST   /admin/sync-berita     # Sync news cache (Auth required)
```

#### News (Berita) Routes
```
GET    /berita                # Get all news (supports ?is_priority=true&limit=10)
POST   /berita                # Create news (Auth required)
PATCH  /berita/:id            # Update news (Auth required)
DELETE /berita/:id            # Delete news (Auth required)
```

#### Scholarship Routes
```
GET    /scholarship           # Get all scholarships
POST   /scholarship           # Create scholarship (Auth required)
PATCH  /scholarship/:id       # Update scholarship (Auth required)
DELETE /scholarship/:id       # Delete scholarship (Auth required)
```

#### Vacancies Routes
```
GET    /vacancies            # Get all vacancies
POST   /vacancies            # Create vacancy (Auth required)
PATCH  /vacancies/:id        # Update vacancy (Auth required)
DELETE /vacancies/:id        # Delete vacancy (Auth required)
```

#### Student Cooperation Routes
```
GET    /kemahasiswaan-kerjasama     # Get all programs
POST   /kemahasiswaan-kerjasama     # Create program (Auth required)
PATCH  /kemahasiswaan-kerjasama/:id # Update program (Auth required)
DELETE /kemahasiswaan-kerjasama/:id # Delete program (Auth required)
```

#### File Routes
```
GET    /file                 # Get all files
POST   /file                 # Upload file (Auth required)
```
