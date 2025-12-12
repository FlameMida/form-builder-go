package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// timepicker_test.go TimePicker组件完整测试

// TestTimePickerCreation 测试TimePicker组件创建
func TestTimePickerCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		tp := NewTimePicker("time", "时间")

		assert.Equal(t, "time", tp.GetField())
		assert.Equal(t, "timePicker", tp.GetType())

		data := tp.GetData()
		assert.Equal(t, "time", data.Field)
		assert.Equal(t, "时间", data.Title)
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		tp := NewTimePicker("time", "时间", "10:00:00")

		data := tp.GetData()
		assert.Equal(t, "10:00:00", data.Value)
	})
}

// TestTimePickerProperties 测试所有属性方法
func TestTimePickerProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*TimePicker) *TimePicker
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "IsRange",
			setup:         func(t *TimePicker) *TimePicker { return t.IsRange(true) },
			propKey:       "is-range",
			expectedValue: true,
		},
		{
			name:          "Format",
			setup:         func(t *TimePicker) *TimePicker { return t.Format("HH:mm:ss") },
			propKey:       "format",
			expectedValue: "HH:mm:ss",
		},
		{
			name:          "ValueFormat",
			setup:         func(t *TimePicker) *TimePicker { return t.ValueFormat("HH:mm:ss") },
			propKey:       "value-format",
			expectedValue: "HH:mm:ss",
		},
		{
			name:          "Placeholder",
			setup:         func(t *TimePicker) *TimePicker { return t.Placeholder("请选择时间") },
			propKey:       "placeholder",
			expectedValue: "请选择时间",
		},
		{
			name:          "Clearable",
			setup:         func(t *TimePicker) *TimePicker { return t.Clearable(true) },
			propKey:       "clearable",
			expectedValue: true,
		},
		{
			name:          "Disabled",
			setup:         func(t *TimePicker) *TimePicker { return t.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
		{
			name:          "Size",
			setup:         func(t *TimePicker) *TimePicker { return t.Size("large") },
			propKey:       "size",
			expectedValue: "large",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := NewTimePicker("test", "测试")
			tp = tt.setup(tp)

			data := tp.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestTimePickerChaining 测试链式调用
func TestTimePickerChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		tp := NewTimePicker("work_time", "工作时间", "09:00:00").
			Format("HH:mm:ss").
			ValueFormat("HH:mm:ss").
			Placeholder("请选择工作时间").
			Clearable(true).
			Size("large").
			Required()

		data := tp.GetData()
		assert.Equal(t, "09:00:00", data.Value)
		assert.Equal(t, "HH:mm:ss", data.Props["format"])
		assert.Equal(t, "HH:mm:ss", data.Props["value-format"])
		assert.Equal(t, "请选择工作时间", data.Props["placeholder"])
		assert.Equal(t, true, data.Props["clearable"])
		assert.Equal(t, "large", data.Props["size"])
		assert.NotEmpty(t, data.Validate)
	})

	t.Run("TimeRangeChain", func(t *testing.T) {
		tp := NewTimePicker("time_range", "时间范围").
			IsRange(true).
			Placeholder("请选择时间范围")

		data := tp.GetData()
		assert.Equal(t, true, data.Props["is-range"])
		assert.Equal(t, "请选择时间范围", data.Props["placeholder"])
	})
}

// TestTimePickerBuild 测试Build方法
func TestTimePickerBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		tp := NewTimePicker("start_time", "开始时间", "08:00:00")

		result := tp.Build()

		assert.Equal(t, "timePicker", result["type"])
		assert.Equal(t, "start_time", result["field"])
		assert.Equal(t, "开始时间", result["title"])
		assert.Equal(t, "08:00:00", result["value"])
	})

	t.Run("BuildWithFormat", func(t *testing.T) {
		tp := NewTimePicker("time", "时间").
			Format("HH:mm").
			ValueFormat("HH:mm:ss")

		result := tp.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "HH:mm", props["format"])
		assert.Equal(t, "HH:mm:ss", props["value-format"])
	})

	t.Run("BuildWithRange", func(t *testing.T) {
		tp := NewTimePicker("time_range", "时间范围").
			IsRange(true)

		result := tp.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, true, props["is-range"])
	})
}

// TestTimePickerEdgeCases 测试边缘情况
func TestTimePickerEdgeCases(t *testing.T) {
	t.Run("RangeMode", func(t *testing.T) {
		tp := NewTimePicker("range", "时间范围").
			IsRange(true)

		data := tp.GetData()
		assert.Equal(t, true, data.Props["is-range"])
	})

	t.Run("CustomFormat", func(t *testing.T) {
		tp := NewTimePicker("time", "时间").
			Format("HH时mm分")

		data := tp.GetData()
		assert.Equal(t, "HH时mm分", data.Props["format"])
	})

	t.Run("SecondsPrecision", func(t *testing.T) {
		tp := NewTimePicker("precise_time", "精确时间").
			Format("HH:mm:ss").
			ValueFormat("HH:mm:ss")

		data := tp.GetData()
		assert.Equal(t, "HH:mm:ss", data.Props["format"])
	})
}

// TestTimePickerWithValidation 测试验证功能
func TestTimePickerWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		tp := NewTimePicker("time", "时间").Required()

		data := tp.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})
}

// TestTimePickerInForm 测试在表单中的使用
func TestTimePickerInForm(t *testing.T) {
	t.Run("FormWithTimePicker", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewTimePicker("start_time", "开始时间").Required(),
			NewTimePicker("end_time", "结束时间").Required(),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 2)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "timePicker", rules[0]["type"])
		assert.Equal(t, "timePicker", rules[1]["type"])
	})
}

// BenchmarkTimePickerCreation 性能测试
func BenchmarkTimePickerCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewTimePicker("test", "测试")
	}
}

// BenchmarkTimePickerWithProperties 性能测试
func BenchmarkTimePickerWithProperties(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewTimePicker("time", "时间").
			Format("HH:mm:ss").
			ValueFormat("HH:mm:ss").
			Required()
	}
}

// BenchmarkTimePickerBuild 性能测试
func BenchmarkTimePickerBuild(b *testing.B) {
	tp := NewTimePicker("time", "时间", "08:00:00").
		Format("HH:mm:ss").
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = tp.Build()
	}
}
