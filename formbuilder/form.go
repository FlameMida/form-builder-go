package formbuilder

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/FlameMida/form-builder-go/contracts"
)

// FormBuilder represents the main form builder
type FormBuilder struct {
	action          string
	method          string
	title           string
	rules           []contracts.Component
	formData        map[string]interface{}
	config          map[string]interface{}
	ui              contracts.BootstrapInterface
	headers         map[string]string
	formContentType string
	dependScript    []string
	template        string
}

// NewForm creates a new form with the specified UI bootstrap
func NewForm(ui contracts.BootstrapInterface, action string, rules []interface{}, config map[string]interface{}) (*FormBuilder, error) {
	// Convert []interface{} to []contracts.Component
	components := make([]contracts.Component, len(rules))
	for i, rule := range rules {
		if comp, ok := rule.(contracts.Component); ok {
			components[i] = comp
		} else {
			return nil, fmt.Errorf("rule at index %d does not implement Component interface", i)
		}
	}

	form := &FormBuilder{
		action:          action,
		method:          "POST",
		title:           "", // 默认为空
		rules:           components,
		formData:        make(map[string]interface{}),
		config:          config,
		ui:              ui,
		headers:         make(map[string]string),
		formContentType: "application/x-www-form-urlencoded",
		dependScript: []string{
			`<script src="https://unpkg.com/jquery@3.3.1/dist/jquery.min.js"></script>`,
			`<script src="https://unpkg.com/vue@2.5.13/dist/vue.min.js"></script>`,
			`<script src="https://unpkg.com/@form-create/data@1.0.0/dist/province_city.js"></script>`,
			`<script src="https://unpkg.com/@form-create/data@1.0.0/dist/province_city_area.js"></script>`,
		},
	}

	if err := ui.Init(form); err != nil {
		return nil, err
	}

	if err := form.checkFieldUnique(); err != nil {
		return nil, err
	}

	return form, nil
}

// SetRule sets the form rules
func (f *FormBuilder) SetRule(rules []contracts.Component) contracts.Form {
	f.rules = rules
	if err := f.checkFieldUnique(); err != nil {
		panic(err) // In production, handle this more gracefully
	}
	return f
}

// Append adds a component to the end of the form
func (f *FormBuilder) Append(component contracts.Component) contracts.Form {
	f.rules = append(f.rules, component)
	if err := f.checkFieldUnique(); err != nil {
		panic(err) // In production, handle this more gracefully
	}
	return f
}

// Prepend adds a component to the beginning of the form
func (f *FormBuilder) Prepend(component contracts.Component) contracts.Form {
	f.rules = append([]contracts.Component{component}, f.rules...)
	if err := f.checkFieldUnique(); err != nil {
		panic(err) // In production, handle this more gracefully
	}
	return f
}

// SetAction sets the form action URL
func (f *FormBuilder) SetAction(action string) contracts.Form {
	f.action = action
	return f
}

// GetAction returns the form action URL
func (f *FormBuilder) GetAction() string {
	return f.action
}

// SetMethod sets the HTTP method
func (f *FormBuilder) SetMethod(method string) contracts.Form {
	f.method = method
	return f
}

// GetMethod returns the HTTP method
func (f *FormBuilder) GetMethod() string {
	return f.method
}

// SetTitle sets the form title
func (f *FormBuilder) SetTitle(title string) contracts.Form {
	f.title = title
	return f
}

// GetTitle returns the form title
func (f *FormBuilder) GetTitle() string {
	return f.title
}

// FormData sets multiple form field values
func (f *FormBuilder) FormData(data map[string]interface{}) contracts.Form {
	f.formData = data
	return f
}

// SetValue sets a single form field value
func (f *FormBuilder) SetValue(field string, value interface{}) contracts.Form {
	f.formData[field] = value
	return f
}

// SetHeader sets an HTTP header
func (f *FormBuilder) SetHeader(name, value string) *FormBuilder {
	f.headers[name] = value
	return f
}

// SetHeaders sets multiple HTTP headers
func (f *FormBuilder) SetHeaders(headers map[string]string) *FormBuilder {
	f.headers = headers
	return f
}

// SetFormContentType sets the form content type
func (f *FormBuilder) SetFormContentType(contentType string) *FormBuilder {
	f.formContentType = contentType
	return f
}

// SetDependScript sets the dependency scripts
func (f *FormBuilder) SetDependScript(scripts []string) *FormBuilder {
	f.dependScript = scripts
	return f
}

// FormRule returns the form rules as a slice of maps
func (f *FormBuilder) FormRule() []map[string]interface{} {
	rules := make([]map[string]interface{}, len(f.rules))
	for i, rule := range f.rules {
		rules[i] = f.parseFormComponent(rule)
	}
	return f.deepSetFormData(f.formData, rules)
}

// FormConfig returns the form configuration
func (f *FormBuilder) FormConfig() map[string]interface{} {
	return f.config
}

// ParseFormRule returns the JSON representation of form rules
func (f *FormBuilder) ParseFormRule() ([]byte, error) {
	return json.Marshal(f.FormRule())
}

// ParseFormConfig returns the JSON representation of form config
func (f *FormBuilder) ParseFormConfig() ([]byte, error) {
	return json.Marshal(f.FormConfig())
}

// ParseHeaders returns the JSON representation of headers
func (f *FormBuilder) ParseHeaders() ([]byte, error) {
	return json.Marshal(f.headers)
}

// ParseDependScript returns the dependency scripts as a string
func (f *FormBuilder) ParseDependScript() string {
	return strings.Join(f.dependScript, "\r\n")
}

// View renders the form to HTML
func (f *FormBuilder) View() (string, error) {
	// This would typically use Go templates
	// For now, return a basic implementation
	return f.renderTemplate()
}

// checkFieldUnique ensures all field names are unique
func (f *FormBuilder) checkFieldUnique() error {
	return f.checkFieldUniqueRecursive(f.rules, make(map[string]bool))
}

// checkFieldUniqueRecursive recursively checks field uniqueness
func (f *FormBuilder) checkFieldUniqueRecursive(rules []contracts.Component, fields map[string]bool) error {
	for _, rule := range rules {
		field := rule.Field()
		if field == "" {
			continue
		}
		if fields[field] {
			return errors.New("组件的 field 不能重复")
		}
		fields[field] = true
	}
	return nil
}

// parseFormComponent parses a component into a map
func (f *FormBuilder) parseFormComponent(component contracts.Component) map[string]interface{} {
	return f.ui.ParseComponent(component)
}

// deepSetFormData recursively sets form data values
func (f *FormBuilder) deepSetFormData(formData map[string]interface{}, rules []map[string]interface{}) []map[string]interface{} {
	if len(formData) == 0 {
		return rules
	}

	for i, rule := range rules {
		if field, ok := rule["field"].(string); ok {
			if value, exists := formData[field]; exists {
				rule["value"] = value
			}
		}
		rules[i] = rule
	}

	return rules
}

// renderTemplate renders the form using templates
func (f *FormBuilder) renderTemplate() (string, error) {
	// Basic template implementation
	// In a real implementation, you'd use html/template
	formRuleJSON, _ := f.ParseFormRule()
	formConfigJSON, _ := f.ParseFormConfig()
	headersJSON, _ := f.ParseHeaders()

	template := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>%s</title>
    %s
</head>
<body>
    <div id="app">
        <form-create 
            v-model="fApi" 
            :rule="rule" 
            :option="option"
            :headers="headers"
            :action="%s"
            :method="%s"
        />
    </div>
    
    <script>
        new Vue({
            el: '#app',
            data: {
                fApi: {},
                rule: %s,
                option: %s,
                headers: %s
            }
        });
    </script>
</body>
</html>`,
		f.title,
		f.ParseDependScript(),
		f.action,
		f.method,
		string(formRuleJSON),
		string(formConfigJSON),
		string(headersJSON),
	)

	return template, nil
}
