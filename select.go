package formbuilder

// select.go 实现Select下拉选择组件

// Select 下拉选择框组件
type Select struct {
	Builder[*Select]
	options []Option // 选项列表
}

// NewSelect 创建一个新的Select组件
func NewSelect(field, title string, value ...interface{}) *Select {
	sel := &Select{}
	sel.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "select",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		sel.data.Value = value[0]
	}
	sel.inst = sel
	sel.options = []Option{}
	return sel
}

// NewIviewSelect 创建iView版本的Select组件
func NewIviewSelect(field, title string, value ...interface{}) *Select {
	sel := NewSelect(field, title, value...)
	sel.data.RuleType = "select"
	return sel
}

// SetOptions 设置选项列表
func (s *Select) SetOptions(options []Option) *Select {
	s.options = options
	return s
}

// AppendOption 追加单个选项
func (s *Select) AppendOption(option Option) *Select {
	s.options = append(s.options, option)
	return s
}

// Multiple 设置是否多选
func (s *Select) Multiple(enable bool) *Select {
	s.data.Props["multiple"] = enable
	return s
}

// Disabled 设置是否禁用
func (s *Select) Disabled(disabled bool) *Select {
	s.data.Props["disabled"] = disabled
	return s
}

// Clearable 设置是否可清空
func (s *Select) Clearable(enable bool) *Select {
	s.data.Props["clearable"] = enable
	return s
}

// Filterable 设置是否可搜索
func (s *Select) Filterable(enable bool) *Select {
	s.data.Props["filterable"] = enable
	return s
}

// Remote 设置是否远程搜索
func (s *Select) Remote(enable bool) *Select {
	s.data.Props["remote"] = enable
	return s
}

// RemoteMethod 设置远程搜索方法（JavaScript函数名）
func (s *Select) RemoteMethod(method string) *Select {
	s.data.Props["remote-method"] = method
	return s
}

// Placeholder 设置占位符
func (s *Select) Placeholder(text string) *Select {
	s.data.Props["placeholder"] = text
	return s
}

// Size 设置尺寸
func (s *Select) Size(size string) *Select {
	s.data.Props["size"] = size
	return s
}

// CollapseTags 设置多选时是否折叠标签
func (s *Select) CollapseTags(enable bool) *Select {
	s.data.Props["collapse-tags"] = enable
	return s
}

// MultipleLimit 设置多选时最多可选项目数
func (s *Select) MultipleLimit(limit int) *Select {
	s.data.Props["multiple-limit"] = limit
	return s
}

// AllowCreate 设置是否允许创建新选项
func (s *Select) AllowCreate(enable bool) *Select {
	s.data.Props["allow-create"] = enable
	return s
}

// DefaultFirstOption 设置是否默认选中第一个选项
func (s *Select) DefaultFirstOption(enable bool) *Select {
	s.data.Props["default-first-option"] = enable
	return s
}

// GetField 实现Component接口
func (s *Select) GetField() string {
	return s.data.Field
}

// GetType 实现Component接口
func (s *Select) GetType() string {
	return s.data.RuleType
}

// Build 实现Component接口
func (s *Select) Build() map[string]interface{} {
	result := buildComponent(s.data)
	// 添加options
	if len(s.options) > 0 {
		opts := make([]map[string]interface{}, len(s.options))
		for i, opt := range s.options {
			opts[i] = opt.ToMap()
		}
		result["options"] = opts
	}
	return result
}
