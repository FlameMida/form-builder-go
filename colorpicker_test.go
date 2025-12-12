package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// colorpicker_test.go ColorPicker组件完整测试

// TestColorPickerCreation 测试ColorPicker组件创建
func TestColorPickerCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		cp := NewColorPicker("theme_color", "主题颜色")

		assert.Equal(t, "theme_color", cp.GetField())
		assert.Equal(t, "colorPicker", cp.GetType())

		data := cp.GetData()
		assert.Equal(t, "theme_color", data.Field)
		assert.Equal(t, "主题颜色", data.Title)
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		cp := NewColorPicker("color", "颜色", "#409EFF")

		data := cp.GetData()
		assert.Equal(t, "#409EFF", data.Value)
	})
}

// TestColorPickerProperties 测试所有属性方法
func TestColorPickerProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*ColorPicker) *ColorPicker
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "ShowAlpha",
			setup:         func(c *ColorPicker) *ColorPicker { return c.ShowAlpha(true) },
			propKey:       "show-alpha",
			expectedValue: true,
		},
		{
			name:          "ColorFormatHex",
			setup:         func(c *ColorPicker) *ColorPicker { return c.ColorFormat("hex") },
			propKey:       "color-format",
			expectedValue: "hex",
		},
		{
			name:          "ColorFormatRgb",
			setup:         func(c *ColorPicker) *ColorPicker { return c.ColorFormat("rgb") },
			propKey:       "color-format",
			expectedValue: "rgb",
		},
		{
			name:          "ColorFormatHsl",
			setup:         func(c *ColorPicker) *ColorPicker { return c.ColorFormat("hsl") },
			propKey:       "color-format",
			expectedValue: "hsl",
		},
		{
			name:          "ColorFormatHsv",
			setup:         func(c *ColorPicker) *ColorPicker { return c.ColorFormat("hsv") },
			propKey:       "color-format",
			expectedValue: "hsv",
		},
		{
			name:          "Disabled",
			setup:         func(c *ColorPicker) *ColorPicker { return c.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
		{
			name:          "Size",
			setup:         func(c *ColorPicker) *ColorPicker { return c.Size("large") },
			propKey:       "size",
			expectedValue: "large",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := NewColorPicker("test", "测试")
			cp = tt.setup(cp)

			data := cp.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestColorPickerChaining 测试链式调用
func TestColorPickerChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		predefine := []string{
			"#ff4500",
			"#ff8c00",
			"#ffd700",
			"#90ee90",
			"#00ced1",
			"#1e90ff",
			"#c71585",
		}

		cp := NewColorPicker("brand_color", "品牌颜色", "#409EFF").
			ShowAlpha(true).
			ColorFormat("hex").
			Predefine(predefine).
			Size("medium").
			Required()

		data := cp.GetData()
		assert.Equal(t, "#409EFF", data.Value)
		assert.Equal(t, true, data.Props["show-alpha"])
		assert.Equal(t, "hex", data.Props["color-format"])
		assert.Equal(t, "medium", data.Props["size"])
		assert.NotEmpty(t, data.Validate)

		colors, ok := data.Props["predefine"].([]string)
		require.True(t, ok)
		assert.Len(t, colors, 7)
	})

	t.Run("RgbFormatChain", func(t *testing.T) {
		cp := NewColorPicker("bg_color", "背景颜色").
			ColorFormat("rgb").
			ShowAlpha(false)

		data := cp.GetData()
		assert.Equal(t, "rgb", data.Props["color-format"])
		assert.Equal(t, false, data.Props["show-alpha"])
	})
}

// TestColorPickerBuild 测试Build方法
func TestColorPickerBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		cp := NewColorPicker("color", "颜色", "#67C23A")

		result := cp.Build()

		assert.Equal(t, "colorPicker", result["type"])
		assert.Equal(t, "color", result["field"])
		assert.Equal(t, "颜色", result["title"])
		assert.Equal(t, "#67C23A", result["value"])
	})

	t.Run("BuildWithShowAlpha", func(t *testing.T) {
		cp := NewColorPicker("color", "颜色").
			ShowAlpha(true)

		result := cp.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, true, props["show-alpha"])
	})

	t.Run("BuildWithColorFormat", func(t *testing.T) {
		cp := NewColorPicker("color", "颜色").
			ColorFormat("rgb")

		result := cp.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, "rgb", props["color-format"])
	})

	t.Run("BuildWithPredefine", func(t *testing.T) {
		predefine := []string{"#ff0000", "#00ff00", "#0000ff"}
		cp := NewColorPicker("color", "颜色").
			Predefine(predefine)

		result := cp.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, predefine, props["predefine"])
	})
}

// TestColorPickerEdgeCases 测试边缘情况
func TestColorPickerEdgeCases(t *testing.T) {
	t.Run("TransparentColor", func(t *testing.T) {
		cp := NewColorPicker("transparent", "透明色", "rgba(0,0,0,0)").
			ShowAlpha(true).
			ColorFormat("rgb")

		data := cp.GetData()
		assert.Equal(t, "rgba(0,0,0,0)", data.Value)
		assert.Equal(t, true, data.Props["show-alpha"])
		assert.Equal(t, "rgb", data.Props["color-format"])
	})

	t.Run("HslFormat", func(t *testing.T) {
		cp := NewColorPicker("hsl_color", "HSL颜色").
			ColorFormat("hsl")

		data := cp.GetData()
		assert.Equal(t, "hsl", data.Props["color-format"])
	})

	t.Run("HsvFormat", func(t *testing.T) {
		cp := NewColorPicker("hsv_color", "HSV颜色").
			ColorFormat("hsv")

		data := cp.GetData()
		assert.Equal(t, "hsv", data.Props["color-format"])
	})

	t.Run("EmptyPredefine", func(t *testing.T) {
		cp := NewColorPicker("color", "颜色").
			Predefine([]string{})

		data := cp.GetData()
		colors, ok := data.Props["predefine"].([]string)
		require.True(t, ok)
		assert.Len(t, colors, 0)
	})

	t.Run("DisabledWithValue", func(t *testing.T) {
		cp := NewColorPicker("readonly_color", "只读颜色", "#E6A23C").
			Disabled(true)

		data := cp.GetData()
		assert.Equal(t, "#E6A23C", data.Value)
		assert.Equal(t, true, data.Props["disabled"])
	})

	t.Run("AllSizes", func(t *testing.T) {
		sizes := []string{"large", "medium", "small", "mini"}
		for _, size := range sizes {
			cp := NewColorPicker("color", "颜色").Size(size)
			data := cp.GetData()
			assert.Equal(t, size, data.Props["size"])
		}
	})
}

// TestColorPickerWithValidation 测试验证功能
func TestColorPickerWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		cp := NewColorPicker("color", "颜色").Required()

		data := cp.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("PatternValidation", func(t *testing.T) {
		cp := NewColorPicker("hex_color", "十六进制颜色").
			Validate(PatternRule{
				Pattern: "^#[0-9A-Fa-f]{6}$",
				Message: "请输入有效的十六进制颜色",
			})

		data := cp.GetData()
		require.Len(t, data.Validate, 1)
	})
}

// TestColorPickerInForm 测试在表单中的使用
func TestColorPickerInForm(t *testing.T) {
	t.Run("FormWithColorPicker", func(t *testing.T) {
		predefineColors := []string{
			"#409EFF", "#67C23A", "#E6A23C", "#F56C6C", "#909399",
		}

		form := NewElmForm("/submit", []Component{
			NewColorPicker("primary_color", "主色").
				Predefine(predefineColors).
				Required(),
			NewColorPicker("secondary_color", "辅色").
				ShowAlpha(true).
				ColorFormat("rgb"),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 2)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "colorPicker", rules[0]["type"])
		assert.Equal(t, "colorPicker", rules[1]["type"])
	})

	t.Run("ColorPickerWithControl", func(t *testing.T) {
		cp := NewColorPicker("use_custom_color", "使用自定义颜色", "#409EFF").
			Control([]ControlRule{
				{
					Value: "#409EFF",
					Rule: []Component{
						NewInput("custom_color_name", "颜色名称").Required(),
					},
				},
			})

		data := cp.GetData()
		assert.Len(t, data.Control, 1)
		assert.Len(t, data.Control[0].Rule, 1)
	})
}

// BenchmarkColorPickerCreation 性能测试
func BenchmarkColorPickerCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewColorPicker("test", "测试")
	}
}

// BenchmarkColorPickerWithProperties 性能测试
func BenchmarkColorPickerWithProperties(b *testing.B) {
	predefine := []string{"#ff0000", "#00ff00", "#0000ff"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewColorPicker("color", "颜色", "#409EFF").
			ShowAlpha(true).
			ColorFormat("hex").
			Predefine(predefine).
			Required()
	}
}

// BenchmarkColorPickerBuild 性能测试
func BenchmarkColorPickerBuild(b *testing.B) {
	predefine := []string{"#ff0000", "#00ff00", "#0000ff"}
	cp := NewColorPicker("color", "颜色", "#409EFF").
		ShowAlpha(true).
		Predefine(predefine).
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cp.Build()
	}
}
