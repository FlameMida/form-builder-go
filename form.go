package formbuilder

import (
	"encoding/json"
	"fmt"
)

// form.go 实现Form核心类
// 对应PHP的Form.php

// Bootstrap UI引导接口
// 用于初始化不同UI框架（Element UI、iView等）的资源
type Bootstrap interface {
	// Init 初始化表单，注入必要的脚本和样式
	Init(form *Form)

	// GetScripts 获取JavaScript脚本列表
	GetScripts() []string

	// GetStyles 获取CSS样式列表
	GetStyles() []string
}

// Form 表单核心类
// 管理表单规则、配置、数据和输出
type Form struct {
	action       string                 // 表单提交地址
	method       string                 // HTTP方法（POST/GET）
	rules        []Component            // 组件规则数组
	config       *Config                // 表单配置
	formData     map[string]interface{} // 表单初始数据
	ui           Bootstrap              // UI引导实例
	dependScript []string               // 依赖脚本列表
	title        string                 // 表单标题
}

// NewElmForm 创建Element UI表单
// 对应PHP的 Form::elm()
func NewElmForm(action string, rules []Component, config *Config) *Form {
	if config == nil {
		config = NewElmConfig()
	}

	form := &Form{
		action:       action,
		method:       "POST",
		rules:        rules,
		config:       config,
		ui:           NewElmBootstrap(),
		formData:     make(map[string]interface{}),
		dependScript: []string{},
	}

	// 初始化UI（注入CDN资源）
	form.ui.Init(form)

	// 验证字段唯一性
	if err := form.checkFieldUnique(); err != nil {
		panic(err) // 或者返回error，取决于API设计偏好
	}

	return form
}

// NewIviewForm 创建iView v3表单
func NewIviewForm(action string, rules []Component, config *Config) *Form {
	if config == nil {
		config = NewIviewConfig()
	}

	form := &Form{
		action:       action,
		method:       "POST",
		rules:        rules,
		config:       config,
		ui:           NewIviewBootstrap(3), // v3
		formData:     make(map[string]interface{}),
		dependScript: []string{},
	}

	form.ui.Init(form)

	if err := form.checkFieldUnique(); err != nil {
		panic(err)
	}

	return form
}

// NewIview4Form 创建iView v4表单
func NewIview4Form(action string, rules []Component, config *Config) *Form {
	if config == nil {
		config = NewIviewConfig()
	}

	form := &Form{
		action:       action,
		method:       "POST",
		rules:        rules,
		config:       config,
		ui:           NewIviewBootstrap(4), // v4
		formData:     make(map[string]interface{}),
		dependScript: []string{},
	}

	form.ui.Init(form)

	if err := form.checkFieldUnique(); err != nil {
		panic(err)
	}

	return form
}

// SetRule 设置组件规则数组
func (f *Form) SetRule(rules []Component) *Form {
	f.rules = rules
	if err := f.checkFieldUnique(); err != nil {
		panic(err)
	}
	return f
}

// Append 追加组件到末尾
func (f *Form) Append(component Component) *Form {
	f.rules = append(f.rules, component)
	if err := f.checkFieldUnique(); err != nil {
		panic(err)
	}
	return f
}

// Prepend 在开头插入组件
func (f *Form) Prepend(component Component) *Form {
	f.rules = append([]Component{component}, f.rules...)
	if err := f.checkFieldUnique(); err != nil {
		panic(err)
	}
	return f
}

// SetAction 设置表单提交地址
func (f *Form) SetAction(action string) *Form {
	f.action = action
	return f
}

// SetMethod 设置HTTP方法
func (f *Form) SetMethod(method string) *Form {
	f.method = method
	return f
}

// SetTitle 设置表单标题
func (f *Form) SetTitle(title string) *Form {
	f.title = title
	return f
}

// SetConfig 设置表单配置
func (f *Form) SetConfig(config *Config) *Form {
	f.config = config
	return f
}

// FormData 设置表单初始数据
// 数据将被应用到对应field的组件上
func (f *Form) FormData(data map[string]interface{}) *Form {
	f.formData = data
	return f
}

// SetValue 设置单个字段的值
func (f *Form) SetValue(field string, value interface{}) *Form {
	if f.formData == nil {
		f.formData = make(map[string]interface{})
	}
	f.formData[field] = value
	return f
}

// checkFieldUnique 检查字段唯一性
// 对应PHP的checkFieldUnique()方法
func (f *Form) checkFieldUnique() error {
	fields := make(map[string]bool)
	return f.checkFieldsRecursive(f.rules, fields)
}

// checkFieldsRecursive 递归检查字段唯一性
// 包括control中的嵌套组件
func (f *Form) checkFieldsRecursive(rules []Component, fields map[string]bool) error {
	for _, rule := range rules {
		field := rule.GetField()
		if field == "" {
			continue // 没有field的组件（如Hidden可能没有）跳过
		}

		if fields[field] {
			return fmt.Errorf("field '%s' is not unique", field)
		}
		fields[field] = true

		// 递归检查control中的组件
		if data := f.getComponentData(rule); data != nil {
			for _, ctrl := range data.Control {
				if err := f.checkFieldsRecursive(ctrl.Rule, fields); err != nil {
					return err
				}
			}
			// 递归检查children中的组件
			if err := f.checkFieldsRecursive(data.Children, fields); err != nil {
				return err
			}
		}
	}
	return nil
}

// getComponentData 获取组件的内部数据
// 通过类型断言获取ComponentData
func (f *Form) getComponentData(c Component) *ComponentData {
	// 定义一个接口来获取内部数据
	type dataGetter interface {
		GetData() *ComponentData
	}

	if dg, ok := c.(dataGetter); ok {
		return dg.GetData()
	}
	return nil
}

// FormRule 获取表单规则数组（应用formData后）
// 对应PHP的formRule()方法
func (f *Form) FormRule() []map[string]interface{} {
	rules := make([]map[string]interface{}, len(f.rules))
	for i, rule := range f.rules {
		rules[i] = rule.Build()
	}

	// 应用表单数据
	f.applyFormData(rules)

	return rules
}

// ParseFormRule 返回JSON格式的表单规则
// 对应PHP的parseFormRule()方法
func (f *Form) ParseFormRule() (string, error) {
	rules := f.FormRule()
	data, err := json.Marshal(rules)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FormConfig 获取表单配置
// 对应PHP的formConfig()方法
func (f *Form) FormConfig() map[string]interface{} {
	config := f.config.ToMap()

	// 添加form属性
	config["form"] = map[string]interface{}{
		"action": f.action,
		"method": f.method,
	}

	return config
}

// ParseFormConfig 返回JSON格式的表单配置
// 对应PHP的parseFormConfig()方法
func (f *Form) ParseFormConfig() (string, error) {
	config := f.FormConfig()
	data, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// applyFormData 递归应用表单数据到规则
// 对应PHP的deepSetFormData()方法
func (f *Form) applyFormData(rules []map[string]interface{}) {
	for _, rule := range rules {
		field, ok := rule["field"].(string)
		if !ok || field == "" {
			continue
		}

		// 设置值
		if value, exists := f.formData[field]; exists {
			rule["value"] = value
		}

		// 递归处理control
		if control, ok := rule["control"].([]map[string]interface{}); ok {
			for _, ctrl := range control {
				if ctrlRules, ok := ctrl["rule"].([]map[string]interface{}); ok {
					f.applyFormData(ctrlRules)
				}
			}
		}

		// 递归处理children
		if children, ok := rule["children"].([]map[string]interface{}); ok {
			f.applyFormData(children)
		}
	}
}

// ShowSubmitBtn 设置是否显示提交按钮
func (f *Form) ShowSubmitBtn(show bool) *Form {
	f.config.SubmitBtn(show)
	return f
}

// ShowResetBtn 设置是否显示重置按钮
func (f *Form) ShowResetBtn(show bool) *Form {
	f.config.ResetBtn(show)
	return f
}

// GetRules 获取组件规则数组
func (f *Form) GetRules() []Component {
	return f.rules
}

// GetConfig 获取配置对象
func (f *Form) GetConfig() *Config {
	return f.config
}

// GetFormData 获取表单数据
func (f *Form) GetFormData() map[string]interface{} {
	return f.formData
}

// GetAction 获取表单提交地址
func (f *Form) GetAction() string {
	return f.action
}

// GetMethod 获取HTTP方法
func (f *Form) GetMethod() string {
	return f.method
}
