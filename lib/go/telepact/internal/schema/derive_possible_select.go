package schema

import (
	"sort"

	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// DerivePossibleSelect derives the possible select structure for a function
func DerivePossibleSelect(fnName string, result *types.TUnion) map[string]interface{} {
	nestedTypes := make(map[string]types.TType)
	okFields := result.Tags["Ok_"].Fields

	okFieldNames := make([]string, 0, len(okFields))
	for name := range okFields {
		okFieldNames = append(okFieldNames, name)
	}
	sort.Strings(okFieldNames)

	for _, fieldDecl := range okFields {
		findNestedTypes(fieldDecl.TypeDeclaration, nestedTypes)
	}

	possibleSelect := make(map[string]interface{})

	possibleSelect["->"] = map[string]interface{}{
		"Ok_": okFieldNames,
	}

	sortedTypeKeys := make([]string, 0, len(nestedTypes))
	for k := range nestedTypes {
		sortedTypeKeys = append(sortedTypeKeys, k)
	}
	sort.Strings(sortedTypeKeys)

	for _, k := range sortedTypeKeys {
		if k[:3] == "fn." {
			continue
		}

		v := nestedTypes[k]
		switch typ := v.(type) {
		case *types.TUnion:
			unionSelect := make(map[string][]string)
			sortedTagKeys := make([]string, 0, len(typ.Tags))
			for tagKey := range typ.Tags {
				sortedTagKeys = append(sortedTagKeys, tagKey)
			}
			sort.Strings(sortedTagKeys)

			for _, c := range sortedTagKeys {
				tagStruct := typ.Tags[c]
				selectedFieldNames := make([]string, 0, len(tagStruct.Fields))
				for fieldName := range tagStruct.Fields {
					selectedFieldNames = append(selectedFieldNames, fieldName)
				}
				sort.Strings(selectedFieldNames)

				if len(selectedFieldNames) > 0 {
					unionSelect[c] = selectedFieldNames
				}
			}

			possibleSelect[k] = unionSelect

		case *types.TStruct:
			structSelect := make([]string, 0, len(typ.Fields))
			for fieldName := range typ.Fields {
				structSelect = append(structSelect, fieldName)
			}
			sort.Strings(structSelect)

			if len(structSelect) > 0 {
				possibleSelect[k] = structSelect
			}
		}
	}

	return possibleSelect
}

// findNestedTypes recursively finds nested types in a type declaration
func findNestedTypes(typeDeclaration *types.TTypeDeclaration, nestedTypes map[string]types.TType) {
	typ := typeDeclaration.Type

	switch t := typ.(type) {
	case *types.TUnion:
		nestedTypes[t.Name] = t
		for _, tag := range t.Tags {
			for _, fieldDecl := range tag.Fields {
				findNestedTypes(fieldDecl.TypeDeclaration, nestedTypes)
			}
		}

	case *types.TStruct:
		nestedTypes[t.Name] = t
		for _, fieldDecl := range t.Fields {
			findNestedTypes(fieldDecl.TypeDeclaration, nestedTypes)
		}

	case *types.TArray, *types.TObject:
		if len(typeDeclaration.TypeParameters) > 0 {
			findNestedTypes(typeDeclaration.TypeParameters[0], nestedTypes)
		}
	}
}
