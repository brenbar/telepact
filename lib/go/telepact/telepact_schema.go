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
	"github.com/brenbar/telepact/lib/go/telepact/internal/schema"
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// TelepactSchema represents a parsed Telepact schema
type TelepactSchema struct {
	Original               []interface{}
	Parsed                 map[string]types.TType
	ParsedRequestHeaders   map[string]*types.TFieldDeclaration
	ParsedResponseHeaders  map[string]*types.TFieldDeclaration
}

// FromJSON creates a TelepactSchema from a JSON string
func (ts *TelepactSchema) FromJSON(json string) (*TelepactSchema, error) {
	fileJSONMap := map[string]string{"auto_": json}
	schemaInternal, err := schema.CreateTelepactSchemaFromFileJSONMap(fileJSONMap)
	if err != nil {
		return nil, err
	}
	return &TelepactSchema{
		Original:              schemaInternal.Original,
		Parsed:                schemaInternal.Parsed,
		ParsedRequestHeaders:  schemaInternal.ParsedRequestHeaders,
		ParsedResponseHeaders: schemaInternal.ParsedResponseHeaders,
	}, nil
}

// FromFileJSONMap creates a TelepactSchema from a map of filenames to JSON
func FromFileJSONMap(fileJSONMap map[string]string) (*TelepactSchema, error) {
	schemaInternal, err := schema.CreateTelepactSchemaFromFileJSONMap(fileJSONMap)
	if err != nil {
		return nil, err
	}
	return &TelepactSchema{
		Original:              schemaInternal.Original,
		Parsed:                schemaInternal.Parsed,
		ParsedRequestHeaders:  schemaInternal.ParsedRequestHeaders,
		ParsedResponseHeaders: schemaInternal.ParsedResponseHeaders,
	}, nil
}

// FromDirectory creates a TelepactSchema from a directory
func FromDirectory(directory string) (*TelepactSchema, error) {
	schemaFileMap, err := schema.GetSchemaFileMap(directory)
	if err != nil {
		return nil, err
	}
	schemaInternal, err := schema.CreateTelepactSchemaFromFileJSONMap(schemaFileMap)
	if err != nil {
		return nil, err
	}
	return &TelepactSchema{
		Original:              schemaInternal.Original,
		Parsed:                schemaInternal.Parsed,
		ParsedRequestHeaders:  schemaInternal.ParsedRequestHeaders,
		ParsedResponseHeaders: schemaInternal.ParsedResponseHeaders,
	}, nil
}
