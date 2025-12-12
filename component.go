package formbuilder

// Component 是所有表单组件必须实现的核心接口
// 与PHP的CustomComponentInterface对应
type Component interface {
	// GetField 返回组件的字段名（对应表单提交的字段）
	GetField() string

	// GetType 返回组件类型（如 "input", "select" 等）
	GetType() string

	// Build 将组件转换为map，用于JSON序列化
	// 对应PHP的build()方法
	Build() map[string]interface{}
}

// ComponentData 存储组件的所有数据
// 这是数据与行为分离设计的核心：数据在ComponentData中，行为在Builder[T]中
// 对应PHP中通过BaseRule等trait混入的属性
type ComponentData struct {
	// Field 字段名，对应表单提交的key
	Field string

	// Title 字段标题/标签
	Title string

	// RuleType 组件类型，如 "input", "select", "radio" 等
	RuleType string

	// Value 字段默认值
	Value interface{}

	// Props 组件属性，对应Vue组件的props
	// 例如：{"placeholder": "请输入", "clearable": true}
	Props map[string]interface{}

	// Validate 验证规则数组
	Validate []ValidateRule

	// Control 条件显示规则数组
	// 根据当前组件的值决定显示哪些其他组件
	Control []ControlRule

	// Children 子组件数组
	// 用于插槽、自定义组件等场景
	Children []Component

	// Emit 事件配置（可选）
	Emit map[string]interface{}

	// AppendRule 自定义规则字段
	// 用于添加标准API之外的自定义配置（如 suffix, prefix 等）
	// 在序列化时会合并到最终输出中，允许覆盖同名的标准字段
	// 对应PHP的appendRule属性
	AppendRule map[string]interface{}
}

// Builder 是泛型构建器，提供所有通用的链式方法
// T 是具体组件类型（如 *Input, *Select），必须是指针类型
//
// 设计理念：
// - 通用方法（Required, Value, Props等）只在Builder[T]中定义一次
// - 每个具体组件嵌入Builder[T]，只需实现特有方法
// - 所有方法返回T类型，保证链式调用的类型安全和IDE支持
//
// 使用示例：
//
//	type Input struct {
//	    Builder[*Input]
//	}
//
//	input := NewInput("username", "用户名").
//	    Placeholder("请输入").  // Input特有方法
//	    Required().            // Builder通用方法，返回*Input
//	    Value("default")       // Builder通用方法，返回*Input
type Builder[T any] struct {
	// data 存储组件数据
	data *ComponentData

	// inst 保存具体组件实例，用于返回正确的类型T
	// 在组件的工厂函数中必须设置：builder.inst = componentInstance
	inst T
}

// GetData 返回内部数据（供Form类等使用）
func (b *Builder[T]) GetData() *ComponentData {
	return b.data
}

// Required 添加必填验证规则
// 这是最常用的验证，因此提供便捷方法
func (b *Builder[T]) Required() T {
	b.data.Validate = append(b.data.Validate, RequiredRule{Message: "此项必填"})
	return b.inst
}

// Value 设置组件默认值
func (b *Builder[T]) Value(v interface{}) T {
	b.data.Value = v
	return b.inst
}

// Title 设置组件标题/标签
func (b *Builder[T]) Title(title string) T {
	b.data.Title = title
	return b.inst
}

// Field 设置字段名
func (b *Builder[T]) Field(field string) T {
	b.data.Field = field
	return b.inst
}

// Props 设置单个属性
// 对应Vue组件的props
//
// 使用示例：
//
//	input.Props("placeholder", "请输入用户名")
//	input.Props("clearable", true)
func (b *Builder[T]) Props(key string, val interface{}) T {
	if b.data.Props == nil {
		b.data.Props = make(map[string]interface{})
	}
	b.data.Props[key] = val
	return b.inst
}

// SetProps 批量设置属性
func (b *Builder[T]) SetProps(props map[string]interface{}) T {
	if b.data.Props == nil {
		b.data.Props = make(map[string]interface{})
	}
	for k, v := range props {
		b.data.Props[k] = v
	}
	return b.inst
}

// Control 设置条件显示规则
// 根据当前组件的值决定显示哪些其他组件
//
// 使用示例：
//
//	radio.Control([]ControlRule{
//	    {Value: "1", Rule: []Component{input1}},
//	    {Value: "2", Rule: []Component{input2}},
//	})
func (b *Builder[T]) Control(rules []ControlRule) T {
	b.data.Control = rules
	return b.inst
}

// AppendControl 追加条件显示规则
func (b *Builder[T]) AppendControl(rule ControlRule) T {
	b.data.Control = append(b.data.Control, rule)
	return b.inst
}

// Validate 添加验证规则
//
// 使用示例：
//
//	input.Validate(PatternRule{Pattern: "^\\d+$", Message: "只能输入数字"})
func (b *Builder[T]) Validate(rules ...ValidateRule) T {
	b.data.Validate = append(b.data.Validate, rules...)
	return b.inst
}

// Children 设置子组件
func (b *Builder[T]) Children(children []Component) T {
	b.data.Children = children
	return b.inst
}

// AppendChild 追加子组件
func (b *Builder[T]) AppendChild(child Component) T {
	b.data.Children = append(b.data.Children, child)
	return b.inst
}

// Emit 设置事件配置
func (b *Builder[T]) Emit(event string, handler interface{}) T {
	if b.data.Emit == nil {
		b.data.Emit = make(map[string]interface{})
	}
	b.data.Emit[event] = handler
	return b.inst
}

// AppendRule 添加自定义规则字段
// 允许添加标准API之外的自定义配置，如 suffix, prefix 等
// 对应PHP的appendRule()方法
//
// 使用示例：
//
//	// 添加后缀说明
//	radio.AppendRule("suffix", map[string]interface{}{
//	    "type": "div",
//	    "style": map[string]interface{}{
//	        "color": "#999999",
//	    },
//	    "domProps": map[string]interface{}{
//	        "innerHTML": "试用期每个用户只能购买一次",
//	    },
//	})
//
//	// 添加前缀图标
//	input.AppendRule("prefix", "¥")
//
// 注意：AppendRule 中的字段会在序列化时覆盖同名的标准字段
func (b *Builder[T]) AppendRule(name string, value interface{}) T {
	if b.data.AppendRule == nil {
		b.data.AppendRule = make(map[string]interface{})
	}
	b.data.AppendRule[name] = value
	return b.inst
}

// Col 设置组件布局规则
// 用于设置组件在表单中的栅格布局
// 对应PHP的col()方法
//
// 使用示例：
//
//	// 设置span为12（占50%宽度）
//	input.Col(12)
//
//	// 自定义布局配置
//	input.Col(map[string]interface{}{
//	    "span": 12,
//	    "offset": 2,
//	})
func (b *Builder[T]) Col(col interface{}) T {
	// 如果传入的是整数，转换为 map[string]interface{}{"span": col}
	switch v := col.(type) {
	case int:
		if b.data.AppendRule == nil {
			b.data.AppendRule = make(map[string]interface{})
		}
		b.data.AppendRule["col"] = map[string]interface{}{"span": v}
	case map[string]interface{}:
		if b.data.AppendRule == nil {
			b.data.AppendRule = make(map[string]interface{})
		}
		b.data.AppendRule["col"] = v
	default:
		if b.data.AppendRule == nil {
			b.data.AppendRule = make(map[string]interface{})
		}
		b.data.AppendRule["col"] = col
	}
	return b.inst
}

// buildComponent 将ComponentData转换为map[string]interface{}
// 这是JSON序列化的核心函数，对应PHP的build()方法
//
// 处理流程：
// 1. 基础字段（type, field, title, value）
// 2. Props属性
// 3. Validate验证规则
// 4. Control条件显示规则（递归处理）
// 5. Children子组件（递归处理）
// 6. Emit事件配置
// 7. AppendRule自定义字段（最后处理，允许覆盖标准字段）
//
// 返回的map将被json.Marshal()转换为JSON字符串
func buildComponent(data *ComponentData) map[string]interface{} {
	result := make(map[string]interface{})

	// 1. 基础字段
	if data.RuleType != "" {
		result["type"] = data.RuleType
	}
	if data.Field != "" {
		result["field"] = data.Field
	}
	if data.Title != "" {
		result["title"] = data.Title
	}
	if data.Value != nil {
		result["value"] = data.Value
	}

	// 2. Props属性
	if len(data.Props) > 0 {
		result["props"] = data.Props
	}

	// 3. Validate验证规则
	if len(data.Validate) > 0 {
		validates := make([]map[string]interface{}, len(data.Validate))
		for i, v := range data.Validate {
			validates[i] = v.ToMap()
		}
		result["validate"] = validates
	}

	// 4. Control条件显示规则（递归处理）
	if len(data.Control) > 0 {
		controls := make([]map[string]interface{}, len(data.Control))
		for i, ctrl := range data.Control {
			// 递归构建control中的组件
			rules := make([]map[string]interface{}, len(ctrl.Rule))
			for j, r := range ctrl.Rule {
				rules[j] = r.Build()
			}
			controls[i] = map[string]interface{}{
				"value": ctrl.Value,
				"rule":  rules,
			}
		}
		result["control"] = controls
	}

	// 5. Children子组件（递归处理）
	if len(data.Children) > 0 {
		children := make([]map[string]interface{}, len(data.Children))
		for i, child := range data.Children {
			children[i] = child.Build()
		}
		result["children"] = children
	}

	// 6. Emit事件配置
	if len(data.Emit) > 0 {
		result["emit"] = data.Emit
	}

	// 7. AppendRule自定义字段（最后处理，允许覆盖标准字段）
	// 对应PHP的 $this->appendRule + $this->getRule() 合并逻辑
	if len(data.AppendRule) > 0 {
		for k, v := range data.AppendRule {
			result[k] = v
		}
	}

	return result
}

// ValidateRule 验证规则接口
// 所有验证规则都必须实现ToMap方法，将规则转换为map用于JSON序列化
// 对应PHP的ValidateInterface
type ValidateRule interface {
	// ToMap 将验证规则转换为map
	ToMap() map[string]interface{}
}

// RequiredRule 必填验证规则
type RequiredRule struct {
	Message string // 验证失败提示信息
	Trigger string // 触发方式：blur, change 等
}

// ToMap 实现ValidateRule接口
func (r RequiredRule) ToMap() map[string]interface{} {
	rule := map[string]interface{}{
		"required": true,
	}
	if r.Message != "" {
		rule["message"] = r.Message
	}
	if r.Trigger != "" {
		rule["trigger"] = r.Trigger
	}
	return rule
}

// ControlRule 条件显示规则
// 当组件值等于Value时，显示Rule中的组件
// 对应PHP的ControlRule
//
// 使用示例：
//
//	ControlRule{
//	    Value: "1",
//	    Rule: []Component{
//	        NewInput("field1", "字段1"),
//	    },
//	}
type ControlRule struct {
	// Value 当组件值等于此值时触发
	Value interface{}

	// Rule 要显示的组件数组
	Rule []Component
}
