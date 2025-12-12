package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// textarea_test.go Textarea组件完整测试

// TestTextareaCreation 测试Textarea组件创建
func TestTextareaCreation(t *testing.T) {
	t.Run("BasicCreationWithTextarea", func(t *testing.T) {
		textarea := Textarea("description", "描述")

		assert.Equal(t, "description", textarea.GetField())
		assert.Equal(t, "input", textarea.GetType())

		data := textarea.GetData()
		assert.Equal(t, "description", data.Field)
		assert.Equal(t, "描述", data.Title)
		assert.Equal(t, "textarea", data.Props["type"])
	})

	t.Run("BasicCreationWithNewTextarea", func(t *testing.T) {
		textarea := NewTextarea("content", "内容")

		assert.Equal(t, "content", textarea.GetField())
		assert.Equal(t, "input", textarea.GetType())

		data := textarea.GetData()
		assert.Equal(t, "content", data.Field)
		assert.Equal(t, "内容", data.Title)
		assert.Equal(t, "textarea", data.Props["type"])
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		textarea := Textarea("bio", "简介", "This is my bio")

		data := textarea.GetData()
		assert.Equal(t, "This is my bio", data.Value)
		assert.Equal(t, "textarea", data.Props["type"])
	})

	t.Run("TextareaIsInput", func(t *testing.T) {
		textarea := Textarea("text", "文本")
		var input *Input = textarea
		assert.NotNil(t, input)
	})
}

// TestTextareaProperties 测试Textarea继承的属性方法
func TestTextareaProperties(t *testing.T) {
	// Textarea继承了Input的所有方法
	t.Run("Placeholder", func(t *testing.T) {
		textarea := Textarea("notes", "备注").
			Placeholder("请输入备注")

		data := textarea.GetData()
		assert.Equal(t, "请输入备注", data.Props["placeholder"])
	})

	t.Run("Rows", func(t *testing.T) {
		textarea := Textarea("content", "内容").
			Rows(5)

		data := textarea.GetData()
		assert.Equal(t, 5, data.Props["rows"])
	})

	t.Run("Disabled", func(t *testing.T) {
		textarea := Textarea("readonly", "只读").
			Disabled(true)

		data := textarea.GetData()
		assert.Equal(t, true, data.Props["disabled"])
	})

	t.Run("Readonly", func(t *testing.T) {
		textarea := Textarea("display", "展示").
			Readonly(true)

		data := textarea.GetData()
		assert.Equal(t, true, data.Props["readonly"])
	})

	t.Run("MaxLength", func(t *testing.T) {
		textarea := Textarea("comment", "评论").
			MaxLength(500)

		data := textarea.GetData()
		assert.Equal(t, 500, data.Props["maxlength"])
	})

	t.Run("ShowWordLimit", func(t *testing.T) {
		textarea := Textarea("content", "内容").
			MaxLength(200).
			ShowWordLimit(true)

		data := textarea.GetData()
		assert.Equal(t, 200, data.Props["maxlength"])
		assert.Equal(t, true, data.Props["show-word-limit"])
	})

	t.Run("Clearable", func(t *testing.T) {
		textarea := Textarea("text", "文本").
			Clearable(true)

		data := textarea.GetData()
		assert.Equal(t, true, data.Props["clearable"])
	})
}

// TestTextareaChaining 测试链式调用
func TestTextareaChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		textarea := Textarea("description", "描述", "Initial text").
			Placeholder("请输入详细描述").
			Rows(8).
			MaxLength(1000).
			ShowWordLimit(true).
			Clearable(true).
			Required()

		data := textarea.GetData()
		assert.Equal(t, "Initial text", data.Value)
		assert.Equal(t, "textarea", data.Props["type"])
		assert.Equal(t, "请输入详细描述", data.Props["placeholder"])
		assert.Equal(t, 8, data.Props["rows"])
		assert.Equal(t, 1000, data.Props["maxlength"])
		assert.Equal(t, true, data.Props["show-word-limit"])
		assert.Equal(t, true, data.Props["clearable"])
		assert.NotEmpty(t, data.Validate)
	})

	t.Run("MinimalChain", func(t *testing.T) {
		textarea := Textarea("note", "备注").
			Rows(3)

		data := textarea.GetData()
		assert.Equal(t, "textarea", data.Props["type"])
		assert.Equal(t, 3, data.Props["rows"])
	})
}

// TestTextareaBuild 测试Build方法
func TestTextareaBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		textarea := Textarea("content", "内容", "Sample content")

		result := textarea.Build()

		assert.Equal(t, "input", result["type"])
		assert.Equal(t, "content", result["field"])
		assert.Equal(t, "内容", result["title"])
		assert.Equal(t, "Sample content", result["value"])

		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "textarea", props["type"])
	})

	t.Run("BuildWithRows", func(t *testing.T) {
		textarea := Textarea("text", "文本").
			Rows(10)

		result := textarea.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, "textarea", props["type"])
		assert.Equal(t, 10, props["rows"])
	})

	t.Run("BuildWithMaxLength", func(t *testing.T) {
		textarea := Textarea("limited", "限制").
			MaxLength(500).
			ShowWordLimit(true)

		result := textarea.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, 500, props["maxlength"])
		assert.Equal(t, true, props["show-word-limit"])
	})
}

// TestTextareaEdgeCases 测试边缘情况
func TestTextareaEdgeCases(t *testing.T) {
	t.Run("EmptyValue", func(t *testing.T) {
		textarea := Textarea("empty", "空", "")

		data := textarea.GetData()
		assert.Equal(t, "", data.Value)
	})

	t.Run("LongText", func(t *testing.T) {
		longText := ""
		for i := 0; i < 1000; i++ {
			longText += "x"
		}
		textarea := Textarea("long", "长文本", longText)

		data := textarea.GetData()
		assert.Equal(t, longText, data.Value)
		assert.Len(t, data.Value.(string), 1000)
	})

	t.Run("TextWithNewlines", func(t *testing.T) {
		textWithNewlines := "Line 1\nLine 2\nLine 3"
		textarea := Textarea("multiline", "多行", textWithNewlines)

		data := textarea.GetData()
		assert.Equal(t, textWithNewlines, data.Value)
	})

	t.Run("ZeroRows", func(t *testing.T) {
		textarea := Textarea("text", "文本").
			Rows(0)

		data := textarea.GetData()
		assert.Equal(t, 0, data.Props["rows"])
	})

	t.Run("LargeRows", func(t *testing.T) {
		textarea := Textarea("text", "文本").
			Rows(50)

		data := textarea.GetData()
		assert.Equal(t, 50, data.Props["rows"])
	})

	t.Run("DisabledAndReadonly", func(t *testing.T) {
		textarea := Textarea("frozen", "冻结").
			Disabled(true).
			Readonly(true)

		data := textarea.GetData()
		assert.Equal(t, true, data.Props["disabled"])
		assert.Equal(t, true, data.Props["readonly"])
	})

	t.Run("MaxLengthWithoutShowWordLimit", func(t *testing.T) {
		textarea := Textarea("text", "文本").
			MaxLength(100)

		data := textarea.GetData()
		assert.Equal(t, 100, data.Props["maxlength"])
		_, hasShowWordLimit := data.Props["show-word-limit"]
		assert.False(t, hasShowWordLimit)
	})
}

// TestTextareaWithValidation 测试验证功能
func TestTextareaWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		textarea := Textarea("content", "内容").Required()

		data := textarea.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("LengthValidation", func(t *testing.T) {
		textarea := Textarea("bio", "简介").
			Validate(LengthRule{
				Min:     10,
				Max:     500,
				Message: "简介长度必须在10-500字符之间",
			})

		data := textarea.GetData()
		require.Len(t, data.Validate, 1)
	})

	t.Run("PatternValidation", func(t *testing.T) {
		textarea := Textarea("code", "代码").
			Validate(PatternRule{
				Pattern: "^[a-zA-Z0-9\\s]+$",
				Message: "只能包含字母、数字和空格",
			})

		data := textarea.GetData()
		require.Len(t, data.Validate, 1)
	})

	t.Run("MultipleValidations", func(t *testing.T) {
		textarea := Textarea("description", "描述").
			Required().
			Validate(LengthRule{
				Min:     20,
				Max:     1000,
				Message: "描述长度必须在20-1000字符之间",
			})

		data := textarea.GetData()
		require.Len(t, data.Validate, 2)
	})
}

// TestTextareaInForm 测试在表单中的使用
func TestTextareaInForm(t *testing.T) {
	t.Run("FormWithTextarea", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("title", "标题").Required(),
			Textarea("content", "内容").
				Rows(10).
				MaxLength(5000).
				ShowWordLimit(true).
				Required(),
			Textarea("notes", "备注").
				Rows(5).
				Placeholder("可选的备注信息"),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 3)

		rules := form.FormRule()
		assert.Len(t, rules, 3)
		assert.Equal(t, "input", rules[0]["type"])
		assert.Equal(t, "input", rules[1]["type"])
		assert.Equal(t, "input", rules[2]["type"])

		// Check that textareas have type="textarea"
		props1 := rules[1]["props"].(map[string]interface{})
		assert.Equal(t, "textarea", props1["type"])

		props2 := rules[2]["props"].(map[string]interface{})
		assert.Equal(t, "textarea", props2["type"])
	})

	t.Run("TextareaWithControl", func(t *testing.T) {
		textarea := Textarea("feedback_type", "反馈类型", "bug").
			Control([]ControlRule{
				{
					Value: "bug",
					Rule: []Component{
						Textarea("bug_description", "问题描述").
							Rows(8).
							Required(),
					},
				},
				{
					Value: "feature",
					Rule: []Component{
						Textarea("feature_description", "功能建议").
							Rows(8).
							Required(),
					},
				},
			})

		data := textarea.GetData()
		assert.Len(t, data.Control, 2)
		assert.Len(t, data.Control[0].Rule, 1)
		assert.Len(t, data.Control[1].Rule, 1)
	})
}

// TestTextareaVsInput 测试Textarea与Input的区别
func TestTextareaVsInput(t *testing.T) {
	t.Run("TypeDifference", func(t *testing.T) {
		input := NewInput("text_input", "输入")
		textarea := Textarea("text_area", "文本框")

		inputData := input.GetData()
		textareaData := textarea.GetData()

		// Both are el-input
		assert.Equal(t, "input", inputData.RuleType)
		assert.Equal(t, "input", textareaData.RuleType)

		// Input has default type="text", textarea has type="textarea"
		assert.Equal(t, "text", inputData.Props["type"])
		assert.Equal(t, "textarea", textareaData.Props["type"])
	})

	t.Run("SameInterface", func(t *testing.T) {
		textarea := Textarea("text", "文本")
		var component Component = textarea
		assert.NotNil(t, component)

		var input *Input = textarea
		assert.NotNil(t, input)
	})
}

// TestTextareaAutosize 测试Autosize功能
func TestTextareaAutosize(t *testing.T) {
	t.Run("AutosizeBoolean", func(t *testing.T) {
		textarea := Textarea("text", "文本").
			Autosize(true)

		data := textarea.GetData()
		assert.Equal(t, true, data.Props["autosize"])
	})

	t.Run("AutosizeWithObject", func(t *testing.T) {
		autosize := map[string]interface{}{
			"minRows": 3,
			"maxRows": 10,
		}
		textarea := Textarea("text", "文本").
			Autosize(autosize)

		data := textarea.GetData()
		autosizeMap, ok := data.Props["autosize"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, 3, autosizeMap["minRows"])
		assert.Equal(t, 10, autosizeMap["maxRows"])
	})
}

// BenchmarkTextareaCreation 性能测试
func BenchmarkTextareaCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Textarea("test", "测试")
	}
}

// BenchmarkTextareaWithProperties 性能测试
func BenchmarkTextareaWithProperties(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Textarea("content", "内容", "Sample text").
			Rows(10).
			MaxLength(1000).
			ShowWordLimit(true).
			Placeholder("请输入内容").
			Required()
	}
}

// BenchmarkTextareaBuild 性能测试
func BenchmarkTextareaBuild(b *testing.B) {
	textarea := Textarea("content", "内容", "Sample text").
		Rows(10).
		MaxLength(1000).
		ShowWordLimit(true).
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = textarea.Build()
	}
}
