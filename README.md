# Online Streaming

The Online Streaming API provides endpoints to manage users and films on an online streaming platform. With this API, users can sign up, log in, search, and view movie details, among other functionalities.

## Installation

Make sure you have Docker and Postgres installed on your system before starting.

### Launch PostgreSQL database in Docker container

```bash
docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgre
```

### Install dependencies

go mod download

### Launch project

```bash
air .
```

### Creating tables

Launching the project itself will create the tables in the database.

An SQL script is attached for creating the database and inserting initial data.

## Postman Collection

A Postman collection is included for testing the API.