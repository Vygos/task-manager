# Task Manager System

## Overview
This project is a **Task Manager** that provides **CRUD operations** for managing tasks. It follows a **Reduced Clean Architecture** approach on the backend, ensuring a well-structured and maintainable codebase, while keeping the domain logic isolated from external dependencies.

## Tech Stack
- **Backend:** Golang (Gin, sqlc)
- **Frontend:** Angular (Angular Material)
- **Database:** PostgreSQL
- **Proxy:** Nginx
- **Containerization:** Docker & Docker Compose

## Architecture
The backend is designed using a **Reduced Clean Architecture** style, with a strong focus on domain logic. It is structured to be modular and independent of external frameworks, making it easier to maintain and extend.

## Getting Started

### Prerequisites
Ensure you have the following installed on your system:
- **Docker** & **Docker Compose**

### Running the Project
To start the application, run the following command:

```bash
docker-compose up --build
```
After the build successfully you could visit localhost.
