// Package components provides form component implementations
package components

import (
	"errors"
	"fmt"

	"github.com/FlameMida/form-builder-go/contracts"
)

// BaseComponent provides common functionality for all components
type BaseComponent struct {
	field         string
	title         string
	value         interface{}
	props         map[string]interface{}
	validateRules []contracts.ValidateRule
}

// NewBaseComponent creates a new base component
func NewBaseComponent(field, title string) *BaseComponent {
	return &BaseComponent{
		field:         field,
		title:         title,
		props:         make(map[string]interface{}),
		validateRules: make([]contracts.ValidateRule, 0),
	}
}

// Field returns the component field name
func (b *BaseComponent) Field() string {
	return b.field
}

// Title returns the component title
func (b *BaseComponent) Title() string {
	return b.title
}

// SetValue sets the component value
func (b *BaseComponent) SetValue(value interface{}) contracts.Component {
	b.value = value
	return b
}

// GetValue returns the component value
func (b *BaseComponent) GetValue() interface{} {
	return b.value
}

// Build returns the component as a map for JSON serialization
func (b *BaseComponent) Build() map[string]interface{} {
	result := map[string]interface{}{
		"field": b.field,
		"title": b.title,
		"props": b.props,
	}

	// 始终包含value字段，如果没有设置则为空字符串（匹配PHP格式）
	if b.value != nil {
		result["value"] = b.value
	} else {
		result["value"] = ""
	}

	if len(b.validateRules) > 0 {
		rules := make([]map[string]interface{}, len(b.validateRules))
		for i, rule := range b.validateRules {
			rules[i] = rule.ToMap()
		}
		result["validate"] = rules
	}

	return result
}

// Validate validates the component value
func (b *BaseComponent) Validate() error {
	for _, rule := range b.validateRules {
		if err := rule.Validate(b.value); err != nil {
			return fmt.Errorf("validation failed for field %s: %w", b.field, err)
		}
	}
	return nil
}

// AddValidateRule adds a validation rule
func (b *BaseComponent) AddValidateRule(rule contracts.ValidateRule) contracts.ValidateComponent {
	b.validateRules = append(b.validateRules, rule)
	return b
}

// GetValidateRules returns all validation rules
func (b *BaseComponent) GetValidateRules() []contracts.ValidateRule {
	return b.validateRules
}

// SetProp sets a component property
func (b *BaseComponent) SetProp(key string, value interface{}) *BaseComponent {
	b.props[key] = value
	return b
}

// GetProp gets a component property
func (b *BaseComponent) GetProp(key string) interface{} {
	return b.props[key]
}

// Input component implementation
type Input struct {
	*BaseComponent
	componentType string
}

// NewInput creates a new input component
func NewInput(field, title string) *Input {
	input := &Input{
		BaseComponent: NewBaseComponent(field, title),
		componentType: "el-input",
	}
	// Set default placeholder
	input.SetProp("placeholder", fmt.Sprintf("请输入%s", title))
	return input
}

// Type sets the input type
func (i *Input) Type(inputType string) *Input {
	i.SetProp("type", inputType)
	return i
}

// Placeholder sets the placeholder text
func (i *Input) Placeholder(text string) *Input {
	i.SetProp("placeholder", text)
	return i
}

// Required makes the input required
func (i *Input) Required() *Input {
	i.AddValidateRule(NewStringRequiredRule(fmt.Sprintf("请输入%s", i.title)))
	return i
}

// Disabled sets the disabled state
func (i *Input) Disabled(disabled bool) *Input {
	i.SetProp("disabled", disabled)
	return i
}

// Clearable makes the input clearable
func (i *Input) Clearable(clearable bool) *Input {
	i.SetProp("clearable", clearable)
	return i
}

// ShowPassword enables password visibility toggle
func (i *Input) ShowPassword(show bool) *Input {
	i.SetProp("show-password", show)
	return i
}

// Size sets the input size
func (i *Input) Size(size string) *Input {
	i.SetProp("size", size)
	return i
}

// Maxlength sets the maximum length
func (i *Input) Maxlength(max int) *Input {
	i.SetProp("maxlength", max)
	return i
}

// Minlength sets the minimum length
func (i *Input) Minlength(min int) *Input {
	i.SetProp("minlength", min)
	return i
}

// ShowWordLimit shows word count
func (i *Input) ShowWordLimit(show bool) *Input {
	i.SetProp("show-word-limit", show)
	return i
}

// Hidden sets the hidden state
func (i *Input) Hidden(hidden bool) *Input {
	if hidden {
		i.SetProp("hidden", true)
	}
	return i
}

// Build returns the input component as a map
func (i *Input) Build() map[string]interface{} {
	result := i.BaseComponent.Build()
	result["type"] = i.componentType

	// Handle hidden property at root level (PHP style)
	if hidden := i.GetProp("hidden"); hidden == true {
		result["hidden"] = true
		// Remove from props
		delete(i.props, "hidden")
	}

	return result
}

// Textarea component implementation
type Textarea struct {
	*Input
}

// NewTextarea creates a new textarea component
func NewTextarea(field, title string) *Textarea {
	input := NewInput(field, title)
	input.componentType = "el-input"
	input.SetProp("type", "textarea")
	// Override placeholder for textarea
	input.SetProp("placeholder", fmt.Sprintf("请输入%s", title))
	return &Textarea{Input: input}
}

// Rows sets the number of rows
func (t *Textarea) Rows(rows int) *Textarea {
	t.SetProp("rows", rows)
	return t
}

// AutoSize enables auto-sizing
func (t *Textarea) AutoSize(minRows, maxRows int) *Textarea {
	autosize := map[string]interface{}{
		"minRows": minRows,
		"maxRows": maxRows,
	}
	t.SetProp("autosize", autosize)
	return t
}

// Switch component implementation
type Switch struct {
	*BaseComponent
}

// NewSwitch creates a new switch component
func NewSwitch(field, title string) *Switch {
	return &Switch{
		BaseComponent: NewBaseComponent(field, title),
	}
}

// ActiveText sets the text for active state
func (s *Switch) ActiveText(text string) *Switch {
	s.SetProp("active-text", text)
	return s
}

// InactiveText sets the text for inactive state
func (s *Switch) InactiveText(text string) *Switch {
	s.SetProp("inactive-text", text)
	return s
}

// ActiveValue sets the value for active state
func (s *Switch) ActiveValue(value interface{}) *Switch {
	s.SetProp("active-value", value)
	return s
}

// InactiveValue sets the value for inactive state
func (s *Switch) InactiveValue(value interface{}) *Switch {
	s.SetProp("inactive-value", value)
	return s
}

// Required makes the switch required
func (s *Switch) Required() *Switch {
	s.AddValidateRule(NewStringRequiredRule(fmt.Sprintf("请设置%s", s.title)))
	return s
}

// Placeholder sets placeholder (not applicable for switch, but required by interface)
func (s *Switch) Placeholder(text string) *Switch {
	return s
}

// Disabled sets the disabled state
func (s *Switch) Disabled(disabled bool) *Switch {
	s.SetProp("disabled", disabled)
	return s
}

// Build returns the switch component as a map
func (s *Switch) Build() map[string]interface{} {
	result := s.BaseComponent.Build()
	result["type"] = "el-switch"
	return result
}

// RequiredRule implements a required validation rule
type RequiredRule struct {
	message string
}

// Type returns the rule type
func (r *RequiredRule) Type() string {
	return "required"
}

// Message returns the validation message
func (r *RequiredRule) Message() string {
	return r.message
}

// Validate validates that the value is not empty
func (r *RequiredRule) Validate(value interface{}) error {
	if value == nil || value == "" {
		return errors.New(r.message)
	}
	return nil
}

// ToMap returns the rule as a map
func (r *RequiredRule) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"required": true,
		"message":  r.message,
	}
}
