# Go CRUD App with Docker and PostgreSQL
#### The go crud app performs crud operations on three different entities. Articles, Categories and Authors. Data for each of them is handled differently using memory, files and postgres database respectively. Each entity has sepreate apis/end points. 

## API Documentation for the Project
- API: [Documentation](https://documenter.getpostman.com/view/26763809/2s946e9DEf)
-- This documentation contains all the end-points and their example requests and their responses that have been tested using postman for this project.

## Prerequisites
- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Postgres (Optional): [Installation Guide](https://www.postgresqltutorial.com/postgresql-getting-started/install-postgresql/)

## Prerequisites
- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Postgres (Optional): [Installation Guide](https://www.postgresqltutorial.com/postgresql-getting-started/install-postgresql/)

## Steps to run the Project

1. Clone the project from the repository.
2. Open a terminal and navigate to the project's root directory.
3. Update the database configuration in `pkg/config/app.go` to use the PostgreSQL host `db`.
4. Open the `docker-compose.yml` file and replace the following environment variables with your PostgreSQL database credentials:
   - `your_postgres_user`
   - `your_postgres_password`
   - `your_postgres_db`
5. Build and run the Docker containers using the following command:
   ```bash 
      docker-compose up -d

## Steps to run the Project without docker

1. Clone the project from the repository.
2. Open a terminal and navigate to the project's root directory.
4. Open the `pkg/config/app.go` file and update the database connection string variables with your PostgreSQL database credentials:
     ```bash 
      	connectionString := "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
5. Run the project  using the following command:
   ```bash 
       go run .\cmd\main.go
