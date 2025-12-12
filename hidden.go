package formbuilder

// hidden.go 实现Hidden隐藏字段组件

// Hidden 隐藏字段组件
type Hidden struct {
	Builder[*Hidden]
}

// NewHidden 创建隐藏字段
func NewHidden(field string, value ...interface{}) *Hidden {
	hidden := &Hidden{}
	hidden.data = &ComponentData{
		Field:    field,
		Title:    "",
		RuleType: "hidden",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		hidden.data.Value = value[0]
	}
	hidden.inst = hidden
	return hidden
}

// GetField 实现Component接口
func (h *Hidden) GetField() string {
	return h.data.Field
}

// GetType 实现Component接口
func (h *Hidden) GetType() string {
	return h.data.RuleType
}

// Build 实现Component接口
func (h *Hidden) Build() map[string]interface{} {
	return buildComponent(h.data)
}
