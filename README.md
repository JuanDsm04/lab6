# MyApplication API
This is an API built with **Go**, using the Chi router and **PostgreSQL** database. It provides endpoints for managing TV series, allowing CRUD operations.

## Getting Started
The repository provides a `docker-compose.yml` file with the necessary instructions to build and run the API inside containers.

### 1. Clone the repository
```sh
git clone https://github.com/JuanDsm04/lab6.git
cd lab6
```

### 2. Build and run the containers
With Docker installed, run:
```sh
docker compose up --build
```
**Note:** If the API does not start correctly on the first attempt, try running the build command again to ensure the database initializes properly before starting the backend.

## API Endpoints
| Method  | Endpoint                  | Description                          |
|---------|--------------------------|--------------------------------------|
| `GET`   | `/api/series`             | Retrieve all series                 |
| `GET`   | `/api/series/{id}`        | Retrieve a series by ID             |
| `POST`  | `/api/series`             | Create a new series                 |
| `PUT`   | `/api/series/{id}`        | Update an existing series           |
| `DELETE`| `/api/series/{id}`        | Delete a series                     |
| `PATCH` | `/api/series/{id}/status` | Update the status of a series       |
| `PATCH` | `/api/series/{id}/episode`| Increment the last episode watched  |
| `PATCH` | `/api/series/{id}/upvote` | Upvote a series                     |
| `PATCH` | `/api/series/{id}/downvote`| Downvote a series                   |

## Swagger Documentation
This API includes **Swagger UI** for easy documentation and testing.

Once the server is running, open in your browser:
```
http://localhost:8080/swagger/index.html
```

## Postman Collection
For testing with Postman, use the following collection:
[Postman Collection](https://elements.getpostman.com/redirect?entityId=19231888-32c5c8cc-5827-4ec9-b58b-ee6f92d4c2a4&entityType=collection)

## Frontend example
![image](https://github.com/user-attachments/assets/3ab9a181-8833-4a3d-bbcb-4f7fb64cd3a0)

## Technologies Used
- **Language:** Go
- **Frameworks:** Chi (Router), GORM (ORM)
- **Database:** PostgreSQL
- **Documentation:** Swagger, Postman
- **Containers:** Docker

## Author
Developed by JuanDsm04.
