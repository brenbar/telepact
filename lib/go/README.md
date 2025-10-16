# Telepact Go Library

This is a Go port of the Python Telepact library.

## Installation

```bash
go get github.com/brenbar/telepact/lib/go/telepact
```

## Usage

### API Definition

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
    "github.com/brenbar/telepact/lib/go/telepact"
)

func main() {
    // Load schema from directory
    schema, err := telepact.NewTelepactSchemaFromDirectory("/path/to/api/files")
    if err != nil {
        panic(err)
    }

    // Create handler
    handler := func(ctx context.Context, requestMessage *telepact.Message) (*telepact.Message, error) {
        functionName := requestMessage.GetBodyTarget()
        arguments := requestMessage.GetBodyPayload()

        // Dispatch to appropriate function
        if functionName == "fn.greet" {
            subject := arguments["subject"].(string)
            return telepact.NewMessage(
                map[string]interface{}{},
                map[string]interface{}{
                    "Ok_": map[string]interface{}{
                        "message": "Hello " + subject + "!",
                    },
                },
            ), nil
        }

        return nil, telepact.NewTelepactError("Function not found")
    }

    // Create server options
    options := &telepact.ServerOptions{
        AuthRequired: true,
        Serialization: telepact.NewDefaultSerialization(),
    }

    // Create server
    server, err := telepact.NewServer(schema, handler, options)
    if err != nil {
        panic(err)
    }

    // Wire up request/response bytes from your transport
    transportHandler := func(requestBytes []byte) ([]byte, error) {
        response, err := server.Process(context.Background(), requestBytes, map[string]interface{}{})
        if err != nil {
            return nil, err
        }
        return response.Bytes, nil
    }

    // Start your transport with transportHandler
    // ...
}
```

### Client

```go
package main

import (
    "context"
    "github.com/brenbar/telepact/lib/go/telepact"
)

func main() {
    // Create adapter
    adapter := func(ctx context.Context, m *telepact.Message, s *telepact.Serializer) (*telepact.Message, error) {
        requestBytes, err := s.Serialize(m)
        if err != nil {
            return nil, err
        }

        // Wire up request/response bytes to your transport
        responseBytes, err := transport.Send(requestBytes)
        if err != nil {
            return nil, err
        }

        return s.Deserialize(responseBytes)
    }

    // Create client options
    options := &telepact.ClientOptions{
        UseBinary: false,
        AlwaysSendJSON: true,
        TimeoutMsDefault: 5000,
        Serialization: telepact.NewDefaultSerialization(),
    }

    // Create client
    client := telepact.NewClient(adapter, options)

    // Make request
    message := telepact.NewMessage(
        map[string]interface{}{},
        map[string]interface{}{
            "fn.greet": map[string]interface{}{
                "subject": "World",
            },
        },
    )

    response, err := client.Request(context.Background(), message)
    if err != nil {
        panic(err)
    }

    // Handle response
    // ...
}
```

## Development

### Building

```bash
make
```

### Testing

```bash
make test
```

## Status

This is a work in progress. See [PROGRESS.md](PROGRESS.md) for details on the porting status.
