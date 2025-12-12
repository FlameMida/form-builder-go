package formbuilder

// radio.go 实现Radio单选框组件

// Radio 单选框组件
type Radio struct {
	Builder[*Radio]
	options []Option
}

// NewRadio 创建一个新的Radio组件
func NewRadio(field, title string, value ...interface{}) *Radio {
	radio := &Radio{}
	radio.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "radio",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		radio.data.Value = value[0]
	}
	radio.inst = radio
	radio.options = []Option{}
	return radio
}

// NewIviewRadio 创建iView版本的Radio组件
// 注意：type值与Element UI版本相同，框架差异由全局配置决定
func NewIviewRadio(field, title string, value ...interface{}) *Radio {
	radio := NewRadio(field, title, value...)
	radio.data.RuleType = "radio"
	return radio
}

// SetOptions 设置选项列表
func (r *Radio) SetOptions(options []Option) *Radio {
	r.options = options
	return r
}

// AppendOption 追加单个选项
func (r *Radio) AppendOption(option Option) *Radio {
	r.options = append(r.options, option)
	return r
}

// Disabled 设置是否禁用
func (r *Radio) Disabled(disabled bool) *Radio {
	r.data.Props["disabled"] = disabled
	return r
}

// Size 设置尺寸
func (r *Radio) Size(size string) *Radio {
	r.data.Props["size"] = size
	return r
}

// Button 设置是否为按钮样式
// 对应PHP: $this->props['type'] = 'button'
func (r *Radio) Button(enable bool) *Radio {
	if enable {
		r.Props("type", "button")
	} else {
		// 删除 props 中的 type
		if r.data.Props != nil {
			delete(r.data.Props, "type")
		}
	}
	return r
}

// GetField 实现Component接口
func (r *Radio) GetField() string {
	return r.data.Field
}

// GetType 实现Component接口
func (r *Radio) GetType() string {
	return r.data.RuleType
}

// Build 实现Component接口
func (r *Radio) Build() map[string]interface{} {
	result := buildComponent(r.data)
	if len(r.options) > 0 {
		opts := make([]map[string]interface{}, len(r.options))
		for i, opt := range r.options {
			opts[i] = opt.ToMap()
		}
		result["options"] = opts
	}
	return result
}
