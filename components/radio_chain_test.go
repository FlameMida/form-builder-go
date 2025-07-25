package components

import (
	"github.com/FlameMida/form-builder-go/contracts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRadioChainedCalls(t *testing.T) {
	t.Run("Radio链式调用测试", func(t *testing.T) {
		radio := NewRadio("gender", "性别")

		// 测试链式调用是否正常工作
		radio.Col(12).
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
		assert.Equal(t, "male", optionsSlice[0].Value, "第一个选项值应该是male")
		assert.Equal(t, "男性", optionsSlice[0].Label, "第一个选项标签应该是男性")
		assert.Equal(t, "female", optionsSlice[1].Value, "第二个选项值应该是female")
		assert.Equal(t, "女性", optionsSlice[1].Label, "第二个选项标签应该是女性")

		// 验证Required设置
		validate, exists := result["validate"]
		assert.True(t, exists, "validate字段应该存在")
		assert.NotEmpty(t, validate, "应该有验证规则")

		// 验证Props设置
		props, exists := result["props"]
		assert.True(t, exists, "props字段应该存在")
		propsMap := props.(map[string]interface{})
		assert.Equal(t, "large", propsMap["size"], "size应该是large")
		assert.Equal(t, false, propsMap["disabled"], "disabled应该是false")

		// 验证类型设置
		assert.Equal(t, "el-radio-group", result["type"], "type应该是el-radio-group")
	})

	t.Run("Radio Col方法单独测试", func(t *testing.T) {
		radio := NewRadio("test", "测试")

		// 测试整数参数
		radio.Col(8)
		result := radio.Build()
		col := result["col"].(map[string]interface{})
		assert.Equal(t, 8, col["span"], "span值应该是8")

		// 测试Map参数
		radio.Col(map[string]interface{}{
			"span":   12,
			"offset": 2,
		})
		result = radio.Build()
		col = result["col"].(map[string]interface{})
		assert.Equal(t, 12, col["span"], "span值应该被更新为12")
		assert.Equal(t, 2, col["offset"], "offset值应该是2")
	})

	t.Run("Radio与其他组件类似的Col行为", func(t *testing.T) {
		// 比较Radio和Input组件的Col行为应该一致
		radio := NewRadio("radio_test", "Radio测试")
		input := NewInput("input_test", "Input测试")

		// 设置相同的Col配置
		colConfig := map[string]interface{}{
			"span":   12,
			"offset": 2,
			"push":   1,
		}

		radio.Col(colConfig)
		input.Col(colConfig)

		radioResult := radio.Build()
		inputResult := input.Build()

		// Col配置应该相同
		assert.Equal(t, inputResult["col"], radioResult["col"], "Radio和Input的col配置应该一致")
	})
}
