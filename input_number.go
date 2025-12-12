package formbuilder

// input_number.go 实现InputNumber数字输入框组件

// InputNumber 数字输入框组件
type InputNumber struct {
	Builder[*InputNumber]
}

// NewInputNumber 创建一个新的InputNumber组件
func NewInputNumber(field, title string, value ...interface{}) *InputNumber {
	num := &InputNumber{}
	num.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "inputNumber",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		num.data.Value = value[0]
	}
	num.inst = num
	return num
}

// NewIviewInputNumber 创建iView版本
func NewIviewInputNumber(field, title string, value ...interface{}) *InputNumber {
	num := NewInputNumber(field, title, value...)
	num.data.RuleType = "inputNumber"
	return num
}

// Min 设置最小值
func (n *InputNumber) Min(min float64) *InputNumber {
	n.data.Props["min"] = min
	return n
}

// Max 设置最大值
func (n *InputNumber) Max(max float64) *InputNumber {
	n.data.Props["max"] = max
	return n
}

// Step 设置步长
func (n *InputNumber) Step(step float64) *InputNumber {
	n.data.Props["step"] = step
	return n
}

// Precision 设置精度（小数位数）
func (n *InputNumber) Precision(precision int) *InputNumber {
	n.data.Props["precision"] = precision
	return n
}

// Controls 设置是否显示控制按钮
func (n *InputNumber) Controls(enable bool) *InputNumber {
	n.data.Props["controls"] = enable
	return n
}

// ControlsPosition 设置控制按钮位置（right）
func (n *InputNumber) ControlsPosition(position string) *InputNumber {
	n.data.Props["controls-position"] = position
	return n
}

// Disabled 设置是否禁用
func (n *InputNumber) Disabled(disabled bool) *InputNumber {
	n.data.Props["disabled"] = disabled
	return n
}

// Placeholder 设置占位符
func (n *InputNumber) Placeholder(text string) *InputNumber {
	n.data.Props["placeholder"] = text
	return n
}

// Size 设置尺寸
func (n *InputNumber) Size(size string) *InputNumber {
	n.data.Props["size"] = size
	return n
}

// GetField 实现Component接口
func (n *InputNumber) GetField() string {
	return n.data.Field
}

// GetType 实现Component接口
func (n *InputNumber) GetType() string {
	return n.data.RuleType
}

// Build 实现Component接口
func (n *InputNumber) Build() map[string]interface{} {
	return buildComponent(n.data)
}

// Number 便捷构造函数别名
func Number(field, title string, value ...interface{}) *InputNumber {
	return NewInputNumber(field, title, value...)
}
