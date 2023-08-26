# website-template

## Setup

To setup this GO project you need to install the following tools:

-   [Go](https://golang.org/dl/)
-   [Air](https://github.com/cosmtrek/air)

### Development

To Start the development server run the following command:

```bash
air
```

Air will start a development server on `port 8080` and create a development build in `tmp` folder.
To change the port run the following command:

```bash
air -- --port :<port> # replace <port> with the port number
```

### Build

Build the app for production run the following command:

```bash
go build -o ./build/
```

This will create a binary file in the `build` folder.

## Folder Structure

```bash
├── build
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
