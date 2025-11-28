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
	"regexp"
	"sort"

	"github.com/brenbar/telepact/lib/go/telepact"
)

// FindSchemaKey finds the schema key in a definition object.
func FindSchemaKey(documentName string, definition map[string]interface{}, index int, documentNamesToJSON map[string]string) (string, error) {
	regexPattern := `^(((fn|errors|headers|info)|((struct|union|_ext)(<[0-2]>)?))\..*)`
	regex := regexp.MustCompile(regexPattern)
	
	matches := []string{}
	keys := make([]string, 0, len(definition))
	for k := range definition {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	
	for _, key := range keys {
		if regex.MatchString(key) {
			matches = append(matches, key)
		}
	}
	
	if len(matches) == 1 {
		return matches[0], nil
	}
	
	parseFailure := NewSchemaParseFailure(
		documentName,
		[]interface{}{index},
		"ObjectKeyRegexMatchCountUnexpected",
		map[string]interface{}{
			"regex":    regexPattern,
			"actual":   len(matches),
			"expected": 1,
			"keys":     keys,
		},
	)
	
	return "", telepact.NewTelepactSchemaParseError([]*SchemaParseFailure{parseFailure}, documentNamesToJSON)
}
