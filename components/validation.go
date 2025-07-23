package components

import (
	"github.com/FlameMida/form-builder-go/contracts"
	formbuildererrors "github.com/FlameMida/form-builder-go/errors"
)

// ValidateComponents validates multiple components and returns the first error encountered
func ValidateComponents(components []contracts.Component) error {
	for _, component := range components {
		if err := component.DoValidate(); err != nil {
			return err
		}
	}
	return nil
}

// ValidateForm validates all components in a form
func ValidateForm(form contracts.Form) error {
	rules := form.FormRule()
	for _, rule := range rules {
		field, ok := rule["field"].(string)
		if !ok || field == "" {
			continue
		}

		validateRules, ok := rule["validate"].([]map[string]interface{})
		if !ok {
			continue
		}

		value := rule["value"]
		for _, vRule := range validateRules {
			if required, ok := vRule["required"].(bool); ok && required {
				if value == nil || value == "" {
					message, _ := vRule["message"].(string)
					if message == "" {
						message = "field is required"
					}
					return formbuildererrors.NewValidationError(field, message, nil)
				}
			}
		}
	}
	return nil
}
