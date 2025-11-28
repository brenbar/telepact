package mock

// IsSubMap checks if actual is a sub-map of expected.
// A sub-map means all keys in expected exist in actual with matching values.
// Uses deep comparison via IsSubMapEntryEqual.
func IsSubMap(expected, actual interface{}) bool {
	expectedMap, expectedOk := expected.(map[string]interface{})
	actualMap, actualOk := actual.(map[string]interface{})

	if !expectedOk || !actualOk {
		return false
	}

	// Check all expected keys exist in actual with matching values
	for key, expectedValue := range expectedMap {
		actualValue, exists := actualMap[key]
		if !exists {
			return false
		}
		if !IsSubMapEntryEqual(expectedValue, actualValue) {
			return false
		}
	}

	return true
}

// IsSubMapEntryEqual performs deep equality checking for map entries.
// Handles maps recursively, arrays element-wise, and primitives directly.
func IsSubMapEntryEqual(expected, actual interface{}) bool {
	// Handle nil cases
	if expected == nil {
		return actual == nil
	}
	if actual == nil {
		return false
	}

	// Handle map[string]interface{} recursively
	if expectedMap, ok := expected.(map[string]interface{}); ok {
		return IsSubMap(expectedMap, actual)
	}

	// Handle []interface{} arrays
	if expectedArr, ok := expected.([]interface{}); ok {
		actualArr, ok := actual.([]interface{})
		if !ok {
			return false
		}
		if len(expectedArr) != len(actualArr) {
			return false
		}
		for i := range expectedArr {
			if !IsSubMapEntryEqual(expectedArr[i], actualArr[i]) {
				return false
			}
		}
		return true
	}

	// Handle primitives (string, int, float64, bool)
	return expected == actual
}
