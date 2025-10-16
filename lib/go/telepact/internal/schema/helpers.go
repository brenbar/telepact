//
//  Copyright The Telepact Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  https://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

package schema

import (
	"encoding/json"
	"fmt"
)

// ParseContext provides context for schema parsing
type ParseContext struct {
	// TODO: Add context fields as needed
}

// SchemaParseFailure represents a failure during schema parsing
type SchemaParseFailure struct {
	DocumentName string
	Path         []string
	ErrorType    string
	Details      map[string]interface{}
}

func (e *SchemaParseFailure) Error() string {
	return fmt.Sprintf("Schema parse failure in %s at %v: %s", e.DocumentName, e.Path, e.ErrorType)
}

// ParseTypeDeclaration parses a type declaration from raw JSON
func ParseTypeDeclaration(typeSpec interface{}) (interface{}, error) {
	// TODO: Implement full type declaration parsing
	// This is a placeholder that just validates it's valid JSON-like structure
	
	switch v := typeSpec.(type) {
	case string:
		// Simple type name like "string", "integer", etc.
		return v, nil
	case map[string]interface{}:
		// Complex type definition
		return v, nil
	default:
		return nil, fmt.Errorf("invalid type specification: %T", typeSpec)
	}
}

// ValidateSchemaJSON validates that a string is valid JSON array
func ValidateSchemaJSON(jsonStr string) error {
	var result interface{}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	
	if _, ok := result.([]interface{}); !ok {
		return fmt.Errorf("schema must be a JSON array")
	}
	
	return nil
}
