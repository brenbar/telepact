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

package telepact

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// MockTelepactSchema represents a mock Telepact schema
type MockTelepactSchema struct {
	Original              []interface{}
	Parsed                map[string]types.TType
	ParsedRequestHeaders  map[string]*types.TFieldDeclaration
	ParsedResponseHeaders map[string]*types.TFieldDeclaration
}

// FromJSON creates a MockTelepactSchema from JSON
func NewMockTelepactSchemaFromJSON(json string) (*MockTelepactSchema, error) {
	// TODO: Implement
	return &MockTelepactSchema{
		Original:              make([]interface{}, 0),
		Parsed:                make(map[string]types.TType),
		ParsedRequestHeaders:  make(map[string]*types.TFieldDeclaration),
		ParsedResponseHeaders: make(map[string]*types.TFieldDeclaration),
	}, nil
}

// FromFileJSONMap creates a MockTelepactSchema from a file JSON map
func NewMockTelepactSchemaFromFileJSONMap(fileJSONMap map[string]string) (*MockTelepactSchema, error) {
	// TODO: Implement
	return &MockTelepactSchema{
		Original:              make([]interface{}, 0),
		Parsed:                make(map[string]types.TType),
		ParsedRequestHeaders:  make(map[string]*types.TFieldDeclaration),
		ParsedResponseHeaders: make(map[string]*types.TFieldDeclaration),
	}, nil
}
