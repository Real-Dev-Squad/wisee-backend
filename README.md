# website-template

## Setup

To setup this GO project you need to install the following tools:

-   [Go](https://golang.org/dl/)
-   [Air](https://github.com/cosmtrek/air)
-   [PostgreSQL](https://www.postgresql.org/download/)
-   [Golang Migrate - CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### Development Setup

1. To Start the development server run the following command:

    ```bash
    air
    ```

    Air will start a development server on `port 8080` and create a development build in `tmp` folder.
    To change the port run the following command:

    ```bash
    air -- --port :<port> # replace <port> with the port number
    ```

2. Create a `.env` using a env file from `environments` folder.

    ```bash
    cp environments/.env.development .env
    ```

    > **Note:** You can also create a `.env` file manually and copy the content from `.env.example` file.

3. Create a database in PostgreSQL and update the database url in `.env` file.
4. Run the following command to run the migrations:

    ```bash
    migrate -path ./migrations -database <database_url> up
    ```

    > **Note:** Replace `<database_url>` with the database url from `.env` file.

### Build

Build the app for production run the following command:

```bash
go build -o ./build/
```

This will create a binary file in the `build` folder.

### Migrations

Migrations are handled using [golang-migrate](https://github.com/golang-migrate/migrate), read the [docs](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md) for more info.

## Folder Structure

```bash
├── build
├── models                # All models
├── migrations            # Database migrations
├── environments          # All env files
    ├── .env.development
    ├── .env.production
    ├── .env.test
    ├── .env.example
├── routes                # All routes
    ├── main.go
    ├── <route group>     # A route group example: users
├── utils                # All utils which are used in the project
    ├── main.go
    ├── <util>           # A util example: formatDate
├── .gitignore
├── .air.toml            # Air config file
├── go.mod
├── go.sum
├── main.go              # Main file
└── README.md
```
