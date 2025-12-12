package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// radio_test.go Radio组件完整测试

// TestRadioCreation 测试Radio组件创建
func TestRadioCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		radio := NewRadio("gender", "性别", "male")

		assert.Equal(t, "gender", radio.GetField())
		assert.Equal(t, "radio", radio.GetType())

		data := radio.GetData()
		assert.Equal(t, "gender", data.Field)
		assert.Equal(t, "性别", data.Title)
		assert.Equal(t, "male", data.Value)
		assert.NotNil(t, radio.options)
	})

	t.Run("CreationWithoutValue", func(t *testing.T) {
		radio := NewRadio("test", "测试")

		data := radio.GetData()
		assert.Nil(t, data.Value)
	})

	t.Run("IviewCreation", func(t *testing.T) {
		radio := NewIviewRadio("test", "测试")

		assert.Equal(t, "radio", radio.GetType())
	})
}

// TestRadioSetOptions 测试SetOptions方法
func TestRadioSetOptions(t *testing.T) {
	t.Run("SetBasicOptions", func(t *testing.T) {
		radio := NewRadio("gender", "性别").SetOptions([]Option{
			{Value: "male", Label: "男"},
			{Value: "female", Label: "女"},
		})

		assert.Len(t, radio.options, 2)
		assert.Equal(t, "male", radio.options[0].Value)
		assert.Equal(t, "男", radio.options[0].Label)
	})

	t.Run("SetOptionsWithDisabled", func(t *testing.T) {
		radio := NewRadio("test", "测试").SetOptions([]Option{
			{Value: "1", Label: "选项1"},
			{Value: "2", Label: "选项2", Disabled: true},
		})

		assert.Equal(t, true, radio.options[1].Disabled)
	})
}

// TestRadioAppendOption 测试AppendOption方法
func TestRadioAppendOption(t *testing.T) {
	t.Run("AppendOptions", func(t *testing.T) {
		radio := NewRadio("test", "测试").
			AppendOption(Option{Value: "1", Label: "选项1"}).
			AppendOption(Option{Value: "2", Label: "选项2"}).
			AppendOption(Option{Value: "3", Label: "选项3"})

		assert.Len(t, radio.options, 3)
	})
}

// TestRadioProperties 测试所有属性方法
func TestRadioProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Radio) *Radio
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "Disabled",
			setup:         func(r *Radio) *Radio { return r.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
		{
			name:          "Size",
			setup:         func(r *Radio) *Radio { return r.Size("large") },
			propKey:       "size",
			expectedValue: "large",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			radio := NewRadio("test", "测试")
			radio = tt.setup(radio)

			data := radio.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestRadioChaining 测试链式调用
func TestRadioChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		radio := NewRadio("status", "状态", "1").
			SetOptions([]Option{
				{Value: "1", Label: "启用"},
				{Value: "0", Label: "禁用"},
			}).
			Size("large").
			Button(true).
			Required()

		data := radio.GetData()
		assert.Equal(t, "1", data.Value)
		assert.Equal(t, "large", data.Props["size"])
		assert.Equal(t, "button", data.Props["type"]) // Button sets props["type"]
		assert.NotEmpty(t, data.Validate)
		assert.Len(t, radio.options, 2)
	})
}

// TestRadioBuild 测试Build方法
func TestRadioBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		radio := NewRadio("gender", "性别", "male").
			SetOptions([]Option{
				{Value: "male", Label: "男"},
				{Value: "female", Label: "女"},
			})

		result := radio.Build()

		assert.Equal(t, "radio", result["type"])
		assert.Equal(t, "gender", result["field"])
		assert.Equal(t, "性别", result["title"])
		assert.Equal(t, "male", result["value"])

		options, ok := result["options"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, options, 2)
	})

	t.Run("BuildWithButton", func(t *testing.T) {
		radio := NewRadio("test", "测试").Button(true)

		result := radio.Build()
		// Button mode sets props["type"] = "button"
		assert.Equal(t, "radio", result["type"])
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "button", props["type"])
	})
}

// TestRadioWithControl 测试Radio的Control功能
func TestRadioWithControl(t *testing.T) {
	t.Run("SimpleControl", func(t *testing.T) {
		radio := NewRadio("type", "类型", "1").
			SetOptions([]Option{
				{Value: "1", Label: "类型1"},
				{Value: "2", Label: "类型2"},
			}).
			Control([]ControlRule{
				{
					Value: "1",
					Rule: []Component{
						NewInput("field1", "字段1").Required(),
					},
				},
			})

		result := radio.Build()

		control, ok := result["control"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, control, 1)
		assert.Equal(t, "1", control[0]["value"])

		rules, ok := control[0]["rule"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, rules, 1)
		assert.Equal(t, "field1", rules[0]["field"])
	})

	t.Run("MultipleControlRules", func(t *testing.T) {
		radio := NewRadio("type", "类型", "1").
			SetOptions([]Option{
				{Value: "1", Label: "试用期"},
				{Value: "2", Label: "正式"},
			}).
			Control([]ControlRule{
				{
					Value: "1",
					Rule: []Component{
						NewInputNumber("days", "试用天数").Required(),
					},
				},
				{
					Value: "2",
					Rule: []Component{
						NewDatePicker("start_date", "入职日期"),
					},
				},
			})

		result := radio.Build()
		control := result["control"].([]map[string]interface{})
		require.Len(t, control, 2)
	})

	t.Run("NestedControl", func(t *testing.T) {
		radio := NewRadio("level1", "一级", "a").
			Control([]ControlRule{
				{
					Value: "a",
					Rule: []Component{
						NewRadio("level2", "二级", "x").
							SetOptions([]Option{
								{Value: "x", Label: "X"},
								{Value: "y", Label: "Y"},
							}).
							Control([]ControlRule{
								{
									Value: "x",
									Rule: []Component{
										NewInput("nested_field", "嵌套字段"),
									},
								},
							}),
					},
				},
			})

		result := radio.Build()
		control := result["control"].([]map[string]interface{})
		require.Len(t, control, 1)

		// 验证嵌套结构
		rules := control[0]["rule"].([]map[string]interface{})
		nestedRadio := rules[0]
		assert.Equal(t, "level2", nestedRadio["field"])

		nestedControl, ok := nestedRadio["control"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, nestedControl, 1)
	})
}

// TestRadioEdgeCases 测试边缘情况
func TestRadioEdgeCases(t *testing.T) {
	t.Run("EmptyOptions", func(t *testing.T) {
		radio := NewRadio("test", "测试")
		result := radio.Build()

		options, exists := result["options"]
		if exists {
			assert.Len(t, options, 0)
		}
	})

	t.Run("NumericValue", func(t *testing.T) {
		radio := NewRadio("test", "测试", 123)

		data := radio.GetData()
		assert.Equal(t, 123, data.Value)
	})

	t.Run("BooleanValue", func(t *testing.T) {
		radio := NewRadio("test", "测试", true).
			SetOptions([]Option{
				{Value: true, Label: "是"},
				{Value: false, Label: "否"},
			})

		assert.Equal(t, true, radio.options[0].Value)
	})

	t.Run("ControlWithEmptyRule", func(t *testing.T) {
		radio := NewRadio("test", "测试").
			Control([]ControlRule{
				{
					Value: "1",
					Rule:  []Component{},
				},
			})

		result := radio.Build()
		control := result["control"].([]map[string]interface{})
		rules := control[0]["rule"].([]map[string]interface{})
		assert.Len(t, rules, 0)
	})
}

// TestRadioWithValidation 测试验证功能
func TestRadioWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		radio := NewRadio("test", "测试").Required()

		data := radio.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("CustomValidation", func(t *testing.T) {
		radio := NewRadio("test", "测试").
			Validate(RequiredRule{Message: "请选择一项"})

		data := radio.GetData()
		require.Len(t, data.Validate, 1)
	})
}

// BenchmarkRadioCreation 性能测试
func BenchmarkRadioCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewRadio("test", "测试", "default")
	}
}

// BenchmarkRadioWithOptions 性能测试
func BenchmarkRadioWithOptions(b *testing.B) {
	options := []Option{
		{Value: "1", Label: "选项1"},
		{Value: "2", Label: "选项2"},
		{Value: "3", Label: "选项3"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewRadio("test", "测试").SetOptions(options)
	}
}

// BenchmarkRadioWithControl 性能测试
func BenchmarkRadioWithControl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewRadio("type", "类型", "1").
			SetOptions([]Option{
				{Value: "1", Label: "类型1"},
				{Value: "2", Label: "类型2"},
			}).
			Control([]ControlRule{
				{
					Value: "1",
					Rule: []Component{
						NewInput("field1", "字段1"),
					},
				},
			}).
			Build()
	}
}
