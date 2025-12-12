package formbuilder

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// component_test.go 测试Component接口和Builder泛型系统
// 这是核心测试文件，覆盖所有基础组件功能

// TestComponentInterface 测试Component接口的实现
func TestComponentInterface(t *testing.T) {
	t.Run("InputImplementsComponent", func(t *testing.T) {
		var _ Component = NewInput("test", "测试")
	})

	t.Run("SelectImplementsComponent", func(t *testing.T) {
		var _ Component = NewSelect("test", "测试")
	})

	t.Run("RadioImplementsComponent", func(t *testing.T) {
		var _ Component = NewRadio("test", "测试")
	})

	t.Run("ComponentInterfaceMethods", func(t *testing.T) {
		input := NewInput("username", "用户名")

		assert.Equal(t, "username", input.GetField())
		assert.Equal(t, "input", input.GetType())

		result := input.Build()
		assert.NotNil(t, result)
		assert.Equal(t, "input", result["type"])
		assert.Equal(t, "username", result["field"])
		assert.Equal(t, "用户名", result["title"])
	})
}

// TestBuilderGenericTypeSafety 测试Builder泛型的类型安全性
func TestBuilderGenericTypeSafety(t *testing.T) {
	t.Run("InputBuilderReturnsInput", func(t *testing.T) {
		input := NewInput("test", "测试").
			Required().
			Value("default")

		// 编译时检查：如果类型不对，这里会编译失败
		assert.Equal(t, "*formbuilder.Input", getTypeName(input))
	})

	t.Run("SelectBuilderReturnsSelect", func(t *testing.T) {
		sel := NewSelect("test", "测试").
			Required().
			Value("default")

		assert.Equal(t, "*formbuilder.Select", getTypeName(sel))
	})

	t.Run("ChainedMethodsPreserveType", func(t *testing.T) {
		// 测试链式调用保持类型
		input := NewInput("test", "测试").
			Placeholder("placeholder"). // Input特有方法
			Required().                 // Builder通用方法
			Value("value").             // Builder通用方法
			Clearable(true)             // Input特有方法

		assert.NotNil(t, input)
	})
}

// TestBuilderRequired 测试Required方法
func TestBuilderRequired(t *testing.T) {
	t.Run("AddsRequiredRule", func(t *testing.T) {
		input := NewInput("test", "测试").Required()

		data := input.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0]
		ruleMap := rule.ToMap()
		assert.Equal(t, true, ruleMap["required"])
		assert.Equal(t, "此项必填", ruleMap["message"])
	})

	t.Run("MultipleRequired", func(t *testing.T) {
		input := NewInput("test", "测试").
			Required().
			Required() // 重复调用

		data := input.GetData()
		assert.Len(t, data.Validate, 2) // 会添加两个规则
	})
}

// TestBuilderValue 测试Value方法
func TestBuilderValue(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected interface{}
	}{
		{"StringValue", "test", "test"},
		{"IntValue", 123, 123},
		{"BoolValue", true, true},
		{"FloatValue", 3.14, 3.14},
		{"NilValue", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := NewInput("test", "测试").Value(tt.value)
			data := input.GetData()
			assert.Equal(t, tt.expected, data.Value)
		})
	}
}

// TestBuilderTitle 测试Title方法
func TestBuilderTitle(t *testing.T) {
	t.Run("SetTitle", func(t *testing.T) {
		input := NewInput("test", "初始标题").Title("新标题")
		data := input.GetData()
		assert.Equal(t, "新标题", data.Title)
	})

	t.Run("EmptyTitle", func(t *testing.T) {
		input := NewInput("test", "").Title("")
		data := input.GetData()
		assert.Equal(t, "", data.Title)
	})
}

// TestBuilderField 测试Field方法
func TestBuilderField(t *testing.T) {
	t.Run("SetField", func(t *testing.T) {
		input := NewInput("oldfield", "测试").Field("newfield")
		data := input.GetData()
		assert.Equal(t, "newfield", data.Field)
		assert.Equal(t, "newfield", input.GetField())
	})
}

// TestBuilderProps 测试Props方法
func TestBuilderProps(t *testing.T) {
	t.Run("SetSingleProp", func(t *testing.T) {
		input := NewInput("test", "测试").
			Props("placeholder", "请输入").
			Props("clearable", true)

		data := input.GetData()
		require.NotNil(t, data.Props)
		assert.Equal(t, "请输入", data.Props["placeholder"])
		assert.Equal(t, true, data.Props["clearable"])
	})

	t.Run("PropsOverwrite", func(t *testing.T) {
		input := NewInput("test", "测试").
			Props("key", "value1").
			Props("key", "value2")

		data := input.GetData()
		assert.Equal(t, "value2", data.Props["key"])
	})

	t.Run("PropsWithNilValue", func(t *testing.T) {
		input := NewInput("test", "测试").Props("key", nil)
		data := input.GetData()
		assert.Nil(t, data.Props["key"])
	})
}

// TestBuilderSetProps 测试SetProps批量设置
func TestBuilderSetProps(t *testing.T) {
	t.Run("BatchSetProps", func(t *testing.T) {
		props := map[string]interface{}{
			"placeholder": "请输入",
			"clearable":   true,
			"maxlength":   50,
		}

		input := NewInput("test", "测试").SetProps(props)

		data := input.GetData()
		assert.Equal(t, "请输入", data.Props["placeholder"])
		assert.Equal(t, true, data.Props["clearable"])
		assert.Equal(t, 50, data.Props["maxlength"])
	})

	t.Run("SetPropsMerge", func(t *testing.T) {
		input := NewInput("test", "测试").
			Props("prop1", "value1").
			SetProps(map[string]interface{}{
				"prop2": "value2",
				"prop3": "value3",
			})

		data := input.GetData()
		assert.Equal(t, "value1", data.Props["prop1"])
		assert.Equal(t, "value2", data.Props["prop2"])
		assert.Equal(t, "value3", data.Props["prop3"])
	})
}

// TestBuilderValidate 测试Validate方法
func TestBuilderValidate(t *testing.T) {
	t.Run("AddSingleValidateRule", func(t *testing.T) {
		input := NewInput("test", "测试").
			Validate(RequiredRule{Message: "必填"})

		data := input.GetData()
		require.Len(t, data.Validate, 1)
	})

	t.Run("AddMultipleValidateRules", func(t *testing.T) {
		input := NewInput("test", "测试").
			Validate(
				RequiredRule{Message: "必填"},
				PatternRule{Pattern: "^\\d+$", Message: "只能数字"},
				LengthRule{Min: 6, Max: 20, Message: "长度6-20"},
			)

		data := input.GetData()
		require.Len(t, data.Validate, 3)
	})

	t.Run("ValidateChaining", func(t *testing.T) {
		input := NewInput("test", "测试").
			Validate(RequiredRule{Message: "必填"}).
			Validate(PatternRule{Pattern: "^\\d+$", Message: "只能数字"})

		data := input.GetData()
		require.Len(t, data.Validate, 2)
	})
}

// TestBuilderControl 测试Control条件显示
func TestBuilderControl(t *testing.T) {
	t.Run("SetControlRules", func(t *testing.T) {
		radio := NewRadio("type", "类型", "1").
			Control([]ControlRule{
				{
					Value: "1",
					Rule: []Component{
						NewInput("field1", "字段1"),
					},
				},
				{
					Value: "2",
					Rule: []Component{
						NewInput("field2", "字段2"),
					},
				},
			})

		data := radio.GetData()
		require.Len(t, data.Control, 2)
		assert.Equal(t, "1", data.Control[0].Value)
		assert.Len(t, data.Control[0].Rule, 1)
	})

	t.Run("ControlRuleOverwrite", func(t *testing.T) {
		radio := NewRadio("type", "类型", "1").
			Control([]ControlRule{{Value: "1", Rule: nil}}).
			Control([]ControlRule{{Value: "2", Rule: nil}})

		data := radio.GetData()
		require.Len(t, data.Control, 1)
		assert.Equal(t, "2", data.Control[0].Value)
	})
}

// TestBuilderAppendControl 测试AppendControl方法
func TestBuilderAppendControl(t *testing.T) {
	t.Run("AppendSingleControl", func(t *testing.T) {
		radio := NewRadio("type", "类型", "1").
			AppendControl(ControlRule{
				Value: "1",
				Rule:  []Component{NewInput("field1", "字段1")},
			}).
			AppendControl(ControlRule{
				Value: "2",
				Rule:  []Component{NewInput("field2", "字段2")},
			})

		data := radio.GetData()
		require.Len(t, data.Control, 2)
	})
}

// TestBuilderChildren 测试Children子组件
func TestBuilderChildren(t *testing.T) {
	t.Run("SetChildren", func(t *testing.T) {
		parent := NewInput("parent", "父组件").
			Children([]Component{
				NewInput("child1", "子组件1"),
				NewInput("child2", "子组件2"),
			})

		data := parent.GetData()
		require.Len(t, data.Children, 2)
	})

	t.Run("ChildrenOverwrite", func(t *testing.T) {
		parent := NewInput("parent", "父组件").
			Children([]Component{NewInput("child1", "子组件1")}).
			Children([]Component{NewInput("child2", "子组件2")})

		data := parent.GetData()
		require.Len(t, data.Children, 1)
		assert.Equal(t, "child2", data.Children[0].GetField())
	})
}

// TestBuilderAppendChild 测试AppendChild方法
func TestBuilderAppendChild(t *testing.T) {
	t.Run("AppendMultipleChildren", func(t *testing.T) {
		parent := NewInput("parent", "父组件").
			AppendChild(NewInput("child1", "子组件1")).
			AppendChild(NewInput("child2", "子组件2")).
			AppendChild(NewInput("child3", "子组件3"))

		data := parent.GetData()
		require.Len(t, data.Children, 3)
	})
}

// TestBuilderEmit 测试Emit事件配置
func TestBuilderEmit(t *testing.T) {
	t.Run("SetSingleEmit", func(t *testing.T) {
		input := NewInput("test", "测试").
			Emit("change", "handleChange")

		data := input.GetData()
		require.NotNil(t, data.Emit)
		assert.Equal(t, "handleChange", data.Emit["change"])
	})

	t.Run("SetMultipleEmits", func(t *testing.T) {
		input := NewInput("test", "测试").
			Emit("change", "handleChange").
			Emit("blur", "handleBlur").
			Emit("focus", "handleFocus")

		data := input.GetData()
		assert.Len(t, data.Emit, 3)
		assert.Equal(t, "handleChange", data.Emit["change"])
		assert.Equal(t, "handleBlur", data.Emit["blur"])
		assert.Equal(t, "handleFocus", data.Emit["focus"])
	})
}

// TestBuildComponent 测试buildComponent函数
func TestBuildComponent(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		input := NewInput("username", "用户名").Value("test")
		result := input.Build()

		assert.Equal(t, "input", result["type"])
		assert.Equal(t, "username", result["field"])
		assert.Equal(t, "用户名", result["title"])
		assert.Equal(t, "test", result["value"])
	})

	t.Run("BuildWithProps", func(t *testing.T) {
		input := NewInput("test", "测试").
			Props("placeholder", "请输入").
			Props("clearable", true)

		result := input.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "请输入", props["placeholder"])
		assert.Equal(t, true, props["clearable"])
	})

	t.Run("BuildWithValidation", func(t *testing.T) {
		input := NewInput("test", "测试").
			Required().
			Validate(PatternRule{Pattern: "^\\d+$", Message: "只能数字"})

		result := input.Build()
		validates, ok := result["validate"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, validates, 2)
	})

	t.Run("BuildWithControl", func(t *testing.T) {
		radio := NewRadio("type", "类型", "1").
			Control([]ControlRule{
				{
					Value: "1",
					Rule: []Component{
						NewInput("field1", "字段1"),
					},
				},
			})

		result := radio.Build()
		controls, ok := result["control"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, controls, 1)
		assert.Equal(t, "1", controls[0]["value"])

		rules, ok := controls[0]["rule"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, rules, 1)
	})

	t.Run("BuildWithChildren", func(t *testing.T) {
		parent := NewInput("parent", "父组件").
			Children([]Component{
				NewInput("child1", "子组件1"),
				NewInput("child2", "子组件2"),
			})

		result := parent.Build()
		children, ok := result["children"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, children, 2)
	})

	t.Run("BuildWithEmit", func(t *testing.T) {
		input := NewInput("test", "测试").
			Emit("change", "handleChange")

		result := input.Build()
		emit, ok := result["emit"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "handleChange", emit["change"])
	})

	t.Run("BuildJSONSerialization", func(t *testing.T) {
		input := NewInput("test", "测试").
			Placeholder("请输入").
			Required()

		result := input.Build()
		jsonBytes, err := json.Marshal(result)
		require.NoError(t, err)

		var decoded map[string]interface{}
		err = json.Unmarshal(jsonBytes, &decoded)
		require.NoError(t, err)
		assert.Equal(t, "input", decoded["type"])
	})
}

// TestComponentDataStructure 测试ComponentData结构
func TestComponentDataStructure(t *testing.T) {
	t.Run("EmptyComponentData", func(t *testing.T) {
		data := &ComponentData{}
		assert.Equal(t, "", data.Field)
		assert.Equal(t, "", data.Title)
		assert.Nil(t, data.Value)
		assert.Nil(t, data.Props)
		assert.Nil(t, data.Validate)
		assert.Nil(t, data.Control)
		assert.Nil(t, data.Children)
		assert.Nil(t, data.Emit)
	})

	t.Run("InitializedComponentData", func(t *testing.T) {
		data := &ComponentData{
			Field:    "test",
			Title:    "测试",
			RuleType: "input",
			Value:    "default",
			Props:    make(map[string]interface{}),
			Validate: []ValidateRule{},
			Control:  []ControlRule{},
			Children: []Component{},
			Emit:     make(map[string]interface{}),
		}

		assert.Equal(t, "test", data.Field)
		assert.NotNil(t, data.Props)
		assert.NotNil(t, data.Validate)
	})
}

// TestRequiredRule 测试RequiredRule
func TestRequiredRule(t *testing.T) {
	t.Run("BasicRequiredRule", func(t *testing.T) {
		rule := RequiredRule{Message: "必填"}
		ruleMap := rule.ToMap()

		assert.Equal(t, true, ruleMap["required"])
		assert.Equal(t, "必填", ruleMap["message"])
	})

	t.Run("RequiredRuleWithTrigger", func(t *testing.T) {
		rule := RequiredRule{Message: "必填", Trigger: "blur"}
		ruleMap := rule.ToMap()

		assert.Equal(t, true, ruleMap["required"])
		assert.Equal(t, "必填", ruleMap["message"])
		assert.Equal(t, "blur", ruleMap["trigger"])
	})

	t.Run("RequiredRuleWithoutMessage", func(t *testing.T) {
		rule := RequiredRule{}
		ruleMap := rule.ToMap()

		assert.Equal(t, true, ruleMap["required"])
		_, hasMessage := ruleMap["message"]
		assert.False(t, hasMessage)
	})
}

// TestControlRule 测试ControlRule
func TestControlRule(t *testing.T) {
	t.Run("BasicControlRule", func(t *testing.T) {
		ctrl := ControlRule{
			Value: "1",
			Rule: []Component{
				NewInput("field1", "字段1"),
			},
		}

		assert.Equal(t, "1", ctrl.Value)
		require.Len(t, ctrl.Rule, 1)
		assert.Equal(t, "field1", ctrl.Rule[0].GetField())
	})

	t.Run("ControlRuleWithMultipleComponents", func(t *testing.T) {
		ctrl := ControlRule{
			Value: "show",
			Rule: []Component{
				NewInput("field1", "字段1"),
				NewSelect("field2", "字段2"),
				NewRadio("field3", "字段3"),
			},
		}

		require.Len(t, ctrl.Rule, 3)
	})

	t.Run("NestedControlRule", func(t *testing.T) {
		ctrl := ControlRule{
			Value: "type1",
			Rule: []Component{
				NewRadio("subtype", "子类型", "a").
					Control([]ControlRule{
						{
							Value: "a",
							Rule: []Component{
								NewInput("nested_field", "嵌套字段"),
							},
						},
					}),
			},
		}

		require.Len(t, ctrl.Rule, 1)
		radio := ctrl.Rule[0]
		builder, ok := radio.(interface{ GetData() *ComponentData })
		require.True(t, ok)

		data := builder.GetData()
		require.Len(t, data.Control, 1)
	})
}

// TestBuilderChaining 测试完整的链式调用
func TestBuilderChaining(t *testing.T) {
	t.Run("ComplexChaining", func(t *testing.T) {
		input := NewInput("username", "用户名").
			Placeholder("请输入用户名").
			Clearable(true).
			MaxLength(50).
			Required().
			Validate(
				PatternRule{Pattern: "^[a-zA-Z0-9]+$", Message: "只能字母数字"},
				LengthRule{Min: 6, Max: 20, Message: "长度6-20"},
			).
			Value("default").
			Props("custom", "value")

		// 验证所有设置都生效
		data := input.GetData()
		assert.Equal(t, "username", data.Field)
		assert.Equal(t, "用户名", data.Title)
		assert.Equal(t, "default", data.Value)
		assert.Equal(t, "请输入用户名", data.Props["placeholder"])
		assert.Equal(t, true, data.Props["clearable"])
		assert.Equal(t, 50, data.Props["maxlength"])
		assert.Equal(t, "value", data.Props["custom"])
		require.Len(t, data.Validate, 3) // Required + 2个自定义规则
	})

	t.Run("ChainingPreservesInstanceType", func(t *testing.T) {
		// 测试链式调用后仍然可以调用组件特有方法
		input := NewInput("test", "测试").
			Required().
			Value("value").
			Placeholder("placeholder"). // Input特有方法
			Clearable(true).            // Input特有方法
			ShowPassword(true)          // Input特有方法

		assert.NotNil(t, input)
	})
}

// getTypeName 辅助函数：获取类型名称（用于测试）
func getTypeName(v interface{}) string {
	switch v.(type) {
	case *Input:
		return "*formbuilder.Input"
	case *Select:
		return "*formbuilder.Select"
	case *Radio:
		return "*formbuilder.Radio"
	default:
		return "unknown"
	}
}

// TestBuilderAppendRule 测试AppendRule方法
func TestBuilderAppendRule(t *testing.T) {
	t.Run("AppendSingleRule", func(t *testing.T) {
		input := NewInput("test", "测试").
			AppendRule("suffix", "后缀内容")

		data := input.GetData()
		require.NotNil(t, data.AppendRule)
		assert.Equal(t, "后缀内容", data.AppendRule["suffix"])
	})

	t.Run("AppendMultipleRules", func(t *testing.T) {
		input := NewInput("test", "测试").
			AppendRule("suffix", "后缀").
			AppendRule("prefix", "前缀").
			AppendRule("customField", "自定义值")

		data := input.GetData()
		require.NotNil(t, data.AppendRule)
		assert.Len(t, data.AppendRule, 3)
		assert.Equal(t, "后缀", data.AppendRule["suffix"])
		assert.Equal(t, "前缀", data.AppendRule["prefix"])
		assert.Equal(t, "自定义值", data.AppendRule["customField"])
	})

	t.Run("AppendComplexObject", func(t *testing.T) {
		input := NewInput("test", "测试").
			AppendRule("suffix", map[string]interface{}{
				"type": "div",
				"style": map[string]interface{}{
					"color": "#999999",
				},
				"domProps": map[string]interface{}{
					"innerHTML": "提示信息",
				},
			})

		data := input.GetData()
		suffix, ok := data.AppendRule["suffix"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "div", suffix["type"])

		style, ok := suffix["style"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "#999999", style["color"])
	})

	t.Run("AppendRuleChaining", func(t *testing.T) {
		input := NewInput("amount", "金额").
			Placeholder("请输入金额").
			Required().
			AppendRule("prefix", "¥").
			AppendRule("suffix", "元").
			Value(100)

		data := input.GetData()
		assert.Equal(t, "¥", data.AppendRule["prefix"])
		assert.Equal(t, "元", data.AppendRule["suffix"])
		assert.Equal(t, 100, data.Value)
		assert.Equal(t, "请输入金额", data.Props["placeholder"])
	})

	t.Run("AppendRuleOverwrite", func(t *testing.T) {
		input := NewInput("test", "测试").
			AppendRule("key", "value1").
			AppendRule("key", "value2")

		data := input.GetData()
		assert.Equal(t, "value2", data.AppendRule["key"])
	})
}

// TestBuildComponentWithAppendRule 测试buildComponent函数对AppendRule的处理
func TestBuildComponentWithAppendRule(t *testing.T) {
	t.Run("BuildWithSingleAppendRule", func(t *testing.T) {
		input := NewInput("test", "测试").
			AppendRule("suffix", "后缀内容")

		result := input.Build()
		assert.Equal(t, "后缀内容", result["suffix"])
	})

	t.Run("BuildWithMultipleAppendRules", func(t *testing.T) {
		input := NewInput("test", "测试").
			AppendRule("suffix", "后缀").
			AppendRule("prefix", "前缀").
			AppendRule("customData", map[string]interface{}{
				"min": 0,
				"max": 100,
			})

		result := input.Build()
		assert.Equal(t, "后缀", result["suffix"])
		assert.Equal(t, "前缀", result["prefix"])

		customData, ok := result["customData"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, 0, customData["min"])
		assert.Equal(t, 100, customData["max"])
	})

	t.Run("AppendRuleOverridesStandardField", func(t *testing.T) {
		// 测试AppendRule可以覆盖标准字段
		input := NewInput("test", "测试").
			Value("original").
			AppendRule("value", "overridden")

		result := input.Build()
		assert.Equal(t, "overridden", result["value"])
	})

	t.Run("BuildWithAppendRuleAndOtherFields", func(t *testing.T) {
		input := NewInput("username", "用户名").
			Placeholder("请输入用户名").
			Required().
			AppendRule("suffix", map[string]interface{}{
				"type": "div",
				"props": map[string]interface{}{
					"innerHTML": "6-20个字符",
				},
			}).
			Value("default")

		result := input.Build()

		// 验证标准字段
		assert.Equal(t, "input", result["type"])
		assert.Equal(t, "username", result["field"])
		assert.Equal(t, "用户名", result["title"])
		assert.Equal(t, "default", result["value"])

		// 验证props
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "请输入用户名", props["placeholder"])

		// 验证validate
		validates, ok := result["validate"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, validates, 1)

		// 验证AppendRule添加的字段
		suffix, ok := result["suffix"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "div", suffix["type"])
	})

	t.Run("AppendRuleJSONSerialization", func(t *testing.T) {
		input := NewInput("test", "测试").
			AppendRule("suffix", "后缀").
			AppendRule("customField", map[string]interface{}{
				"key": "value",
			})

		result := input.Build()
		jsonBytes, err := json.Marshal(result)
		require.NoError(t, err)

		var decoded map[string]interface{}
		err = json.Unmarshal(jsonBytes, &decoded)
		require.NoError(t, err)

		assert.Equal(t, "后缀", decoded["suffix"])
		customField, ok := decoded["customField"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "value", customField["key"])
	})
}

// TestAppendRuleWithPHPExample 测试复现PHP示例的用法
func TestAppendRuleWithPHPExample(t *testing.T) {
	t.Run("PHPRadioSuffixExample", func(t *testing.T) {
		// 对应PHP示例：
		// Elm::radio('svip_type', '会员类别：', '2')
		//     ->appendRule('suffix', [
		//         'type' => 'div',
		//         'style' => ['color' => '#999999'],
		//         'domProps' => [
		//             'innerHTML' =>'试用期每个用户只能购买一次',
		//         ]
		//     ])

		radio := NewRadio("svip_type", "会员类别：", "2").
			SetOptions([]Option{
				{Value: "1", Label: "试用期"},
				{Value: "2", Label: "有限期"},
				{Value: "3", Label: "永久期"},
			}).
			AppendRule("suffix", map[string]interface{}{
				"type": "div",
				"style": map[string]interface{}{
					"color": "#999999",
				},
				"domProps": map[string]interface{}{
					"innerHTML": "试用期每个用户只能购买一次",
				},
			})

		result := radio.Build()

		// 验证基础字段
		assert.Equal(t, "radio", result["type"])
		assert.Equal(t, "svip_type", result["field"])
		assert.Equal(t, "会员类别：", result["title"])
		assert.Equal(t, "2", result["value"])

		// 验证suffix
		suffix, ok := result["suffix"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "div", suffix["type"])

		style, ok := suffix["style"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "#999999", style["color"])

		domProps, ok := suffix["domProps"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "试用期每个用户只能购买一次", domProps["innerHTML"])
	})

	t.Run("MultipleAppendRulesWithOtherMethods", func(t *testing.T) {
		input := NewInput("price", "价格").
			Placeholder("请输入价格").
			Required().
			Validate(NewRange(0, 99999, "价格必须在0-99999之间")).
			AppendRule("prefix", "¥").
			AppendRule("suffix", "元")

		result := input.Build()

		assert.Equal(t, "¥", result["prefix"])
		assert.Equal(t, "元", result["suffix"])
		assert.Equal(t, "input", result["type"])
		assert.Equal(t, "price", result["field"])

		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "请输入价格", props["placeholder"])

		validates, ok := result["validate"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, validates, 2) // Required + Range
	})
}
