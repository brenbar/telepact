package schema

import (
	"fmt"
	"regexp"

	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// ParseUnionType parses a union type definition from pseudo-JSON
func ParseUnionType(
	path []interface{},
	unionDefinitionAsPseudoJSON map[string]interface{},
	schemaKey string,
	ignoreKeys []string,
	requiredKeys []string,
	ctx *ParseContext,
) (*types.TUnion, error) {
	parseFailures := []*SchemaParseFailure{}

	otherKeys := make(map[string]bool)
	for k := range unionDefinitionAsPseudoJSON {
		otherKeys[k] = true
	}

	delete(otherKeys, schemaKey)
	delete(otherKeys, "///")
	for _, ignoreKey := range ignoreKeys {
		delete(otherKeys, ignoreKey)
	}

	if len(otherKeys) > 0 {
		for k := range otherKeys {
			loopPath := append(path, k)
			parseFailures = append(parseFailures, &SchemaParseFailure{
				DocumentName: ctx.DocumentName,
				Path:         loopPath,
				Reason:       "ObjectKeyDisallowed",
				Data:         map[string]interface{}{},
			})
		}
	}

	thisPath := append(path, schemaKey)
	defInit := unionDefinitionAsPseudoJSON[schemaKey]

	definitionSlice, ok := defInit.([]interface{})
	if !ok {
		finalParseFailures := GetTypeUnexpectedParseFailure(
			ctx.DocumentName, thisPath, defInit, "Array")
		parseFailures = append(parseFailures, finalParseFailures...)
		return nil, &TelepactSchemaParseErrorType{
			SchemaParseFailures:                     parseFailures,
			TelepactSchemaDocumentNamesToPseudoJSON: ctx.TelepactSchemaDocumentNamesToPseudoJSON,
		}
	}

	definition := []map[string]interface{}{}
	for index, element := range definitionSlice {
		loopPath := append(thisPath, index)
		elementMap, ok := element.(map[string]interface{})
		if !ok {
			thisParseFailures := GetTypeUnexpectedParseFailure(
				ctx.DocumentName, loopPath, element, "Object")
			parseFailures = append(parseFailures, thisParseFailures...)
			continue
		}
		definition = append(definition, elementMap)
	}

	if len(parseFailures) > 0 {
		return nil, &TelepactSchemaParseErrorType{
			SchemaParseFailures:                     parseFailures,
			TelepactSchemaDocumentNamesToPseudoJSON: ctx.TelepactSchemaDocumentNamesToPseudoJSON,
		}
	}

	if len(definition) == 0 {
		parseFailures = append(parseFailures, &SchemaParseFailure{
			DocumentName: ctx.DocumentName,
			Path:         thisPath,
			Reason:       "EmptyArrayDisallowed",
			Data:         map[string]interface{}{},
		})
	} else {
		// Check for required keys
		for _, requiredKey := range requiredKeys {
			found := false
			for _, element := range definition {
				tagKeys := make(map[string]bool)
				for k := range element {
					if k != "///" {
						tagKeys[k] = true
					}
				}
				if tagKeys[requiredKey] {
					found = true
					break
				}
			}
			if !found {
				branchPath := append(thisPath, 0)
				parseFailures = append(parseFailures, &SchemaParseFailure{
					DocumentName: ctx.DocumentName,
					Path:         branchPath,
					Reason:       "RequiredObjectKeyMissing",
					Data:         map[string]interface{}{"key": requiredKey},
				})
			}
		}
	}

	tags := make(map[string]*types.TStruct)
	tagIndices := make(map[string]int)

	regexString := `^([A-Z][a-zA-Z0-9_]*)$`
	regex := regexp.MustCompile(regexString)

	for i, element := range definition {
		loopPath := append(thisPath, i)
		mapInit := make(map[string]interface{})
		for k, v := range element {
			mapInit[k] = v
		}
		delete(mapInit, "///")
		keys := make([]string, 0, len(mapInit))
		for k := range mapInit {
			keys = append(keys, k)
		}

		matches := []string{}
		for _, k := range keys {
			if regex.MatchString(k) {
				matches = append(matches, k)
			}
		}

		if len(matches) != 1 {
			parseFailures = append(parseFailures, &SchemaParseFailure{
				DocumentName: ctx.DocumentName,
				Path:         loopPath,
				Reason:       "ObjectKeyRegexMatchCountUnexpected",
				Data: map[string]interface{}{
					"regex":    regexString,
					"actual":   len(matches),
					"expected": 1,
					"keys":     keys,
				},
			})
			continue
		}
		if len(mapInit) != 1 {
			parseFailures = append(parseFailures, &SchemaParseFailure{
				DocumentName: ctx.DocumentName,
				Path:         loopPath,
				Reason:       "ObjectSizeUnexpected",
				Data: map[string]interface{}{
					"expected": 1,
					"actual":   len(mapInit),
				},
			})
			continue
		}

		// Get the single tag entry
		var unionTag string
		var unionTagStruct map[string]interface{}
		for k, v := range mapInit {
			unionTag = k
			unionKeyPath := append(loopPath, unionTag)
			unionTagStructMap, ok := v.(map[string]interface{})
			if !ok {
				thisParseFailures := GetTypeUnexpectedParseFailure(
					ctx.DocumentName, unionKeyPath, v, "Object")
				parseFailures = append(parseFailures, thisParseFailures...)
				continue
			}
			unionTagStruct = unionTagStructMap
			break
		}

		if unionTagStruct == nil {
			continue
		}

		unionKeyPath := append(loopPath, unionTag)
		fields, err := ParseStructFields(unionKeyPath, unionTagStruct, false, ctx)
		if err != nil {
			if schemaErr, ok := err.(*TelepactSchemaParseErrorType); ok {
				parseFailures = append(parseFailures, schemaErr.SchemaParseFailures...)
				continue
			}
			return nil, err
		}

		unionStruct := types.NewTStruct(
			fmt.Sprintf("%s.%s", schemaKey, unionTag), fields)

		tags[unionTag] = unionStruct
		tagIndices[unionTag] = i
	}

	if len(parseFailures) > 0 {
		return nil, &TelepactSchemaParseErrorType{
			SchemaParseFailures:                     parseFailures,
			TelepactSchemaDocumentNamesToPseudoJSON: ctx.TelepactSchemaDocumentNamesToPseudoJSON,
		}
	}

	return types.NewTUnion(schemaKey, tags, tagIndices), nil
}
