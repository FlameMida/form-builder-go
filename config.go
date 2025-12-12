package formbuilder

// config.go 实现表单配置类
// 对应PHP的Config类

// Config 表单配置
type Config struct {
	submitBtn map[string]interface{} // 提交按钮配置
	resetBtn  map[string]interface{} // 重置按钮配置
	formStyle map[string]interface{} // 表单样式配置
	row       map[string]interface{} // 行布局配置
	info      map[string]interface{} // 提示信息配置
	global    map[string]interface{} // 全局配置
}

// NewConfig 创建默认配置
func NewConfig() *Config {
	return &Config{
		submitBtn: map[string]interface{}{
			"show":      true,
			"innerText": "提交",
		},
		resetBtn: map[string]interface{}{
			"show":      true,
			"innerText": "重置",
		},
		global: make(map[string]interface{}),
	}
}

// NewElmConfig 创建Element UI配置
func NewElmConfig() *Config {
	return NewConfig()
}

// NewIviewConfig 创建iView配置
func NewIviewConfig() *Config {
	return NewConfig()
}

// SubmitBtn 配置提交按钮
// show: 是否显示
// text: 按钮文本（可选）
func (c *Config) SubmitBtn(show bool, text ...string) *Config {
	c.submitBtn["show"] = show
	if len(text) > 0 {
		c.submitBtn["innerText"] = text[0]
	}
	return c
}

// SetSubmitBtnProps 设置提交按钮属性
func (c *Config) SetSubmitBtnProps(props map[string]interface{}) *Config {
	for k, v := range props {
		c.submitBtn[k] = v
	}
	return c
}

// ResetBtn 配置重置按钮
func (c *Config) ResetBtn(show bool, text ...string) *Config {
	c.resetBtn["show"] = show
	if len(text) > 0 {
		c.resetBtn["innerText"] = text[0]
	}
	return c
}

// SetResetBtnProps 设置重置按钮属性
func (c *Config) SetResetBtnProps(props map[string]interface{}) *Config {
	for k, v := range props {
		c.resetBtn[k] = v
	}
	return c
}

// FormStyle 设置表单样式
func (c *Config) FormStyle(style map[string]interface{}) *Config {
	c.formStyle = style
	return c
}

// Row 设置行布局
func (c *Config) Row(row map[string]interface{}) *Config {
	c.row = row
	return c
}

// Info 设置提示信息
// type: success/warning/error/info
func (c *Config) Info(infoType string, title string, show bool) *Config {
	c.info = map[string]interface{}{
		"type":  infoType,
		"title": title,
		"show":  show,
	}
	return c
}

// Global 设置全局配置
func (c *Config) Global(key string, value interface{}) *Config {
	if c.global == nil {
		c.global = make(map[string]interface{})
	}
	c.global[key] = value
	return c
}

// SetGlobal 批量设置全局配置
func (c *Config) SetGlobal(global map[string]interface{}) *Config {
	if c.global == nil {
		c.global = make(map[string]interface{})
	}
	for k, v := range global {
		c.global[k] = v
	}
	return c
}

// ToMap 将配置转换为map，用于JSON序列化
func (c *Config) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	if c.submitBtn != nil && len(c.submitBtn) > 0 {
		result["submitBtn"] = c.submitBtn
	}

	if c.resetBtn != nil && len(c.resetBtn) > 0 {
		result["resetBtn"] = c.resetBtn
	}

	if c.formStyle != nil && len(c.formStyle) > 0 {
		result["formStyle"] = c.formStyle
	}

	if c.row != nil && len(c.row) > 0 {
		result["row"] = c.row
	}

	if c.info != nil && len(c.info) > 0 {
		result["info"] = c.info
	}

	if c.global != nil && len(c.global) > 0 {
		result["global"] = c.global
	}

	return result
}
