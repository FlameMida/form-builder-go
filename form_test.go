package formbuilder

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// form_test.go 测试Form核心类
// 重构并增强formbuilder_test.go中的表单测试

// TestFormCreation 测试表单创建
func TestFormCreation(t *testing.T) {
	t.Run("NewElmForm", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("test", "测试"),
		}, nil)

		assert.NotNil(t, form)
		assert.Equal(t, "/submit", form.action)
		assert.Equal(t, "POST", form.method)
		assert.Len(t, form.rules, 1)
		assert.NotNil(t, form.config)
		assert.NotNil(t, form.ui)
	})

	t.Run("NewElmFormWithConfig", func(t *testing.T) {
		config := NewElmConfig()
		config.SubmitBtn(true, "提交")

		form := NewElmForm("/submit", []Component{}, config)

		assert.NotNil(t, form)
		assert.Equal(t, config, form.config)
	})

	t.Run("NewIviewForm", func(t *testing.T) {
		form := NewIviewForm("/submit", []Component{}, nil)
		assert.NotNil(t, form)
		assert.NotNil(t, form.ui)
	})

	t.Run("NewIview4Form", func(t *testing.T) {
		form := NewIview4Form("/submit", []Component{}, nil)
		assert.NotNil(t, form)
	})

	t.Run("EmptyRules", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{}, nil)
		assert.Len(t, form.rules, 0)
	})
}

// TestFormFieldUniqueness 测试字段唯一性验证
func TestFormFieldUniqueness(t *testing.T) {
	t.Run("UniqueFields", func(t *testing.T) {
		// 不应该panic
		form := NewElmForm("/submit", []Component{
			NewInput("field1", "字段1"),
			NewInput("field2", "字段2"),
			NewSelect("field3", "字段3"),
		}, nil)

		assert.NotNil(t, form)
	})

	t.Run("DuplicateFields", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic for duplicate field, but didn't panic")
			}
		}()

		// 这应该会panic
		NewElmForm("/submit", []Component{
			NewInput("username", "用户名"),
			NewInput("username", "用户名2"), // 重复
		}, nil)
	})

	t.Run("DuplicateFieldsInControl", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic for duplicate field in control, but didn't panic")
			}
		}()

		// Control中的字段重复也应该panic
		NewElmForm("/submit", []Component{
			NewInput("username", "用户名"),
			NewRadio("type", "类型", "1").Control([]ControlRule{
				{
					Value: "1",
					Rule: []Component{
						NewInput("username", "用户名2"), // 与外层重复
					},
				},
			}),
		}, nil)
	})

	t.Run("EmptyFieldAllowed", func(t *testing.T) {
		// 空字段名应该被跳过，不检查唯一性
		form := NewElmForm("/submit", []Component{
			NewInput("", "无字段1"),
			NewInput("", "无字段2"),
		}, nil)

		assert.NotNil(t, form)
	})
}

// TestFormSetRule 测试SetRule方法
func TestFormSetRule(t *testing.T) {
	t.Run("SetNewRules", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("old", "旧字段"),
		}, nil)

		form.SetRule([]Component{
			NewInput("new1", "新字段1"),
			NewInput("new2", "新字段2"),
		})

		assert.Len(t, form.rules, 2)
		assert.Equal(t, "new1", form.rules[0].GetField())
	})

	t.Run("SetRulesPanicOnDuplicate", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic for duplicate field")
			}
		}()

		form := NewElmForm("/submit", []Component{
			NewInput("test", "测试"),
		}, nil)

		form.SetRule([]Component{
			NewInput("dup", "重复"),
			NewInput("dup", "重复2"),
		})
	})
}

// TestFormAppend 测试Append方法
func TestFormAppend(t *testing.T) {
	t.Run("AppendComponent", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("field1", "字段1"),
		}, nil)

		form.Append(NewInput("field2", "字段2")).
			Append(NewInput("field3", "字段3"))

		assert.Len(t, form.rules, 3)
		assert.Equal(t, "field3", form.rules[2].GetField())
	})

	t.Run("AppendPanicOnDuplicate", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic")
			}
		}()

		form := NewElmForm("/submit", []Component{
			NewInput("test", "测试"),
		}, nil)

		form.Append(NewInput("test", "重复"))
	})
}

// TestFormPrepend 测试Prepend方法
func TestFormPrepend(t *testing.T) {
	t.Run("PrependComponent", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("field2", "字段2"),
		}, nil)

		form.Prepend(NewInput("field1", "字段1"))

		assert.Len(t, form.rules, 2)
		assert.Equal(t, "field1", form.rules[0].GetField())
		assert.Equal(t, "field2", form.rules[1].GetField())
	})
}

// TestFormSetters 测试各种setter方法
func TestFormSetters(t *testing.T) {
	t.Run("SetAction", func(t *testing.T) {
		form := NewElmForm("/old", []Component{}, nil)
		form.SetAction("/new")
		assert.Equal(t, "/new", form.action)
	})

	t.Run("SetMethod", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{}, nil)
		form.SetMethod("GET")
		assert.Equal(t, "GET", form.method)
	})

	t.Run("SetTitle", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{}, nil)
		form.SetTitle("用户注册")
		assert.Equal(t, "用户注册", form.title)
	})

	t.Run("SetConfig", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{}, nil)
		newConfig := NewElmConfig()
		form.SetConfig(newConfig)
		assert.Equal(t, newConfig, form.config)
	})
}

// TestFormData 测试表单数据应用
func TestFormData(t *testing.T) {
	t.Run("SetFormData", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("username", "用户名"),
			NewInput("email", "邮箱"),
		}, nil)

		form.FormData(map[string]interface{}{
			"username": "john_doe",
			"email":    "john@example.com",
		})

		rules := form.FormRule()
		assert.Equal(t, "john_doe", rules[0]["value"])
		assert.Equal(t, "john@example.com", rules[1]["value"])
	})

	t.Run("PartialFormData", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("field1", "字段1"),
			NewInput("field2", "字段2"),
			NewInput("field3", "字段3"),
		}, nil)

		form.FormData(map[string]interface{}{
			"field1": "value1",
			// field2 没有值
			"field3": "value3",
		})

		rules := form.FormRule()
		assert.Equal(t, "value1", rules[0]["value"])
		assert.Nil(t, rules[1]["value"]) // field2应该是nil
		assert.Equal(t, "value3", rules[2]["value"])
	})

	t.Run("SetValue", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("test", "测试"),
		}, nil)

		form.SetValue("test", "value")

		rules := form.FormRule()
		assert.Equal(t, "value", rules[0]["value"])
	})

	t.Run("SetValueChaining", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("field1", "字段1"),
			NewInput("field2", "字段2"),
		}, nil)

		form.SetValue("field1", "value1").
			SetValue("field2", "value2")

		assert.Equal(t, "value1", form.formData["field1"])
		assert.Equal(t, "value2", form.formData["field2"])
	})
}

// TestFormRule 测试FormRule方法
func TestFormRule(t *testing.T) {
	t.Run("BasicFormRule", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("username", "用户名").Required(),
			NewSelect("role", "角色").SetOptions([]Option{
				{Value: "admin", Label: "管理员"},
			}),
		}, nil)

		rules := form.FormRule()

		require.Len(t, rules, 2)
		assert.Equal(t, "input", rules[0]["type"])
		assert.Equal(t, "username", rules[0]["field"])
		assert.Equal(t, "select", rules[1]["type"])
	})

	t.Run("FormRuleWithData", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("test", "测试"),
		}, nil)

		form.FormData(map[string]interface{}{
			"test": "预填值",
		})

		rules := form.FormRule()
		assert.Equal(t, "预填值", rules[0]["value"])
	})

	t.Run("EmptyFormRule", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{}, nil)
		rules := form.FormRule()
		assert.Len(t, rules, 0)
	})
}

// TestParseFormRule 测试JSON序列化
func TestParseFormRule(t *testing.T) {
	t.Run("BasicJSONParsing", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("username", "用户名").Required(),
			Password("password", "密码").Required(),
		}, nil)

		jsonStr, err := form.ParseFormRule()
		require.NoError(t, err)

		// 验证JSON是否有效
		var rules []interface{}
		err = json.Unmarshal([]byte(jsonStr), &rules)
		require.NoError(t, err)
		assert.Len(t, rules, 2)
	})

	t.Run("JSONStructure", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("test", "测试").
				Placeholder("请输入").
				Required(),
		}, nil)

		jsonStr, err := form.ParseFormRule()
		require.NoError(t, err)

		var rules []map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &rules)
		require.NoError(t, err)

		assert.Equal(t, "input", rules[0]["type"])
		assert.Equal(t, "test", rules[0]["field"])
		assert.Equal(t, "测试", rules[0]["title"])

		props := rules[0]["props"].(map[string]interface{})
		assert.Equal(t, "请输入", props["placeholder"])
	})

	t.Run("JSONWithComplexStructure", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewRadio("type", "类型", "1").
				SetOptions([]Option{
					{Value: "1", Label: "类型1"},
					{Value: "2", Label: "类型2"},
				}).
				Control([]ControlRule{
					{
						Value: "1",
						Rule: []Component{
							NewInput("extra", "额外字段"),
						},
					},
				}),
		}, nil)

		jsonStr, err := form.ParseFormRule()
		require.NoError(t, err)

		var rules []map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &rules)
		require.NoError(t, err)

		// 验证control结构
		control := rules[0]["control"].([]interface{})
		assert.Len(t, control, 1)
	})
}

// TestFormChaining 测试Form的链式调用
func TestFormChaining(t *testing.T) {
	t.Run("CompleteChaining", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("field1", "字段1"),
		}, nil).
			SetAction("/new-submit").
			SetMethod("POST").
			SetTitle("测试表单").
			Append(NewInput("field2", "字段2")).
			FormData(map[string]interface{}{
				"field1": "value1",
			})

		assert.Equal(t, "/new-submit", form.action)
		assert.Equal(t, "POST", form.method)
		assert.Equal(t, "测试表单", form.title)
		assert.Len(t, form.rules, 2)
		assert.Equal(t, "value1", form.formData["field1"])
	})
}

// TestFormWithMultipleComponents 测试包含多种组件的表单
func TestFormWithMultipleComponents(t *testing.T) {
	t.Run("ComplexForm", func(t *testing.T) {
		form := NewElmForm("/api/user/create", []Component{
			NewInput("username", "用户名").
				Placeholder("请输入用户名").
				Required(),

			Password("password", "密码").
				MinLength(6).
				Required(),

			Email("email", "邮箱").Required(),

			NewSelect("role", "角色").
				SetOptions([]Option{
					{Value: "admin", Label: "管理员"},
					{Value: "user", Label: "用户"},
				}).
				Required(),

			NewRadio("status", "状态", "1").
				SetOptions([]Option{
					{Value: "1", Label: "启用"},
					{Value: "0", Label: "禁用"},
				}),

			NewSwitch("is_active", "是否激活").
				ActiveText("是").
				InactiveText("否"),
		}, nil)

		rules := form.FormRule()
		assert.Len(t, rules, 6)

		// 验证每个组件的类型
		assert.Equal(t, "input", rules[0]["type"])
		assert.Equal(t, "input", rules[1]["type"])
		assert.Equal(t, "input", rules[2]["type"])
		assert.Equal(t, "select", rules[3]["type"])
		assert.Equal(t, "radio", rules[4]["type"])
		assert.Equal(t, "switch", rules[5]["type"])

		// 验证JSON序列化
		jsonStr, err := form.ParseFormRule()
		require.NoError(t, err)
		assert.True(t, strings.Contains(jsonStr, "username"))
		assert.True(t, strings.Contains(jsonStr, "password"))
	})
}

// TestFormGetUI 测试GetUI方法
func TestFormGetUI(t *testing.T) {
	t.Run("GetElmUI", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{}, nil)
		ui := form.GetUI()
		assert.NotNil(t, ui)

		scripts := ui.GetScripts()
		assert.NotEmpty(t, scripts)

		styles := ui.GetStyles()
		assert.NotEmpty(t, styles)
	})
}

// BenchmarkFormCreation 性能测试：表单创建
func BenchmarkFormCreation(b *testing.B) {
	rules := []Component{
		NewInput("field1", "字段1").Required(),
		NewInput("field2", "字段2").Required(),
		NewSelect("field3", "字段3").SetOptions([]Option{
			{Value: "1", Label: "选项1"},
			{Value: "2", Label: "选项2"},
		}),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewElmForm("/submit", rules, nil)
	}
}

// BenchmarkFormRule 性能测试：FormRule生成
func BenchmarkFormRule(b *testing.B) {
	form := createComplexTestForm()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = form.FormRule()
	}
}

// BenchmarkParseFormRule 性能测试：JSON序列化
func BenchmarkParseFormRule(b *testing.B) {
	form := createComplexTestForm()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = form.ParseFormRule()
	}
}
