package formbuilder

// switch.go 实现Switch开关组件

// Switch 开关组件
type Switch struct {
	Builder[*Switch]
}

// NewSwitch 创建开关
func NewSwitch(field, title string, value ...interface{}) *Switch {
	sw := &Switch{}
	sw.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "switch",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		sw.data.Value = value[0]
	}
	sw.inst = sw
	return sw
}

// ActiveText 设置打开时的文字
func (s *Switch) ActiveText(text string) *Switch {
	s.data.Props["active-text"] = text
	return s
}

// InactiveText 设置关闭时的文字
func (s *Switch) InactiveText(text string) *Switch {
	s.data.Props["inactive-text"] = text
	return s
}

// ActiveValue 设置打开时的值
func (s *Switch) ActiveValue(value interface{}) *Switch {
	s.data.Props["active-value"] = value
	return s
}

// InactiveValue 设置关闭时的值
func (s *Switch) InactiveValue(value interface{}) *Switch {
	s.data.Props["inactive-value"] = value
	return s
}

// ActiveColor 设置打开时的背景色
func (s *Switch) ActiveColor(color string) *Switch {
	s.data.Props["active-color"] = color
	return s
}

// InactiveColor 设置关闭时的背景色
func (s *Switch) InactiveColor(color string) *Switch {
	s.data.Props["inactive-color"] = color
	return s
}

// Disabled 设置是否禁用
func (s *Switch) Disabled(disabled bool) *Switch {
	s.data.Props["disabled"] = disabled
	return s
}

// GetField 实现Component接口
func (s *Switch) GetField() string {
	return s.data.Field
}

// GetType 实现Component接口
func (s *Switch) GetType() string {
	return s.data.RuleType
}

// Build 实现Component接口
func (s *Switch) Build() map[string]interface{} {
	return buildComponent(s.data)
}
