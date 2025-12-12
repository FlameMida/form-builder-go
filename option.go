package formbuilder

// option.go 定义选项结构，用于Select、Radio、Checkbox、Cascader等组件
// 对应PHP的Option类

// Option 表示一个选项
// 用于Select、Radio、Checkbox、Cascader等组件
//
// 使用示例：
//
//	select.SetOptions([]Option{
//	    {Value: "1", Label: "选项1"},
//	    {Value: "2", Label: "选项2", Disabled: true},
//	})
//
// 级联选项示例：
//
//	cascader.SetOptions([]Option{
//	    {
//	        Value: "beijing",
//	        Label: "北京",
//	        Children: []Option{
//	            {Value: "chaoyang", Label: "朝阳区"},
//	            {Value: "haidian", Label: "海淀区"},
//	        },
//	    },
//	})
type Option struct {
	// Value 选项值（提交到后端的值）
	Value interface{}

	// Label 选项标签（显示给用户的文本）
	Label string

	// Disabled 是否禁用此选项
	Disabled bool

	// Children 子选项（用于级联选择）
	Children []Option

	// Extra 额外的自定义字段
	// 某些UI组件可能需要额外的字段，如icon、color等
	Extra map[string]interface{}
}

// ToMap 将Option转换为map，用于JSON序列化
func (o Option) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"value": o.Value,
		"label": o.Label,
	}

	if o.Disabled {
		m["disabled"] = true
	}

	// 递归处理子选项
	if len(o.Children) > 0 {
		children := make([]map[string]interface{}, len(o.Children))
		for i, child := range o.Children {
			children[i] = child.ToMap()
		}
		m["children"] = children
	}

	// 添加额外字段
	if len(o.Extra) > 0 {
		for k, v := range o.Extra {
			m[k] = v
		}
	}

	return m
}

// NewOption 创建一个新选项（便捷构造函数）
func NewOption(value interface{}, label string) Option {
	return Option{
		Value: value,
		Label: label,
	}
}

// NewOptions 批量创建选项（便捷构造函数）
// 从map创建选项数组
//
// 使用示例：
//
//	options := NewOptions(map[interface{}]string{
//	    1: "选项1",
//	    2: "选项2",
//	    3: "选项3",
//	})
func NewOptions(data map[interface{}]string) []Option {
	options := make([]Option, 0, len(data))
	for value, label := range data {
		options = append(options, Option{
			Value: value,
			Label: label,
		})
	}
	return options
}

// NewOptionsFromSlice 从字符串切片创建选项
// value和label相同
//
// 使用示例：
//
//	options := NewOptionsFromSlice([]string{"选项1", "选项2", "选项3"})
func NewOptionsFromSlice(labels []string) []Option {
	options := make([]Option, len(labels))
	for i, label := range labels {
		options[i] = Option{
			Value: label,
			Label: label,
		}
	}
	return options
}

// NewOptionsFromPairs 从键值对数组创建选项
// 每对[0]是value，[1]是label
//
// 使用示例：
//
//	options := NewOptionsFromPairs([][2]interface{}{
//	    {1, "选项1"},
//	    {2, "选项2"},
//	})
func NewOptionsFromPairs(pairs [][2]interface{}) []Option {
	options := make([]Option, len(pairs))
	for i, pair := range pairs {
		options[i] = Option{
			Value: pair[0],
			Label: pair[1].(string),
		}
	}
	return options
}
