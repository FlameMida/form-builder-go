// Package elm provides ElementUI specific implementations
package elm

import (
	"github.com/FlameMida/form-builder-go/contracts"
)

// Bootstrap implements ElementUI bootstrap functionality
type Bootstrap struct {
	version      int
	dependScript []string
}

// NewBootstrap creates a new ElementUI bootstrap
func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		version: 2,
		dependScript: []string{
			`<script src="https://unpkg.com/element-ui@2.15.6/lib/index.js"></script>`,
			`<link rel="stylesheet" href="https://unpkg.com/element-ui@2.15.6/lib/theme-chalk/index.css">`,
		},
	}
}

// Init initializes the bootstrap with a form
func (b *Bootstrap) Init(form contracts.Form) error {
	// ElementUI specific initialization
	return nil
}

// GetDependScript returns the dependency scripts for ElementUI
func (b *Bootstrap) GetDependScript() []string {
	return b.dependScript
}

// ParseComponent parses a component for ElementUI
func (b *Bootstrap) ParseComponent(component contracts.Component) map[string]interface{} {
	result := component.Build()

	if componentType, ok := result["type"].(string); ok {
		switch componentType {
		case "el-input":
			result["type"] = "input"
		case "el-input-number":
			result["type"] = "inputNumber"
		case "el-switch":
			result["type"] = "switch"
		case "el-radio-group":
			result["type"] = "radio"
		case "el-checkbox-group":
			result["type"] = "checkbox"
		case "el-select":
			result["type"] = "select"
		case "el-date-picker":
			result["type"] = "datePicker"
		case "el-time-picker":
			result["type"] = "timePicker"
		case "el-color-picker":
			result["type"] = "colorPicker"
		case "el-upload":
			result["type"] = "upload"
		case "el-rate":
			result["type"] = "rate"
		case "el-slider":
			result["type"] = "slider"
		case "el-cascader":
			result["type"] = "cascader"
		case "el-tree":
			result["type"] = "tree"
		case "el-button":
			result["type"] = "button"
		}
	}

	if validate, ok := result["validate"].([]map[string]interface{}); ok {
		for i, rule := range validate {
			if rule["type"] == "required" {
				validate[i]["type"] = "string"
				validate[i]["trigger"] = "change"
				// 修改message格式
				if title, ok := result["title"].(string); ok {
					validate[i]["message"] = "请输入" + title
				}
			}
		}
	}

	// 为input类型组件添加text类型属性
	if result["type"] == "input" {
		if props, ok := result["props"].(map[string]interface{}); ok {
			if _, hasType := props["type"]; !hasType {
				props["type"] = "text"
			}
		}
	}

	// 为数字类型组件添加数字验证
	if result["type"] == "inputNumber" {
		if validate, ok := result["validate"].([]map[string]interface{}); ok {
			for i, rule := range validate {
				if rule["type"] == "string" {
					validate[i]["type"] = "number"
				}
			}
		}
	}

	// 为switch组件转换属性名称
	if result["type"] == "switch" {
		if props, ok := result["props"].(map[string]interface{}); ok {
			// 转换连字符属性名为驼峰命名
			if activeText, exists := props["active-text"]; exists {
				props["activeText"] = activeText
				delete(props, "active-text")
			}
			if inactiveText, exists := props["inactive-text"]; exists {
				props["inactiveText"] = inactiveText
				delete(props, "inactive-text")
			}
			if activeValue, exists := props["active-value"]; exists {
				props["activeValue"] = activeValue
				delete(props, "active-value")
			}
			if inactiveValue, exists := props["inactive-value"]; exists {
				props["inactiveValue"] = inactiveValue
				delete(props, "inactive-value")
			}
		}
	}

	// 递归处理control中的组件
	if control, ok := result["control"].([]map[string]interface{}); ok {
		for _, ctrl := range control {
			if rules, ok := ctrl["rule"].([]map[string]interface{}); ok {
				for j, rule := range rules {
					// 递归处理control中的每个组件
					rules[j] = b.parseControlComponent(rule)
				}
			}
		}
	}

	return result
}

// parseControlComponent recursively parses components in control rules
func (b *Bootstrap) parseControlComponent(component map[string]interface{}) map[string]interface{} {
	// 转换组件类型
	if componentType, ok := component["type"].(string); ok {
		switch componentType {
		case "el-input":
			component["type"] = "input"
		case "el-input-number":
			component["type"] = "inputNumber"
		case "el-switch":
			component["type"] = "switch"
		case "el-radio-group":
			component["type"] = "radio"
		case "el-checkbox-group":
			component["type"] = "checkbox"
		case "el-select":
			component["type"] = "select"
		case "el-date-picker":
			component["type"] = "datePicker"
		case "el-time-picker":
			component["type"] = "timePicker"
		case "el-color-picker":
			component["type"] = "colorPicker"
		case "el-upload":
			component["type"] = "upload"
		case "el-rate":
			component["type"] = "rate"
		case "el-slider":
			component["type"] = "slider"
		case "el-cascader":
			component["type"] = "cascader"
		case "el-tree":
			component["type"] = "tree"
		case "el-button":
			component["type"] = "button"
		}
	}

	// 转换验证规则格式
	if validate, ok := component["validate"].([]map[string]interface{}); ok {
		for i, rule := range validate {
			if rule["type"] == "required" {
				validate[i]["type"] = "string"
				validate[i]["trigger"] = "change"
				// 修改message格式
				if title, ok := component["title"].(string); ok {
					validate[i]["message"] = "请输入" + title
				}
			}
		}
	}

	// 为input类型组件添加text类型属性
	if component["type"] == "input" {
		if props, ok := component["props"].(map[string]interface{}); ok {
			if _, hasType := props["type"]; !hasType {
				props["type"] = "text"
			}
		}
	}

	// 为数字类型组件添加数字验证
	if component["type"] == "inputNumber" {
		if validate, ok := component["validate"].([]map[string]interface{}); ok {
			for i, rule := range validate {
				if rule["type"] == "string" {
					validate[i]["type"] = "number"
				}
			}
		}
	}

	return component
}

// Config implements ElementUI configuration
type Config struct {
	config map[string]interface{}
}

// NewConfig creates a new ElementUI configuration
func NewConfig(config map[string]interface{}) *Config {
	if config == nil {
		config = make(map[string]interface{})
	}

	// Set ElementUI defaults
	if _, exists := config["submitBtn"]; !exists {
		config["submitBtn"] = true
	}
	if _, exists := config["resetBtn"]; !exists {
		config["resetBtn"] = false
	}
	if _, exists := config["form"]; !exists {
		config["form"] = map[string]interface{}{
			"labelPosition": "right",
			"labelWidth":    "125px",
			"size":          "medium",
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
