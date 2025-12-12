package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// switch_test.go Switch组件完整测试

// TestSwitchCreation 测试Switch组件创建
func TestSwitchCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		sw := NewSwitch("is_active", "是否激活")

		assert.Equal(t, "is_active", sw.GetField())
		assert.Equal(t, "switch", sw.GetType())

		data := sw.GetData()
		assert.Equal(t, "is_active", data.Field)
		assert.Equal(t, "是否激活", data.Title)
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		sw := NewSwitch("enabled", "启用", true)

		data := sw.GetData()
		assert.Equal(t, true, data.Value)
	})

	t.Run("CreationWithFalseValue", func(t *testing.T) {
		sw := NewSwitch("enabled", "启用", false)

		data := sw.GetData()
		assert.Equal(t, false, data.Value)
	})
}

// TestSwitchProperties 测试所有属性方法
func TestSwitchProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Switch) *Switch
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "ActiveText",
			setup:         func(s *Switch) *Switch { return s.ActiveText("启用") },
			propKey:       "active-text",
			expectedValue: "启用",
		},
		{
			name:          "InactiveText",
			setup:         func(s *Switch) *Switch { return s.InactiveText("禁用") },
			propKey:       "inactive-text",
			expectedValue: "禁用",
		},
		{
			name:          "ActiveValue",
			setup:         func(s *Switch) *Switch { return s.ActiveValue(1) },
			propKey:       "active-value",
			expectedValue: 1,
		},
		{
			name:          "InactiveValue",
			setup:         func(s *Switch) *Switch { return s.InactiveValue(0) },
			propKey:       "inactive-value",
			expectedValue: 0,
		},
		{
			name:          "ActiveColor",
			setup:         func(s *Switch) *Switch { return s.ActiveColor("#13ce66") },
			propKey:       "active-color",
			expectedValue: "#13ce66",
		},
		{
			name:          "InactiveColor",
			setup:         func(s *Switch) *Switch { return s.InactiveColor("#ff4949") },
			propKey:       "inactive-color",
			expectedValue: "#ff4949",
		},
		{
			name:          "Disabled",
			setup:         func(s *Switch) *Switch { return s.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := NewSwitch("test", "测试")
			sw = tt.setup(sw)

			data := sw.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestSwitchChaining 测试链式调用
func TestSwitchChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		sw := NewSwitch("is_active", "是否激活", true).
			ActiveText("启用").
			InactiveText("禁用").
			ActiveValue(1).
			InactiveValue(0).
			ActiveColor("#13ce66").
			InactiveColor("#ff4949").
			Required()

		data := sw.GetData()
		assert.Equal(t, true, data.Value)
		assert.Equal(t, "启用", data.Props["active-text"])
		assert.Equal(t, "禁用", data.Props["inactive-text"])
		assert.Equal(t, 1, data.Props["active-value"])
		assert.Equal(t, 0, data.Props["inactive-value"])
		assert.Equal(t, "#13ce66", data.Props["active-color"])
		assert.Equal(t, "#ff4949", data.Props["inactive-color"])
		assert.NotEmpty(t, data.Validate)
	})

	t.Run("ChainPreservesType", func(t *testing.T) {
		sw := NewSwitch("test", "测试").
			ActiveText("是").
			InactiveText("否")

		// 验证链式调用返回的仍然是*Switch类型
		assert.IsType(t, &Switch{}, sw)
	})
}

// TestSwitchBuild 测试Build方法
func TestSwitchBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		sw := NewSwitch("enabled", "启用", true)

		result := sw.Build()

		assert.Equal(t, "switch", result["type"])
		assert.Equal(t, "enabled", result["field"])
		assert.Equal(t, "启用", result["title"])
		assert.Equal(t, true, result["value"])
	})

	t.Run("BuildWithTexts", func(t *testing.T) {
		sw := NewSwitch("is_active", "是否激活").
			ActiveText("开启").
			InactiveText("关闭")

		result := sw.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "开启", props["active-text"])
		assert.Equal(t, "关闭", props["inactive-text"])
	})

	t.Run("BuildWithCustomValues", func(t *testing.T) {
		sw := NewSwitch("status", "状态").
			ActiveValue("on").
			InactiveValue("off")

		result := sw.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, "on", props["active-value"])
		assert.Equal(t, "off", props["inactive-value"])
	})

	t.Run("BuildWithColors", func(t *testing.T) {
		sw := NewSwitch("test", "测试").
			ActiveColor("#13ce66").
			InactiveColor("#ff4949")

		result := sw.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, "#13ce66", props["active-color"])
		assert.Equal(t, "#ff4949", props["inactive-color"])
	})

	t.Run("BuildWithDisabled", func(t *testing.T) {
		sw := NewSwitch("test", "测试").Disabled(true)

		result := sw.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, true, props["disabled"])
	})
}

// TestSwitchEdgeCases 测试边缘情况
func TestSwitchEdgeCases(t *testing.T) {
	t.Run("BooleanValue", func(t *testing.T) {
		sw := NewSwitch("flag", "标志", true)

		data := sw.GetData()
		assert.Equal(t, true, data.Value)
	})

	t.Run("NumericActiveInactiveValues", func(t *testing.T) {
		sw := NewSwitch("status", "状态").
			ActiveValue(1).
			InactiveValue(0)

		data := sw.GetData()
		assert.Equal(t, 1, data.Props["active-value"])
		assert.Equal(t, 0, data.Props["inactive-value"])
	})

	t.Run("StringActiveInactiveValues", func(t *testing.T) {
		sw := NewSwitch("mode", "模式").
			ActiveValue("light").
			InactiveValue("dark")

		data := sw.GetData()
		assert.Equal(t, "light", data.Props["active-value"])
		assert.Equal(t, "dark", data.Props["inactive-value"])
	})

	t.Run("EmptyTexts", func(t *testing.T) {
		sw := NewSwitch("test", "测试").
			ActiveText("").
			InactiveText("")

		data := sw.GetData()
		assert.Equal(t, "", data.Props["active-text"])
		assert.Equal(t, "", data.Props["inactive-text"])
	})
}

// TestSwitchWithValidation 测试验证功能
func TestSwitchWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		sw := NewSwitch("agree", "同意条款").Required()

		data := sw.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("CustomValidation", func(t *testing.T) {
		sw := NewSwitch("test", "测试").
			Validate(RequiredRule{Message: "必须同意才能继续"})

		data := sw.GetData()
		require.Len(t, data.Validate, 1)
	})
}

// TestSwitchInForm 测试在表单中的使用
func TestSwitchInForm(t *testing.T) {
	t.Run("FormWithSwitch", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("username", "用户名").Required(),
			NewSwitch("is_active", "是否激活").
				ActiveText("启用").
				InactiveText("禁用").
				Required(),
			NewSwitch("agree", "同意条款").
				ActiveText("同意").
				InactiveText("不同意"),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 3)

		rules := form.FormRule()
		assert.Len(t, rules, 3)
		assert.Equal(t, "input", rules[0]["type"])
		assert.Equal(t, "switch", rules[1]["type"])
		assert.Equal(t, "switch", rules[2]["type"])
	})

	t.Run("FormDataWithSwitch", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewSwitch("is_active", "是否激活"),
			NewSwitch("agree", "同意条款"),
		}, nil)

		form.FormData(map[string]interface{}{
			"is_active": true,
			"agree":     false,
		})

		rules := form.FormRule()
		assert.Equal(t, true, rules[0]["value"])
		assert.Equal(t, false, rules[1]["value"])
	})
}

// TestSwitchWithControl 测试Switch的Control功能
func TestSwitchWithControl(t *testing.T) {
	t.Run("SwitchControlsField", func(t *testing.T) {
		sw := NewSwitch("enable_feature", "启用功能", false).
			Control([]ControlRule{
				{
					Value: true,
					Rule: []Component{
						NewInput("feature_config", "功能配置").Required(),
					},
				},
			})

		result := sw.Build()

		control, ok := result["control"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, control, 1)
		assert.Equal(t, true, control[0]["value"])

		rules, ok := control[0]["rule"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, rules, 1)
		assert.Equal(t, "feature_config", rules[0]["field"])
	})
}

// BenchmarkSwitchCreation 性能测试
func BenchmarkSwitchCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSwitch("test", "测试")
	}
}

// BenchmarkSwitchWithProperties 性能测试
func BenchmarkSwitchWithProperties(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSwitch("is_active", "是否激活", true).
			ActiveText("启用").
			InactiveText("禁用").
			ActiveValue(1).
			InactiveValue(0).
			Required()
	}
}

// BenchmarkSwitchBuild 性能测试
func BenchmarkSwitchBuild(b *testing.B) {
	sw := NewSwitch("is_active", "是否激活", true).
		ActiveText("启用").
		InactiveText("禁用").
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sw.Build()
	}
}
