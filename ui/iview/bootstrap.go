// Package iview provides IView specific implementations
package iview

import (
	"github.com/FlameMida/form-builder-go/contracts"
)

// Bootstrap implements IView bootstrap functionality
type Bootstrap struct {
	version      int
	dependScript []string
}

// NewBootstrap creates a new IView bootstrap
func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		version: 3,
		dependScript: []string{
			`<script src="https://unpkg.com/iview@3.5.4/dist/iview.min.js"></script>`,
			`<link rel="stylesheet" href="https://unpkg.com/iview@3.5.4/dist/styles/iview.css">`,
		},
	}
}

// NewBootstrapV4 creates a new IView v4 bootstrap
func NewBootstrapV4() *Bootstrap {
	return &Bootstrap{
		version: 4,
		dependScript: []string{
			`<script src="https://unpkg.com/view-design@4.7.0/dist/iview.min.js"></script>`,
			`<link rel="stylesheet" href="https://unpkg.com/view-design@4.7.0/dist/styles/iview.css">`,
		},
	}
}

// Init initializes the bootstrap with a form
func (b *Bootstrap) Init(form contracts.Form) error {
	// IView specific initialization
	return nil
}

// GetDependScript returns the dependency scripts for IView
func (b *Bootstrap) GetDependScript() []string {
	return b.dependScript
}

// ParseComponent parses a component for IView
func (b *Bootstrap) ParseComponent(component contracts.Component) map[string]interface{} {
	// IView might need different component parsing logic
	result := component.Build()

	// Convert ElementUI component types to IView equivalents
	if componentType, ok := result["type"].(string); ok {
		switch componentType {
		case "el-input":
			result["type"] = "input"
		case "el-switch":
			result["type"] = "i-switch"
		case "el-select":
			result["type"] = "i-select"
		case "el-radio-group":
			result["type"] = "radio"
		case "el-checkbox-group":
			result["type"] = "checkbox"
		case "el-input-number":
			result["type"] = "input-number"
		case "el-date-picker":
			result["type"] = "date-picker"
		case "el-time-picker":
			result["type"] = "time-picker"
		case "el-slider":
			result["type"] = "slider"
		case "el-rate":
			result["type"] = "rate"
		case "el-color-picker":
			result["type"] = "color-picker"
		case "el-upload":
			result["type"] = "upload"
		case "el-cascader":
			result["type"] = "cascader"
		case "el-tree":
			result["type"] = "tree"
		case "el-button":
			result["type"] = "i-button"
		}
	}

	return result
}

// Config implements IView configuration
type Config struct {
	config map[string]interface{}
}

// NewConfig creates a new IView configuration
func NewConfig(config map[string]interface{}) *Config {
	if config == nil {
		config = make(map[string]interface{})
	}

	// Set IView defaults
	if _, exists := config["submitBtn"]; !exists {
		config["submitBtn"] = true
	}
	if _, exists := config["resetBtn"]; !exists {
		config["resetBtn"] = false
	}
	if _, exists := config["form"]; !exists {
		config["form"] = map[string]interface{}{
			"labelPosition": "right",
			"labelWidth":    125,
			"size":          "default",
		}
	}

	return &Config{config: config}
}

// GetConfig returns the configuration map
func (c *Config) GetConfig() map[string]interface{} {
	return c.config
}

// SetSubmitBtn sets the submit button visibility
func (c *Config) SetSubmitBtn(show bool) contracts.ConfigInterface {
	c.config["submitBtn"] = show
	return c
}

// SetResetBtn sets the reset button visibility
func (c *Config) SetResetBtn(show bool) contracts.ConfigInterface {
	c.config["resetBtn"] = show
	return c
}

// ComponentGlobalConfig sets global configuration for a component
func (c *Config) ComponentGlobalConfig(componentName string, config map[string]interface{}) contracts.ConfigInterface {
	if _, exists := c.config["global"]; !exists {
		c.config["global"] = make(map[string]interface{})
	}

	global := c.config["global"].(map[string]interface{})
	global[componentName] = config

	return c
}
