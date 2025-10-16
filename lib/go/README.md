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
    "fmt"
    
    "github.com/brenbar/telepact/lib/go/telepact"
)

func main() {
    schema, err := telepact.FromDirectory("/directory/containing/api/files")
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
                Headers: make(map[string]interface{}),
                Body: map[string]interface{}{
                    "Ok_": map[string]interface{}{
                        "message": fmt.Sprintf("Hello %s!", subject),
                    },
                },
            }, nil
        }
        
        return nil, fmt.Errorf("function not found")
    }
    
    options := telepact.NewServerOptions()
    server, err := telepact.NewServer(schema, handler, options)
    if err != nil {
        panic(err)
    }
    
    // Wire up request/response bytes from your transport of choice
    transportHandler := func(ctx context.Context, requestBytes []byte) ([]byte, error) {
        response, err := server.Process(ctx, requestBytes, nil)
        if err != nil {
            return nil, err
        }
        return response.Bytes, nil
    }
    
    // Use transportHandler with your transport
    _ = transportHandler
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
        responseBytes, err := transport.Send(ctx, requestBytes)
        if err != nil {
            return nil, err
        }
        
        return s.Deserialize(responseBytes)
    }
    
    options := telepact.NewClientOptions()
    client := telepact.NewClient(adapter, options)
    
    // Use client
    _ = client
}
```

For more concrete usage examples, see the tests.
