# Admission Portal Backend

A RESTful backend API for an Admission Portal built using Go, Gin, GORM, and PostgreSQL.  
This project manages students, courses, enrollments, and student authentication. All services run in containers using Docker Compose for zero manual setup.

---

## Table of Contents

- [Objective](#objective)
- [Features](#features)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [Setup & Installation](#setup--installation)
- [Running the Project](#running-the-project)
- [API Endpoints](#api-endpoints)
- [Database Models](#database-models)
- [Testing (Postman)](#testing-postman)
- [License](#license)

---

## Objective

Design and implement a backend system for an Admission Portal. 
Features include student signup and info management, course creation/listing, and student enrollment to courses.  
Focus is on containerization, automation, ease of setup, and code clarity.

---

## Features

- **Student Management**
  - Register a new student (with hashed password)
  - Login and receive JWT authentication token
  - View all students
  - View own student profile (JWT protected)

- **Course Management**
  - Create a new course
  - Retrieve all courses

- **Enrollment Management**
  - Enroll a student in a course
  - Retrieve all enrollments

- Automatic database migration (GORM AutoMigrate)
- PostgreSQL runs in container, managed by Docker Compose
- Environment variable-based configuration via `.env` and Docker
- Easy, one-command setup
- API endpoints tested and sharable using a Postman collection

---

## Project Structure

```
admission-portal/
├── config/              # Database connection and setup logic
├── controllers/         # API handlers
├── middlewares/         # JWT/auth middleware
├── models/              # GORM models
├── routes/              # Route definitions
├── utils/               # JWT helper
├── .env                 # Environment variables
├── docker-compose.yml   # Docker Compose config (API + Postgres)
├── Dockerfile           # Backend build
├── main.go              # Main server entrypoint
├── go.mod / go.sum      # Go modules
├── admission-management-system.postman_collection.json # Postman Collection (for demo/testing)
└── README.md            # This file
```

---

## Technologies Used

- Go  
- Gin  
- GORM  
- PostgreSQL  
- Docker & Docker Compose  
- JWT (token-based authentication)  
- godotenv  

---

## Setup & Installation

### Prerequisites

- Docker and Docker Compose installed

### Quickstart

1. Clone the repository:
    ```sh
    git clone https://github.com/anupamsingh2/ADDMISSION-MANAGEMENT-SYSTEM.git
    cd admission-portal
    ```

2. Start the system (API + DB) in one command:
    ```sh
    docker-compose up --build
    ```
    - API available at [http://localhost:8080](http://localhost:8080)

*No manual DB setup is needed. Password hashing, JWT secrets, and database config are handled via Docker Compose + .env.*

---

## Running the Project (Dev mode)

If you want to run outside Docker:

1. Create a `.env` in the root:
    ```
    DATABASE_URL=postgres://admin:admin123@localhost:5432/admissiondb
    JWT_SECRET=S1@Anupam
    ```
2. Run Postgres separately (or use Docker Compose for DB only).
3. In project directory:
    ```sh
    go mod download
    go run main.go
    ```

---

## API Endpoints

### Students

| Method | Endpoint             | Description                       |
|--------|--------------------- |-----------------------------------|
| POST   | `/students/`         | Register a new student            |
| POST   | `/students/login`    | Login as student, receive JWT     |
| GET    | `/students/`         | List all students                 |
| GET    | `/students/profile`  | Get profile (JWT required)        |

### Courses

| Method | Endpoint     | Description             |
|--------|--------------|------------------------|
| POST   | `/courses/`  | Create a new course    |
| GET    | `/courses/`  | List all courses       |

### Enrollments

| Method | Endpoint         | Description                   |
|--------|-----------------|-------------------------------|
| POST   | `/enrollments/`  | Enroll student in a course   |
| GET    | `/enrollments/`  | List all enrollments         |

---

## Database Models

- **Student**  
  - `id`, `name`, `email`, `phone`, `password (hashed)`

- **Course**  
  - `id`, `title`, `code`

- **Enrollment**
  - `id`, `student_id`, `course_id`

Tables are created automatically on first run.

---

## Testing (Postman)

1. Start the system:
    ```sh
    docker-compose up --build
    ```

2. Open [Postman](https://www.postman.com/).

3. Import the file `admission-management-system.postman_collection.json` from this repository.

4. Use the pre-built requests:
    - Register a student.
    - **Login as that student:** copy the `token` from the response.
    - For protected endpoint(s) ("Get Student Profile"):  
      Select the request, go to the **Authorization** tab, set type to **Bearer Token**, and paste your JWT token.

5. All endpoints can be tested in sequence.

---


*Created by Anupam Singh*
