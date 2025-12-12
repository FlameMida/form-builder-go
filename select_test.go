package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// select_test.go Select组件完整测试

// TestSelectCreation 测试Select组件创建
func TestSelectCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		sel := NewSelect("role", "角色")

		assert.Equal(t, "role", sel.GetField())
		assert.Equal(t, "select", sel.GetType())

		data := sel.GetData()
		assert.Equal(t, "role", data.Field)
		assert.Equal(t, "角色", data.Title)
		assert.NotNil(t, sel.options)
		assert.Len(t, sel.options, 0) // 初始为空
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		sel := NewSelect("test", "测试", "default")

		data := sel.GetData()
		assert.Equal(t, "default", data.Value)
	})

	t.Run("IviewCreation", func(t *testing.T) {
		sel := NewIviewSelect("test", "测试")

		assert.Equal(t, "select", sel.GetType())
	})
}

// TestSelectSetOptions 测试SetOptions方法
func TestSelectSetOptions(t *testing.T) {
	t.Run("SetBasicOptions", func(t *testing.T) {
		sel := NewSelect("role", "角色").SetOptions([]Option{
			{Value: "admin", Label: "管理员"},
			{Value: "user", Label: "用户"},
		})

		assert.Len(t, sel.options, 2)
		assert.Equal(t, "admin", sel.options[0].Value)
		assert.Equal(t, "管理员", sel.options[0].Label)
	})

	t.Run("SetEmptyOptions", func(t *testing.T) {
		sel := NewSelect("test", "测试").SetOptions([]Option{})

		assert.Len(t, sel.options, 0)
	})

	t.Run("OptionsOverwrite", func(t *testing.T) {
		sel := NewSelect("test", "测试").
			SetOptions([]Option{{Value: "1", Label: "选项1"}}).
			SetOptions([]Option{{Value: "2", Label: "选项2"}})

		assert.Len(t, sel.options, 1)
		assert.Equal(t, "2", sel.options[0].Value)
	})
}

// TestSelectAppendOption 测试AppendOption方法
func TestSelectAppendOption(t *testing.T) {
	t.Run("AppendSingleOption", func(t *testing.T) {
		sel := NewSelect("test", "测试").
			AppendOption(Option{Value: "1", Label: "选项1"}).
			AppendOption(Option{Value: "2", Label: "选项2"})

		assert.Len(t, sel.options, 2)
	})
}

// TestSelectProperties 测试所有属性方法
func TestSelectProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Select) *Select
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "Multiple",
			setup:         func(s *Select) *Select { return s.Multiple(true) },
			propKey:       "multiple",
			expectedValue: true,
		},
		{
			name:          "Disabled",
			setup:         func(s *Select) *Select { return s.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
		{
			name:          "Clearable",
			setup:         func(s *Select) *Select { return s.Clearable(true) },
			propKey:       "clearable",
			expectedValue: true,
		},
		{
			name:          "Filterable",
			setup:         func(s *Select) *Select { return s.Filterable(true) },
			propKey:       "filterable",
			expectedValue: true,
		},
		{
			name:          "Remote",
			setup:         func(s *Select) *Select { return s.Remote(true) },
			propKey:       "remote",
			expectedValue: true,
		},
		{
			name:          "RemoteMethod",
			setup:         func(s *Select) *Select { return s.RemoteMethod("fetchData") },
			propKey:       "remote-method",
			expectedValue: "fetchData",
		},
		{
			name:          "Placeholder",
			setup:         func(s *Select) *Select { return s.Placeholder("请选择") },
			propKey:       "placeholder",
			expectedValue: "请选择",
		},
		{
			name:          "Size",
			setup:         func(s *Select) *Select { return s.Size("large") },
			propKey:       "size",
			expectedValue: "large",
		},
		{
			name:          "CollapseTags",
			setup:         func(s *Select) *Select { return s.CollapseTags(true) },
			propKey:       "collapse-tags",
			expectedValue: true,
		},
		{
			name:          "MultipleLimit",
			setup:         func(s *Select) *Select { return s.MultipleLimit(3) },
			propKey:       "multiple-limit",
			expectedValue: 3,
		},
		{
			name:          "AllowCreate",
			setup:         func(s *Select) *Select { return s.AllowCreate(true) },
			propKey:       "allow-create",
			expectedValue: true,
		},
		{
			name:          "DefaultFirstOption",
			setup:         func(s *Select) *Select { return s.DefaultFirstOption(true) },
			propKey:       "default-first-option",
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sel := NewSelect("test", "测试")
			sel = tt.setup(sel)

			data := sel.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestSelectChaining 测试链式调用
func TestSelectChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		sel := NewSelect("role", "角色").
			Placeholder("请选择角色").
			Clearable(true).
			Filterable(true).
			Multiple(true).
			SetOptions([]Option{
				{Value: "admin", Label: "管理员"},
				{Value: "user", Label: "用户"},
			}).
			Required().
			Value("admin")

		data := sel.GetData()
		assert.Equal(t, "请选择角色", data.Props["placeholder"])
		assert.Equal(t, true, data.Props["clearable"])
		assert.Equal(t, true, data.Props["filterable"])
		assert.Equal(t, true, data.Props["multiple"])
		assert.Equal(t, "admin", data.Value)
		assert.NotEmpty(t, data.Validate)
		assert.Len(t, sel.options, 2)
	})
}

// TestSelectBuild 测试Build方法
func TestSelectBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		sel := NewSelect("role", "角色").
			SetOptions([]Option{
				{Value: "1", Label: "选项1"},
				{Value: "2", Label: "选项2"},
			})

		result := sel.Build()

		assert.Equal(t, "select", result["type"])
		assert.Equal(t, "role", result["field"])
		assert.Equal(t, "角色", result["title"])

		options, ok := result["options"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, options, 2)
		assert.Equal(t, "1", options[0]["value"])
		assert.Equal(t, "选项1", options[0]["label"])
	})

	t.Run("BuildWithProps", func(t *testing.T) {
		sel := NewSelect("test", "测试").
			Placeholder("请选择").
			Multiple(true)

		result := sel.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "请选择", props["placeholder"])
		assert.Equal(t, true, props["multiple"])
	})

	t.Run("BuildWithOptionsDisabled", func(t *testing.T) {
		sel := NewSelect("test", "测试").SetOptions([]Option{
			{Value: "1", Label: "选项1"},
			{Value: "2", Label: "选项2", Disabled: true},
		})

		result := sel.Build()
		options := result["options"].([]map[string]interface{})
		// 第一个选项没有设置disabled，所以不应该有这个键或者为false
		if disabled, exists := options[0]["disabled"]; exists {
			assert.Equal(t, false, disabled)
		}
		// 第二个选项明确禁用
		assert.Equal(t, true, options[1]["disabled"])
	})
}

// TestSelectWithControl 测试带Control的Select
func TestSelectWithControl(t *testing.T) {
	t.Run("SelectWithControl", func(t *testing.T) {
		sel := NewSelect("type", "类型").
			SetOptions([]Option{
				{Value: "1", Label: "类型1"},
				{Value: "2", Label: "类型2"},
			}).
			Control([]ControlRule{
				{
					Value: "1",
					Rule: []Component{
						NewInput("extra", "额外字段"),
					},
				},
			})

		result := sel.Build()
		control, ok := result["control"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, control, 1)
		assert.Equal(t, "1", control[0]["value"])
	})
}

// TestSelectEdgeCases 测试边缘情况
func TestSelectEdgeCases(t *testing.T) {
	t.Run("EmptyOptions", func(t *testing.T) {
		sel := NewSelect("test", "测试")
		result := sel.Build()

		// 空选项应该不出现在Build结果中，或者是空数组
		options, exists := result["options"]
		if exists {
			assert.Len(t, options, 0)
		}
	})

	t.Run("MultipleWithSingleValue", func(t *testing.T) {
		sel := NewSelect("test", "测试").
			Multiple(true).
			Value("single")

		data := sel.GetData()
		assert.Equal(t, "single", data.Value)
	})

	t.Run("MultipleWithArrayValue", func(t *testing.T) {
		sel := NewSelect("test", "测试").
			Multiple(true).
			Value([]string{"1", "2", "3"})

		data := sel.GetData()
		values, ok := data.Value.([]string)
		require.True(t, ok)
		assert.Len(t, values, 3)
	})

	t.Run("ZeroMultipleLimit", func(t *testing.T) {
		sel := NewSelect("test", "测试").MultipleLimit(0)

		data := sel.GetData()
		assert.Equal(t, 0, data.Props["multiple-limit"])
	})
}

// BenchmarkSelectCreation 性能测试
func BenchmarkSelectCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSelect("test", "测试")
	}
}

// BenchmarkSelectWithOptions 性能测试
func BenchmarkSelectWithOptions(b *testing.B) {
	options := []Option{
		{Value: "1", Label: "选项1"},
		{Value: "2", Label: "选项2"},
		{Value: "3", Label: "选项3"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSelect("test", "测试").SetOptions(options)
	}
}

// BenchmarkSelectBuild 性能测试
func BenchmarkSelectBuild(b *testing.B) {
	sel := NewSelect("role", "角色").
		SetOptions([]Option{
			{Value: "admin", Label: "管理员"},
			{Value: "user", Label: "用户"},
		}).
		Placeholder("请选择").
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sel.Build()
	}
}
