package formbuilder

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/FlameMida/form-builder-go/contracts"
	"github.com/FlameMida/form-builder-go/errors"
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
	
	// Performance optimization fields
	mu               sync.RWMutex // Protects cache fields
	cachedFormRule   []byte       // Cached JSON for FormRule
	cachedFormConfig []byte       // Cached JSON for FormConfig
	cacheValid       bool         // Whether cache is valid
}

// NewForm creates a new form with the specified UI bootstrap (optimized for memory efficiency)
func NewForm(ui contracts.BootstrapInterface, action string, rules []interface{}, config map[string]interface{}) (*FormBuilder, error) {
	// Pre-allocate components slice with exact capacity to avoid reallocations
	components := make([]contracts.Component, 0, len(rules))
	for i, rule := range rules {
		if comp, ok := rule.(contracts.Component); ok {
			components = append(components, comp)
		} else {
			return nil, fmt.Errorf("rule at index %d does not implement Component interface", i)
		}
	}

	// Pre-allocate maps with estimated capacity
	form := &FormBuilder{
		action:          action,
		method:          "POST",
		title:           "", // 默认为空
		rules:           components,
		formData:        make(map[string]interface{}, 16), // Pre-allocate reasonable capacity
		config:          config,
		ui:              ui,
		headers:         make(map[string]string, 8),      // Pre-allocate reasonable capacity
		formContentType: "application/x-www-form-urlencoded",
		dependScript: []string{
			`<script src="https://unpkg.com/jquery@3.3.1/dist/jquery.min.js"></script>`,
			`<script src="https://unpkg.com/vue@2.5.13/dist/vue.min.js"></script>`,
			`<script src="https://unpkg.com/@form-create/data@1.0.0/dist/province_city.js"></script>`,
			`<script src="https://unpkg.com/@form-create/data@1.0.0/dist/province_city_area.js"></script>`,
		},
		cacheValid: false, // Initialize cache as invalid
	}

	if err := ui.Init(form); err != nil {
		return nil, err
	}

	if err := form.checkFieldUnique(); err != nil {
		return nil, err
	}

	return form, nil
}

// invalidateCache marks the cache as invalid (should be called when form structure changes)
func (f *FormBuilder) invalidateCache() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.invalidateCacheUnsafe()
}

// invalidateCacheUnsafe is the non-locking version of invalidateCache
func (f *FormBuilder) invalidateCacheUnsafe() {
	f.cacheValid = false
	f.cachedFormRule = nil
	f.cachedFormConfig = nil
}

// SetRule sets the form rules (thread-safe)
func (f *FormBuilder) SetRule(rules []contracts.Component) (contracts.Form, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	
	f.rules = rules
	if err := f.checkFieldUniqueUnsafe(); err != nil {
		return nil, err
	}
	f.invalidateCacheUnsafe()
	return f, nil
}

// Append adds a component to the end of the form (thread-safe)
func (f *FormBuilder) Append(component contracts.Component) (contracts.Form, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	
	f.rules = append(f.rules, component)
	if err := f.checkFieldUniqueUnsafe(); err != nil {
		// Remove the component that was just added to maintain consistency
		f.rules = f.rules[:len(f.rules)-1]
		return nil, err
	}
	f.invalidateCacheUnsafe()
	return f, nil
}

// Prepend adds a component to the beginning of the form (thread-safe)
func (f *FormBuilder) Prepend(component contracts.Component) (contracts.Form, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	
	f.rules = append([]contracts.Component{component}, f.rules...)
	if err := f.checkFieldUniqueUnsafe(); err != nil {
		// Remove the component that was just prepended to maintain consistency
		f.rules = f.rules[1:]
		return nil, err
	}
	f.invalidateCacheUnsafe()
	return f, nil
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

// FormData sets multiple form field values (thread-safe)
func (f *FormBuilder) FormData(data map[string]interface{}) contracts.Form {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.formData = data
	f.invalidateCacheUnsafe() // Form data changes invalidate rule cache
	return f
}

// SetValue sets a single form field value (thread-safe)
func (f *FormBuilder) SetValue(field string, value interface{}) contracts.Form {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.formData[field] = value
	f.invalidateCacheUnsafe() // Form data changes invalidate rule cache
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

// FormRule returns the form rules as a slice of maps (with memory optimization and concurrency safety)
func (f *FormBuilder) FormRule() []map[string]interface{} {
	f.mu.RLock()
	defer f.mu.RUnlock()
	
	// Pre-allocate slice with known capacity to avoid reallocations
	rules := make([]map[string]interface{}, 0, len(f.rules))
	for _, rule := range f.rules {
		rules = append(rules, f.parseFormComponent(rule))
	}
	return f.deepSetFormData(f.formData, rules)
}

// FormConfig returns the form configuration (thread-safe)
func (f *FormBuilder) FormConfig() map[string]interface{} {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.config
}

// ParseFormRule returns the JSON representation of form rules with caching
func (f *FormBuilder) ParseFormRule() ([]byte, error) {
	f.mu.RLock()
	if f.cacheValid && f.cachedFormRule != nil {
		defer f.mu.RUnlock()
		return f.cachedFormRule, nil
	}
	f.mu.RUnlock()
	
	// Generate fresh JSON
	data, err := json.Marshal(f.FormRule())
	if err != nil {
		return nil, err
	}
	
	// Cache the result
	f.mu.Lock()
	f.cachedFormRule = data
	f.cacheValid = true
	f.mu.Unlock()
	
	return data, nil
}

// ParseFormConfig returns the JSON representation of form config with caching
func (f *FormBuilder) ParseFormConfig() ([]byte, error) {
	f.mu.RLock()
	if f.cacheValid && f.cachedFormConfig != nil {
		defer f.mu.RUnlock()
		return f.cachedFormConfig, nil
	}
	f.mu.RUnlock()
	
	// Generate fresh JSON
	data, err := json.Marshal(f.FormConfig())
	if err != nil {
		return nil, err
	}
	
	// Cache the result
	f.mu.Lock()
	f.cachedFormConfig = data
	f.cacheValid = true
	f.mu.Unlock()
	
	return data, nil
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

// checkFieldUnique ensures all field names are unique using O(n) algorithm
func (f *FormBuilder) checkFieldUnique() error {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.checkFieldUniqueUnsafe()
}

// checkFieldUniqueUnsafe is the non-locking version of checkFieldUnique
func (f *FormBuilder) checkFieldUniqueUnsafe() error {
	if len(f.rules) == 0 {
		return nil
	}
	
	// Pre-allocate map with estimated capacity to reduce allocations
	fields := make(map[string]bool, len(f.rules))
	
	for _, rule := range f.rules {
		field := rule.Field()
		if field == "" {
			continue
		}
		if fields[field] {
			return errors.NewDuplicateFieldError(field)
		}
		fields[field] = true
	}
	return nil
}

// parseFormComponent parses a component into a map
func (f *FormBuilder) parseFormComponent(component contracts.Component) map[string]interface{} {
	return f.ui.ParseComponent(component)
}

// deepSetFormData recursively sets form data values (optimized for memory efficiency)
func (f *FormBuilder) deepSetFormData(formData map[string]interface{}, rules []map[string]interface{}) []map[string]interface{} {
	if len(formData) == 0 {
		return rules
	}

	// Optimize: modify rules in place instead of creating new maps
	for i := range rules {
		if field, ok := rules[i]["field"].(string); ok {
			if value, exists := formData[field]; exists {
				rules[i]["value"] = value
			}
		}
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
