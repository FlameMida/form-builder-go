package formbuilder

// timepicker.go 实现TimePicker时间选择器组件

// TimePicker 时间选择器组件
type TimePicker struct {
	Builder[*TimePicker]
}

// NewTimePicker 创建时间选择器
func NewTimePicker(field, title string, value ...interface{}) *TimePicker {
	tp := &TimePicker{}
	tp.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "timePicker",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		tp.data.Value = value[0]
	}
	tp.inst = tp
	return tp
}

// IsRange 设置是否为时间范围选择
func (t *TimePicker) IsRange(enable bool) *TimePicker {
	t.data.Props["is-range"] = enable
	return t
}

// Format 设置显示格式
func (t *TimePicker) Format(format string) *TimePicker {
	t.data.Props["format"] = format
	return t
}

// ValueFormat 设置绑定值的格式
func (t *TimePicker) ValueFormat(format string) *TimePicker {
	t.data.Props["value-format"] = format
	return t
}

// Placeholder 设置占位符
func (t *TimePicker) Placeholder(text string) *TimePicker {
	t.data.Props["placeholder"] = text
	return t
}

// Clearable 设置是否可清空
func (t *TimePicker) Clearable(enable bool) *TimePicker {
	t.data.Props["clearable"] = enable
	return t
}

// Disabled 设置是否禁用
func (t *TimePicker) Disabled(disabled bool) *TimePicker {
	t.data.Props["disabled"] = disabled
	return t
}

// Size 设置尺寸
func (t *TimePicker) Size(size string) *TimePicker {
	t.data.Props["size"] = size
	return t
}

// GetField 实现Component接口
func (t *TimePicker) GetField() string {
	return t.data.Field
}

// GetType 实现Component接口
func (t *TimePicker) GetType() string {
	return t.data.RuleType
}

// Build 实现Component接口
func (t *TimePicker) Build() map[string]interface{} {
	return buildComponent(t.data)
}
