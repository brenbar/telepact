package telepact

import (
	"context"
)

// MockServerOptions contains configuration options for MockServer
type MockServerOptions struct {
	OnError                              func(error)
	EnableMessageResponseGeneration      bool
	EnableOptionalFieldGeneration        bool
	RandomizeOptionalFieldGeneration     bool
	GeneratedCollectionLengthMin         int
	GeneratedCollectionLengthMax         int
}

// NewMockServerOptions creates default options for MockServer
func NewMockServerOptions() *MockServerOptions {
	return &MockServerOptions{
		OnError:                              func(error) {},
		EnableMessageResponseGeneration:      true,
		EnableOptionalFieldGeneration:        true,
		RandomizeOptionalFieldGeneration:     true,
		GeneratedCollectionLengthMin:         0,
		GeneratedCollectionLengthMax:         3,
	}
}

// MockServer is a mock instance of a Telepact server for testing
type MockServer struct {
	random                           *RandomGenerator
	enableGeneratedDefaultStub       bool
	enableOptionalFieldGeneration    bool
	randomizeOptionalFieldGeneration bool
	stubs                            []interface{} // TODO: Define MockStub type
	invocations                      []interface{} // TODO: Define MockInvocation type
	server                           *Server
}

// NewMockServer creates a new MockServer instance
func NewMockServer(mockSchema *MockTelepactSchema, options *MockServerOptions) (*MockServer, error) {
	if options == nil {
		options = NewMockServerOptions()
	}

	random := NewRandomGenerator(
		options.GeneratedCollectionLengthMin,
		options.GeneratedCollectionLengthMax,
	)

	// Create underlying TelepactSchema from MockTelepactSchema
	schema := &TelepactSchema{
		Original:                mockSchema.Original,
		Parsed:                  mockSchema.Parsed,
		ParsedRequestHeaders:    mockSchema.ParsedRequestHeaders,
		ParsedResponseHeaders:   mockSchema.ParsedResponseHeaders,
	}

	serverOptions := &ServerOptions{
		OnError:       options.OnError,
		OnRequest:     func(*Message) {},
		OnResponse:    func(*Message) {},
		AuthRequired:  false,
		Serialization: NewDefaultSerialization(),
	}

	mockServer := &MockServer{
		random:                           random,
		enableGeneratedDefaultStub:       options.EnableMessageResponseGeneration,
		enableOptionalFieldGeneration:    options.EnableOptionalFieldGeneration,
		randomizeOptionalFieldGeneration: options.RandomizeOptionalFieldGeneration,
		stubs:                            make([]interface{}, 0),
		invocations:                      make([]interface{}, 0),
	}

	// Create server with mock handler
	server, err := NewServer(schema, mockServer.handle, serverOptions)
	if err != nil {
		return nil, err
	}
	mockServer.server = server

	return mockServer, nil
}

// Process processes a Telepact request message into a response message
func (ms *MockServer) Process(ctx context.Context, message []byte) (*Response, error) {
	return ms.server.Process(ctx, message, nil)
}

// handle is the internal handler function that processes mock requests
func (ms *MockServer) handle(ctx context.Context, requestMessage *Message) (*Message, error) {
	// TODO: Implement mock_handle logic
	// For now, return a stub response
	return &Message{
		Headers: make(map[string]interface{}),
		Body:    map[string]interface{}{"Ok_": map[string]interface{}{}},
	}, nil
}
