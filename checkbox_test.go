package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// checkbox_test.go Checkbox组件完整测试

// TestCheckboxCreation 测试Checkbox组件创建
func TestCheckboxCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		checkbox := NewCheckbox("hobbies", "爱好")

		assert.Equal(t, "hobbies", checkbox.GetField())
		assert.Equal(t, "checkbox", checkbox.GetType())

		data := checkbox.GetData()
		assert.Equal(t, "hobbies", data.Field)
		assert.Equal(t, "爱好", data.Title)
		assert.NotNil(t, checkbox.options)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		checkbox := NewCheckbox("test", "测试", []string{"option1"})

		data := checkbox.GetData()
		values, ok := data.Value.([]string)
		require.True(t, ok)
		assert.Len(t, values, 1)
		assert.Equal(t, "option1", values[0])
	})

	t.Run("IviewCreation", func(t *testing.T) {
		checkbox := NewIviewCheckbox("test", "测试")

		assert.Equal(t, "checkbox", checkbox.GetType())
	})
}

// TestCheckboxSetOptions 测试SetOptions方法
func TestCheckboxSetOptions(t *testing.T) {
	t.Run("SetBasicOptions", func(t *testing.T) {
		checkbox := NewCheckbox("hobbies", "爱好").SetOptions([]Option{
			{Value: "reading", Label: "阅读"},
			{Value: "music", Label: "音乐"},
			{Value: "sports", Label: "运动"},
		})

		assert.Len(t, checkbox.options, 3)
		assert.Equal(t, "reading", checkbox.options[0].Value)
		assert.Equal(t, "阅读", checkbox.options[0].Label)
	})

	t.Run("SetOptionsWithDisabled", func(t *testing.T) {
		checkbox := NewCheckbox("test", "测试").SetOptions([]Option{
			{Value: "1", Label: "选项1"},
			{Value: "2", Label: "选项2", Disabled: true},
		})

		assert.Equal(t, true, checkbox.options[1].Disabled)
	})
}

// TestCheckboxAppendOption 测试AppendOption方法
func TestCheckboxAppendOption(t *testing.T) {
	t.Run("AppendOptions", func(t *testing.T) {
		checkbox := NewCheckbox("test", "测试").
			AppendOption(Option{Value: "1", Label: "选项1"}).
			AppendOption(Option{Value: "2", Label: "选项2"}).
			AppendOption(Option{Value: "3", Label: "选项3"})

		assert.Len(t, checkbox.options, 3)
	})
}

// TestCheckboxProperties 测试所有属性方法
func TestCheckboxProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Checkbox) *Checkbox
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "Disabled",
			setup:         func(c *Checkbox) *Checkbox { return c.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
		{
			name:          "Size",
			setup:         func(c *Checkbox) *Checkbox { return c.Size("large") },
			propKey:       "size",
			expectedValue: "large",
		},
		{
			name:          "Min",
			setup:         func(c *Checkbox) *Checkbox { return c.Min(1) },
			propKey:       "min",
			expectedValue: 1,
		},
		{
			name:          "Max",
			setup:         func(c *Checkbox) *Checkbox { return c.Max(3) },
			propKey:       "max",
			expectedValue: 3,
		},
		{
			name:          "CheckedColor",
			setup:         func(c *Checkbox) *Checkbox { return c.CheckedColor("#409EFF") },
			propKey:       "checked-color",
			expectedValue: "#409EFF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkbox := NewCheckbox("test", "测试")
			checkbox = tt.setup(checkbox)

			data := checkbox.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestCheckboxChaining 测试链式调用
func TestCheckboxChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		checkbox := NewCheckbox("hobbies", "爱好", []string{"reading"}).
			SetOptions([]Option{
				{Value: "reading", Label: "阅读"},
				{Value: "music", Label: "音乐"},
			}).
			Size("large").
			Min(1).
			Max(2).
			Required()

		data := checkbox.GetData()
		assert.Equal(t, "large", data.Props["size"])
		assert.Equal(t, 1, data.Props["min"])
		assert.Equal(t, 2, data.Props["max"])
		assert.NotEmpty(t, data.Validate)
		assert.Len(t, checkbox.options, 2)
	})
}

// TestCheckboxBuild 测试Build方法
func TestCheckboxBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		checkbox := NewCheckbox("hobbies", "爱好").
			SetOptions([]Option{
				{Value: "reading", Label: "阅读"},
				{Value: "music", Label: "音乐"},
			})

		result := checkbox.Build()

		assert.Equal(t, "checkbox", result["type"])
		assert.Equal(t, "hobbies", result["field"])
		assert.Equal(t, "爱好", result["title"])

		options, ok := result["options"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, options, 2)
	})

	t.Run("BuildWithProps", func(t *testing.T) {
		checkbox := NewCheckbox("test", "测试").
			Size("large").
			Min(1).
			Max(3)

		result := checkbox.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "large", props["size"])
		assert.Equal(t, 1, props["min"])
		assert.Equal(t, 3, props["max"])
	})

	t.Run("BuildWithDisabledOption", func(t *testing.T) {
		checkbox := NewCheckbox("test", "测试").SetOptions([]Option{
			{Value: "1", Label: "选项1"},
			{Value: "2", Label: "选项2", Disabled: true},
		})

		result := checkbox.Build()
		options := result["options"].([]map[string]interface{})
		assert.Equal(t, true, options[1]["disabled"])
	})
}

// TestCheckboxEdgeCases 测试边缘情况
func TestCheckboxEdgeCases(t *testing.T) {
	t.Run("EmptyOptions", func(t *testing.T) {
		checkbox := NewCheckbox("test", "测试")
		result := checkbox.Build()

		options, exists := result["options"]
		if exists {
			assert.Len(t, options, 0)
		}
	})

	t.Run("ArrayValue", func(t *testing.T) {
		checkbox := NewCheckbox("test", "测试", []string{"1", "2", "3"})

		data := checkbox.GetData()
		values, ok := data.Value.([]string)
		require.True(t, ok)
		assert.Len(t, values, 3)
	})

	t.Run("MinMaxValidation", func(t *testing.T) {
		checkbox := NewCheckbox("test", "测试").
			Min(1).
			Max(3).
			SetOptions([]Option{
				{Value: "1", Label: "选项1"},
				{Value: "2", Label: "选项2"},
				{Value: "3", Label: "选项3"},
				{Value: "4", Label: "选项4"},
			})

		data := checkbox.GetData()
		assert.Equal(t, 1, data.Props["min"])
		assert.Equal(t, 3, data.Props["max"])
	})
}

// TestCheckboxWithValidation 测试验证功能
func TestCheckboxWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		checkbox := NewCheckbox("test", "测试").Required()

		data := checkbox.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("CustomValidation", func(t *testing.T) {
		checkbox := NewCheckbox("test", "测试").
			Validate(RequiredRule{Message: "请至少选择一项"})

		data := checkbox.GetData()
		require.Len(t, data.Validate, 1)
	})
}

// BenchmarkCheckboxCreation 性能测试
func BenchmarkCheckboxCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewCheckbox("test", "测试")
	}
}

// BenchmarkCheckboxWithOptions 性能测试
func BenchmarkCheckboxWithOptions(b *testing.B) {
	options := []Option{
		{Value: "1", Label: "选项1"},
		{Value: "2", Label: "选项2"},
		{Value: "3", Label: "选项3"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewCheckbox("test", "测试").SetOptions(options)
	}
}

// BenchmarkCheckboxBuild 性能测试
func BenchmarkCheckboxBuild(b *testing.B) {
	checkbox := NewCheckbox("hobbies", "爱好").
		SetOptions([]Option{
			{Value: "reading", Label: "阅读"},
			{Value: "music", Label: "音乐"},
		}).
		Min(1).
		Max(2).
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = checkbox.Build()
	}
}
