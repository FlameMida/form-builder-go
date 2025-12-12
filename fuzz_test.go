package formbuilder

import (
	"encoding/json"
	"strings"
	"testing"
)

// fuzz_test.go 模糊测试
// 使用Go 1.18+的原生fuzzing功能测试代码健壮性

// FuzzInputCreation 模糊测试Input组件创建
func FuzzInputCreation(f *testing.F) {
	// 添加种子数据
	f.Add("username", "用户名", "default")
	f.Add("", "", "")
	f.Add("field with space", "Title with 中文", "value")
	f.Add("<script>alert('xss')</script>", "SQL'; DROP TABLE users--", "';--")

	f.Fuzz(func(t *testing.T, field string, title string, value string) {
		// 测试创建不会panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Panic with field=%q, title=%q, value=%q: %v", field, title, value, r)
			}
		}()

		input := NewInput(field, title).Value(value)

		// 验证基本属性
		if input == nil {
			t.Error("NewInput returned nil")
			return
		}

		if input.GetField() != field {
			t.Errorf("Field mismatch: got %q, want %q", input.GetField(), field)
		}

		// 验证Build不会panic
		result := input.Build()
		if result == nil {
			t.Error("Build returned nil")
			return
		}

		// 验证JSON序列化
		_, err := json.Marshal(result)
		if err != nil {
			t.Errorf("JSON marshal failed: %v", err)
		}
	})
}

// FuzzSelectOptions 模糊测试Select组件选项
func FuzzSelectOptions(f *testing.F) {
	// 种子数据
	f.Add("role", "角色", "admin", "管理员")
	f.Add("", "", "", "")
	f.Add("<tag>", "'; DROP--", "0", "zero")

	f.Fuzz(func(t *testing.T, field string, title string, value string, label string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Panic: %v", r)
			}
		}()

		sel := NewSelect(field, title).SetOptions([]Option{
			{Value: value, Label: label},
		})

		if sel == nil {
			t.Error("NewSelect returned nil")
			return
		}

		result := sel.Build()
		if result == nil {
			t.Error("Build returned nil")
			return
		}

		// 验证JSON序列化
		jsonBytes, err := json.Marshal(result)
		if err != nil {
			t.Errorf("JSON marshal failed: %v", err)
		}

		// 验证JSON可以反序列化
		var decoded map[string]interface{}
		err = json.Unmarshal(jsonBytes, &decoded)
		if err != nil {
			t.Errorf("JSON unmarshal failed: %v", err)
		}
	})
}

// FuzzFormCreation 模糊测试表单创建
func FuzzFormCreation(f *testing.F) {
	// 种子数据
	f.Add("/submit", "field1", "title1")
	f.Add("", "", "")
	f.Add("/api/user/<script>", "'; DROP TABLE", "XSS<>")
	f.Add("https://example.com/submit", "normal_field", "Normal Title")

	f.Fuzz(func(t *testing.T, action string, field string, title string) {
		defer func() {
			if r := recover(); r != nil {
				// 字段重复是预期的panic，跳过
				if errMsg, ok := r.(error); ok {
					if strings.Contains(errMsg.Error(), "not unique") {
						return
					}
				}
				t.Errorf("Unexpected panic with action=%q, field=%q, title=%q: %v",
					action, field, title, r)
			}
		}()

		form := NewElmForm(action, []Component{
			NewInput(field, title),
		}, nil)

		if form == nil {
			t.Error("NewElmForm returned nil")
			return
		}

		// 验证FormRule不panic
		rules := form.FormRule()
		if rules == nil {
			t.Error("FormRule returned nil")
			return
		}

		// 验证JSON序列化
		_, err := form.ParseFormRule()
		if err != nil {
			t.Errorf("ParseFormRule failed: %v", err)
		}
	})
}

// FuzzValidationRules 模糊测试验证规则
func FuzzValidationRules(f *testing.F) {
	// 种子数据
	f.Add("必填", 6, 20)
	f.Add("", 0, 0)
	f.Add("Long message with special chars: <>\"'&", -100, 1000)

	f.Fuzz(func(t *testing.T, message string, min int, max int) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Panic: %v", r)
			}
		}()

		// 测试LengthRule
		rule := LengthRule{
			Min:     min,
			Max:     max,
			Message: message,
		}

		ruleMap := rule.ToMap()
		if ruleMap == nil {
			t.Error("ToMap returned nil")
			return
		}

		// 验证JSON序列化
		_, err := json.Marshal(ruleMap)
		if err != nil {
			t.Errorf("JSON marshal failed: %v", err)
		}
	})
}

// FuzzPatternRule 模糊测试正则表达式规则
func FuzzPatternRule(f *testing.F) {
	// 种子数据
	f.Add("^[a-zA-Z0-9]+$", "只能字母数字")
	f.Add("", "")
	f.Add("(invalid[regex", "Invalid regex")
	f.Add(".*", "Any character")
	f.Add("^\\d+$", "Only digits")

	f.Fuzz(func(t *testing.T, pattern string, message string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Panic: %v", r)
			}
		}()

		rule := PatternRule{
			Pattern: pattern,
			Message: message,
		}

		ruleMap := rule.ToMap()
		if ruleMap == nil {
			t.Error("ToMap returned nil")
			return
		}

		// 验证JSON序列化
		jsonBytes, err := json.Marshal(ruleMap)
		if err != nil {
			t.Errorf("JSON marshal failed: %v", err)
			return
		}

		// 验证反序列化
		var decoded map[string]interface{}
		err = json.Unmarshal(jsonBytes, &decoded)
		if err != nil {
			t.Errorf("JSON unmarshal failed: %v", err)
		}
	})
}

// FuzzFormData 模糊测试表单数据应用
func FuzzFormData(f *testing.F) {
	// 种子数据
	f.Add("field1", "value1")
	f.Add("", "")
	f.Add("<script>", "'; DROP TABLE users--")
	f.Add("normal", "12345")

	f.Fuzz(func(t *testing.T, field string, value string) {
		defer func() {
			if r := recover(); r != nil {
				// 字段重复是预期的，跳过
				if errMsg, ok := r.(error); ok {
					if strings.Contains(errMsg.Error(), "not unique") {
						return
					}
				}
				t.Errorf("Panic: %v", r)
			}
		}()

		form := NewElmForm("/submit", []Component{
			NewInput(field, "测试"),
		}, nil)

		// 应用数据
		form.FormData(map[string]interface{}{
			field: value,
		})

		rules := form.FormRule()
		if len(rules) > 0 && field != "" {
			// 验证值被正确应用
			if rules[0]["value"] != value {
				t.Errorf("Value not applied correctly: got %v, want %v",
					rules[0]["value"], value)
			}
		}

		// 验证JSON序列化
		_, err := form.ParseFormRule()
		if err != nil {
			t.Errorf("ParseFormRule failed: %v", err)
		}
	})
}

// FuzzControlRules 模糊测试Control规则
func FuzzControlRules(f *testing.F) {
	// 种子数据
	f.Add("1", "field1", "title1")
	f.Add("", "", "")
	f.Add("true", "nested", "Nested Field")

	f.Fuzz(func(t *testing.T, controlValue string, field string, title string) {
		defer func() {
			if r := recover(); r != nil {
				// 字段重复是预期的
				if errMsg, ok := r.(error); ok {
					if strings.Contains(errMsg.Error(), "not unique") {
						return
					}
				}
				t.Errorf("Panic: %v", r)
			}
		}()

		radio := NewRadio("type", "类型", "1").
			SetOptions([]Option{
				{Value: "1", Label: "选项1"},
				{Value: "2", Label: "选项2"},
			}).
			Control([]ControlRule{
				{
					Value: controlValue,
					Rule: []Component{
						NewInput(field, title),
					},
				},
			})

		result := radio.Build()
		if result == nil {
			t.Error("Build returned nil")
			return
		}

		// 验证JSON序列化
		jsonBytes, err := json.Marshal(result)
		if err != nil {
			t.Errorf("JSON marshal failed: %v", err)
			return
		}

		// 验证反序列化
		var decoded map[string]interface{}
		err = json.Unmarshal(jsonBytes, &decoded)
		if err != nil {
			t.Errorf("JSON unmarshal failed: %v", err)
		}
	})
}

// FuzzNumberRange 模糊测试数值范围
func FuzzNumberRange(f *testing.F) {
	// 种子数据
	f.Add(0.0, 100.0)
	f.Add(-1000.0, 1000.0)
	f.Add(3.14, 2.71)

	f.Fuzz(func(t *testing.T, min float64, max float64) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Panic with min=%f, max=%f: %v", min, max, r)
			}
		}()

		rule := RangeRule{
			Min:     min,
			Max:     max,
			Message: "范围错误",
		}

		ruleMap := rule.ToMap()
		if ruleMap == nil {
			t.Error("ToMap returned nil")
			return
		}

		// 验证JSON序列化
		_, err := json.Marshal(ruleMap)
		if err != nil {
			t.Errorf("JSON marshal failed: %v", err)
		}
	})
}

// FuzzOptionToMap 模糊测试Option.ToMap
func FuzzOptionToMap(f *testing.F) {
	// 种子数据
	f.Add("value1", "label1", false)
	f.Add("", "", true)
	f.Add("<script>", "'; DROP--", false)

	f.Fuzz(func(t *testing.T, value string, label string, disabled bool) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Panic: %v", r)
			}
		}()

		option := Option{
			Value:    value,
			Label:    label,
			Disabled: disabled,
		}

		optionMap := option.ToMap()
		if optionMap == nil {
			t.Error("ToMap returned nil")
			return
		}

		// 验证JSON序列化
		jsonBytes, err := json.Marshal(optionMap)
		if err != nil {
			t.Errorf("JSON marshal failed: %v", err)
			return
		}

		// 验证反序列化
		var decoded map[string]interface{}
		err = json.Unmarshal(jsonBytes, &decoded)
		if err != nil {
			t.Errorf("JSON unmarshal failed: %v", err)
		}
	})
}
