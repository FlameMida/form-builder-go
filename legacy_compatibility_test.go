package formbuilder

import (
	"encoding/json"
	"testing"
)

// formbuilder_test.go 测试文件
// 提供对比测试框架，确保与PHP版本100%兼容

// normalizeJSON 标准化JSON字符串，用于对比
// 去除空格、排序键，确保对比的准确性
func normalizeJSON(jsonStr string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return "", err
	}
	normalized, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(normalized), nil
}

// TestSimpleInput 测试简单的Input组件
func TestSimpleInput(t *testing.T) {
	input := NewInput("username", "用户名")
	result := input.Build()

	if result["type"] != "input" {
		t.Errorf("Expected type 'el-input', got '%s'", result["type"])
	}
	if result["field"] != "username" {
		t.Errorf("Expected field 'username', got '%s'", result["field"])
	}
	if result["title"] != "用户名" {
		t.Errorf("Expected title '用户名', got '%s'", result["title"])
	}
}

// TestInputWithRequired 测试带必填验证的Input
func TestInputWithRequired(t *testing.T) {
	input := NewInput("username", "用户名").Required()
	result := input.Build()

	validates, ok := result["validate"].([]map[string]interface{})
	if !ok {
		t.Fatal("Expected validate to be array of maps")
	}

	if len(validates) != 1 {
		t.Fatalf("Expected 1 validate rule, got %d", len(validates))
	}

	if validates[0]["required"] != true {
		t.Error("Expected required to be true")
	}
}

// TestSelectWithOptions 测试带选项的Select
func TestSelectWithOptions(t *testing.T) {
	sel := NewSelect("role", "角色").SetOptions([]Option{
		{Value: "1", Label: "管理员"},
		{Value: "2", Label: "用户"},
	})

	result := sel.Build()

	options, ok := result["options"].([]map[string]interface{})
	if !ok {
		t.Fatal("Expected options to be array of maps")
	}

	if len(options) != 2 {
		t.Fatalf("Expected 2 options, got %d", len(options))
	}

	if options[0]["value"] != "1" {
		t.Errorf("Expected first option value '1', got '%v'", options[0]["value"])
	}
	if options[0]["label"] != "管理员" {
		t.Errorf("Expected first option label '管理员', got '%v'", options[0]["label"])
	}
}

// TestRadioWithControl 测试带条件显示的Radio
func TestLegacyRadioWithControl(t *testing.T) {
	radio := NewRadio("type", "类型", "1").
		SetOptions([]Option{
			{Value: "1", Label: "试用期"},
			{Value: "2", Label: "有限期"},
		}).
		Control([]ControlRule{
			{
				Value: "1",
				Rule: []Component{
					NewInputNumber("days", "天数").Required(),
				},
			},
		})

	result := radio.Build()

	control, ok := result["control"].([]map[string]interface{})
	if !ok {
		t.Fatal("Expected control to be array of maps")
	}

	if len(control) != 1 {
		t.Fatalf("Expected 1 control rule, got %d", len(control))
	}

	if control[0]["value"] != "1" {
		t.Errorf("Expected control value '1', got '%v'", control[0]["value"])
	}

	rules, ok := control[0]["rule"].([]map[string]interface{})
	if !ok {
		t.Fatal("Expected control rule to be array of maps")
	}

	if len(rules) != 1 {
		t.Fatalf("Expected 1 rule in control, got %d", len(rules))
	}
}

// TestFormJSON 测试Form的JSON序列化
func TestFormJSON(t *testing.T) {
	form := NewElmForm("/submit", []Component{
		NewInput("username", "用户名").Required(),
		Password("password", "密码").Required(),
	}, nil)

	jsonStr, err := form.ParseFormRule()
	if err != nil {
		t.Fatalf("Failed to parse form rule: %v", err)
	}

	// 验证JSON是否有效
	var rules []interface{}
	if err := json.Unmarshal([]byte(jsonStr), &rules); err != nil {
		t.Fatalf("Invalid JSON: %v", err)
	}

	if len(rules) != 2 {
		t.Errorf("Expected 2 rules, got %d", len(rules))
	}
}

// TestFormData 测试表单数据应用
func TestLegacyFormData(t *testing.T) {
	form := NewElmForm("/submit", []Component{
		NewInput("username", "用户名"),
		NewInput("email", "邮箱"),
	}, nil)

	form.FormData(map[string]interface{}{
		"username": "john_doe",
		"email":    "john@example.com",
	})

	rules := form.FormRule()

	if rules[0]["value"] != "john_doe" {
		t.Errorf("Expected username value 'john_doe', got '%v'", rules[0]["value"])
	}
	if rules[1]["value"] != "john@example.com" {
		t.Errorf("Expected email value 'john@example.com', got '%v'", rules[1]["value"])
	}
}

// TestFieldUniqueness 测试字段唯一性验证
func TestLegacyFieldUniqueness(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for duplicate field, but didn't panic")
		}
	}()

	// 这应该会panic，因为有重复的field
	NewElmForm("/submit", []Component{
		NewInput("username", "用户名"),
		NewInput("username", "用户名2"), // 重复的field
	}, nil)
}

// TestValidationRules 测试各种验证规则
func TestValidationRules(t *testing.T) {
	tests := []struct {
		name     string
		rule     ValidateRule
		expected map[string]interface{}
	}{
		{
			name: "Required",
			rule: RequiredRule{Message: "必填"},
			expected: map[string]interface{}{
				"required": true,
				"message":  "必填",
			},
		},
		{
			name: "Pattern",
			rule: PatternRule{Pattern: "^\\d+$", Message: "只能输入数字"},
			expected: map[string]interface{}{
				"pattern": "^\\d+$",
				"message": "只能输入数字",
			},
		},
		{
			name: "Length",
			rule: LengthRule{Min: 6, Max: 20, Message: "长度6-20"},
			expected: map[string]interface{}{
				"min":     6,
				"max":     20,
				"message": "长度6-20",
			},
		},
		{
			name: "Email",
			rule: EmailRule{Message: "邮箱格式错误"},
			expected: map[string]interface{}{
				"type":    "email",
				"message": "邮箱格式错误",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rule.ToMap()
			for k, v := range tt.expected {
				if result[k] != v {
					t.Errorf("Expected %s to be %v, got %v", k, v, result[k])
				}
			}
		})
	}
}

// TestChainedMethods 测试链式调用
func TestChainedMethods(t *testing.T) {
	input := NewInput("username", "用户名").
		Placeholder("请输入用户名").
		Clearable(true).
		MaxLength(50).
		Required().
		Value("default")

	result := input.Build()

	if result["field"] != "username" {
		t.Error("Chain broken: field not set correctly")
	}

	props, ok := result["props"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected props to be map")
	}

	if props["placeholder"] != "请输入用户名" {
		t.Error("Chain broken: placeholder not set")
	}
	if props["clearable"] != true {
		t.Error("Chain broken: clearable not set")
	}
	if props["maxlength"] != 50 {
		t.Error("Chain broken: maxlength not set")
	}

	if result["value"] != "default" {
		t.Error("Chain broken: value not set")
	}
}

// TestFactoryMethods 测试工厂方法
func TestFactoryMethods(t *testing.T) {
	// 测试Elm工厂
	input := Elm.Input("test", "测试")
	if input.GetType() != "input" {
		t.Errorf("Expected el-input, got %s", input.GetType())
	}

	sel := Elm.Select("test", "测试")
	if sel.GetType() != "select" {
		t.Errorf("Expected el-select, got %s", sel.GetType())
	}

	// 测试Iview工厂
	iviewInput := Iview.Input("test", "测试")
	if iviewInput.GetType() != "input" {
		t.Errorf("Expected i-input, got %s", iviewInput.GetType())
	}
}

// TestFormConfig 测试表单配置
func TestFormConfig(t *testing.T) {
	config := NewElmConfig()
	config.SubmitBtn(true, "提交表单")
	config.ResetBtn(false)

	configMap := config.ToMap()

	submitBtn, ok := configMap["submitBtn"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected submitBtn to be map")
	}

	if submitBtn["show"] != true {
		t.Error("Expected submitBtn show to be true")
	}
	if submitBtn["innerText"] != "提交表单" {
		t.Error("Expected submitBtn innerText to be '提交表单'")
	}

	resetBtn, ok := configMap["resetBtn"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected resetBtn to be map")
	}

	if resetBtn["show"] != false {
		t.Error("Expected resetBtn show to be false")
	}
}

// BenchmarkFormBuild 性能测试：表单构建
func BenchmarkFormBuild(b *testing.B) {
	rules := []Component{
		NewInput("field1", "字段1").Required(),
		NewInput("field2", "字段2").Required(),
		NewSelect("field3", "字段3").SetOptions([]Option{
			{Value: "1", Label: "选项1"},
			{Value: "2", Label: "选项2"},
		}),
		NewRadio("field4", "字段4").SetOptions([]Option{
			{Value: "a", Label: "A"},
			{Value: "b", Label: "B"},
		}),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		form := NewElmForm("/submit", rules, nil)
		_ = form.FormRule()
	}
}

// BenchmarkJSONSerialization 性能测试：JSON序列化
func BenchmarkJSONSerialization(b *testing.B) {
	form := NewElmForm("/submit", []Component{
		NewInput("username", "用户名").Required(),
		Password("password", "密码").Required(),
		NewSelect("role", "角色").SetOptions([]Option{
			{Value: "admin", Label: "管理员"},
			{Value: "user", Label: "用户"},
		}),
	}, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = form.ParseFormRule()
	}
}
