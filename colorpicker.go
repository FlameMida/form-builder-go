package formbuilder

// colorpicker.go 实现ColorPicker颜色选择器组件

// ColorPicker 颜色选择器组件
type ColorPicker struct {
	Builder[*ColorPicker]
}

// NewColorPicker 创建颜色选择器
func NewColorPicker(field, title string, value ...interface{}) *ColorPicker {
	cp := &ColorPicker{}
	cp.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "colorPicker",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		cp.data.Value = value[0]
	}
	cp.inst = cp
	return cp
}

// ShowAlpha 设置是否支持透明度选择
func (c *ColorPicker) ShowAlpha(show bool) *ColorPicker {
	c.data.Props["show-alpha"] = show
	return c
}

// ColorFormat 设置颜色格式（hsl/hsv/hex/rgb）
func (c *ColorPicker) ColorFormat(format string) *ColorPicker {
	c.data.Props["color-format"] = format
	return c
}

// Predefine 设置预定义颜色
func (c *ColorPicker) Predefine(colors []string) *ColorPicker {
	c.data.Props["predefine"] = colors
	return c
}

// Disabled 设置是否禁用
func (c *ColorPicker) Disabled(disabled bool) *ColorPicker {
	c.data.Props["disabled"] = disabled
	return c
}

// Size 设置尺寸
func (c *ColorPicker) Size(size string) *ColorPicker {
	c.data.Props["size"] = size
	return c
}

// GetField 实现Component接口
func (c *ColorPicker) GetField() string {
	return c.data.Field
}

// GetType 实现Component接口
func (c *ColorPicker) GetType() string {
	return c.data.RuleType
}

// Build 实现Component接口
func (c *ColorPicker) Build() map[string]interface{} {
	return buildComponent(c.data)
}
