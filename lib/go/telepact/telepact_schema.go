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
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// TelepactSchema represents a parsed telepact schema
type TelepactSchema struct {
	Original              []interface{}
	Parsed                map[string]types.TType
	ParsedRequestHeaders  map[string]*types.TFieldDeclaration
	ParsedResponseHeaders map[string]*types.TFieldDeclaration
}

// FromJSON creates a TelepactSchema from a JSON string
func FromJSON(json string) (*TelepactSchema, error) {
	// TODO: Implementation
	return nil, &TelepactError{Message: "Not implemented"}
}

// FromFileJSONMap creates a TelepactSchema from a map of filenames to JSON strings
func FromFileJSONMap(fileJSONMap map[string]string) (*TelepactSchema, error) {
	// TODO: Implementation
	return nil, &TelepactError{Message: "Not implemented"}
}

// FromDirectory creates a TelepactSchema from a directory containing .telepact.json files
func FromDirectory(directory string) (*TelepactSchema, error) {
	// TODO: Implementation
	return nil, &TelepactError{Message: "Not implemented"}
}
