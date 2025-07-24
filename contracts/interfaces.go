// Package contracts defines the core interfaces for the form builder system
package contracts

import "encoding/json"

// Component represents a form component that can be rendered
type Component interface {
	Field() string
	Title() string
	Build() map[string]interface{}
	SetValue(value interface{}) Component
	GetValue() interface{}
	DoValidate() error
	Validate(rule ValidateRule) ValidateComponent
	Col(col interface{}) Component
	SetType(iType string) Component
	SetName(name string) Component
	Style(style interface{}) Component
}

// FormComponent represents a form input component
type FormComponent interface {
	Component
	Required() FormComponent
	Placeholder(text string) FormComponent
	Disabled(disabled bool) FormComponent
}

// OptionComponent represents a component with selectable options
type OptionComponent interface {
	Component
	Options(options []Option) OptionComponent
	AddOption(value interface{}, label string) OptionComponent
}

// ValidateComponent represents a component with validation rules
type ValidateComponent interface {
	Component
	AddValidateRule(rule ValidateRule) ValidateComponent
	GetValidateRules() []ValidateRule
}

// BootstrapInterface defines UI framework integration
type BootstrapInterface interface {
	Init(form Form) error
	GetDependScript() []string
	ParseComponent(component Component) map[string]interface{}
}

// ConfigInterface defines form configuration
type ConfigInterface interface {
	GetConfig() map[string]interface{}
	SetSubmitBtn(show bool) ConfigInterface
	SetResetBtn(show bool) ConfigInterface
	ComponentGlobalConfig(componentName string, config map[string]interface{}) ConfigInterface
}

// Form represents the main form builder
type Form interface {
	SetRule(rules []Component) (Form, error)
	Append(component Component) (Form, error)
	Prepend(component Component) (Form, error)
	SetAction(action string) Form
	GetAction() string
	SetMethod(method string) Form
	GetMethod() string
	SetTitle(title string) Form
	GetTitle() string
	FormData(data map[string]interface{}) Form
	SetValue(field string, value interface{}) Form
	FormRule() []map[string]interface{}
	FormConfig() map[string]interface{}
	ParseFormRule() ([]byte, error)
	ParseFormConfig() ([]byte, error)
	View() (string, error)
}

// StyleInterface defines component styling
type StyleInterface interface {
	GetStyle() map[string]interface{}
	SetStyle(style map[string]interface{}) StyleInterface
}

// ValidateRule represents a validation rule
type ValidateRule interface {
	Type() string
	Message() string
	Validate(value interface{}) error
	ToMap() map[string]interface{}
}

// Option represents a selectable option
type Option struct {
	Value    interface{} `json:"value"`
	Label    string      `json:"label"`
	Disabled bool        `json:"disabled,omitempty"`
}

// MarshalJSON implements json.Marshaler for Option
func (o Option) MarshalJSON() ([]byte, error) {
	type Alias Option
	return json.Marshal(Alias(o))
}
