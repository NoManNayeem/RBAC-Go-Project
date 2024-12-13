
# RBAC-Go-Project

## 1. Project Overview

This project implements a **Role-Based Access Control (RBAC)** system with:
- **Go Backend** for handling APIs and middleware.
- **React Frontend** for the user interface.
- **Docker** and **Docker Compose** for containerized deployment.
- Support for **local development** with environment-specific configurations.

### Features:
- Role-Based Routing: `admin` and `public` users with protected routes.
- API Endpoints:
  - `/login`: Authenticate user and return their role (`admin` or `public`).
  - `/public`: Accessible to all users.
  - `/admin`: Restricted to admin users.
- Environment-based DB connection for both **local** and **dockerized** setups.

---

## 2. How to Clone from GitHub

1. Clone the repository:
    ```bash
    git clone https://github.com/NoManNayeem/RBAC-Go-Project.git
    ```

2. Navigate to the project directory:
    ```bash
    cd RBAC-Go-Project
    ```

---

## 3. How to Run Locally

### Prerequisites:
- **Go** (version 1.22 or later)
- **Node.js** (version 20 or later)
- **MS SQL Server** running locally or on a remote server

### Steps:

1. **Setup Backend:**
   - Navigate to the backend directory:
     ```bash
     cd Go-Backend
     ```
   - Copy the `env.example` file to `.env`:
     ```bash
     cp env.example .env
     ```
   - Update the `.env` file with your local database credentials.
   - Run the backend:
     ```bash
     go run cmd/main.go
     ```

2. **Setup Frontend:**
   - Navigate to the frontend directory:
     ```bash
     cd react-frontend
     ```
   - Copy the `env.example` file to `.env`:
     ```bash
     cp env.example .env
     ```
   - Install dependencies and run the frontend:
     ```bash
     npm install
     npm start
     ```

3. **Test Locally:**
   - Access the frontend at `http://localhost:3000`.
   - The backend API runs on `http://localhost:8080`.

---

## 4. How to Run with Docker and Docker Compose

### Prerequisites:
- **Docker** and **Docker Compose** installed.
- A running database instance accessible by the backend (update global `.env` for DB connection).

### Steps:

1. **Prepare Global Environment File:**
   - At the root of the project, copy `env.example` to `.env`:
     ```bash
     cp env.example .env
     ```
   - Update the `.env` file with your database credentials and other environment variables.

2. **Run with Docker Compose:**
   - Build and start the services:
     ```bash
     docker-compose up --build
     ```

3. **Test with Docker:**
   - Access the frontend at `http://localhost:3000`.
   - The backend API runs on `http://localhost:8080`.

---

## 5. Additional Information

### API Endpoints:
- **Login API:** `/login` (POST)
  - Request: `{ "username": "Nayeem", "password": "Password" }`
  - Response: `{ "role": "admin" }` (or `"public"` for other users).
- **Public API:** `/public` (GET)
- **Admin API:** `/admin` (GET; restricted to `admin`).

### Folder Structure:
```
RBAC-Go-Project/
├── .env
├── docker-compose.yml
├── Go-Backend/
│   ├── .env
│   ├── Dockerfile
│   ├── cmd/
│   ├── config/
│   ├── internal/
│   └── go.mod
└── react-frontend/
    ├── .env
    ├── Dockerfile
    ├── public/
    └── src/
```

---

## Enjoy using the RBAC-Go-Project! 🚀
