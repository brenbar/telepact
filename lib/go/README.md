## Telepact

### Installation

```bash
go get github.com/brenbar/telepact/lib/go/telepact
```

### Usage

API:

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

Server:

```go
package main

import (
    "context"
    "github.com/brenbar/telepact/lib/go/telepact"
)

func main() {
    files, err := telepact.NewTelepactSchemaFiles("./api")
    if err != nil {
        panic(err)
    }

    schema, err := telepact.FromFileJSONMap(files.FilenameToJSON)
    if err != nil {
        panic(err)
    }

    handler := func(ctx context.Context, requestMessage *telepact.Message) (*telepact.Message, error) {
        functionName := requestMessage.GetBodyTarget()
        arguments := requestMessage.GetBodyPayload()

        // Dispatch request to appropriate function handling code
        if functionName == "fn.greet" {
            subject := arguments["subject"].(string)
            return &telepact.Message{
                Headers: map[string]interface{}{},
                Body: map[string]interface{}{
                    "Ok_": map[string]interface{}{
                        "message": "Hello " + subject + "!",
                    },
                },
            }, nil
        }

        return nil, &telepact.TelepactError{Message: "Function not found"}
    }

    options := telepact.NewServerOptions()
    server, err := telepact.NewServer(schema, handler, options)
    if err != nil {
        panic(err)
    }

    // Wire up request/response bytes from your transport of choice
    // Example: HTTP handler
    // response, err := server.Process(ctx, requestBytes, nil)
    // responseBytes := response.Bytes
}
```

Client:

```go
package main

import (
    "context"
    "github.com/brenbar/telepact/lib/go/telepact"
)

func main() {
    adapter := func(ctx context.Context, m *telepact.Message, s *telepact.Serializer) (*telepact.Message, error) {
        requestBytes, err := s.Serialize(m)
        if err != nil {
            return nil, err
        }

        // Wire up request/response bytes to your transport of choice
        // responseBytes := transport.Send(requestBytes)

        return s.Deserialize(responseBytes)
    }

    options := telepact.NewClientOptions()
    client := telepact.NewClient(adapter, options)

    // Make requests
    // response, err := client.Request(ctx, requestMessage)
}
```

For more concrete usage examples, see the tests in the repository.
