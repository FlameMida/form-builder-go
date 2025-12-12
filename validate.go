package formbuilder

// validate.go 包含所有验证规则的实现
// 对应PHP的ValidateInterface和各种验证规则类

// PatternRule 正则表达式验证规则
// 用于验证字符串是否匹配指定的正则表达式
//
// 使用示例：
//
//	input.Validate(PatternRule{
//	    Pattern: "^1[3-9]\\d{9}$",
//	    Message: "请输入正确的手机号",
//	})
type PatternRule struct {
	Pattern string // 正则表达式
	Message string // 验证失败提示信息
	Trigger string // 触发方式：blur, change
}

// ToMap 实现ValidateRule接口
func (r PatternRule) ToMap() map[string]interface{} {
	rule := map[string]interface{}{
		"pattern": r.Pattern,
	}
	if r.Message != "" {
		rule["message"] = r.Message
	}
	if r.Trigger != "" {
		rule["trigger"] = r.Trigger
	}
	return rule
}

// LengthRule 长度验证规则
// 用于验证字符串长度或数组元素个数
//
// 使用示例：
//
//	input.Validate(LengthRule{
//	    Min: 6,
//	    Max: 20,
//	    Message: "密码长度必须在6-20个字符之间",
//	})
type LengthRule struct {
	Min     int    // 最小长度
	Max     int    // 最大长度
	Message string // 验证失败提示信息
	Trigger string // 触发方式：blur, change
}

// ToMap 实现ValidateRule接口
func (r LengthRule) ToMap() map[string]interface{} {
	rule := make(map[string]interface{})
	if r.Min > 0 {
		rule["min"] = r.Min
	}
	if r.Max > 0 {
		rule["max"] = r.Max
	}
	if r.Message != "" {
		rule["message"] = r.Message
	}
	if r.Trigger != "" {
		rule["trigger"] = r.Trigger
	}
	return rule
}

// RangeRule 数值范围验证规则
// 用于验证数字是否在指定范围内
//
// 使用示例：
//
//	inputNumber.Validate(RangeRule{
//	    Min: 0,
//	    Max: 100,
//	    Message: "请输入0-100之间的数字",
//	})
type RangeRule struct {
	Min     float64 // 最小值
	Max     float64 // 最大值
	Message string  // 验证失败提示信息
	Trigger string  // 触发方式：blur, change
}

// ToMap 实现ValidateRule接口
func (r RangeRule) ToMap() map[string]interface{} {
	rule := make(map[string]interface{})
	rule["type"] = "number"
	if r.Min != 0 {
		rule["min"] = r.Min
	}
	if r.Max != 0 {
		rule["max"] = r.Max
	}
	if r.Message != "" {
		rule["message"] = r.Message
	}
	if r.Trigger != "" {
		rule["trigger"] = r.Trigger
	}
	return rule
}

// EmailRule 邮箱验证规则
// 使用内置的email类型验证
//
// 使用示例：
//
//	input.Validate(EmailRule{Message: "请输入正确的邮箱地址"})
type EmailRule struct {
	Message string // 验证失败提示信息
	Trigger string // 触发方式：blur, change
}

// ToMap 实现ValidateRule接口
func (r EmailRule) ToMap() map[string]interface{} {
	rule := map[string]interface{}{
		"type": "email",
	}
	if r.Message != "" {
		rule["message"] = r.Message
	}
	if r.Trigger != "" {
		rule["trigger"] = r.Trigger
	}
	return rule
}

// URLRule URL验证规则
// 使用内置的url类型验证
//
// 使用示例：
//
//	input.Validate(URLRule{Message: "请输入正确的URL地址"})
type URLRule struct {
	Message string // 验证失败提示信息
	Trigger string // 触发方式：blur, change
}

// ToMap 实现ValidateRule接口
func (r URLRule) ToMap() map[string]interface{} {
	rule := map[string]interface{}{
		"type": "url",
	}
	if r.Message != "" {
		rule["message"] = r.Message
	}
	if r.Trigger != "" {
		rule["trigger"] = r.Trigger
	}
	return rule
}

// DateRule 日期验证规则
// 验证是否为有效日期
//
// 使用示例：
//
//	datePicker.Validate(DateRule{Message: "请选择正确的日期"})
type DateRule struct {
	Message string // 验证失败提示信息
	Trigger string // 触发方式：blur, change
}

// ToMap 实现ValidateRule接口
func (r DateRule) ToMap() map[string]interface{} {
	rule := map[string]interface{}{
		"type": "date",
	}
	if r.Message != "" {
		rule["message"] = r.Message
	}
	if r.Trigger != "" {
		rule["trigger"] = r.Trigger
	}
	return rule
}

// EnumRule 枚举验证规则
// 验证值是否在允许的枚举列表中
//
// 使用示例：
//
//	select.Validate(EnumRule{
//	    Enum: []interface{}{"male", "female"},
//	    Message: "请选择正确的性别",
//	})
type EnumRule struct {
	Enum    []interface{} // 允许的值列表
	Message string        // 验证失败提示信息
	Trigger string        // 触发方式：blur, change
}

// ToMap 实现ValidateRule接口
func (r EnumRule) ToMap() map[string]interface{} {
	rule := map[string]interface{}{
		"enum": r.Enum,
	}
	if r.Message != "" {
		rule["message"] = r.Message
	}
	if r.Trigger != "" {
		rule["trigger"] = r.Trigger
	}
	return rule
}

// CustomRule 自定义验证规则
// 允许使用JavaScript函数字符串或完全自定义的规则对象
//
// 使用示例（JavaScript验证函数）：
//
//	input.Validate(CustomRule{
//	    Validator: "function(rule, value, callback) { ... }",
//	    Message: "验证失败",
//	})
//
// 使用示例（完全自定义规则对象）：
//
//	input.Validate(CustomRule{
//	    Rule: map[string]interface{}{
//	        "validator": "customValidatorName",
//	        "custom": true,
//	    },
//	})
type CustomRule struct {
	Validator string                 // JavaScript验证函数字符串
	Message   string                 // 验证失败提示信息
	Trigger   string                 // 触发方式：blur, change
	Rule      map[string]interface{} // 完全自定义的规则对象（优先级高于Validator）
}

// ToMap 实现ValidateRule接口
func (r CustomRule) ToMap() map[string]interface{} {
	// 如果提供了完全自定义的规则对象，直接返回
	if r.Rule != nil {
		return r.Rule
	}

	// 否则使用Validator字符串
	rule := make(map[string]interface{})
	if r.Validator != "" {
		rule["validator"] = r.Validator
	}
	if r.Message != "" {
		rule["message"] = r.Message
	}
	if r.Trigger != "" {
		rule["trigger"] = r.Trigger
	}
	return rule
}

// WhitespaceRule 空白字符验证规则
// 验证字符串是否只包含空白字符（默认不允许）
//
// 使用示例：
//
//	input.Validate(WhitespaceRule{
//	    Whitespace: false,  // 不允许只有空白字符
//	    Message: "不能只输入空格",
//	})
type WhitespaceRule struct {
	Whitespace bool   // true=允许空白字符，false=不允许
	Message    string // 验证失败提示信息
	Trigger    string // 触发方式：blur, change
}

// ToMap 实现ValidateRule接口
func (r WhitespaceRule) ToMap() map[string]interface{} {
	rule := map[string]interface{}{
		"whitespace": r.Whitespace,
	}
	if r.Message != "" {
		rule["message"] = r.Message
	}
	if r.Trigger != "" {
		rule["trigger"] = r.Trigger
	}
	return rule
}

// 便捷验证规则构造函数

// NewRequired 创建必填验证规则
func NewRequired(message ...string) RequiredRule {
	rule := RequiredRule{Message: "此项必填"}
	if len(message) > 0 {
		rule.Message = message[0]
	}
	return rule
}

// NewPattern 创建正则验证规则
func NewPattern(pattern, message string) PatternRule {
	return PatternRule{
		Pattern: pattern,
		Message: message,
	}
}

// NewLength 创建长度验证规则
func NewLength(min, max int, message string) LengthRule {
	return LengthRule{
		Min:     min,
		Max:     max,
		Message: message,
	}
}

// NewMin 创建最小长度验证规则
func NewMin(min int, message string) LengthRule {
	return LengthRule{
		Min:     min,
		Message: message,
	}
}

// NewMax 创建最大长度验证规则
func NewMax(max int, message string) LengthRule {
	return LengthRule{
		Max:     max,
		Message: message,
	}
}

// NewRange 创建数值范围验证规则
func NewRange(min, max float64, message string) RangeRule {
	return RangeRule{
		Min:     min,
		Max:     max,
		Message: message,
	}
}

// NewEmail 创建邮箱验证规则
func NewEmail(message ...string) EmailRule {
	rule := EmailRule{Message: "请输入正确的邮箱地址"}
	if len(message) > 0 {
		rule.Message = message[0]
	}
	return rule
}

// NewURL 创建URL验证规则
func NewURL(message ...string) URLRule {
	rule := URLRule{Message: "请输入正确的URL地址"}
	if len(message) > 0 {
		rule.Message = message[0]
	}
	return rule
}

// NewEnum 创建枚举验证规则
func NewEnum(enum []interface{}, message string) EnumRule {
	return EnumRule{
		Enum:    enum,
		Message: message,
	}
}
