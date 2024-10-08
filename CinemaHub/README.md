
```markdown
# CinemaHub

CinemaHub is an API for managing movies and directors, built with Go using the Gin framework, GORM for PostgreSQL database interaction, and Swagger for API documentation.

## Features

- **Manage Movies:** Add, view, update, and delete movies.
- **Manage Directors:** Add and view director details.
- **Object-Oriented Architecture:** Designed using structs, methods, and interfaces for an object-oriented approach.
- **PostgreSQL Database:** Stores information in a robust and reliable PostgreSQL database.
- **Swagger Documentation:** Interactive API documentation with Swagger UI.

## Prerequisites

- [Go](https://golang.org/dl/) version 1.16 or higher
- [PostgreSQL](https://www.postgresql.org/download/) version 12 or higher

## Installation and Setup

Clone the repository:

```bash
git clone https://github.com/yourusername/cinemahub.git
cd cinemahub
```

### Install Dependencies

To install the project dependencies, run:

```bash
go mod tidy
```

### Database Configuration

Ensure that PostgreSQL is installed and running on your system. Update the database connection settings in the `main.go` file with your own credentials:

```go
dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable TimeZone=Asia/Tehran"
```

### Migrate Models

GORM's auto-migration feature will create the necessary tables in your database. This will happen automatically when you run the application.

### Generate Swagger Documentation

To generate Swagger documentation, make sure you have Swag installed. If not, install it using:

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

Then generate the Swagger files:

```bash
swag init
```

### Run the Project

To start the project, execute:

```bash
go run main.go
```

The API will be available on port 8080.

### Access Swagger Documentation

Once the server is running, you can access the interactive Swagger documentation at:

```
http://localhost:8080/swagger/index.html
```

## API Endpoints

- **GET /movies** - Retrieve a list of all movies
- **GET /movies/:id** - Retrieve details of a specific movie by ID
- **POST /movies** - Create a new movie
- **PUT /movies/:id** - Update an existing movie
- **DELETE /movies/:id** - Delete a movie

## Contributing

If you would like to contribute to CinemaHub:

1. Fork the repository.
2. Create a new branch for your feature or bug fix (`git checkout -b feature/YourFeatureName`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to your branch (`git push origin feature/YourFeatureName`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Contact

For questions or suggestions, you can reach out via email:

[Mohsen.Khorrami1990@gmail.com](mailto:Mohsen.Khorrami1990@gmail.com)
```

### How to Use the Documentation

- Place the `README.md` file in the root directory of your project.
- Update the project name, database connection string (`dsn`), GitHub link, and contact email with your details.
- Upload the file to your GitHub repository to serve as the project's documentation.

This README now includes instructions for setting up and accessing Swagger documentation, providing a complete overview for users to install, run, and contribute to the CinemaHub project.