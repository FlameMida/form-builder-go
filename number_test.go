package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// number_test.go InputNumber组件完整测试

// TestInputNumberCreation 测试InputNumber组件创建
func TestInputNumberCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		num := NewInputNumber("age", "年龄")

		assert.Equal(t, "age", num.GetField())
		assert.Equal(t, "inputNumber", num.GetType())

		data := num.GetData()
		assert.Equal(t, "age", data.Field)
		assert.Equal(t, "年龄", data.Title)
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		num := NewInputNumber("age", "年龄", 25)

		data := num.GetData()
		assert.Equal(t, 25, data.Value)
	})

	t.Run("IviewCreation", func(t *testing.T) {
		num := NewIviewInputNumber("test", "测试")

		assert.Equal(t, "inputNumber", num.GetType())
	})

	t.Run("NumberHelper", func(t *testing.T) {
		num := Number("count", "数量", 10)

		assert.Equal(t, "count", num.GetField())
		assert.Equal(t, "inputNumber", num.GetType())
		assert.Equal(t, 10, num.GetData().Value)
	})
}

// TestInputNumberProperties 测试所有属性方法
func TestInputNumberProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*InputNumber) *InputNumber
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "Min",
			setup:         func(n *InputNumber) *InputNumber { return n.Min(0) },
			propKey:       "min",
			expectedValue: float64(0),
		},
		{
			name:          "Max",
			setup:         func(n *InputNumber) *InputNumber { return n.Max(100) },
			propKey:       "max",
			expectedValue: float64(100),
		},
		{
			name:          "Step",
			setup:         func(n *InputNumber) *InputNumber { return n.Step(5) },
			propKey:       "step",
			expectedValue: float64(5),
		},
		{
			name:          "Precision",
			setup:         func(n *InputNumber) *InputNumber { return n.Precision(2) },
			propKey:       "precision",
			expectedValue: 2,
		},
		{
			name:          "Controls",
			setup:         func(n *InputNumber) *InputNumber { return n.Controls(false) },
			propKey:       "controls",
			expectedValue: false,
		},
		{
			name:          "ControlsPosition",
			setup:         func(n *InputNumber) *InputNumber { return n.ControlsPosition("right") },
			propKey:       "controls-position",
			expectedValue: "right",
		},
		{
			name:          "Disabled",
			setup:         func(n *InputNumber) *InputNumber { return n.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
		{
			name:          "Placeholder",
			setup:         func(n *InputNumber) *InputNumber { return n.Placeholder("请输入数字") },
			propKey:       "placeholder",
			expectedValue: "请输入数字",
		},
		{
			name:          "Size",
			setup:         func(n *InputNumber) *InputNumber { return n.Size("large") },
			propKey:       "size",
			expectedValue: "large",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := NewInputNumber("test", "测试")
			num = tt.setup(num)

			data := num.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestInputNumberChaining 测试链式调用
func TestInputNumberChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		num := NewInputNumber("price", "价格", 99.99).
			Min(0).
			Max(999.99).
			Step(0.01).
			Precision(2).
			Placeholder("请输入价格").
			Size("large").
			Required()

		data := num.GetData()
		assert.Equal(t, 99.99, data.Value)
		assert.Equal(t, float64(0), data.Props["min"])
		assert.Equal(t, 999.99, data.Props["max"])
		assert.Equal(t, 0.01, data.Props["step"])
		assert.Equal(t, 2, data.Props["precision"])
		assert.Equal(t, "请输入价格", data.Props["placeholder"])
		assert.Equal(t, "large", data.Props["size"])
		assert.NotEmpty(t, data.Validate)
	})

	t.Run("ChainPreservesType", func(t *testing.T) {
		num := NewInputNumber("test", "测试").
			Min(0).
			Max(100).
			Step(1)

		// 验证链式调用返回的仍然是*InputNumber类型
		assert.IsType(t, &InputNumber{}, num)
	})
}

// TestInputNumberBuild 测试Build方法
func TestInputNumberBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		num := NewInputNumber("age", "年龄", 25)

		result := num.Build()

		assert.Equal(t, "inputNumber", result["type"])
		assert.Equal(t, "age", result["field"])
		assert.Equal(t, "年龄", result["title"])
		assert.Equal(t, 25, result["value"])
	})

	t.Run("BuildWithMinMax", func(t *testing.T) {
		num := NewInputNumber("score", "分数").
			Min(0).
			Max(100)

		result := num.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, float64(0), props["min"])
		assert.Equal(t, float64(100), props["max"])
	})

	t.Run("BuildWithStepAndPrecision", func(t *testing.T) {
		num := NewInputNumber("price", "价格").
			Step(0.01).
			Precision(2)

		result := num.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, 0.01, props["step"])
		assert.Equal(t, 2, props["precision"])
	})

	t.Run("BuildWithControls", func(t *testing.T) {
		num := NewInputNumber("test", "测试").
			Controls(false).
			ControlsPosition("right")

		result := num.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, false, props["controls"])
		assert.Equal(t, "right", props["controls-position"])
	})
}

// TestInputNumberEdgeCases 测试边缘情况
func TestInputNumberEdgeCases(t *testing.T) {
	t.Run("NegativeNumbers", func(t *testing.T) {
		num := NewInputNumber("temperature", "温度", -10).
			Min(-100).
			Max(100)

		data := num.GetData()
		assert.Equal(t, -10, data.Value)
		assert.Equal(t, float64(-100), data.Props["min"])
	})

	t.Run("FloatNumbers", func(t *testing.T) {
		num := NewInputNumber("rate", "比率", 0.5).
			Min(0.0).
			Max(1.0).
			Step(0.1).
			Precision(1)

		data := num.GetData()
		assert.Equal(t, 0.5, data.Value)
		assert.Equal(t, float64(0.0), data.Props["min"])
		assert.Equal(t, 1.0, data.Props["max"])
		assert.Equal(t, 0.1, data.Props["step"])
		assert.Equal(t, 1, data.Props["precision"])
	})

	t.Run("ZeroValue", func(t *testing.T) {
		num := NewInputNumber("count", "计数", 0)

		data := num.GetData()
		assert.Equal(t, 0, data.Value)
	})

	t.Run("LargePrecision", func(t *testing.T) {
		num := NewInputNumber("precise", "精确值").
			Precision(6).
			Step(0.000001)

		data := num.GetData()
		assert.Equal(t, 6, data.Props["precision"])
		assert.Equal(t, 0.000001, data.Props["step"])
	})
}

// TestInputNumberWithValidation 测试验证功能
func TestInputNumberWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		num := NewInputNumber("age", "年龄").Required()

		data := num.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("RangeValidation", func(t *testing.T) {
		num := NewInputNumber("age", "年龄").
			Validate(RangeRule{
				Min:     18,
				Max:     100,
				Message: "年龄必须在18-100之间",
			})

		data := num.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, "number", rule["type"])
		assert.Equal(t, float64(18), rule["min"])
		assert.Equal(t, float64(100), rule["max"])
		assert.Equal(t, "年龄必须在18-100之间", rule["message"])
	})

	t.Run("MultipleValidations", func(t *testing.T) {
		num := NewInputNumber("score", "分数").
			Required().
			Validate(RangeRule{Min: 0, Max: 100, Message: "分数必须在0-100之间"})

		data := num.GetData()
		require.Len(t, data.Validate, 2)
	})
}

// TestInputNumberInForm 测试在表单中的使用
func TestInputNumberInForm(t *testing.T) {
	t.Run("FormWithInputNumber", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInputNumber("age", "年龄").
				Min(0).
				Max(150).
				Required(),
			NewInputNumber("salary", "薪资").
				Min(0).
				Precision(2).
				Placeholder("请输入薪资"),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 2)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "inputNumber", rules[0]["type"])
		assert.Equal(t, "inputNumber", rules[1]["type"])
	})
}

// BenchmarkInputNumberCreation 性能测试
func BenchmarkInputNumberCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewInputNumber("test", "测试")
	}
}

// BenchmarkInputNumberWithProperties 性能测试
func BenchmarkInputNumberWithProperties(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewInputNumber("price", "价格").
			Min(0).
			Max(999.99).
			Step(0.01).
			Precision(2).
			Required()
	}
}

// BenchmarkInputNumberBuild 性能测试
func BenchmarkInputNumberBuild(b *testing.B) {
	num := NewInputNumber("age", "年龄", 25).
		Min(0).
		Max(150).
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = num.Build()
	}
}
