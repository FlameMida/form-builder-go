// Package components provides specialized validation rules for form components
package components

import (
	"fmt"
	"strings"

	"github.com/FlameMida/form-builder-go/contracts"
)

// ComponentRequiredRule implements a required validation rule for components
// with appropriate type based on component characteristics
type ComponentRequiredRule struct {
	message        string
	validationType string
}

// NewArrayRequiredRule creates a required rule for array-type components (always array validation)
func NewArrayRequiredRule(message string) contracts.ValidateRule {
	return &ComponentRequiredRule{
		message:        message,
		validationType: "array",
	}
}

// NewStringRequiredRule creates a required rule for string-type components (always string validation)
func NewStringRequiredRule(message string) contracts.ValidateRule {
	return &ComponentRequiredRule{
		message:        message,
		validationType: "string",
	}
}

// NewNumberRequiredRule creates a required rule for number-type components (always number validation)
func NewNumberRequiredRule(message string) contracts.ValidateRule {
	return &ComponentRequiredRule{
		message:        message,
		validationType: "number",
	}
}

// Type returns the validation type
func (r *ComponentRequiredRule) Type() string {
	return r.validationType
}

// Message returns the validation message
func (r *ComponentRequiredRule) Message() string {
	return r.message
}

// Validate validates that the value is not empty according to the validation type
func (r *ComponentRequiredRule) Validate(value interface{}) error {
	if value == nil {
		return fmt.Errorf("%s", r.message)
	}

	switch r.validationType {
	case "array":
		// For array validation, value should be a non-empty array
		switch v := value.(type) {
		case []interface{}:
			if len(v) == 0 {
				return fmt.Errorf("%s", r.message)
			}
		case []string:
			if len(v) == 0 {
				return fmt.Errorf("%s", r.message)
			}
		case []int:
			if len(v) == 0 {
				return fmt.Errorf("%s", r.message)
			}
		default:
			return fmt.Errorf("%s", r.message)
		}
	case "string":
		// For string validation, value should be a non-empty string
		switch v := value.(type) {
		case string:
			if strings.TrimSpace(v) == "" {
				return fmt.Errorf("%s", r.message)
			}
		default:
			if v == "" || v == nil {
				return fmt.Errorf("%s", r.message)
			}
		}
	case "number":
		// For number validation, value should be a valid number
		switch v := value.(type) {
		case int, float64, float32, int32, int64:
			// Numbers are always valid if not nil
		case string:
			if strings.TrimSpace(v) == "" {
				return fmt.Errorf("%s", r.message)
			}
			// Could add number parsing validation here if needed
		default:
			return fmt.Errorf("%s", r.message)
		}
	default:
		// Default to string validation
		if v, ok := value.(string); ok {
			if strings.TrimSpace(v) == "" {
				return fmt.Errorf("%s", r.message)
			}
		} else if value == "" || value == nil {
			return fmt.Errorf("%s", r.message)
		}
	}

	return nil
}

// ToMap returns the rule as a map for JSON serialization
func (r *ComponentRequiredRule) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"required": true,
		"message":  r.message,
		"type":     r.Type(),
		"trigger":  "change",
	}
}

// ConditionalRequiredRule implements a required validation rule with conditional type logic
type ConditionalRequiredRule struct {
	message       string
	component     interface{}              // Reference to the component for checking properties
	typeCheckFunc func(interface{}) string // Function to determine validation type
}

// NewConditionalRequiredRule creates a conditional required rule
func NewConditionalRequiredRule(message string, component interface{}, typeCheckFunc func(interface{}) string) contracts.ValidateRule {
	return &ConditionalRequiredRule{
		message:       message,
		component:     component,
		typeCheckFunc: typeCheckFunc,
	}
}

// Type returns the validation type based on component state
func (r *ConditionalRequiredRule) Type() string {
	return r.typeCheckFunc(r.component)
}

// Message returns the validation message
func (r *ConditionalRequiredRule) Message() string {
	return r.message
}

// Validate validates that the value is not empty according to the dynamic validation type
func (r *ConditionalRequiredRule) Validate(value interface{}) error {
	validationType := r.Type()

	if value == nil {
		return fmt.Errorf("%s", r.message)
	}

	switch validationType {
	case "array":
		switch v := value.(type) {
		case []interface{}:
			if len(v) == 0 {
				return fmt.Errorf("%s", r.message)
			}
		case []string:
			if len(v) == 0 {
				return fmt.Errorf("%s", r.message)
			}
		case []int:
			if len(v) == 0 {
				return fmt.Errorf("%s", r.message)
			}
		default:
			return fmt.Errorf("%s", r.message)
		}
	case "string":
		switch v := value.(type) {
		case string:
			if strings.TrimSpace(v) == "" {
				return fmt.Errorf("%s", r.message)
			}
		default:
			if v == "" || v == nil {
				return fmt.Errorf("%s", r.message)
			}
		}
	case "number":
		switch v := value.(type) {
		case int, float64, float32, int32, int64:
			// Numbers are always valid if not nil
		case string:
			if strings.TrimSpace(v) == "" {
				return fmt.Errorf("%s", r.message)
			}
		default:
			return fmt.Errorf("%s", r.message)
		}
	default:
		// Default to string validation
		if v, ok := value.(string); ok {
			if strings.TrimSpace(v) == "" {
				return fmt.Errorf("%s", r.message)
			}
		} else if value == "" || value == nil {
			return fmt.Errorf("%s", r.message)
		}
	}

	return nil
}

// ToMap returns the rule as a map for JSON serialization
func (r *ConditionalRequiredRule) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"required": true,
		"message":  r.message,
		"type":     r.Type(),
		"trigger":  "change",
	}
}
