# Telepact Go Library

Go implementation of the Telepact API library.

## Installation

```bash
go get github.com/brenbar/telepact/lib/go
```

## Usage

### API Schema

Define your API in `.telepact.json` files:

```json
[
    {
        "fn.greet": {
            "subject": "string"
        },
        "->": {
            "Ok_": {
                "message": "string"
            }
        }
    }
]
```

### Server

```go
package main

import (
    "context"
    "log"
    "net/http"
    
    telepact "github.com/brenbar/telepact/lib/go"
)

func main() {
    // Load schema from directory
    schema, err := telepact.FromDirectory("./api")
    if err != nil {
        log.Fatal(err)
    }

    // Create handler
    handler := func(ctx context.Context, requestMessage *telepact.Message) (*telepact.Message, error) {
        functionName := requestMessage.GetBodyTarget()
        arguments := requestMessage.GetBodyPayload()

        switch functionName {
        case "fn.greet":
            subject := arguments["subject"].(string)
            return telepact.NewMessage(
                map[string]interface{}{},
                map[string]interface{}{
                    "Ok_": map[string]interface{}{
                        "message": "Hello " + subject + "!",
                    },
                },
            ), nil
        default:
            return nil, telepact.NewTelepactError("Function not found", nil)
        }
    }

    // Create server options
    options := telepact.NewServerOptions()
    options.AuthRequired = false

    // Create server
    server, err := telepact.NewServer(schema, handler, options)
    if err != nil {
        log.Fatal(err)
    }

    // HTTP handler
    http.HandleFunc("/api/telepact", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        // Read request body
        requestBytes, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Process request
        response, err := server.Process(r.Context(), requestBytes, nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Determine content type
        contentType := "application/json"
        if _, hasBin := response.Headers["bin_"]; hasBin {
            contentType = "application/octet-stream"
        }

        // Write response
        w.Header().Set("Content-Type", contentType)
        w.Write(response.Bytes)
    })

    log.Println("Server starting on :8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
```

### Client

```go
package main

import (
    "context"
    "fmt"
    "io"
    "log"
    "net/http"
    
    telepact "github.com/brenbar/telepact/lib/go"
)

func main() {
    // Create adapter function
    adapter := func(ctx context.Context, message *telepact.Message, serializer *telepact.Serializer) (*telepact.Message, error) {
        // Serialize request
        requestBytes, err := serializer.Serialize(message)
        if err != nil {
            return nil, err
        }

        // Send HTTP request
        req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8000/api/telepact", bytes.NewReader(requestBytes))
        if err != nil {
            return nil, err
        }
        req.Header.Set("Content-Type", "application/json")

        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()

        // Read response
        responseBytes, err := io.ReadAll(resp.Body)
        if err != nil {
            return nil, err
        }

        // Deserialize response
        return serializer.Deserialize(responseBytes)
    }

    // Create client
    options := telepact.NewClientOptions()
    client := telepact.NewClient(adapter, options)

    // Make request
    requestMessage := telepact.NewMessage(
        map[string]interface{}{},
        map[string]interface{}{
            "fn.greet": map[string]interface{}{
                "subject": "World",
            },
        },
    )

    responseMessage, err := client.Request(context.Background(), requestMessage)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Response: %+v\n", responseMessage)
}
```

## Development Status

This is an initial implementation of the Telepact library for Go. The core functionality for server and client operations is in place, but some advanced features are still being implemented:

- ✅ Basic message handling
- ✅ Server and Client implementations
- ✅ Schema loading from files
- ✅ JSON serialization
- 🚧 Binary encoding (placeholder)
- 🚧 Full schema parsing and validation
- 🚧 Type generation
- 🚧 Advanced query features

## License

Licensed under the Apache License, Version 2.0. See LICENSE for details.
