# Online Streaming

The Online Streaming API provides endpoints to manage users and films on an online streaming platform. With this API, users can sign up, log in, search, and view movie details, among other functionalities.

## Instalación

Asegúrate de tener Docker y Postgres instalados en tu sistema antes de comenzar.

### Lanza base de datos postgresql en contenedor docker

```bash
docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgre
```

### Lanzar proyecto

```bash
air .
```

### Creación de tablas

El lanzamiento del propio proyecto creará las tablas en la base de datos.

Se adjunta script sql para la creación de la base de datos e inserción de datos iniciales.