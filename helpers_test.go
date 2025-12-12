package formbuilder

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 测试辅助工具函数

// assertJSONEqual 比较两个JSON字符串是否相等（忽略格式和键顺序）
func assertJSONEqual(t *testing.T, expected, actual string) {
	var expectedData, actualData interface{}

	err := json.Unmarshal([]byte(expected), &expectedData)
	require.NoError(t, err, "Expected JSON is invalid")

	err = json.Unmarshal([]byte(actual), &actualData)
	require.NoError(t, err, "Actual JSON is invalid")

	assert.Equal(t, expectedData, actualData)
}

// normalizeJSONString 标准化JSON字符串，用于对比
// 去除空格、排序键，确保对比的准确性
func normalizeJSONString(jsonStr string) (string, error) {
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

// createTestForm 创建一个标准的测试表单
func createTestForm() *Form {
	return NewElmForm("/test/submit", []Component{
		NewInput("username", "用户名").Required(),
		NewInput("email", "邮箱"),
		NewSelect("role", "角色").SetOptions([]Option{
			{Value: "admin", Label: "管理员"},
			{Value: "user", Label: "用户"},
		}),
	}, nil)
}

// createComplexTestForm 创建一个复杂的测试表单，包含多种组件
func createComplexTestForm() *Form {
	return NewElmForm("/test/complex", []Component{
		NewInput("username", "用户名").
			Placeholder("请输入用户名").
			MaxLength(50).
			Required(),

		Password("password", "密码").
			MinLength(6).
			Required(),

		Email("email", "邮箱").Required(),

		NewSelect("role", "角色").
			SetOptions(mockOptions(3)).
			Required(),

		NewRadio("status", "状态", "1").
			SetOptions([]Option{
				{Value: "1", Label: "启用"},
				{Value: "0", Label: "禁用"},
			}),

		NewCheckbox("permissions", "权限").
			SetOptions([]Option{
				{Value: "read", Label: "读取"},
				{Value: "write", Label: "写入"},
				{Value: "delete", Label: "删除"},
			}),

		NewInputNumber("age", "年龄").
			Min(0).
			Max(150),

		NewSwitch("is_active", "是否激活").
			ActiveText("激活").
			InactiveText("禁用"),
	}, nil)
}

// mockOptions 创建测试用选项
func mockOptions(count int) []Option {
	options := make([]Option, count)
	for i := 0; i < count; i++ {
		options[i] = Option{
			Value: i + 1,
			Label: "选项" + string(rune('A'+i)),
		}
	}
	return options
}

// mockOptionsWithValues 创建带自定义值的测试选项
func mockOptionsWithValues(values map[interface{}]string) []Option {
	options := make([]Option, 0, len(values))
	for value, label := range values {
		options = append(options, Option{
			Value: value,
			Label: label,
		})
	}
	return options
}

// assertComponentField 断言组件的字段名
func assertComponentField(t *testing.T, component Component, expectedField string) {
	assert.Equal(t, expectedField, component.GetField(), "Component field mismatch")
}

// assertComponentType 断言组件的类型
func assertComponentType(t *testing.T, component Component, expectedType string) {
	assert.Equal(t, expectedType, component.GetType(), "Component type mismatch")
}

// assertComponentHasProp 断言组件包含指定的prop属性
func assertComponentHasProp(t *testing.T, component Component, propName string, expectedValue interface{}) {
	builder, ok := component.(interface{ GetData() *ComponentData })
	require.True(t, ok, "Component does not implement GetData()")

	data := builder.GetData()
	require.NotNil(t, data.Props, "Component props is nil")

	actualValue, exists := data.Props[propName]
	require.True(t, exists, "Prop %s does not exist", propName)
	assert.Equal(t, expectedValue, actualValue, "Prop %s value mismatch", propName)
}

// assertComponentHasValidation 断言组件包含验证规则
func assertComponentHasValidation(t *testing.T, component Component) {
	builder, ok := component.(interface{ GetData() *ComponentData })
	require.True(t, ok, "Component does not implement GetData()")

	data := builder.GetData()
	assert.NotEmpty(t, data.Validate, "Component should have validation rules")
}

// assertBuildMapHasKey 断言Build()结果包含指定的键
func assertBuildMapHasKey(t *testing.T, component Component, key string) {
	result := component.Build()
	_, exists := result[key]
	assert.True(t, exists, "Build result should contain key: %s", key)
}

// assertBuildMapValue 断言Build()结果中指定键的值
func assertBuildMapValue(t *testing.T, component Component, key string, expectedValue interface{}) {
	result := component.Build()
	actualValue, exists := result[key]
	require.True(t, exists, "Build result should contain key: %s", key)
	assert.Equal(t, expectedValue, actualValue, "Build result key %s value mismatch", key)
}

// createTestControlRules 创建测试用Control规则
func createTestControlRules() []ControlRule {
	return []ControlRule{
		{
			Value: "1",
			Rule: []Component{
				NewInput("extra_field", "额外字段").Required(),
			},
		},
		{
			Value: "2",
			Rule: []Component{
				NewInputNumber("number_field", "数字字段"),
			},
		},
	}
}

// createNestedControlRules 创建嵌套的Control规则（用于测试复杂场景）
func createNestedControlRules() []ControlRule {
	return []ControlRule{
		{
			Value: "type1",
			Rule: []Component{
				NewRadio("subtype", "子类型", "a").
					SetOptions([]Option{
						{Value: "a", Label: "类型A"},
						{Value: "b", Label: "类型B"},
					}).
					Control([]ControlRule{
						{
							Value: "a",
							Rule: []Component{
								NewInput("type_a_field", "类型A字段"),
							},
						},
					}),
			},
		},
	}
}

// TestHelpersUtility 测试辅助函数本身的正确性
func TestHelpersUtility(t *testing.T) {
	t.Run("assertJSONEqual", func(t *testing.T) {
		json1 := `{"name":"test","value":123}`
		json2 := `{"value":123,"name":"test"}` // 键顺序不同
		assertJSONEqual(t, json1, json2)       // 应该相等
	})

	t.Run("normalizeJSONString", func(t *testing.T) {
		input := `{"b":2,"a":1}`
		normalized, err := normalizeJSONString(input)
		require.NoError(t, err)
		assert.NotEmpty(t, normalized)
	})

	t.Run("createTestForm", func(t *testing.T) {
		form := createTestForm()
		assert.NotNil(t, form)
		assert.Equal(t, "/test/submit", form.action)
		assert.Len(t, form.rules, 3)
	})

	t.Run("mockOptions", func(t *testing.T) {
		options := mockOptions(5)
		assert.Len(t, options, 5)
		assert.Equal(t, 1, options[0].Value)
		assert.Equal(t, "选项A", options[0].Label)
	})
}
