package formbuilder

// slider.go 实现Slider滑块组件

// Slider 滑块组件
type Slider struct {
	Builder[*Slider]
}

// NewSlider 创建滑块
func NewSlider(field, title string, value ...interface{}) *Slider {
	slider := &Slider{}
	slider.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "slider",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		slider.data.Value = value[0]
	}
	slider.inst = slider
	return slider
}

// Min 设置最小值
func (s *Slider) Min(min float64) *Slider {
	s.data.Props["min"] = min
	return s
}

// Max 设置最大值
func (s *Slider) Max(max float64) *Slider {
	s.data.Props["max"] = max
	return s
}

// Step 设置步长
func (s *Slider) Step(step float64) *Slider {
	s.data.Props["step"] = step
	return s
}

// Range 设置是否为范围选择
func (s *Slider) Range(enable bool) *Slider {
	s.data.Props["range"] = enable
	return s
}

// ShowStops 设置是否显示间断点
func (s *Slider) ShowStops(show bool) *Slider {
	s.data.Props["show-stops"] = show
	return s
}

// ShowInput 设置是否显示输入框
func (s *Slider) ShowInput(show bool) *Slider {
	s.data.Props["show-input"] = show
	return s
}

// Disabled 设置是否禁用
func (s *Slider) Disabled(disabled bool) *Slider {
	s.data.Props["disabled"] = disabled
	return s
}

// GetField 实现Component接口
func (s *Slider) GetField() string {
	return s.data.Field
}

// GetType 实现Component接口
func (s *Slider) GetType() string {
	return s.data.RuleType
}

// Build 实现Component接口
func (s *Slider) Build() map[string]interface{} {
	return buildComponent(s.data)
}
