package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// datepicker_test.go DatePicker组件完整测试

// TestDatePickerCreation 测试DatePicker组件创建
func TestDatePickerCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		dp := NewDatePicker("birth_date", "出生日期")

		assert.Equal(t, "birth_date", dp.GetField())
		assert.Equal(t, "datePicker", dp.GetType())

		data := dp.GetData()
		assert.Equal(t, "birth_date", data.Field)
		assert.Equal(t, "出生日期", data.Title)
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		dp := NewDatePicker("date", "日期", "2024-01-01")

		data := dp.GetData()
		assert.Equal(t, "2024-01-01", data.Value)
	})
}

// TestDatePickerProperties 测试所有属性方法
func TestDatePickerProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*DatePicker) *DatePicker
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "DateType",
			setup:         func(d *DatePicker) *DatePicker { return d.DateType("date") },
			propKey:       "type",
			expectedValue: "date",
		},
		{
			name:          "DateTypeDateTime",
			setup:         func(d *DatePicker) *DatePicker { return d.DateType("datetime") },
			propKey:       "type",
			expectedValue: "datetime",
		},
		{
			name:          "DateTypeRange",
			setup:         func(d *DatePicker) *DatePicker { return d.DateType("daterange") },
			propKey:       "type",
			expectedValue: "daterange",
		},
		{
			name:          "Format",
			setup:         func(d *DatePicker) *DatePicker { return d.Format("yyyy-MM-dd") },
			propKey:       "format",
			expectedValue: "yyyy-MM-dd",
		},
		{
			name:          "ValueFormat",
			setup:         func(d *DatePicker) *DatePicker { return d.ValueFormat("yyyy-MM-dd") },
			propKey:       "value-format",
			expectedValue: "yyyy-MM-dd",
		},
		{
			name:          "Placeholder",
			setup:         func(d *DatePicker) *DatePicker { return d.Placeholder("请选择日期") },
			propKey:       "placeholder",
			expectedValue: "请选择日期",
		},
		{
			name:          "RangeSeparator",
			setup:         func(d *DatePicker) *DatePicker { return d.RangeSeparator("至") },
			propKey:       "range-separator",
			expectedValue: "至",
		},
		{
			name:          "StartPlaceholder",
			setup:         func(d *DatePicker) *DatePicker { return d.StartPlaceholder("开始日期") },
			propKey:       "start-placeholder",
			expectedValue: "开始日期",
		},
		{
			name:          "EndPlaceholder",
			setup:         func(d *DatePicker) *DatePicker { return d.EndPlaceholder("结束日期") },
			propKey:       "end-placeholder",
			expectedValue: "结束日期",
		},
		{
			name:          "Clearable",
			setup:         func(d *DatePicker) *DatePicker { return d.Clearable(true) },
			propKey:       "clearable",
			expectedValue: true,
		},
		{
			name:          "Disabled",
			setup:         func(d *DatePicker) *DatePicker { return d.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
		{
			name:          "Editable",
			setup:         func(d *DatePicker) *DatePicker { return d.Editable(false) },
			propKey:       "editable",
			expectedValue: false,
		},
		{
			name:          "Size",
			setup:         func(d *DatePicker) *DatePicker { return d.Size("large") },
			propKey:       "size",
			expectedValue: "large",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dp := NewDatePicker("test", "测试")
			dp = tt.setup(dp)

			data := dp.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestDatePickerChaining 测试链式调用
func TestDatePickerChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		dp := NewDatePicker("event_date", "活动日期", "2024-12-01").
			DateType("datetime").
			Format("yyyy-MM-dd HH:mm:ss").
			ValueFormat("yyyy-MM-dd HH:mm:ss").
			Placeholder("请选择活动日期").
			Clearable(true).
			Size("large").
			Required()

		data := dp.GetData()
		assert.Equal(t, "2024-12-01", data.Value)
		assert.Equal(t, "datetime", data.Props["type"])
		assert.Equal(t, "yyyy-MM-dd HH:mm:ss", data.Props["format"])
		assert.Equal(t, "yyyy-MM-dd HH:mm:ss", data.Props["value-format"])
		assert.Equal(t, "请选择活动日期", data.Props["placeholder"])
		assert.Equal(t, true, data.Props["clearable"])
		assert.Equal(t, "large", data.Props["size"])
		assert.NotEmpty(t, data.Validate)
	})

	t.Run("DateRangeChain", func(t *testing.T) {
		dp := NewDatePicker("date_range", "日期范围").
			DateType("daterange").
			RangeSeparator("至").
			StartPlaceholder("开始日期").
			EndPlaceholder("结束日期")

		data := dp.GetData()
		assert.Equal(t, "daterange", data.Props["type"])
		assert.Equal(t, "至", data.Props["range-separator"])
		assert.Equal(t, "开始日期", data.Props["start-placeholder"])
		assert.Equal(t, "结束日期", data.Props["end-placeholder"])
	})
}

// TestDatePickerBuild 测试Build方法
func TestDatePickerBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		dp := NewDatePicker("birth_date", "出生日期", "1990-01-01")

		result := dp.Build()

		assert.Equal(t, "datePicker", result["type"])
		assert.Equal(t, "birth_date", result["field"])
		assert.Equal(t, "出生日期", result["title"])
		assert.Equal(t, "1990-01-01", result["value"])
	})

	t.Run("BuildWithDateType", func(t *testing.T) {
		dp := NewDatePicker("datetime", "日期时间").
			DateType("datetime")

		result := dp.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "datetime", props["type"])
	})

	t.Run("BuildWithFormat", func(t *testing.T) {
		dp := NewDatePicker("date", "日期").
			Format("yyyy/MM/dd").
			ValueFormat("yyyy-MM-dd")

		result := dp.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, "yyyy/MM/dd", props["format"])
		assert.Equal(t, "yyyy-MM-dd", props["value-format"])
	})
}

// TestDatePickerEdgeCases 测试边��情况
func TestDatePickerEdgeCases(t *testing.T) {
	t.Run("DateRange", func(t *testing.T) {
		dp := NewDatePicker("range", "日期范围").
			DateType("daterange")

		data := dp.GetData()
		assert.Equal(t, "daterange", data.Props["type"])
	})

	t.Run("MonthRange", func(t *testing.T) {
		dp := NewDatePicker("month_range", "月份范围").
			DateType("monthrange")

		data := dp.GetData()
		assert.Equal(t, "monthrange", data.Props["type"])
	})

	t.Run("Week", func(t *testing.T) {
		dp := NewDatePicker("week", "周").
			DateType("week")

		data := dp.GetData()
		assert.Equal(t, "week", data.Props["type"])
	})

	t.Run("Year", func(t *testing.T) {
		dp := NewDatePicker("year", "年份").
			DateType("year")

		data := dp.GetData()
		assert.Equal(t, "year", data.Props["type"])
	})
}

// TestDatePickerWithValidation 测试验证功能
func TestDatePickerWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		dp := NewDatePicker("date", "日期").Required()

		data := dp.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("DateValidation", func(t *testing.T) {
		dp := NewDatePicker("date", "日期").
			Validate(DateRule{Message: "请选择正确的日期"})

		data := dp.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, "date", rule["type"])
	})
}

// TestDatePickerInForm 测试在表单中的使用
func TestDatePickerInForm(t *testing.T) {
	t.Run("FormWithDatePicker", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewDatePicker("birth_date", "出生日期").
				DateType("date").
				Required(),
			NewDatePicker("event_time", "活动时间").
				DateType("datetime").
				Placeholder("请选择活动时间"),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 2)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "datePicker", rules[0]["type"])
		assert.Equal(t, "datePicker", rules[1]["type"])
	})
}

// BenchmarkDatePickerCreation 性能测试
func BenchmarkDatePickerCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewDatePicker("test", "测试")
	}
}

// BenchmarkDatePickerWithProperties 性能测试
func BenchmarkDatePickerWithProperties(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewDatePicker("date", "日期").
			DateType("datetime").
			Format("yyyy-MM-dd HH:mm:ss").
			ValueFormat("yyyy-MM-dd HH:mm:ss").
			Required()
	}
}

// BenchmarkDatePickerBuild 性能测试
func BenchmarkDatePickerBuild(b *testing.B) {
	dp := NewDatePicker("date", "日期", "2024-01-01").
		DateType("date").
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dp.Build()
	}
}
