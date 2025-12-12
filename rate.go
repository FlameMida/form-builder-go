package formbuilder

// rate.go 实现Rate评分组件

// Rate 评分组件
type Rate struct {
	Builder[*Rate]
}

// NewRate 创建评分组件
func NewRate(field, title string, value ...interface{}) *Rate {
	rate := &Rate{}
	rate.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "rate",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		rate.data.Value = value[0]
	}
	rate.inst = rate
	return rate
}

// Max 设置最大分值
func (r *Rate) Max(max int) *Rate {
	r.data.Props["max"] = max
	return r
}

// AllowHalf 设置是否允许半选
func (r *Rate) AllowHalf(allow bool) *Rate {
	r.data.Props["allow-half"] = allow
	return r
}

// ShowText 设置是否显示辅助文字
func (r *Rate) ShowText(show bool) *Rate {
	r.data.Props["show-text"] = show
	return r
}

// ShowScore 设置是否显示当前分数
func (r *Rate) ShowScore(show bool) *Rate {
	r.data.Props["show-score"] = show
	return r
}

// Colors 设置icon颜色数组
func (r *Rate) Colors(colors []string) *Rate {
	r.data.Props["colors"] = colors
	return r
}

// Texts 设置辅助文字数组
func (r *Rate) Texts(texts []string) *Rate {
	r.data.Props["texts"] = texts
	return r
}

// Disabled 设置是否禁用
func (r *Rate) Disabled(disabled bool) *Rate {
	r.data.Props["disabled"] = disabled
	return r
}

// GetField 实现Component接口
func (r *Rate) GetField() string {
	return r.data.Field
}

// GetType 实现Component接口
func (r *Rate) GetType() string {
	return r.data.RuleType
}

// Build 实现Component接口
func (r *Rate) Build() map[string]interface{} {
	return buildComponent(r.data)
}
