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
	"encoding/json"
	"sort"

	"github.com/brenbar/telepact/lib/go/telepact"
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// ParseTelepactSchema parses the Telepact schema from JSON documents.
// This is a stub implementation that will be expanded.
func ParseTelepactSchema(telepactSchemaDocumentNamesToJSON map[string]string) (*telepact.TelepactSchema, error) {
	// Initialize state
	originalSchema := make(map[string]interface{})
	parsedTypes := make(map[string]types.TType)
	fnErrorRegexes := make(map[string]string) // Will be used in full implementation
	parseFailures := []*SchemaParseFailure{}
	failedTypes := make(map[string]bool) // Will be used in full implementation
	schemaKeysToDocumentNames := make(map[string]string)
	schemaKeysToIndex := make(map[string]int)
	schemaKeys := make(map[string]bool)
	
	_ = fnErrorRegexes // Used in full implementation
	_ = failedTypes    // Used in full implementation
	
	// Sort document names for deterministic order
	orderedDocumentNames := make([]string, 0, len(telepactSchemaDocumentNamesToJSON))
	for name := range telepactSchemaDocumentNamesToJSON {
		orderedDocumentNames = append(orderedDocumentNames, name)
	}
	sort.Strings(orderedDocumentNames)
	
	// Parse JSON documents to pseudo-JSON
	telepactSchemaDocumentNameToPseudoJSON := make(map[string][]interface{})
	for documentName, telepactSchemaJSON := range telepactSchemaDocumentNamesToJSON {
		var telepactSchemaPseudoJSONInit interface{}
		err := json.Unmarshal([]byte(telepactSchemaJSON), &telepactSchemaPseudoJSONInit)
		if err != nil {
			return nil, telepact.NewTelepactSchemaParseError(
				[]*SchemaParseFailure{
					NewSchemaParseFailure(documentName, []interface{}{}, "JsonInvalid", map[string]interface{}{}),
				},
				telepactSchemaDocumentNamesToJSON,
			)
		}
		
		// Check if it's an array
		telepactSchemaPseudoJSON, ok := telepactSchemaPseudoJSONInit.([]interface{})
		if !ok {
			thisParseFailure := GetTypeUnexpectedParseFailure(
				documentName, []interface{}{}, telepactSchemaPseudoJSONInit, "Array",
			)
			return nil, telepact.NewTelepactSchemaParseError(
				thisParseFailure,
				telepactSchemaDocumentNamesToJSON,
			)
		}
		
		telepactSchemaDocumentNameToPseudoJSON[documentName] = telepactSchemaPseudoJSON
	}
	
	// Process each document to find schema keys
	for _, documentName := range orderedDocumentNames {
		telepactSchemaPseudoJSON := telepactSchemaDocumentNameToPseudoJSON[documentName]
		
		for index, definition := range telepactSchemaPseudoJSON {
			loopPath := []interface{}{index}
			
			// Check if definition is an object
			def, ok := definition.(map[string]interface{})
			if !ok {
				thisParseFailures := GetTypeUnexpectedParseFailure(
					documentName, loopPath, definition, "Object",
				)
				parseFailures = append(parseFailures, thisParseFailures...)
				continue
			}
			
			// Find schema key
			schemaKey, err := FindSchemaKey(documentName, def, index, telepactSchemaDocumentNamesToJSON)
			if err != nil {
				if tspErr, ok := err.(*telepact.TelepactSchemaParseError); ok {
					if failures, ok := tspErr.SchemaParseFailures.([]*SchemaParseFailure); ok {
						parseFailures = append(parseFailures, failures...)
					}
				}
				continue
			}
			
			// Check for collisions
			if matchingSchemaKey := FindMatchingSchemaKey(schemaKeys, schemaKey); matchingSchemaKey != nil {
				// TODO: Report collision
				continue
			}
			
			schemaKeys[schemaKey] = true
			schemaKeysToIndex[schemaKey] = index
			schemaKeysToDocumentNames[schemaKey] = documentName
			if documentName == "auto_" || documentName == "auth_" || !endsWith(documentName, "_") {
				originalSchema[schemaKey] = def
			}
		}
	}
	
	if len(parseFailures) > 0 {
		return nil, telepact.NewTelepactSchemaParseError(parseFailures, telepactSchemaDocumentNamesToJSON)
	}
	
	// TODO: Parse all types, errors, and headers
	// For now, create an empty schema
	
	// Build final schema
	finalOriginalSchema := make([]interface{}, 0, len(originalSchema))
	schemaKeysList := make([]string, 0, len(originalSchema))
	for k := range originalSchema {
		schemaKeysList = append(schemaKeysList, k)
	}
	sort.Slice(schemaKeysList, func(i, j int) bool {
		ki, kj := schemaKeysList[i], schemaKeysList[j]
		// Sort with info. first
		if startsWithInfo(ki) != startsWithInfo(kj) {
			return startsWithInfo(ki)
		}
		return ki < kj
	})
	
	for _, k := range schemaKeysList {
		finalOriginalSchema = append(finalOriginalSchema, originalSchema[k])
	}
	
	requestHeaders := make(map[string]*types.TFieldDeclaration)
	responseHeaders := make(map[string]*types.TFieldDeclaration)
	
	return telepact.NewTelepactSchema(
		finalOriginalSchema,
		parsedTypes,
		requestHeaders,
		responseHeaders,
	), nil
}

func endsWith(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func startsWithInfo(s string) bool {
	return len(s) >= 5 && s[:5] == "info."
}
