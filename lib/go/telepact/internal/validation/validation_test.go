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

package validation_test

import (
	"testing"

	"github.com/brenbar/telepact/lib/go/telepact/internal/validation"
)

func TestValidateString(t *testing.T) {
	// Valid string
	failures := validation.ValidateString("hello")
	if len(failures) != 0 {
		t.Errorf("Expected no failures for valid string, got %d", len(failures))
	}

	// Invalid: number
	failures = validation.ValidateString(123)
	if len(failures) == 0 {
		t.Error("Expected failure for number instead of string")
	}
}

func TestValidateInteger(t *testing.T) {
	// Valid integer
	failures := validation.ValidateInteger(42)
	if len(failures) != 0 {
		t.Errorf("Expected no failures for valid integer, got %d", len(failures))
	}

	// Valid: float64 that's actually an integer
	failures = validation.ValidateInteger(float64(42))
	if len(failures) != 0 {
		t.Errorf("Expected no failures for integer-valued float64, got %d", len(failures))
	}

	// Invalid: float64 that's not an integer
	failures = validation.ValidateInteger(42.5)
	if len(failures) == 0 {
		t.Error("Expected failure for non-integer float64")
	}

	// Invalid: string
	failures = validation.ValidateInteger("not a number")
	if len(failures) == 0 {
		t.Error("Expected failure for string instead of integer")
	}
}

func TestValidateBoolean(t *testing.T) {
	// Valid boolean
	failures := validation.ValidateBoolean(true)
	if len(failures) != 0 {
		t.Errorf("Expected no failures for valid boolean, got %d", len(failures))
	}

	// Invalid: string
	failures = validation.ValidateBoolean("true")
	if len(failures) == 0 {
		t.Error("Expected failure for string instead of boolean")
	}
}

func TestValidateNumber(t *testing.T) {
	// Valid number (integer)
	failures := validation.ValidateNumber(42)
	if len(failures) != 0 {
		t.Errorf("Expected no failures for valid integer number, got %d", len(failures))
	}

	// Valid number (float)
	failures = validation.ValidateNumber(3.14)
	if len(failures) != 0 {
		t.Errorf("Expected no failures for valid float number, got %d", len(failures))
	}

	// Invalid: string
	failures = validation.ValidateNumber("123")
	if len(failures) == 0 {
		t.Error("Expected failure for string instead of number")
	}
}

func TestValidateContext(t *testing.T) {
	ctx := validation.NewValidateContext(nil, "fn.test", false)
	
	if ctx == nil {
		t.Fatal("Expected non-nil ValidateContext")
	}

	if ctx.Fn != "fn.test" {
		t.Errorf("Expected Fn to be 'fn.test', got '%s'", ctx.Fn)
	}

	if ctx.CoerceBase64 != false {
		t.Error("Expected CoerceBase64 to be false")
	}

	if len(ctx.Path) != 0 {
		t.Errorf("Expected empty Path, got length %d", len(ctx.Path))
	}
}
