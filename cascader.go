package formbuilder

// cascader.go 实现Cascader级联选择器组件

// Cascader 级联选择器组件
type Cascader struct {
	Builder[*Cascader]
	options []Option
}

// NewCascader 创建级联选择器
func NewCascader(field, title string, value ...interface{}) *Cascader {
	cascader := &Cascader{}
	cascader.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "cascader",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		cascader.data.Value = value[0]
	}
	cascader.inst = cascader
	cascader.options = []Option{}
	return cascader
}

// SetOptions 设置级联选项
func (c *Cascader) SetOptions(options []Option) *Cascader {
	c.options = options
	return c
}

// CascaderProps 设置配置选项
// props示例: map[string]interface{}{"value": "id", "label": "name", "children": "children"}
func (c *Cascader) CascaderProps(props map[string]interface{}) *Cascader {
	c.data.Props["props"] = props
	return c
}

// Separator 设置选项分隔符
func (c *Cascader) Separator(separator string) *Cascader {
	c.data.Props["separator"] = separator
	return c
}

// Filterable 设置是否可搜索
func (c *Cascader) Filterable(enable bool) *Cascader {
	c.data.Props["filterable"] = enable
	return c
}

// Clearable 设置是否可清空
func (c *Cascader) Clearable(enable bool) *Cascader {
	c.data.Props["clearable"] = enable
	return c
}

// ShowAllLevels 设置是否显示完整路径
func (c *Cascader) ShowAllLevels(show bool) *Cascader {
	c.data.Props["show-all-levels"] = show
	return c
}

// Placeholder 设置占位符
func (c *Cascader) Placeholder(text string) *Cascader {
	c.data.Props["placeholder"] = text
	return c
}

// Disabled 设置是否禁用
func (c *Cascader) Disabled(disabled bool) *Cascader {
	c.data.Props["disabled"] = disabled
	return c
}

// Size 设置尺寸
func (c *Cascader) Size(size string) *Cascader {
	c.data.Props["size"] = size
	return c
}

// GetField 实现Component接口
func (c *Cascader) GetField() string {
	return c.data.Field
}

// GetType 实现Component接口
func (c *Cascader) GetType() string {
	return c.data.RuleType
}

// Build 实现Component接口
func (c *Cascader) Build() map[string]interface{} {
	result := buildComponent(c.data)
	if len(c.options) > 0 {
		opts := make([]map[string]interface{}, len(c.options))
		for i, opt := range c.options {
			opts[i] = opt.ToMap()
		}

		// 确保 props 对象存在
		if result["props"] == nil {
			result["props"] = make(map[string]interface{})
		}

		// 将 options 添加到 props 内部
		if props, ok := result["props"].(map[string]interface{}); ok {
			props["options"] = opts
		}
	}
	return result
}
