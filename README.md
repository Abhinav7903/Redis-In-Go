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
- **Simple TCP Server Interface**: Easy-to-use communication protocol over TCP.
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

### Example Code

Below is an example of how to integrate Go-idis into a Go project:

```go
package main

import (
    "go-idis/internal/idis"
    "go-idis/server"
    "log"
    "time"
)

func main() {
    // Initialize the in-memory repository
    store := idis.NewInMemoryRepository()

    // Create a new TCP server
    srv := server.NewServer("0.0.0.0:1234", store)

    // Enable periodic data dumps
    filepath := "/path/to/dump.json"
    store.StartAutoDump(filepath, 2*time.Hour)

    // Start the server
    if err := srv.Run(); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
```

## API Reference

### Commands

- **Set a Key**:
  ```bash
  SET key value
  ```
- **Get a Key**:
  ```bash
  GET key
  ```
- **Delete a Key**:
  ```bash
  DELETE key
  ```
- **Help Command**:
  ```bash
  HELP
  ```

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/Abhinav7903/Go-idis/blob/main/LICENSE) file for details.
