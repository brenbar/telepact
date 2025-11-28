// Copyright The Telepact Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/util"
)

// GetTypeUnexpectedParseFailure creates a parse failure for unexpected types.
func GetTypeUnexpectedParseFailure(documentName string, path []interface{}, value interface{}, expectedTypeName string) []*SchemaParseFailure {
	actualTypeName := util.GetType(value)
	
	return []*SchemaParseFailure{
		NewSchemaParseFailure(
			documentName,
			path,
			"TypeUnexpected",
			map[string]interface{}{
				"expected": expectedTypeName,
				"actual":   actualTypeName,
			},
		),
	}
}
