package components

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStyleMethod(t *testing.T) {
	t.Run("Style方法 - 字符串样式", func(t *testing.T) {
		input := NewInput("test", "测试")
		input.Style("color: red; font-size: 14px;")

		result := input.Build()
		style, exists := result["style"]
		assert.True(t, exists, "style字段应该存在")
		assert.Equal(t, "color: red; font-size: 14px;", style, "style值应该正确")
	})

	t.Run("Style方法 - Map样式", func(t *testing.T) {
		input := NewInput("test", "测试")
		styleMap := map[string]interface{}{
			"color":      "blue",
			"fontSize":   "16px",
			"fontWeight": "bold",
		}
		input.Style(styleMap)

		result := input.Build()
		style := result["style"].(map[string]interface{})
		assert.Equal(t, "blue", style["color"], "color值应该正确")
		assert.Equal(t, "16px", style["fontSize"], "fontSize值应该正确")
		assert.Equal(t, "bold", style["fontWeight"], "fontWeight值应该正确")
	})

	t.Run("Style方法 - 链式调用", func(t *testing.T) {
		input := NewInput("test", "测试")
		input.Style("color: green;")
		input.Required()

		buildResult := input.Build()
		style := buildResult["style"]
		assert.Equal(t, "color: green;", style, "Style值应该正确设置")

		// 验证Required()也生效了
		validate, exists := buildResult["validate"]
		assert.True(t, exists, "validate字段应该存在")
		assert.NotEmpty(t, validate, "应该有验证规则")
	})

	t.Run("Style方法 - 覆盖样式", func(t *testing.T) {
		input := NewInput("test", "测试")
		input.Style("color: red;")
		input.Style("color: blue; background: white;") // 覆盖之前的样式

		result := input.Build()
		style := result["style"]
		assert.Equal(t, "color: blue; background: white;", style, "最后一次设置的样式应该生效")
	})

	t.Run("GetStyle方法", func(t *testing.T) {
		input := NewInput("test", "测试")
		input.Style("test-style")

		assert.Equal(t, "test-style", input.GetStyle(), "GetStyle应该返回正确的样式")
	})
}

func TestRowsMethod(t *testing.T) {
	t.Run("Input组件 - Rows方法", func(t *testing.T) {
		input := NewInput("test", "测试")
		input.Type("textarea").Rows(5)

		result := input.Build()
		props, exists := result["props"]
		assert.True(t, exists, "props字段应该存在")

		propsMap := props.(map[string]interface{})
		assert.Equal(t, "textarea", propsMap["type"], "type应该设置为textarea")
		assert.Equal(t, 5, propsMap["rows"], "rows值应该是5")
	})

	t.Run("Textarea组件 - Rows方法", func(t *testing.T) {
		textarea := NewTextarea("test", "测试")
		textarea.Rows(8)

		result := textarea.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, 8, props["rows"], "rows值应该是8")
		assert.Equal(t, "textarea", props["type"], "type应该是textarea")
	})

	t.Run("Rows方法 - 链式调用", func(t *testing.T) {
		textarea := NewTextarea("content", "内容")
		textarea.Rows(6)
		textarea.Placeholder("请输入内容")
		textarea.Required()

		buildResult := textarea.Build()
		props := buildResult["props"].(map[string]interface{})
		assert.Equal(t, 6, props["rows"], "Rows值应该正确设置")
		assert.Equal(t, "请输入内容", props["placeholder"], "Placeholder值应该正确设置")

		// 验证Required()也生效了
		validate, exists := buildResult["validate"]
		assert.True(t, exists, "validate字段应该存在")
		assert.NotEmpty(t, validate, "应该有验证规则")
	})

	t.Run("Rows方法 - 覆盖设置", func(t *testing.T) {
		textarea := NewTextarea("test", "测试")
		textarea.Rows(3)
		textarea.Rows(10) // 覆盖之前的设置

		result := textarea.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, 10, props["rows"], "最后一次设置的rows值应该生效")
	})
}

func TestCombinedStyleAndRows(t *testing.T) {
	t.Run("组合使用Style和Rows方法", func(t *testing.T) {
		textarea := NewTextarea("description", "描述")
		textarea.Style("border: 1px solid #ccc;")
		textarea.Rows(5)

		result := textarea.Build()

		// 检查style
		style := result["style"]
		assert.Equal(t, "border: 1px solid #ccc;", style, "style应该正确设置")

		// 检查rows
		props := result["props"].(map[string]interface{})
		assert.Equal(t, 5, props["rows"], "rows应该正确设置")
	})

	t.Run("复杂链式调用", func(t *testing.T) {
		input := NewInput("email", "邮箱")
		input.Type("email")
		input.Style(map[string]interface{}{
			"width":  "100%",
			"border": "2px solid #007bff",
		})
		input.Placeholder("请输入邮箱地址")
		input.Required()
		input.Col(12)

		result := input.Build()

		// 检查所有设置的属性
		props := result["props"].(map[string]interface{})
		assert.Equal(t, "email", props["type"], "type应该是email")
		assert.Equal(t, "请输入邮箱地址", props["placeholder"], "placeholder应该正确")

		style := result["style"].(map[string]interface{})
		assert.Equal(t, "100%", style["width"], "style width应该正确")
		assert.Equal(t, "2px solid #007bff", style["border"], "style border应该正确")

		col := result["col"].(map[string]interface{})
		assert.Equal(t, 12, col["span"], "col span应该是12")

		validate, exists := result["validate"]
		assert.True(t, exists, "应该有验证规则")
		assert.NotEmpty(t, validate, "验证规则不应该为空")
	})
}
