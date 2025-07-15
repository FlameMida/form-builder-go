// Package components provides additional form components
package components

import (
	"fmt"
	"strings"

	"github.com/FlameMida/form-builder-go/contracts"
)

// Select component implementation
type Select struct {
	*BaseComponent
	options []contracts.Option
}

// NewSelect creates a new select component
func NewSelect(field, title string) *Select {
	select_ := &Select{
		BaseComponent: NewBaseComponent(field, title),
		options:       make([]contracts.Option, 0),
	}
	// Set default placeholder
	select_.SetProp("placeholder", fmt.Sprintf("请选择%s", title))
	return select_
}

// Options sets multiple options at once
func (s *Select) Options(options []contracts.Option) *Select {
	s.options = options
	return s
}

// AddOption adds a single option
func (s *Select) AddOption(value interface{}, label string) *Select {
	s.options = append(s.options, contracts.Option{
		Value: value,
		Label: label,
	})
	return s
}

// AddOptions adds multiple options from a struct slice
func (s *Select) AddOptions(options []contracts.Option) *Select {
	s.options = append(s.options, options...)
	return s
}

// Multiple enables multiple selection
func (s *Select) Multiple(multiple bool) *Select {
	s.SetProp("multiple", multiple)
	return s
}

// Clearable makes the select clearable
func (s *Select) Clearable(clearable bool) *Select {
	s.SetProp("clearable", clearable)
	return s
}

// Filterable enables filtering
func (s *Select) Filterable(filterable bool) *Select {
	s.SetProp("filterable", filterable)
	return s
}

// Remote enables remote data loading
func (s *Select) Remote(remote bool) *Select {
	s.SetProp("remote", remote)
	return s
}

// RemoteMethod sets the remote method
func (s *Select) RemoteMethod(method string) *Select {
	s.SetProp("remote-method", method)
	return s
}

// Loading sets the loading state
func (s *Select) Loading(loading bool) *Select {
	s.SetProp("loading", loading)
	return s
}

// Size sets the select size
func (s *Select) Size(size string) *Select {
	s.SetProp("size", size)
	return s
}

// Required makes the select required
func (s *Select) Required() *Select {
	// Create required rule with appropriate type based on multiple selection
	rule := &SelectRequiredRule{
		message:    fmt.Sprintf("请选择%s", s.title),
		isMultiple: s.isMultiple(),
	}
	s.AddValidateRule(rule)
	return s
}

// isMultiple checks if the select is in multiple mode
func (s *Select) isMultiple() bool {
	if multiple, exists := s.props["multiple"]; exists {
		if val, ok := multiple.(bool); ok {
			return val
		}
	}
	return false
}

// Placeholder sets placeholder (not typically used for select, but required by interface)
func (s *Select) Placeholder(text string) *Select {
	s.SetProp("placeholder", text)
	return s
}

// Disabled sets the disabled state
func (s *Select) Disabled(disabled bool) *Select {
	s.SetProp("disabled", disabled)
	return s
}

// Build returns the select component as a map
func (s *Select) Build() map[string]interface{} {
	result := s.BaseComponent.Build()
	result["type"] = "el-select"

	// Add options to props
	if len(s.options) > 0 {
		result["options"] = s.options
	}

	// Override value based on multiple selection mode
	if s.value == nil {
		if s.isMultiple() {
			result["value"] = []interface{}{} // Empty array for multiple selection
		} else {
			result["value"] = "" // Empty string for single selection
		}
	}

	return result
}

// ControlRule represents a control rule for conditional rendering
type ControlRule struct {
	Value interface{}           `json:"value"`
	Rule  []contracts.Component `json:"rule"`
}

// SuffixElement represents a suffix element
type SuffixElement struct {
	Type     string                 `json:"type"`
	Style    map[string]interface{} `json:"style,omitempty"`
	DomProps map[string]interface{} `json:"domProps,omitempty"`
}

// Radio component implementation
type Radio struct {
	*BaseComponent
	options []contracts.Option
	control []ControlRule
	suffix  *SuffixElement
}

// NewRadio creates a new radio component
func NewRadio(field, title string) *Radio {
	return &Radio{
		BaseComponent: NewBaseComponent(field, title),
		options:       make([]contracts.Option, 0),
		control:       make([]ControlRule, 0),
		suffix:        nil,
	}
}

// Options sets multiple options at once
func (r *Radio) Options(options []contracts.Option) *Radio {
	r.options = options
	return r
}

// AddOption adds a single option
func (r *Radio) AddOption(value interface{}, label string) *Radio {
	r.options = append(r.options, contracts.Option{
		Value: value,
		Label: label,
	})
	return r
}

// AddOptions adds multiple options from a struct slice
func (r *Radio) AddOptions(options []contracts.Option) *Radio {
	r.options = append(r.options, options...)
	return r
}

// Control adds control rules for conditional rendering
func (r *Radio) Control(controlRules []ControlRule) *Radio {
	r.control = controlRules
	return r
}

// AddControl adds a single control rule
func (r *Radio) AddControl(value interface{}, rules []contracts.Component) *Radio {
	r.control = append(r.control, ControlRule{
		Value: value,
		Rule:  rules,
	})
	return r
}

// AddControls adds multiple control rules from a struct slice
func (r *Radio) AddControls(controlRules []ControlRule) *Radio {
	r.control = append(r.control, controlRules...)
	return r
}

// AppendRule adds a suffix element
func (r *Radio) AppendRule(key string, element SuffixElement) *Radio {
	if key == "suffix" {
		r.suffix = &element
	}
	return r
}

// Size sets the radio size
func (r *Radio) Size(size string) *Radio {
	r.SetProp("size", size)
	return r
}

// Required makes the radio required
func (r *Radio) Required() *Radio {
	r.AddValidateRule(NewStringRequiredRule(fmt.Sprintf("请选择%s", r.title)))
	return r
}

// Placeholder sets placeholder (not applicable for radio, but required by interface)
func (r *Radio) Placeholder(text string) *Radio {
	return r
}

// Disabled sets the disabled state
func (r *Radio) Disabled(disabled bool) *Radio {
	r.SetProp("disabled", disabled)
	return r
}

// Build returns the radio component as a map
func (r *Radio) Build() map[string]interface{} {
	result := r.BaseComponent.Build()
	result["type"] = "el-radio-group"

	// Add options to props
	if len(r.options) > 0 {
		result["options"] = r.options
	}

	// Add control rules if present
	if len(r.control) > 0 {
		controlData := make([]map[string]interface{}, len(r.control))
		for i, ctrl := range r.control {
			// Build each component in the rule
			ruleData := make([]map[string]interface{}, len(ctrl.Rule))
			for j, component := range ctrl.Rule {
				ruleData[j] = component.Build()
			}

			controlData[i] = map[string]interface{}{
				"value": ctrl.Value,
				"rule":  ruleData,
			}
		}
		result["control"] = controlData
	}

	// Add suffix if present
	if r.suffix != nil {
		result["suffix"] = map[string]interface{}{
			"type": r.suffix.Type,
		}
		if r.suffix.Style != nil {
			result["suffix"].(map[string]interface{})["style"] = r.suffix.Style
		}
		if r.suffix.DomProps != nil {
			result["suffix"].(map[string]interface{})["domProps"] = r.suffix.DomProps
		}
	}

	return result
}

// Checkbox component implementation
type Checkbox struct {
	*BaseComponent
	options []contracts.Option
}

// NewCheckbox creates a new checkbox component
func NewCheckbox(field, title string) *Checkbox {
	return &Checkbox{
		BaseComponent: NewBaseComponent(field, title),
		options:       make([]contracts.Option, 0),
	}
}

// Options sets multiple options at once
func (c *Checkbox) Options(options []contracts.Option) *Checkbox {
	c.options = options
	return c
}

// AddOption adds a single option
func (c *Checkbox) AddOption(value interface{}, label string) *Checkbox {
	c.options = append(c.options, contracts.Option{
		Value: value,
		Label: label,
	})
	return c
}

// AddOptions adds multiple options from a struct slice
func (c *Checkbox) AddOptions(options []contracts.Option) *Checkbox {
	c.options = append(c.options, options...)
	return c
}

// Min sets minimum number of checked items
func (c *Checkbox) Min(min int) *Checkbox {
	c.SetProp("min", min)
	return c
}

// Max sets maximum number of checked items
func (c *Checkbox) Max(max int) *Checkbox {
	c.SetProp("max", max)
	return c
}

// Size sets the checkbox size
func (c *Checkbox) Size(size string) *Checkbox {
	c.SetProp("size", size)
	return c
}

// Required makes the checkbox required
func (c *Checkbox) Required() *Checkbox {
	c.AddValidateRule(NewArrayRequiredRule(fmt.Sprintf("请选择%s", c.title)))
	return c
}

// Placeholder sets placeholder (not applicable for checkbox, but required by interface)
func (c *Checkbox) Placeholder(text string) *Checkbox {
	return c
}

// Disabled sets the disabled state
func (c *Checkbox) Disabled(disabled bool) *Checkbox {
	c.SetProp("disabled", disabled)
	return c
}

// Build returns the checkbox component as a map
func (c *Checkbox) Build() map[string]interface{} {
	result := c.BaseComponent.Build()
	result["type"] = "el-checkbox-group"

	// Add options to props
	if len(c.options) > 0 {
		result["options"] = c.options
	}

	// Override value - checkbox always uses array
	if c.value == nil {
		result["value"] = []interface{}{} // Empty array for checkbox
	}

	return result
}

// InputNumber component implementation
type InputNumber struct {
	*BaseComponent
}

// NewInputNumber creates a new input number component
func NewInputNumber(field, title string) *InputNumber {
	inputNumber := &InputNumber{
		BaseComponent: NewBaseComponent(field, title),
	}
	// Set default placeholder
	inputNumber.SetProp("placeholder", fmt.Sprintf("请输入%s", title))
	return inputNumber
}

// Min sets the minimum value
func (i *InputNumber) Min(min float64) *InputNumber {
	i.SetProp("min", min)
	return i
}

// Max sets the maximum value
func (i *InputNumber) Max(max float64) *InputNumber {
	i.SetProp("max", max)
	return i
}

// Step sets the step value
func (i *InputNumber) Step(step float64) *InputNumber {
	i.SetProp("step", step)
	return i
}

// StepStrictly enables step strictly mode
func (i *InputNumber) StepStrictly(strictly bool) *InputNumber {
	i.SetProp("step-strictly", strictly)
	return i
}

// Precision sets the precision
func (i *InputNumber) Precision(precision int) *InputNumber {
	i.SetProp("precision", precision)
	return i
}

// Size sets the input number size
func (i *InputNumber) Size(size string) *InputNumber {
	i.SetProp("size", size)
	return i
}

// Controls enables/disables controls
func (i *InputNumber) Controls(controls bool) *InputNumber {
	i.SetProp("controls", controls)
	return i
}

// ControlsPosition sets the controls position
func (i *InputNumber) ControlsPosition(position string) *InputNumber {
	i.SetProp("controls-position", position)
	return i
}

// Required makes the input number required
func (i *InputNumber) Required() *InputNumber {
	i.AddValidateRule(NewNumberRequiredRule(fmt.Sprintf("请输入%s", i.title)))
	return i
}

// Placeholder sets placeholder
func (i *InputNumber) Placeholder(text string) *InputNumber {
	i.SetProp("placeholder", text)
	return i
}

// Disabled sets the disabled state
func (i *InputNumber) Disabled(disabled bool) *InputNumber {
	i.SetProp("disabled", disabled)
	return i
}

// Build returns the input number component as a map
func (i *InputNumber) Build() map[string]interface{} {
	result := i.BaseComponent.Build()
	result["type"] = "el-input-number"
	return result
}

// Slider component implementation
type Slider struct {
	*BaseComponent
}

// NewSlider creates a new slider component
func NewSlider(field, title string) *Slider {
	return &Slider{
		BaseComponent: NewBaseComponent(field, title),
	}
}

// Min sets the minimum value
func (s *Slider) Min(min float64) *Slider {
	s.SetProp("min", min)
	return s
}

// Max sets the maximum value
func (s *Slider) Max(max float64) *Slider {
	s.SetProp("max", max)
	return s
}

// Step sets the step value
func (s *Slider) Step(step float64) *Slider {
	s.SetProp("step", step)
	return s
}

// ShowInput shows input box
func (s *Slider) ShowInput(show bool) *Slider {
	s.SetProp("show-input", show)
	return s
}

// ShowInputControls shows input controls
func (s *Slider) ShowInputControls(show bool) *Slider {
	s.SetProp("show-input-controls", show)
	return s
}

// ShowStops shows step stops
func (s *Slider) ShowStops(show bool) *Slider {
	s.SetProp("show-stops", show)
	return s
}

// ShowTooltip shows tooltip
func (s *Slider) ShowTooltip(show bool) *Slider {
	s.SetProp("show-tooltip", show)
	return s
}

// Range enables range mode
func (s *Slider) Range(range_ bool) *Slider {
	s.SetProp("range", range_)
	return s
}

// Vertical enables vertical mode
func (s *Slider) Vertical(vertical bool) *Slider {
	s.SetProp("vertical", vertical)
	return s
}

// Height sets height for vertical slider
func (s *Slider) Height(height string) *Slider {
	s.SetProp("height", height)
	return s
}

// Required makes the slider required
func (s *Slider) Required() *Slider {
	// Create conditional required rule based on range mode
	rule := NewConditionalRequiredRule(
		fmt.Sprintf("请设置%s", s.title),
		s,
		func(component interface{}) string {
			slider := component.(*Slider)
			if range_, exists := slider.props["range"]; exists {
				if val, ok := range_.(bool); ok && val {
					return "array" // Range mode uses array validation
				}
			}
			return "number" // Single value mode uses number validation
		},
	)
	s.AddValidateRule(rule)
	return s
}

// Placeholder sets placeholder (not applicable for slider, but required by interface)
func (s *Slider) Placeholder(text string) *Slider {
	return s
}

// Disabled sets the disabled state
func (s *Slider) Disabled(disabled bool) *Slider {
	s.SetProp("disabled", disabled)
	return s
}

// Build returns the slider component as a map
func (s *Slider) Build() map[string]interface{} {
	result := s.BaseComponent.Build()
	result["type"] = "el-slider"

	// Override value based on range mode
	if s.value == nil {
		if range_, exists := s.props["range"]; exists {
			if val, ok := range_.(bool); ok && val {
				result["value"] = []interface{}{} // Empty array for range mode
			} else {
				result["value"] = 0 // Default number for single value mode
			}
		} else {
			result["value"] = 0 // Default number for single value mode
		}
	}

	return result
}

// Rate component implementation
type Rate struct {
	*BaseComponent
}

// NewRate creates a new rate component
func NewRate(field, title string) *Rate {
	return &Rate{
		BaseComponent: NewBaseComponent(field, title),
	}
}

// Max sets the maximum rating
func (r *Rate) Max(max int) *Rate {
	r.SetProp("max", max)
	return r
}

// AllowHalf enables half star
func (r *Rate) AllowHalf(allow bool) *Rate {
	r.SetProp("allow-half", allow)
	return r
}

// LowThreshold sets the low threshold
func (r *Rate) LowThreshold(threshold int) *Rate {
	r.SetProp("low-threshold", threshold)
	return r
}

// HighThreshold sets the high threshold
func (r *Rate) HighThreshold(threshold int) *Rate {
	r.SetProp("high-threshold", threshold)
	return r
}

// Colors sets the colors for different thresholds
func (r *Rate) Colors(colors []string) *Rate {
	r.SetProp("colors", colors)
	return r
}

// VoidColor sets the void color
func (r *Rate) VoidColor(color string) *Rate {
	r.SetProp("void-color", color)
	return r
}

// DisabledVoidColor sets the disabled void color
func (r *Rate) DisabledVoidColor(color string) *Rate {
	r.SetProp("disabled-void-color", color)
	return r
}

// IconClasses sets the icon classes
func (r *Rate) IconClasses(classes []string) *Rate {
	r.SetProp("icon-classes", classes)
	return r
}

// VoidIconClass sets the void icon class
func (r *Rate) VoidIconClass(class string) *Rate {
	r.SetProp("void-icon-class", class)
	return r
}

// DisabledVoidIconClass sets the disabled void icon class
func (r *Rate) DisabledVoidIconClass(class string) *Rate {
	r.SetProp("disabled-void-icon-class", class)
	return r
}

// ShowText shows text
func (r *Rate) ShowText(show bool) *Rate {
	r.SetProp("show-text", show)
	return r
}

// ShowScore shows score
func (r *Rate) ShowScore(show bool) *Rate {
	r.SetProp("show-score", show)
	return r
}

// TextColor sets the text color
func (r *Rate) TextColor(color string) *Rate {
	r.SetProp("text-color", color)
	return r
}

// Texts sets the text array
func (r *Rate) Texts(texts []string) *Rate {
	r.SetProp("texts", texts)
	return r
}

// ScoreTemplate sets the score template
func (r *Rate) ScoreTemplate(template string) *Rate {
	r.SetProp("score-template", template)
	return r
}

// Required makes the rate required
func (r *Rate) Required() *Rate {
	r.AddValidateRule(NewNumberRequiredRule(fmt.Sprintf("请选择%s", r.title)))
	return r
}

// Placeholder sets placeholder (not applicable for rate, but required by interface)
func (r *Rate) Placeholder(text string) *Rate {
	return r
}

// Disabled sets the disabled state
func (r *Rate) Disabled(disabled bool) *Rate {
	r.SetProp("disabled", disabled)
	return r
}

// Build returns the rate component as a map
func (r *Rate) Build() map[string]interface{} {
	result := r.BaseComponent.Build()
	result["type"] = "el-rate"
	return result
}

// SelectRequiredRule implements a required validation rule for select components
// with appropriate type based on multiple selection mode
type SelectRequiredRule struct {
	message    string
	isMultiple bool
}

// Type returns the appropriate validation type based on multiple selection mode
func (r *SelectRequiredRule) Type() string {
	if r.isMultiple {
		return "array"
	}
	return "string"
}

// Message returns the validation message
func (r *SelectRequiredRule) Message() string {
	return r.message
}

// Validate validates that the value is not empty
func (r *SelectRequiredRule) Validate(value interface{}) error {
	if value == nil {
		return fmt.Errorf(r.message)
	}

	if r.isMultiple {
		// For multiple selection, value should be an array
		switch v := value.(type) {
		case []interface{}:
			if len(v) == 0 {
				return fmt.Errorf(r.message)
			}
		case []string:
			if len(v) == 0 {
				return fmt.Errorf(r.message)
			}
		case []int:
			if len(v) == 0 {
				return fmt.Errorf(r.message)
			}
		default:
			return fmt.Errorf(r.message)
		}
	} else {
		// For single selection, value should be a non-empty string or number
		switch v := value.(type) {
		case string:
			if strings.TrimSpace(v) == "" {
				return fmt.Errorf(r.message)
			}
		case int, float64:
			// Numbers are always valid if not nil
		default:
			if v == "" || v == nil {
				return fmt.Errorf(r.message)
			}
		}
	}

	return nil
}

// ToMap returns the rule as a map for JSON serialization
func (r *SelectRequiredRule) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"required": true,
		"message":  r.message,
		"type":     r.Type(),
		"trigger":  "change",
	}
}
