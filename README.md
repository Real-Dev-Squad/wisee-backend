# website-template

## Setup

To setup this GO project you need to install the following tools:

-   [Go](https://golang.org/dl/)
-   [Air](https://github.com/cosmtrek/air)
-   [PostgreSQL](https://www.postgresql.org/download/)

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

### Build

Build the app for production run the following command:

```bash
go build -o ./build/
```

This will create a binary file in the `build` folder.

## Folder Structure

```bash
├── build
├── models                # All models
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
