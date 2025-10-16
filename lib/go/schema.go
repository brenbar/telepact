//|
//|  Copyright The Telepact Authors
//|
//|  Licensed under the Apache License, Version 2.0 (the "License");
//|  you may not use this file except in compliance with the License.
//|  You may obtain a copy of the License at
//|
//|  https://www.apache.org/licenses/LICENSE-2.0
//|
//|  Unless required by applicable law or agreed to in writing, software
//|  distributed under the License is distributed on an "AS IS" BASIS,
//|  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//|  See the License for the specific language governing permissions and
//|  limitations under the License.
//|

package telepact

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/brenbar/telepact/lib/go/internal/schema"
)

// TelepactSchema represents a parsed Telepact schema
type TelepactSchema struct {
	Original              []interface{}
	Parsed                map[string]*schema.TType
	ParsedRequestHeaders  map[string]*schema.TFieldDeclaration
	ParsedResponseHeaders map[string]*schema.TFieldDeclaration
}

// NewTelepactSchema creates a new TelepactSchema
func NewTelepactSchema(original []interface{}, parsed map[string]*schema.TType,
	parsedRequestHeaders, parsedResponseHeaders map[string]*schema.TFieldDeclaration) *TelepactSchema {
	return &TelepactSchema{
		Original:              original,
		Parsed:                parsed,
		ParsedRequestHeaders:  parsedRequestHeaders,
		ParsedResponseHeaders: parsedResponseHeaders,
	}
}

// FromJSON creates a TelepactSchema from a JSON string
func FromJSON(jsonStr string) (*TelepactSchema, error) {
	fileMap := map[string]string{
		"auto_": jsonStr,
	}
	return FromFileJSONMap(fileMap)
}

// FromFileJSONMap creates a TelepactSchema from a map of filename to JSON content
func FromFileJSONMap(fileJSONMap map[string]string) (*TelepactSchema, error) {
	internalSchema, err := schema.CreateTelepactSchemaFromFileJSONMap(fileJSONMap)
	if err != nil {
		return nil, err
	}

	// Convert internal schema to public schema
	return &TelepactSchema{
		Original:              internalSchema.Original,
		Parsed:                internalSchema.Parsed,
		ParsedRequestHeaders:  internalSchema.ParsedRequestHeaders,
		ParsedResponseHeaders: internalSchema.ParsedResponseHeaders,
	}, nil
}

// FromDirectory creates a TelepactSchema from a directory of .telepact.json files
func FromDirectory(directory string) (*TelepactSchema, error) {
	fileMap, err := getSchemaFileMap(directory)
	if err != nil {
		return nil, err
	}
	return FromFileJSONMap(fileMap)
}

// getSchemaFileMap reads all .telepact.json files from a directory
func getSchemaFileMap(directory string) (map[string]string, error) {
	fileMap := make(map[string]string)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".telepact.json") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(directory, path)
			if err != nil {
				return err
			}

			fileMap[relPath] = string(content)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileMap, nil
}

// TelepactSchemaFiles manages schema files from a directory
type TelepactSchemaFiles struct {
	FilenamesToJSON map[string]string
}

// NewTelepactSchemaFiles creates a TelepactSchemaFiles from a directory
func NewTelepactSchemaFiles(directory string) (*TelepactSchemaFiles, error) {
	fileMap, err := getSchemaFileMap(directory)
	if err != nil {
		return nil, err
	}

	return &TelepactSchemaFiles{
		FilenamesToJSON: fileMap,
	}, nil
}

// ToTelepactSchema converts the schema files to a TelepactSchema
func (tsf *TelepactSchemaFiles) ToTelepactSchema() (*TelepactSchema, error) {
	return FromFileJSONMap(tsf.FilenamesToJSON)
}

// Unmarshal parses JSON data into a Go value
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Marshal converts a Go value to JSON
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
