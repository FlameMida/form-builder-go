package formbuilder

// checkbox.go 实现Checkbox复选框组件

// Checkbox 复选框组件
type Checkbox struct {
	Builder[*Checkbox]
	options []Option
}

// NewCheckbox 创建一个新的Checkbox组件
func NewCheckbox(field, title string, value ...interface{}) *Checkbox {
	checkbox := &Checkbox{}
	checkbox.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "checkbox",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		checkbox.data.Value = value[0]
	}
	checkbox.inst = checkbox
	checkbox.options = []Option{}
	return checkbox
}

// NewIviewCheckbox 创建iView版本的Checkbox组件
func NewIviewCheckbox(field, title string, value ...interface{}) *Checkbox {
	checkbox := NewCheckbox(field, title, value...)
	checkbox.data.RuleType = "checkbox"
	return checkbox
}

// SetOptions 设置选项列表
func (c *Checkbox) SetOptions(options []Option) *Checkbox {
	c.options = options
	return c
}

// AppendOption 追加单个选项
func (c *Checkbox) AppendOption(option Option) *Checkbox {
	c.options = append(c.options, option)
	return c
}

// Disabled 设置是否禁用
func (c *Checkbox) Disabled(disabled bool) *Checkbox {
	c.data.Props["disabled"] = disabled
	return c
}

// Size 设置尺寸
func (c *Checkbox) Size(size string) *Checkbox {
	c.data.Props["size"] = size
	return c
}

// Min 设置最小选中数量
func (c *Checkbox) Min(min int) *Checkbox {
	c.data.Props["min"] = min
	return c
}

// Max 设置最大选中数量
func (c *Checkbox) Max(max int) *Checkbox {
	c.data.Props["max"] = max
	return c
}

// CheckedColor 设置选中颜色
func (c *Checkbox) CheckedColor(color string) *Checkbox {
	c.data.Props["checked-color"] = color
	return c
}

// GetField 实现Component接口
func (c *Checkbox) GetField() string {
	return c.data.Field
}

// GetType 实现Component接口
func (c *Checkbox) GetType() string {
	return c.data.RuleType
}

// Build 实现Component接口
func (c *Checkbox) Build() map[string]interface{} {
	result := buildComponent(c.data)
	if len(c.options) > 0 {
		opts := make([]map[string]interface{}, len(c.options))
		for i, opt := range c.options {
			opts[i] = opt.ToMap()
		}
		result["options"] = opts
	}
	return result
}
