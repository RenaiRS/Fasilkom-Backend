
# Backend API - Redesign Challenge Website GDGoC Universitas Sriwijaya

This is a backend service built using Go (Fiber & GORM) to manage news articles and downloadable files for the Redesign Challenge website.

---

## Features

- Admin authentication using secure token-based cookies
- CRUD operations for news articles
- File upload and download (PDF and image cover)
- Caching for news listing (with Go cache)
- Static file hosting for cover images and uploaded files

---

## Setup Instructions

1. **Clone the repository**
   ```bash
   git clone https://github.com/Safmica/redesign_challenge.git
   cd redesign_challenge
   ```

2. **Copy the environment config**
   ```bash
   cp .env.example .env
   ```

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Create asset folders**
   ```bash
   mkdir -p assets/covers assets/files
   ```

5. **Run the server**
   ```bash
   go run main.go
   ```

---

## Project Structure

| Folder/File       | Description                          |
|-------------------|--------------------------------------|
| `main.go`         | Application entry point              |
| `controllers/`    | HTTP handler logic                   |
| `models/`         | GORM models                          |
| `routes/`         | Route definitions                    |
| `middlewares/`    | Auth middleware                      |
| `utils/`          | Helper functions                     |
| `assets/`         | Folder for file and cover uploads    |
| `.env`            | Environment configuration            |

---
## Example `.env` File
    DB_USER=your_db_user
    DB_PASS=your_db_pass
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=your_db_name
    BASE_URL=http://localhost:8080
---

## Authentication

- Admin-only access
- Token is stored in cookie: `admin_token`
- Use middleware to protect endpoints

---

## API Endpoints

| Method | Endpoint              | Description                       | Auth Required |
|--------|-----------------------|-----------------------------------|---------------|
| POST   | `/api/admin/login`              | Admin login and set token cookie  | No            |
|POST    |`/api/admin/signup`|Create new admin account|Yes|
|POST    |`/api/admin/logout`|Logout and Delete admin auth token|Yes|
|POST    |`/api/admin/sync-berita`|Sync berita cache with the latest data from database (flush + refresh)|Yes|
| GET    | `/api/berita`         | Get all posted news (optional: `?is_priority=true`) | No |
| POST   | `/api/berita`         | Create a news article             | Yes           |
| PATCH    | `/api/berita/:id`     | Update news article               | Yes           |
| DELETE | `/api/berita/:id`     | Delete news article               | Yes           |
| GET    | `/api/files`      | Get All file list                   | No            |
| POST   | `/api/files`          | Upload a file (PDF)               | Yes           |

---

## Static File Access

- Cover images: `http://<BASE_URL>/assets/covers/<filename>`
- Uploaded files: `http://<BASE_URL>/assets/files/<filename>`

---

## Notes

- Uses in-memory cache for news listing
- Token expiration and secure cookie settings are applied
- `.env.example` provided for easy setup
- Full Postman Docs:
> https://documenter.getpostman.com/view/36503501/2sB3BBqrwt

---
