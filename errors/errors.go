package errors

import "fmt"

// FormBuilderError represents various errors that can occur in form building
type FormBuilderError struct {
	Type    string
	Field   string
	Message string
	Cause   error
}

func (e *FormBuilderError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s error for field '%s': %s", e.Type, e.Field, e.Message)
	}
	return fmt.Sprintf("%s error: %s", e.Type, e.Message)
}

func (e *FormBuilderError) Unwrap() error {
	return e.Cause
}

// Error type constants
const (
	ErrTypeDuplicateField = "DuplicateField"
	ErrTypeValidation     = "Validation"
	ErrTypeComponent      = "Component"
	ErrTypeConfig         = "Config"
	ErrTypeInitialization = "Initialization"
)

// NewDuplicateFieldError creates a new duplicate field error
func NewDuplicateFieldError(field string) *FormBuilderError {
	return &FormBuilderError{
		Type:    ErrTypeDuplicateField,
		Field:   field,
		Message: "field name must be unique within the form",
	}
}

// NewValidationError creates a new validation error
func NewValidationError(field, message string, cause error) *FormBuilderError {
	return &FormBuilderError{
		Type:    ErrTypeValidation,
		Field:   field,
		Message: message,
		Cause:   cause,
	}
}

// NewComponentError creates a new component error
func NewComponentError(message string, cause error) *FormBuilderError {
	return &FormBuilderError{
		Type:    ErrTypeComponent,
		Message: message,
		Cause:   cause,
	}
}

// NewConfigError creates a new configuration error
func NewConfigError(message string, cause error) *FormBuilderError {
	return &FormBuilderError{
		Type:    ErrTypeConfig,
		Message: message,
		Cause:   cause,
	}
}

// NewInitializationError creates a new initialization error
func NewInitializationError(message string, cause error) *FormBuilderError {
	return &FormBuilderError{
		Type:    ErrTypeInitialization,
		Message: message,
		Cause:   cause,
	}
}
