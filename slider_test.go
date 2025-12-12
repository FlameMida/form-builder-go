package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// slider_test.go Slider组件完整测试

// TestSliderCreation 测试Slider组件创建
func TestSliderCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		slider := NewSlider("volume", "音量")

		assert.Equal(t, "volume", slider.GetField())
		assert.Equal(t, "slider", slider.GetType())

		data := slider.GetData()
		assert.Equal(t, "volume", data.Field)
		assert.Equal(t, "音量", data.Title)
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		slider := NewSlider("brightness", "亮度", 50)

		data := slider.GetData()
		assert.Equal(t, 50, data.Value)
	})
}

// TestSliderProperties 测试所有属性方法
func TestSliderProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Slider) *Slider
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "Min",
			setup:         func(s *Slider) *Slider { return s.Min(0) },
			propKey:       "min",
			expectedValue: float64(0),
		},
		{
			name:          "Max",
			setup:         func(s *Slider) *Slider { return s.Max(100) },
			propKey:       "max",
			expectedValue: float64(100),
		},
		{
			name:          "Step",
			setup:         func(s *Slider) *Slider { return s.Step(5) },
			propKey:       "step",
			expectedValue: float64(5),
		},
		{
			name:          "Range",
			setup:         func(s *Slider) *Slider { return s.Range(true) },
			propKey:       "range",
			expectedValue: true,
		},
		{
			name:          "ShowStops",
			setup:         func(s *Slider) *Slider { return s.ShowStops(true) },
			propKey:       "show-stops",
			expectedValue: true,
		},
		{
			name:          "ShowInput",
			setup:         func(s *Slider) *Slider { return s.ShowInput(true) },
			propKey:       "show-input",
			expectedValue: true,
		},
		{
			name:          "Disabled",
			setup:         func(s *Slider) *Slider { return s.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slider := NewSlider("test", "测试")
			slider = tt.setup(slider)

			data := slider.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestSliderChaining 测试链式调用
func TestSliderChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		slider := NewSlider("score", "评分", 75).
			Min(0).
			Max(100).
			Step(5).
			ShowStops(true).
			ShowInput(true).
			Required()

		data := slider.GetData()
		assert.Equal(t, 75, data.Value)
		assert.Equal(t, float64(0), data.Props["min"])
		assert.Equal(t, float64(100), data.Props["max"])
		assert.Equal(t, float64(5), data.Props["step"])
		assert.Equal(t, true, data.Props["show-stops"])
		assert.Equal(t, true, data.Props["show-input"])
		assert.NotEmpty(t, data.Validate)
	})

	t.Run("RangeSliderChain", func(t *testing.T) {
		slider := NewSlider("price_range", "价格范围").
			Range(true).
			Min(0).
			Max(1000).
			Step(10)

		data := slider.GetData()
		assert.Equal(t, true, data.Props["range"])
		assert.Equal(t, float64(0), data.Props["min"])
		assert.Equal(t, float64(1000), data.Props["max"])
		assert.Equal(t, float64(10), data.Props["step"])
	})
}

// TestSliderBuild 测试Build方法
func TestSliderBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		slider := NewSlider("volume", "音量", 50)

		result := slider.Build()

		assert.Equal(t, "slider", result["type"])
		assert.Equal(t, "volume", result["field"])
		assert.Equal(t, "音量", result["title"])
		assert.Equal(t, 50, result["value"])
	})

	t.Run("BuildWithMinMax", func(t *testing.T) {
		slider := NewSlider("temperature", "温度").
			Min(18).
			Max(30)

		result := slider.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, float64(18), props["min"])
		assert.Equal(t, float64(30), props["max"])
	})

	t.Run("BuildWithStep", func(t *testing.T) {
		slider := NewSlider("test", "测试").
			Step(0.5)

		result := slider.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, 0.5, props["step"])
	})

	t.Run("BuildWithShowStops", func(t *testing.T) {
		slider := NewSlider("test", "测试").
			Step(10).
			ShowStops(true)

		result := slider.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, true, props["show-stops"])
	})
}

// TestSliderEdgeCases 测试边缘情况
func TestSliderEdgeCases(t *testing.T) {
	t.Run("RangeMode", func(t *testing.T) {
		slider := NewSlider("range", "范围").
			Range(true)

		data := slider.GetData()
		assert.Equal(t, true, data.Props["range"])
	})

	t.Run("FloatStep", func(t *testing.T) {
		slider := NewSlider("precision", "精度").
			Min(0).
			Max(1).
			Step(0.01)

		data := slider.GetData()
		assert.Equal(t, float64(0), data.Props["min"])
		assert.Equal(t, float64(1), data.Props["max"])
		assert.Equal(t, 0.01, data.Props["step"])
	})

	t.Run("NegativeRange", func(t *testing.T) {
		slider := NewSlider("temp", "温度").
			Min(-20).
			Max(50)

		data := slider.GetData()
		assert.Equal(t, float64(-20), data.Props["min"])
		assert.Equal(t, float64(50), data.Props["max"])
	})

	t.Run("ShowInputWithRange", func(t *testing.T) {
		slider := NewSlider("test", "测试").
			Range(true).
			ShowInput(true)

		data := slider.GetData()
		assert.Equal(t, true, data.Props["range"])
		assert.Equal(t, true, data.Props["show-input"])
	})
}

// TestSliderWithValidation 测试验证功能
func TestSliderWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		slider := NewSlider("rating", "评分").Required()

		data := slider.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("RangeValidation", func(t *testing.T) {
		slider := NewSlider("score", "分数").
			Validate(RangeRule{
				Min:     0,
				Max:     100,
				Message: "分数必须在0-100之间",
			})

		data := slider.GetData()
		require.Len(t, data.Validate, 1)
	})
}

// TestSliderInForm 测试在表单中的使用
func TestSliderInForm(t *testing.T) {
	t.Run("FormWithSlider", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewSlider("satisfaction", "满意度").
				Min(0).
				Max(10).
				ShowStops(true).
				Required(),
			NewSlider("price_range", "价格范围").
				Range(true).
				Min(0).
				Max(1000),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 2)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "slider", rules[0]["type"])
		assert.Equal(t, "slider", rules[1]["type"])
	})
}

// BenchmarkSliderCreation 性能测试
func BenchmarkSliderCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSlider("test", "测试")
	}
}

// BenchmarkSliderWithProperties 性能测试
func BenchmarkSliderWithProperties(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSlider("score", "分数", 75).
			Min(0).
			Max(100).
			Step(5).
			ShowStops(true).
			Required()
	}
}

// BenchmarkSliderBuild 性能测试
func BenchmarkSliderBuild(b *testing.B) {
	slider := NewSlider("volume", "音量", 50).
		Min(0).
		Max(100).
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = slider.Build()
	}
}
