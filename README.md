# Go API template

This is a template for a Go API project. It uses [Fiber](https://gofiber.io/) as the web framework and [GORM](https://gorm.io/) as the ORM.

## Getting started

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Running the project

1. Clone the repository

```bash
git clone
```

2. Build the project

```bash
docker-compose build
```

3. Run the project

```bash
docker-compose up
```

4. Open the browser and go to [http://localhost:3000/ping](http://localhost:3000/ping)

If you see the message `pong`, then everything is working correctly.

## Documentation

This project uses [Swagger](https://swagger.io/) for API documentation. You can access the documentation at [http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html).

### Updating the documentation

To update the documentation, you need to run the following command:

```bash
swag init --dir ./cmd/http,./internal/todo --parseDependency --parseInternal
```

> **Note:** The documentation is only available when the project is running.

> **Note:** The documentation is automatically generated based on the code. You don't need to manually update it.

## Authentication

This project uses [JWT](https://jwt.io/) for authentication. The JWT secret is stored in the `.env` file.

The JWT is store on a Cookie.

## Database

This project uses [PostgreSQL](https://www.postgresql.org/) as the database. The database configuration is stored in the `.env` file.

## Architecture

This project uses the [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) as the architecture.

### Project structure

```
.
├── cmd                     # Application commands (entry points)
│   └── http                # HTTP server
│       └── main.go         # Main file
├── config                  # Configuration files
│   └── viper.go            # Configuration file for environment variables
├── docs                    # Swagger files (auto generated)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── Internal                # Business logic
│   ├── todo                # Domain
│   │   ├── middleware.go   # Middleware layer (if needed)
│   │   ├── controller.go   # Controller layer (more details in file)
│   │   ├── router.go       # Router layer
│   │   └── storage.go      # Storage layer (more details in file)
│   └── storage             # Databases
│       └── postgres.go
├── pkg                     # Packages
│   └── shutdown
│       └── gracefully.go   # Gracefully shutdown
├── .env                    # Environment variables
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── README.md
└── .air.toml
```
