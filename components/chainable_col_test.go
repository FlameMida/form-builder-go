package components

import (
	"github.com/FlameMida/form-builder-go/contracts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChainableColDesign(t *testing.T) {
	t.Run("Radio完整链式调用测试", func(t *testing.T) {
		// 测试完整的链式调用，仿照PHP模式
		radio := NewRadio("gender", "性别").
			Col(12).
			AddOption("male", "男性").
			AddOption("female", "女性").
			Required().
			Size("large").
			Disabled(false)

		result := radio.Build()

		// 验证Col设置
		col, exists := result["col"]
		assert.True(t, exists, "col字段应该存在")
		colMap, ok := col.(map[string]interface{})
		assert.True(t, ok, "col应该是map类型")
		assert.Equal(t, 12, colMap["span"], "span值应该是12")

		// 验证Options设置
		options, exists := result["options"]
		assert.True(t, exists, "options字段应该存在")
		optionsSlice, ok := options.([]contracts.Option)
		assert.True(t, ok, "options应该是Option切片类型")
		assert.Len(t, optionsSlice, 2, "应该有2个选项")

		// 验证Props设置
		props, exists := result["props"]
		assert.True(t, exists, "props字段应该存在")
		propsMap := props.(map[string]interface{})
		assert.Equal(t, "large", propsMap["size"], "size应该是large")
		assert.Equal(t, false, propsMap["disabled"], "disabled应该是false")

		// 验证类型设置
		assert.Equal(t, "el-radio-group", result["type"], "type应该是el-radio-group")
	})

	t.Run("Input完整链式调用测试", func(t *testing.T) {
		// 测试Input组件的链式调用
		input := NewInput("email", "邮箱").
			Col(map[string]interface{}{
				"span":   12,
				"offset": 2,
			}).
			Type("email").
			Placeholder("请输入邮箱地址").
			Required().
			Clearable(true).
			Size("large")

		result := input.Build()

		// 验证Col设置
		col := result["col"].(map[string]interface{})
		assert.Equal(t, 12, col["span"], "span应该是12")
		assert.Equal(t, 2, col["offset"], "offset应该是2")

		// 验证Props设置
		props := result["props"].(map[string]interface{})
		assert.Equal(t, "email", props["type"], "type应该是email")
		assert.Equal(t, "请输入邮箱地址", props["placeholder"], "placeholder应该正确")
		assert.Equal(t, true, props["clearable"], "clearable应该是true")
		assert.Equal(t, "large", props["size"], "size应该是large")
	})

	t.Run("Textarea完整链式调用测试", func(t *testing.T) {
		// 测试Textarea组件的链式调用
		textarea := NewTextarea("content", "内容").
			Col(8).
			Rows(6).
			Placeholder("请输入内容").
			Required().
			AutoSize(2, 10)

		result := textarea.Build()

		// 验证Col设置
		col := result["col"].(map[string]interface{})
		assert.Equal(t, 8, col["span"], "span应该是8")

		// 验证Props设置
		props := result["props"].(map[string]interface{})
		assert.Equal(t, 6, props["rows"], "rows应该是6")
		assert.Equal(t, "请输入内容", props["placeholder"], "placeholder应该正确")

		// 验证AutoSize设置
		autosize, exists := props["autosize"]
		assert.True(t, exists, "autosize应该存在")
		autosizeMap := autosize.(map[string]interface{})
		assert.Equal(t, 2, autosizeMap["minRows"], "minRows应该是2")
		assert.Equal(t, 10, autosizeMap["maxRows"], "maxRows应该是10")
	})

	t.Run("Switch完整链式调用测试", func(t *testing.T) {
		// 测试Switch组件的链式调用
		switchComp := NewSwitch("enabled", "启用状态").
			Col(6).
			ActiveText("开启").
			InactiveText("关闭").
			ActiveValue(1).
			InactiveValue(0).
			Required()

		result := switchComp.Build()

		// 验证Col设置
		col := result["col"].(map[string]interface{})
		assert.Equal(t, 6, col["span"], "span应该是6")

		// 验证Props设置
		props := result["props"].(map[string]interface{})
		assert.Equal(t, "开启", props["active-text"], "active-text应该正确")
		assert.Equal(t, "关闭", props["inactive-text"], "inactive-text应该正确")
		assert.Equal(t, 1, props["active-value"], "active-value应该是1")
		assert.Equal(t, 0, props["inactive-value"], "inactive-value应该是0")
	})

	t.Run("Select完整链式调用测试", func(t *testing.T) {
		// 测试Select组件的链式调用
		selectComp := NewSelect("city", "城市").
			Col(10).
			AddOption("bj", "北京").
			AddOption("sh", "上海").
			AddOption("gz", "广州").
			Multiple(true).
			Clearable(true).
			Required()

		result := selectComp.Build()

		// 验证Col设置
		col := result["col"].(map[string]interface{})
		assert.Equal(t, 10, col["span"], "span应该是10")

		// 验证Options设置
		options, exists := result["options"]
		assert.True(t, exists, "options字段应该存在")
		optionsSlice := options.([]contracts.Option)
		assert.Len(t, optionsSlice, 3, "应该有3个选项")

		// 验证Props设置
		props := result["props"].(map[string]interface{})
		assert.Equal(t, true, props["multiple"], "multiple应该是true")
		assert.Equal(t, true, props["clearable"], "clearable应该是true")
	})

	t.Run("Checkbox完整链式调用测试", func(t *testing.T) {
		// 测试Checkbox组件的链式调用
		checkbox := NewCheckbox("hobbies", "爱好").
			Col(12).
			AddOption("reading", "阅读").
			AddOption("music", "音乐").
			AddOption("sports", "运动").
			Min(1).
			Max(2).
			Required()

		result := checkbox.Build()

		// 验证Col设置
		col := result["col"].(map[string]interface{})
		assert.Equal(t, 12, col["span"], "span应该是12")

		// 验证Options设置
		options, exists := result["options"]
		assert.True(t, exists, "options字段应该存在")
		optionsSlice := options.([]contracts.Option)
		assert.Len(t, optionsSlice, 3, "应该有3个选项")

		// 验证Props设置
		props := result["props"].(map[string]interface{})
		assert.Equal(t, 1, props["min"], "min应该是1")
		assert.Equal(t, 2, props["max"], "max应该是2")
	})
}
