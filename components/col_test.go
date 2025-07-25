package components

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColMethod(t *testing.T) {
	t.Run("基本Col功能 - 整数参数", func(t *testing.T) {
		input := NewInput("test", "测试")
		input.Col(12)

		result := input.Build()
		col, exists := result["col"]
		assert.True(t, exists, "col字段应该存在")

		colMap, ok := col.(map[string]interface{})
		assert.True(t, ok, "col应该是map类型")
		assert.Equal(t, 12, colMap["span"], "span值应该是12")
	})

	t.Run("Col功能 - Map参数", func(t *testing.T) {
		input := NewInput("test", "测试")
		colConfig := map[string]interface{}{
			"span":   12,
			"offset": 2,
			"push":   1,
		}
		input.Col(colConfig)

		result := input.Build()
		col := result["col"].(map[string]interface{})
		assert.Equal(t, 12, col["span"], "span值应该正确")
		assert.Equal(t, 2, col["offset"], "offset值应该正确")
		assert.Equal(t, 1, col["push"], "push值应该正确")
	})

	t.Run("Col功能 - Map参数合并", func(t *testing.T) {
		input := NewInput("test", "测试")

		// 先设置基本配置
		input.Col(map[string]interface{}{
			"span":   12,
			"offset": 2,
		})

		// 再添加更多配置
		input.Col(map[string]interface{}{
			"push": 1,
			"span": 8, // 覆盖之前的span值
		})

		result := input.Build()
		col := result["col"].(map[string]interface{})
		assert.Equal(t, 8, col["span"], "span值应该被覆盖")
		assert.Equal(t, 2, col["offset"], "offset值应该保留")
		assert.Equal(t, 1, col["push"], "push值应该添加")
	})

	t.Run("多次调用Col - 整数覆盖", func(t *testing.T) {
		input := NewInput("test", "测试")
		input.Col(6)
		input.Col(12) // 应该覆盖之前的值

		result := input.Build()
		col := result["col"].(map[string]interface{})
		assert.Equal(t, 12, col["span"], "最后一次调用的值应该生效")
	})

	t.Run("链式调用", func(t *testing.T) {
		input := NewInput("test", "测试")
		input.Col(8)
		input.Required()

		buildResult := input.Build()
		col := buildResult["col"].(map[string]interface{})
		assert.Equal(t, 8, col["span"], "Col值应该正确设置")

		// 验证Required()也生效了
		validate, exists := buildResult["validate"]
		assert.True(t, exists, "validate字段应该存在")
		assert.NotEmpty(t, validate, "应该有验证规则")
	})

	t.Run("Col功能 - 其他类型参数", func(t *testing.T) {
		input := NewInput("test", "测试")
		input.Col("12") // 字符串参数

		result := input.Build()
		col := result["col"].(map[string]interface{})
		assert.Equal(t, "12", col["span"], "非整数参数应该直接作为span值")
	})
}
