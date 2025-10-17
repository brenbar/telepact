package telepact

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// MockTelepactSchema is a parsed Telepact schema with mock types included
// It includes the standard types plus _ext.Call_ and _ext.Stub_ for testing
type MockTelepactSchema struct {
	// Original is the original pseudo-JSON array before parsing
	Original []interface{}
	
	// Parsed is a map of type names to their parsed TType definitions
	Parsed map[string]types.TType
	
	// ParsedRequestHeaders is a map of header names to their field declarations for request headers
	ParsedRequestHeaders map[string]*types.TFieldDeclaration
	
	// ParsedResponseHeaders is a map of header names to their field declarations for response headers
	ParsedResponseHeaders map[string]*types.TFieldDeclaration
}

// NewMockTelepactSchemaFromJson creates a MockTelepactSchema from a single JSON string
func NewMockTelepactSchemaFromJson(json string) (*MockTelepactSchema, error) {
	return NewMockTelepactSchemaFromFileJsonMap(map[string]string{"auto_": json})
}

// NewMockTelepactSchemaFromFileJsonMap creates a MockTelepactSchema from a map of document names to JSON strings
func NewMockTelepactSchemaFromFileJsonMap(fileJsonMap map[string]string) (*MockTelepactSchema, error) {
	// This will be implemented by calling internal/schema/create_mock_telepact_schema_from_file_json_map.go
	// For now, return a placeholder error
	// TODO: Wire up to internal/schema.CreateMockTelepactSchemaFromFileJsonMap
	return nil, &TelepactError{Message: "NewMockTelepactSchemaFromFileJsonMap not yet implemented"}
}

// NewMockTelepactSchemaFromDirectory creates a MockTelepactSchema by loading .telepact.json files from a directory
func NewMockTelepactSchemaFromDirectory(directory string) (*MockTelepactSchema, error) {
	// This will load files using internal/schema.GetSchemaFileMap
	// and then call NewMockTelepactSchemaFromFileJsonMap
	// TODO: Wire up to internal/schema.GetSchemaFileMap and CreateMockTelepactSchemaFromFileJsonMap
	return nil, &TelepactError{Message: "NewMockTelepactSchemaFromDirectory not yet implemented"}
}
