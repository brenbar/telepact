package schema

import (
	"github.com/brenbar/telepact/lib/go/telepact"
)

// CreateMockTelepactSchemaFromFileJsonMap creates a mock telepact schema from a map of JSON documents
// This adds the mock schema (which includes _ext.Call_ and _ext.Stub_ types) in addition to
// the internal and auth schemas
func CreateMockTelepactSchemaFromFileJsonMap(jsonDocuments map[string]string) (*telepact.MockTelepactSchema, error) {
	// Copy the input map
	finalJsonDocuments := make(map[string]string, len(jsonDocuments)+1)
	for k, v := range jsonDocuments {
		finalJsonDocuments[k] = v
	}
	
	// Add the mock schema
	finalJsonDocuments["mock_"] = GetMockTelepactJson()
	
	// Create the base telepact schema (which will also add internal_ and conditionally auth_)
	telepactSchema, err := CreateTelepactSchemaFromFileJsonMap(finalJsonDocuments)
	if err != nil {
		return nil, err
	}
	
	// Return a MockTelepactSchema with the same fields
	return &telepact.MockTelepactSchema{
		Original:              telepactSchema.Original,
		Parsed:                telepactSchema.Parsed,
		ParsedRequestHeaders:  telepactSchema.ParsedRequestHeaders,
		ParsedResponseHeaders: telepactSchema.ParsedResponseHeaders,
	}, nil
}
