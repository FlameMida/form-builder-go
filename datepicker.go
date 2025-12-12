package formbuilder

// datepicker.go 实现DatePicker日期选择器组件

// DatePicker 日期选择器组件
type DatePicker struct {
	Builder[*DatePicker]
}

// NewDatePicker 创建日期选择器
func NewDatePicker(field, title string, value ...interface{}) *DatePicker {
	dp := &DatePicker{}
	dp.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "datePicker",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		dp.data.Value = value[0]
	}
	dp.inst = dp
	return dp
}

// DateType 设置日期类型
// year/month/date/dates/week/datetime/datetimerange/daterange/monthrange
func (d *DatePicker) DateType(dtype string) *DatePicker {
	d.data.Props["type"] = dtype
	return d
}

// Format 设置显示格式
func (d *DatePicker) Format(format string) *DatePicker {
	d.data.Props["format"] = format
	return d
}

// ValueFormat 设置绑定值的格式
func (d *DatePicker) ValueFormat(format string) *DatePicker {
	d.data.Props["value-format"] = format
	return d
}

// Placeholder 设置占位符
func (d *DatePicker) Placeholder(text string) *DatePicker {
	d.data.Props["placeholder"] = text
	return d
}

// RangeSeparator 设置范围分隔符
func (d *DatePicker) RangeSeparator(separator string) *DatePicker {
	d.data.Props["range-separator"] = separator
	return d
}

// StartPlaceholder 设置开始日期占位符
func (d *DatePicker) StartPlaceholder(text string) *DatePicker {
	d.data.Props["start-placeholder"] = text
	return d
}

// EndPlaceholder 设置结束日期占位符
func (d *DatePicker) EndPlaceholder(text string) *DatePicker {
	d.data.Props["end-placeholder"] = text
	return d
}

// Clearable 设置是否可清空
func (d *DatePicker) Clearable(enable bool) *DatePicker {
	d.data.Props["clearable"] = enable
	return d
}

// Disabled 设置是否禁用
func (d *DatePicker) Disabled(disabled bool) *DatePicker {
	d.data.Props["disabled"] = disabled
	return d
}

// Editable 设置是否可输入
func (d *DatePicker) Editable(enable bool) *DatePicker {
	d.data.Props["editable"] = enable
	return d
}

// Size 设置尺寸
func (d *DatePicker) Size(size string) *DatePicker {
	d.data.Props["size"] = size
	return d
}

// GetField 实现Component接口
func (d *DatePicker) GetField() string {
	return d.data.Field
}

// GetType 实现Component接口
func (d *DatePicker) GetType() string {
	return d.data.RuleType
}

// Build 实现Component接口
func (d *DatePicker) Build() map[string]interface{} {
	return buildComponent(d.data)
}
