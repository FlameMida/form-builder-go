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
	return &DatePicker{
		BaseComponent: NewBaseComponent(field, title),
	}
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
	return &TimePicker{
		BaseComponent: NewBaseComponent(field, title),
	}
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
	return &ColorPicker{
		BaseComponent: NewBaseComponent(field, title),
	}
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

// Upload component implementation
type Upload struct {
	*BaseComponent
}

// NewUpload creates a new upload component
func NewUpload(field, title string) *Upload {
	return &Upload{
		BaseComponent: NewBaseComponent(field, title),
	}
}

// Action sets the upload action URL
func (u *Upload) Action(action string) *Upload {
	u.SetProp("action", action)
	return u
}

// Headers sets upload headers
func (u *Upload) Headers(headers map[string]string) *Upload {
	u.SetProp("headers", headers)
	return u
}

// Multiple enables multiple file upload
func (u *Upload) Multiple(multiple bool) *Upload {
	u.SetProp("multiple", multiple)
	return u
}

// Data sets additional data
func (u *Upload) Data(data map[string]interface{}) *Upload {
	u.SetProp("data", data)
	return u
}

// Name sets the file field name
func (u *Upload) Name(name string) *Upload {
	u.SetProp("name", name)
	return u
}

// WithCredentials enables credentials
func (u *Upload) WithCredentials(with bool) *Upload {
	u.SetProp("with-credentials", with)
	return u
}

// ShowFileList shows file list
func (u *Upload) ShowFileList(show bool) *Upload {
	u.SetProp("show-file-list", show)
	return u
}

// Drag enables drag upload
func (u *Upload) Drag(drag bool) *Upload {
	u.SetProp("drag", drag)
	return u
}

// Accept sets accepted file types
func (u *Upload) Accept(accept string) *Upload {
	u.SetProp("accept", accept)
	return u
}

// OnPreview sets preview handler
func (u *Upload) OnPreview(handler string) *Upload {
	u.SetProp("on-preview", handler)
	return u
}

// OnRemove sets remove handler
func (u *Upload) OnRemove(handler string) *Upload {
	u.SetProp("on-remove", handler)
	return u
}

// OnSuccess sets success handler
func (u *Upload) OnSuccess(handler string) *Upload {
	u.SetProp("on-success", handler)
	return u
}

// OnError sets error handler
func (u *Upload) OnError(handler string) *Upload {
	u.SetProp("on-error", handler)
	return u
}

// OnProgress sets progress handler
func (u *Upload) OnProgress(handler string) *Upload {
	u.SetProp("on-progress", handler)
	return u
}

// OnChange sets change handler
func (u *Upload) OnChange(handler string) *Upload {
	u.SetProp("on-change", handler)
	return u
}

// BeforeUpload sets before upload handler
func (u *Upload) BeforeUpload(handler string) *Upload {
	u.SetProp("before-upload", handler)
	return u
}

// BeforeRemove sets before remove handler
func (u *Upload) BeforeRemove(handler string) *Upload {
	u.SetProp("before-remove", handler)
	return u
}

// ListType sets list type
func (u *Upload) ListType(listType string) *Upload {
	u.SetProp("list-type", listType)
	return u
}

// AutoUpload enables auto upload
func (u *Upload) AutoUpload(auto bool) *Upload {
	u.SetProp("auto-upload", auto)
	return u
}

// FileList sets file list
func (u *Upload) FileList(files []interface{}) *Upload {
	u.SetProp("file-list", files)
	return u
}

// HttpRequest sets custom HTTP request
func (u *Upload) HttpRequest(request string) *Upload {
	u.SetProp("http-request", request)
	return u
}

// Limit sets file limit
func (u *Upload) Limit(limit int) *Upload {
	u.SetProp("limit", limit)
	return u
}

// OnExceed sets exceed handler
func (u *Upload) OnExceed(handler string) *Upload {
	u.SetProp("on-exceed", handler)
	return u
}

// Required makes the upload required
func (u *Upload) Required() contracts.FormComponent {
	// Create conditional required rule based on upload limit
	rule := NewConditionalRequiredRule(
		fmt.Sprintf("请上传%s", u.title),
		u,
		func(component interface{}) string {
			upload := component.(*Upload)
			if limit, exists := upload.props["limit"]; exists {
				if val, ok := limit.(int); ok && val == 1 {
					return "string" // Single file upload uses string validation
				}
			}
			return "array" // Multiple file upload uses array validation
		},
	)
	u.AddValidateRule(rule)
	return u
}

// Placeholder sets placeholder (not applicable for upload, but required by interface)
func (u *Upload) Placeholder(text string) contracts.FormComponent {
	return u
}

// Disabled sets the disabled state
func (u *Upload) Disabled(disabled bool) contracts.FormComponent {
	u.SetProp("disabled", disabled)
	return u
}

// Build returns the upload component as a map
func (u *Upload) Build() map[string]interface{} {
	result := u.BaseComponent.Build()
	result["type"] = "el-upload"

	// Override value based on upload limit
	if u.value == nil {
		if limit, exists := u.props["limit"]; exists {
			if val, ok := limit.(int); ok && val == 1 {
				result["value"] = "" // Empty string for single file upload
			} else {
				result["value"] = []interface{}{} // Empty array for multiple file upload
			}
		} else {
			result["value"] = []interface{}{} // Default to empty array
		}
	}

	return result
}
