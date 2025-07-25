package components

import (
	"github.com/FlameMida/form-builder-go/contracts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceCompatibility(t *testing.T) {
	t.Run("所有组件都能正确实现Component接口", func(t *testing.T) {
		// 创建各种组件
		input := NewInput("name", "姓名")
		textarea := NewTextarea("content", "内容")
		switchComp := NewSwitch("enabled", "启用")
		radio := NewRadio("gender", "性别")
		select_ := NewSelect("city", "城市")
		checkbox := NewCheckbox("hobbies", "爱好")

		// 测试所有组件都能被存储为Component接口类型
		components := []contracts.Component{
			input,
			textarea,
			switchComp,
			radio,
			select_,
			checkbox,
		}

		// 测试接口方法调用
		for i, component := range components {
			// 测试基本方法
			assert.NotEmpty(t, component.Field(), "Component %d should have field", i)
			assert.NotEmpty(t, component.Title(), "Component %d should have title", i)

			// 测试Style方法
			styleResult := component.Style("color: red")
			assert.NotNil(t, styleResult, "Style method should return non-nil for component %d", i)
			assert.Implements(t, (*contracts.Component)(nil), styleResult, "Style result should implement Component interface for component %d", i)

			// 测试SetValue方法
			valueResult := component.SetValue("test")
			assert.NotNil(t, valueResult, "SetValue method should return non-nil for component %d", i)
			assert.Implements(t, (*contracts.Component)(nil), valueResult, "SetValue result should implement Component interface for component %d", i)

			// 测试Build方法
			buildResult := component.Build()
			assert.NotNil(t, buildResult, "Build method should return non-nil for component %d", i)
			assert.IsType(t, map[string]interface{}{}, buildResult, "Build should return map for component %d", i)
		}
	})

	t.Run("具体类型的链式调用仍然有效", func(t *testing.T) {
		// 测试具体类型的Col方法仍然有效
		radio := NewRadio("gender", "性别").
			Col(12).
			AddOption("male", "男性").
			AddOption("female", "女性").
			Required().
			Size("large")

		result := radio.Build()

		// 验证Col设置
		col, exists := result["col"]
		assert.True(t, exists, "col字段应该存在")
		colMap := col.(map[string]interface{})
		assert.Equal(t, 12, colMap["span"], "span值应该是12")

		// 验证Options设置
		options, exists := result["options"]
		assert.True(t, exists, "options字段应该存在")
		optionsSlice := options.([]contracts.Option)
		assert.Len(t, optionsSlice, 2, "应该有2个选项")
	})

	t.Run("在FormBuilder中的接口兼容性", func(t *testing.T) {
		// 模拟FormBuilder中的使用场景
		components := []contracts.Component{
			NewInput("name", "姓名").Col(12),
			NewRadio("gender", "性别").Col(6),
			NewSelect("city", "城市").Col(8),
		}

		// 验证所有组件都可以正常使用
		for i, component := range components {
			result := component.Build()
			assert.NotEmpty(t, result["field"], "Component %d should have field", i)
			assert.NotEmpty(t, result["title"], "Component %d should have title", i)
			// 注意：现在Col不在接口中，但组件仍然可以使用Col方法
		}
	})

	t.Run("AsComponent方法兼容性", func(t *testing.T) {
		// 测试所有组件的AsComponent方法
		input := NewInput("test", "测试")
		textarea := NewTextarea("test", "测试")
		radio := NewRadio("test", "测试")

		// 通过AsComponent方法获取接口类型
		var component1 contracts.Component = input.AsComponent()
		var component2 contracts.Component = textarea.AsComponent()
		var component3 contracts.Component = radio.AsComponent()

		// 验证接口方法可用
		assert.Equal(t, "test", component1.Field())
		assert.Equal(t, "test", component2.Field())
		assert.Equal(t, "test", component3.Field())
	})
}
