# Go-idis

Go-idis is an open-source in-memory database written in Go. This project aims to provide a lightweight and simple solution for managing key-value data with support for persistent storage.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [API Reference](#api-reference)
- [Contributing](#contributing)
- [License](#license)
- [Contact Information](#contact-information)

## Features

- **In-Memory Data Storage**: Fast, lightweight storage for key-value pairs.
- **Persistent Data Dump**: Supports saving and reloading data from disk.
- **Simple HTTP and TCP Server Interfaces**: Easy-to-use communication protocols.
- **Minimalistic Design**: No reliance on complex data structures (yet).

Future updates will incorporate advanced data structures for enhanced functionality.

## Installation

Ensure Go is installed on your system. Then, clone the repository and build the project:

```bash
git clone https://github.com/Abhinav7903/Go-idis.git
cd Go-idis
go build -o go-idis ./cmd/main.go
```

## Usage

To start the Go-idis server, execute the following:

```bash
./go-idis
```

By default, the server listens on `0.0.0.0:1234`. Modify the address and port in the `main.go` file if needed.

## Examples

### Running the Server

Run the server with:

```bash
./go-idis
```

### Using the Script

The provided script can populate a key with values from 1 to 100:

```bash
go run cmd/script/main.go mykey
```

### HTTP Requests

You can interact with the server using HTTP requests. Below are examples of available routes:

- **Set a Key** (POST)
    ```bash
    curl -X POST http://localhost:1234/set/a -H "Content-Type: application/json" -d '["45", "56"]'
    ```

- **Get a Key** (GET)
    ```bash
    curl -X GET http://localhost:1234/get/a
    ```

- **Get Unique Values for a Key** (GET)
    ```bash
    curl -X GET http://localhost:1234/getuq/a
    ```

- **Delete a Key** (DELETE)
    ```bash
    curl -X DELETE http://localhost:1234/delete/a
    ```

## API Reference

- **Set a Key** (POST): `/set/{key}`
    - Example: `curl -X POST http://localhost:1234/set/a -H "Content-Type: application/json" -d '["45", "56"]'`

- **Get a Key** (GET): `/get/{key}`
    - Example: `curl -X GET http://localhost:1234/get/a`

- **Get Unique Values for a Key** (GET): `/getuq/{key}`
    - Example: `curl -X GET http://localhost:1234/getuq/a`

- **Delete a Key** (DELETE): `/delete/{key}`
    - Example: `curl -X DELETE http://localhost:1234/delete/a`

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/Abhinav7903/Go-idis/blob/main/LICENSE) file for details.

- ```bash
  :)
  ```