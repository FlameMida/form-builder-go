// Package components provides date/time and other complex components
package components

import (
	"fmt"
	"strings"

	"github.com/FlameMida/form-builder-go/contracts"
)

// DatePicker component implementation
type DatePicker struct {
	*BaseComponent
}

// NewDatePicker creates a new date picker component
func NewDatePicker(field, title string) *DatePicker {
	datePicker := &DatePicker{
		BaseComponent: NewBaseComponent(field, title),
	}
	// Set default placeholder
	datePicker.SetProp("placeholder", fmt.Sprintf("请选择%s", title))
	return datePicker
}

// Type sets the date picker type
func (d *DatePicker) Type(pickerType string) *DatePicker {
	d.SetProp("type", pickerType)
	return d
}

// Format sets the display format
func (d *DatePicker) Format(format string) *DatePicker {
	d.SetProp("format", format)
	return d
}

// ValueFormat sets the value format
func (d *DatePicker) ValueFormat(format string) *DatePicker {
	d.SetProp("value-format", format)
	return d
}

// Readonly sets readonly state
func (d *DatePicker) Readonly(readonly bool) *DatePicker {
	d.SetProp("readonly", readonly)
	return d
}

// Editable sets editable state
func (d *DatePicker) Editable(editable bool) *DatePicker {
	d.SetProp("editable", editable)
	return d
}

// Clearable makes the date picker clearable
func (d *DatePicker) Clearable(clearable bool) *DatePicker {
	d.SetProp("clearable", clearable)
	return d
}

// Size sets the date picker size
func (d *DatePicker) Size(size string) *DatePicker {
	d.SetProp("size", size)
	return d
}

// StartPlaceholder sets start placeholder for range picker
func (d *DatePicker) StartPlaceholder(text string) *DatePicker {
	d.SetProp("start-placeholder", text)
	return d
}

// EndPlaceholder sets end placeholder for range picker
func (d *DatePicker) EndPlaceholder(text string) *DatePicker {
	d.SetProp("end-placeholder", text)
	return d
}

// RangeSeparator sets range separator
func (d *DatePicker) RangeSeparator(separator string) *DatePicker {
	d.SetProp("range-separator", separator)
	return d
}

// DefaultValue sets default value
func (d *DatePicker) DefaultValue(value interface{}) *DatePicker {
	d.SetProp("default-value", value)
	return d
}

// DefaultTime sets default time
func (d *DatePicker) DefaultTime(time interface{}) *DatePicker {
	d.SetProp("default-time", time)
	return d
}

// Required makes the date picker required
func (d *DatePicker) Required() contracts.FormComponent {
	// Create conditional required rule based on date picker mode
	rule := NewConditionalRequiredRule(
		fmt.Sprintf("请选择%s", d.title),
		d,
		func(component interface{}) string {
			picker := component.(*DatePicker)
			// Check if it's range mode or multiple mode
			if dateType, exists := picker.props["type"]; exists {
				if typeStr, ok := dateType.(string); ok {
					if strings.Contains(typeStr, "range") {
						return "array" // Range modes use array validation
					}
				}
			}
			if multiple, exists := picker.props["multiple"]; exists {
				if val, ok := multiple.(bool); ok && val {
					return "array" // Multiple selection uses array validation
				}
			}
			return "date" // Single date selection uses date validation (fallback to string)
		},
	)
	d.AddValidateRule(rule)
	return d
}

// Placeholder sets placeholder
func (d *DatePicker) Placeholder(text string) contracts.FormComponent {
	d.SetProp("placeholder", text)
	return d
}

// Disabled sets the disabled state
func (d *DatePicker) Disabled(disabled bool) contracts.FormComponent {
	d.SetProp("disabled", disabled)
	return d
}

// Build returns the date picker component as a map
func (d *DatePicker) Build() map[string]interface{} {
	result := d.BaseComponent.Build()
	result["type"] = "el-date-picker"

	// Override value based on date picker mode
	if d.value == nil {
		// Check if it's range mode or multiple mode
		if dateType, exists := d.props["type"]; exists {
			if typeStr, ok := dateType.(string); ok {
				if strings.Contains(typeStr, "range") {
					result["value"] = []interface{}{} // Empty array for range modes
				} else {
					result["value"] = "" // Empty string for single date
				}
			} else {
				result["value"] = "" // Default to empty string
			}
		} else if multiple, exists := d.props["multiple"]; exists {
			if val, ok := multiple.(bool); ok && val {
				result["value"] = []interface{}{} // Empty array for multiple selection
			} else {
				result["value"] = "" // Empty string for single date
			}
		} else {
			result["value"] = "" // Default to empty string
		}
	}

	return result
}

// TimePicker component implementation
type TimePicker struct {
	*BaseComponent
}

// NewTimePicker creates a new time picker component
func NewTimePicker(field, title string) *TimePicker {
	timePicker := &TimePicker{
		BaseComponent: NewBaseComponent(field, title),
	}
	// Set default placeholder
	timePicker.SetProp("placeholder", fmt.Sprintf("请选择%s", title))
	return timePicker
}

// IsRange enables range mode
func (t *TimePicker) IsRange(isRange bool) *TimePicker {
	t.SetProp("is-range", isRange)
	return t
}

// ArrowControl enables arrow control
func (t *TimePicker) ArrowControl(arrow bool) *TimePicker {
	t.SetProp("arrow-control", arrow)
	return t
}

// Format sets the time format
func (t *TimePicker) Format(format string) *TimePicker {
	t.SetProp("format", format)
	return t
}

// ValueFormat sets the value format
func (t *TimePicker) ValueFormat(format string) *TimePicker {
	t.SetProp("value-format", format)
	return t
}

// Readonly sets readonly state
func (t *TimePicker) Readonly(readonly bool) *TimePicker {
	t.SetProp("readonly", readonly)
	return t
}

// Editable sets editable state
func (t *TimePicker) Editable(editable bool) *TimePicker {
	t.SetProp("editable", editable)
	return t
}

// Clearable makes the time picker clearable
func (t *TimePicker) Clearable(clearable bool) *TimePicker {
	t.SetProp("clearable", clearable)
	return t
}

// Size sets the time picker size
func (t *TimePicker) Size(size string) *TimePicker {
	t.SetProp("size", size)
	return t
}

// StartPlaceholder sets start placeholder for range picker
func (t *TimePicker) StartPlaceholder(text string) *TimePicker {
	t.SetProp("start-placeholder", text)
	return t
}

// EndPlaceholder sets end placeholder for range picker
func (t *TimePicker) EndPlaceholder(text string) *TimePicker {
	t.SetProp("end-placeholder", text)
	return t
}

// RangeSeparator sets range separator
func (t *TimePicker) RangeSeparator(separator string) *TimePicker {
	t.SetProp("range-separator", separator)
	return t
}

// Required makes the time picker required
func (t *TimePicker) Required() contracts.FormComponent {
	// Create conditional required rule based on range mode
	rule := NewConditionalRequiredRule(
		fmt.Sprintf("请选择%s", t.title),
		t,
		func(component interface{}) string {
			picker := component.(*TimePicker)
			if isRange, exists := picker.props["is-range"]; exists {
				if val, ok := isRange.(bool); ok && val {
					return "array" // Range mode uses array validation
				}
			}
			return "string" // Single time selection uses string validation
		},
	)
	t.AddValidateRule(rule)
	return t
}

// Placeholder sets placeholder
func (t *TimePicker) Placeholder(text string) contracts.FormComponent {
	t.SetProp("placeholder", text)
	return t
}

// Disabled sets the disabled state
func (t *TimePicker) Disabled(disabled bool) contracts.FormComponent {
	t.SetProp("disabled", disabled)
	return t
}

// Build returns the time picker component as a map
func (t *TimePicker) Build() map[string]interface{} {
	result := t.BaseComponent.Build()
	result["type"] = "el-time-picker"

	// Override value based on range mode
	if t.value == nil {
		if isRange, exists := t.props["is-range"]; exists {
			if val, ok := isRange.(bool); ok && val {
				result["value"] = []interface{}{} // Empty array for range mode
			} else {
				result["value"] = "" // Empty string for single time
			}
		} else {
			result["value"] = "" // Default to empty string
		}
	}

	return result
}

// ColorPicker component implementation
type ColorPicker struct {
	*BaseComponent
}

// NewColorPicker creates a new color picker component
func NewColorPicker(field, title string) *ColorPicker {
	colorPicker := &ColorPicker{
		BaseComponent: NewBaseComponent(field, title),
	}
	// Set default placeholder
	colorPicker.SetProp("placeholder", fmt.Sprintf("请选择%s", title))
	return colorPicker
}

// ShowAlpha enables alpha selection
func (c *ColorPicker) ShowAlpha(show bool) *ColorPicker {
	c.SetProp("show-alpha", show)
	return c
}

// ColorFormat sets the color format
func (c *ColorPicker) ColorFormat(format string) *ColorPicker {
	c.SetProp("color-format", format)
	return c
}

// Predefine sets predefined colors
func (c *ColorPicker) Predefine(colors []string) *ColorPicker {
	c.SetProp("predefine", colors)
	return c
}

// Size sets the color picker size
func (c *ColorPicker) Size(size string) *ColorPicker {
	c.SetProp("size", size)
	return c
}

// Required makes the color picker required
func (c *ColorPicker) Required() contracts.FormComponent {
	c.AddValidateRule(NewStringRequiredRule(fmt.Sprintf("请选择%s", c.title)))
	return c
}

// Placeholder sets placeholder (not applicable for color picker, but required by interface)
func (c *ColorPicker) Placeholder(text string) contracts.FormComponent {
	return c
}

// Disabled sets the disabled state
func (c *ColorPicker) Disabled(disabled bool) contracts.FormComponent {
	c.SetProp("disabled", disabled)
	return c
}

// Build returns the color picker component as a map
func (c *ColorPicker) Build() map[string]interface{} {
	result := c.BaseComponent.Build()
	result["type"] = "el-color-picker"
	return result
}
