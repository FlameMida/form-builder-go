package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// rate_test.go Rate组件完整测试

// TestRateCreation 测试Rate组件创建
func TestRateCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		rate := NewRate("rating", "评分")

		assert.Equal(t, "rating", rate.GetField())
		assert.Equal(t, "rate", rate.GetType())

		data := rate.GetData()
		assert.Equal(t, "rating", data.Field)
		assert.Equal(t, "评分", data.Title)
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		rate := NewRate("score", "分数", 4)

		data := rate.GetData()
		assert.Equal(t, 4, data.Value)
	})
}

// TestRateProperties 测试所有属性方法
func TestRateProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Rate) *Rate
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "Max",
			setup:         func(r *Rate) *Rate { return r.Max(10) },
			propKey:       "max",
			expectedValue: 10,
		},
		{
			name:          "AllowHalf",
			setup:         func(r *Rate) *Rate { return r.AllowHalf(true) },
			propKey:       "allow-half",
			expectedValue: true,
		},
		{
			name:          "ShowText",
			setup:         func(r *Rate) *Rate { return r.ShowText(true) },
			propKey:       "show-text",
			expectedValue: true,
		},
		{
			name:          "ShowScore",
			setup:         func(r *Rate) *Rate { return r.ShowScore(true) },
			propKey:       "show-score",
			expectedValue: true,
		},
		{
			name:          "Disabled",
			setup:         func(r *Rate) *Rate { return r.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rate := NewRate("test", "测试")
			rate = tt.setup(rate)

			data := rate.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestRateChaining 测试链式调用
func TestRateChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		rate := NewRate("product_rating", "商品评分", 3).
			Max(10).
			AllowHalf(true).
			ShowText(true).
			ShowScore(false).
			Required()

		data := rate.GetData()
		assert.Equal(t, 3, data.Value)
		assert.Equal(t, 10, data.Props["max"])
		assert.Equal(t, true, data.Props["allow-half"])
		assert.Equal(t, true, data.Props["show-text"])
		assert.Equal(t, false, data.Props["show-score"])
		assert.NotEmpty(t, data.Validate)
	})

	t.Run("WithColors", func(t *testing.T) {
		colors := []string{"#99A9BF", "#F7BA2A", "#FF9900"}
		rate := NewRate("rating", "评分").
			Colors(colors)

		data := rate.GetData()
		colorArray, ok := data.Props["colors"].([]string)
		require.True(t, ok)
		assert.Len(t, colorArray, 3)
		assert.Equal(t, "#99A9BF", colorArray[0])
	})

	t.Run("WithTexts", func(t *testing.T) {
		texts := []string{"极差", "差", "一般", "好", "极好"}
		rate := NewRate("feedback", "反馈").
			Texts(texts).
			ShowText(true)

		data := rate.GetData()
		textArray, ok := data.Props["texts"].([]string)
		require.True(t, ok)
		assert.Len(t, textArray, 5)
		assert.Equal(t, "极差", textArray[0])
		assert.Equal(t, "极好", textArray[4])
	})
}

// TestRateBuild 测试Build方法
func TestRateBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		rate := NewRate("satisfaction", "满意度", 5)

		result := rate.Build()

		assert.Equal(t, "rate", result["type"])
		assert.Equal(t, "satisfaction", result["field"])
		assert.Equal(t, "满意度", result["title"])
		assert.Equal(t, 5, result["value"])
	})

	t.Run("BuildWithMax", func(t *testing.T) {
		rate := NewRate("rating", "评分").
			Max(10)

		result := rate.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, 10, props["max"])
	})

	t.Run("BuildWithAllowHalf", func(t *testing.T) {
		rate := NewRate("score", "分数").
			AllowHalf(true)

		result := rate.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, true, props["allow-half"])
	})

	t.Run("BuildWithColors", func(t *testing.T) {
		colors := []string{"#F56C6C", "#E6A23C", "#67C23A"}
		rate := NewRate("quality", "质量").
			Colors(colors)

		result := rate.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, colors, props["colors"])
	})
}

// TestRateEdgeCases 测试边缘情况
func TestRateEdgeCases(t *testing.T) {
	t.Run("MaxValue10", func(t *testing.T) {
		rate := NewRate("rating", "评分").
			Max(10)

		data := rate.GetData()
		assert.Equal(t, 10, data.Props["max"])
	})

	t.Run("HalfStar", func(t *testing.T) {
		rate := NewRate("score", "分数", 3.5).
			AllowHalf(true)

		data := rate.GetData()
		assert.Equal(t, 3.5, data.Value)
		assert.Equal(t, true, data.Props["allow-half"])
	})

	t.Run("ZeroRating", func(t *testing.T) {
		rate := NewRate("rating", "评分", 0)

		data := rate.GetData()
		assert.Equal(t, 0, data.Value)
	})

	t.Run("ShowTextAndScore", func(t *testing.T) {
		rate := NewRate("rating", "评分").
			ShowText(true).
			ShowScore(true)

		data := rate.GetData()
		assert.Equal(t, true, data.Props["show-text"])
		assert.Equal(t, true, data.Props["show-score"])
	})

	t.Run("DisabledWithValue", func(t *testing.T) {
		rate := NewRate("display_rating", "展示评分", 4).
			Disabled(true)

		data := rate.GetData()
		assert.Equal(t, 4, data.Value)
		assert.Equal(t, true, data.Props["disabled"])
	})
}

// TestRateWithValidation 测试验证功能
func TestRateWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		rate := NewRate("rating", "评分").Required()

		data := rate.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("RangeValidation", func(t *testing.T) {
		rate := NewRate("score", "分数").
			Validate(RangeRule{
				Min:     1,
				Max:     5,
				Message: "评分必须在1-5之间",
			})

		data := rate.GetData()
		require.Len(t, data.Validate, 1)
	})
}

// TestRateInForm 测试在表单中的使用
func TestRateInForm(t *testing.T) {
	t.Run("FormWithRate", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewRate("product_quality", "商品质量").
				Max(5).
				AllowHalf(true).
				ShowText(true).
				Required(),
			NewRate("service_quality", "服务质量").
				Max(5).
				ShowScore(true),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 2)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "rate", rules[0]["type"])
		assert.Equal(t, "rate", rules[1]["type"])
	})

	t.Run("RateWithControl", func(t *testing.T) {
		rate := NewRate("satisfaction", "满意度", 0).
			Max(5).
			Control([]ControlRule{
				{
					Value: 3,
					Rule: []Component{
						NewInput("feedback", "请提供建议").Required(),
					},
				},
			})

		data := rate.GetData()
		assert.Len(t, data.Control, 1)
		assert.Len(t, data.Control[0].Rule, 1)
	})
}

// BenchmarkRateCreation 性能测试
func BenchmarkRateCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewRate("test", "测试")
	}
}

// BenchmarkRateWithProperties 性能测试
func BenchmarkRateWithProperties(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewRate("rating", "评分", 4).
			Max(10).
			AllowHalf(true).
			ShowText(true).
			Required()
	}
}

// BenchmarkRateBuild 性能测试
func BenchmarkRateBuild(b *testing.B) {
	colors := []string{"#F56C6C", "#E6A23C", "#67C23A"}
	rate := NewRate("quality", "质量").
		Max(5).
		Colors(colors).
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rate.Build()
	}
}
