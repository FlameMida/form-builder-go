package formbuilder

// input.go 实现Input输入框组件
// 对应PHP的 src/UI/Elm/Components/Input.php
//
// Input组件作为所有组件的模板，包含完整的注释说明
// 展示了如何使用泛型Builder[T]模式实现链式调用

// Input 输入框组件
// 对应Element UI和iView的Input组件
//
// 使用示例：
//
//	input := NewInput("username", "用户名").
//	    Placeholder("请输入用户名").
//	    Clearable(true).
//	    MaxLength(50).
//	    Required().
//	    Value("default")
type Input struct {
	Builder[*Input] // 嵌入泛型Builder，自引用类型为*Input
}

// NewInput 创建一个新的Input组件
// field: 字段名（表单提交的key）
// title: 字段标题/标签
// value: 默认值（可选）
func NewInput(field, title string, value ...interface{}) *Input {
	input := &Input{}

	// 初始化ComponentData
	input.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "input", // 组件类型
		Props:    make(map[string]interface{}),
	}

	// 设置默认值（如果提供）
	if len(value) > 0 {
		input.data.Value = value[0]
	}

	// 设置默认type为text
	input.data.Props["type"] = "text"

	// 关键步骤：设置inst为自己，使Builder的方法能返回*Input类型
	input.inst = input

	return input
}

// NewIviewInput 创建iView版本的Input组件
// 注意：type值与Element UI版本相同，框架差异由全局配置决定
func NewIviewInput(field, title string, value ...interface{}) *Input {
	input := NewInput(field, title, value...)
	// iView版本的type与Element UI相同
	input.data.RuleType = "input"
	return input
}

// 以下是Input组件的特有方法
// 每个方法都返回*Input类型，保证链式调用的类型安全

// Type 设置输入框类型
// 可选值：text, password, textarea, number, email, url, tel, search
func (i *Input) Type(inputType string) *Input {
	i.data.Props["type"] = inputType
	return i
}

// Placeholder 设置占位符文本
func (i *Input) Placeholder(text string) *Input {
	i.data.Props["placeholder"] = text
	return i
}

// Clearable 设置是否显示清空按钮
func (i *Input) Clearable(enable bool) *Input {
	i.data.Props["clearable"] = enable
	return i
}

// ShowPassword 设置是否显示密码切换按钮（仅在type="password"时有效）
func (i *Input) ShowPassword(enable bool) *Input {
	i.data.Props["show-password"] = enable
	return i
}

// Disabled 设置是否禁用
func (i *Input) Disabled(disabled bool) *Input {
	i.data.Props["disabled"] = disabled
	return i
}

// Readonly 设置是否只读
func (i *Input) Readonly(readonly bool) *Input {
	i.data.Props["readonly"] = readonly
	return i
}

// MaxLength 设置最大输入长度
func (i *Input) MaxLength(length int) *Input {
	i.data.Props["maxlength"] = length
	return i
}

// MinLength 设置最小输入长度
func (i *Input) MinLength(length int) *Input {
	i.data.Props["minlength"] = length
	return i
}

// ShowWordLimit 设置是否显示字数统计（需要设置maxlength）
func (i *Input) ShowWordLimit(show bool) *Input {
	i.data.Props["show-word-limit"] = show
	return i
}

// PrefixIcon 设置输入框头部图标
func (i *Input) PrefixIcon(icon string) *Input {
	i.data.Props["prefix-icon"] = icon
	return i
}

// SuffixIcon 设置输入框尾部图标
func (i *Input) SuffixIcon(icon string) *Input {
	i.data.Props["suffix-icon"] = icon
	return i
}

// Size 设置输入框尺寸
// 可选值：large, default, small, mini
func (i *Input) Size(size string) *Input {
	i.data.Props["size"] = size
	return i
}

// Autocomplete 设置原生autocomplete属性
func (i *Input) Autocomplete(value string) *Input {
	i.data.Props["autocomplete"] = value
	return i
}

// Autofocus 设置是否自动获取焦点
func (i *Input) Autofocus(enable bool) *Input {
	i.data.Props["autofocus"] = enable
	return i
}

// Rows 设置textarea的行数（仅在type="textarea"时有效）
func (i *Input) Rows(rows int) *Input {
	i.data.Props["rows"] = rows
	return i
}

// Autosize 设置textarea自适应高度（仅在type="textarea"时有效）
// 可以传入对象 {minRows: 2, maxRows: 6}
func (i *Input) Autosize(autosize interface{}) *Input {
	i.data.Props["autosize"] = autosize
	return i
}

// ValidateEvent 设置触发表单验证的时机
// 默认为true，设为false后只在调用表单validateField方法时验证
func (i *Input) ValidateEvent(enable bool) *Input {
	i.data.Props["validate-event"] = enable
	return i
}

// 实现Component接口

// GetField 返回字段名
func (i *Input) GetField() string {
	return i.data.Field
}

// GetType 返回组件类型
func (i *Input) GetType() string {
	return i.data.RuleType
}

// Build 将组件转换为map，用于JSON序列化
func (i *Input) Build() map[string]interface{} {
	return buildComponent(i.data)
}

// 便捷构造函数

// Password 创建密码输入框
func Password(field, title string, value ...interface{}) *Input {
	input := NewInput(field, title, value...)
	input.Type("password")
	return input
}

// Email 创建邮箱输入框（自动添加邮箱验证）
func Email(field, title string, value ...interface{}) *Input {
	input := NewInput(field, title, value...)
	input.Type("email")
	input.Validate(NewEmail())
	return input
}

// URL 创建URL输入框（自动添加URL验证）
func URL(field, title string, value ...interface{}) *Input {
	input := NewInput(field, title, value...)
	input.Type("url")
	input.Validate(NewURL())
	return input
}

// Tel 创建电话输入框
func Tel(field, title string, value ...interface{}) *Input {
	input := NewInput(field, title, value...)
	input.Type("tel")
	return input
}

// Search 创建搜索输入框
func Search(field, title string, value ...interface{}) *Input {
	input := NewInput(field, title, value...)
	input.Type("search")
	return input
}
