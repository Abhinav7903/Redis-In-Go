# Go-idis

Go-idis is a Redis-like feature implemented in Go. This project aims to provide a simple in-memory database with persistent storage capabilities, similar to Redis.

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

- In-memory data storage
- Persistent data dump
- Simple TCP server interface

## Installation

To install Go-idis, make sure you have Go installed and set up on your machine. Then, clone the repository and build the project:

```bash
git clone https://github.com/Abhinav7903/Go-idis.git
cd Go-idis
go build -o go-idis ./cmd/main.go
```

## Usage

To start the Go-idis server, run the following command:

```bash
./go-idis
```

By default, the server will run on `0.0.0.0:1234`. You can modify the address and port in the `main.go` file.

## Examples

### Running the Server

To run the server:

```bash
./go-idis
```

### Using the Script

To use the provided script to set a key with values from 1 to 100:

```bash
go run cmd/script/main.go mykey
```

### Example Code

Here is an example of how to use Go-idis in your Go project:

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

    // Dump file every 2 hours
    filepath := "/path/to/dump.json"
    store.StartAutoDump(filepath, 2*time.Minute) //change the time according to your need

    // Run the server
    if err := srv.Run(); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
```

## API Reference

### Setting a Key

To set a key with the Go-idis server:

```bash
SET key value
```

### Getting a Key

To get a key from the Go-idis server:

```bash
GET key
```

### Deleting a Key

To delete a key from the Go-idis server:

```bash
DELETE key
```
### Help command
```bash
HELP
```

:)

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Abhinav7903/Go-idis/blob/main/LICENSE) file for details.

