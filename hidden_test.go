package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// hidden_test.go Hidden组件完整测试

// TestHiddenCreation 测试Hidden组件创建
func TestHiddenCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		hidden := NewHidden("token")

		assert.Equal(t, "token", hidden.GetField())
		assert.Equal(t, "hidden", hidden.GetType())

		data := hidden.GetData()
		assert.Equal(t, "token", data.Field)
		assert.Equal(t, "", data.Title) // Hidden fields have empty title
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		hidden := NewHidden("csrf_token", "abc123")

		data := hidden.GetData()
		assert.Equal(t, "abc123", data.Value)
	})

	t.Run("CreationWithNumericValue", func(t *testing.T) {
		hidden := NewHidden("user_id", 12345)

		data := hidden.GetData()
		assert.Equal(t, 12345, data.Value)
	})

	t.Run("CreationWithComplexValue", func(t *testing.T) {
		complexValue := map[string]interface{}{
			"timestamp": 1234567890,
			"signature": "abc123",
		}
		hidden := NewHidden("metadata", complexValue)

		data := hidden.GetData()
		assert.Equal(t, complexValue, data.Value)
	})
}

// TestHiddenBuild 测试Build方法
func TestHiddenBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		hidden := NewHidden("id", 123)

		result := hidden.Build()

		assert.Equal(t, "hidden", result["type"])
		assert.Equal(t, "id", result["field"])
		// Hidden has empty title, might be omitted or empty in build
		title, hasTitle := result["title"]
		if hasTitle {
			assert.Equal(t, "", title)
		}
		assert.Equal(t, 123, result["value"])
	})

	t.Run("BuildWithStringValue", func(t *testing.T) {
		hidden := NewHidden("token", "secret-token-value")

		result := hidden.Build()

		assert.Equal(t, "hidden", result["type"])
		assert.Equal(t, "token", result["field"])
		assert.Equal(t, "secret-token-value", result["value"])
	})

	t.Run("BuildWithoutValue", func(t *testing.T) {
		hidden := NewHidden("optional_field")

		result := hidden.Build()

		assert.Equal(t, "hidden", result["type"])
		assert.Equal(t, "optional_field", result["field"])
		_, hasValue := result["value"]
		assert.False(t, hasValue)
	})
}

// TestHiddenChaining 测试链式调用
func TestHiddenChaining(t *testing.T) {
	t.Run("HiddenWithControl", func(t *testing.T) {
		hidden := NewHidden("enable_feature", "yes").
			Control([]ControlRule{
				{
					Value: "yes",
					Rule: []Component{
						NewInput("feature_name", "功能名称").Required(),
					},
				},
			})

		data := hidden.GetData()
		assert.Equal(t, "yes", data.Value)
		assert.Len(t, data.Control, 1)
		assert.Len(t, data.Control[0].Rule, 1)
	})

	t.Run("HiddenWithEmit", func(t *testing.T) {
		hidden := NewHidden("trigger", "init").
			Emit("mounted", "function() { console.log('loaded'); }")

		data := hidden.GetData()
		assert.NotNil(t, data.Emit)
		assert.Equal(t, "function() { console.log('loaded'); }", data.Emit["mounted"])
	})
}

// TestHiddenEdgeCases 测试边缘情况
func TestHiddenEdgeCases(t *testing.T) {
	t.Run("EmptyStringValue", func(t *testing.T) {
		hidden := NewHidden("empty", "")

		data := hidden.GetData()
		assert.Equal(t, "", data.Value)
	})

	t.Run("ZeroValue", func(t *testing.T) {
		hidden := NewHidden("zero", 0)

		data := hidden.GetData()
		assert.Equal(t, 0, data.Value)
	})

	t.Run("BooleanValue", func(t *testing.T) {
		hidden := NewHidden("flag", false)

		data := hidden.GetData()
		assert.Equal(t, false, data.Value)
	})

	t.Run("NilValue", func(t *testing.T) {
		hidden := NewHidden("nullable", nil)

		data := hidden.GetData()
		assert.Nil(t, data.Value)
	})

	t.Run("ArrayValue", func(t *testing.T) {
		arr := []int{1, 2, 3}
		hidden := NewHidden("array", arr)

		data := hidden.GetData()
		assert.Equal(t, arr, data.Value)
	})

	t.Run("MapValue", func(t *testing.T) {
		m := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}
		hidden := NewHidden("map", m)

		data := hidden.GetData()
		assert.Equal(t, m, data.Value)
	})
}

// TestHiddenInForm 测试在表单中的使用
func TestHiddenInForm(t *testing.T) {
	t.Run("FormWithHidden", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewHidden("user_id", 123),
			NewHidden("csrf_token", "abc123"),
			NewInput("username", "用户名").Required(),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 3)

		rules := form.FormRule()
		assert.Len(t, rules, 3)
		assert.Equal(t, "hidden", rules[0]["type"])
		assert.Equal(t, "hidden", rules[1]["type"])
		assert.Equal(t, "input", rules[2]["type"])
	})

	t.Run("MultipleHiddenFields", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewHidden("id", 1),
			NewHidden("timestamp", 1234567890),
			NewHidden("signature", "abc"),
			NewHidden("version", "1.0"),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 4)

		rules := form.FormRule()
		for _, rule := range rules {
			assert.Equal(t, "hidden", rule["type"])
		}
	})

	t.Run("HiddenWithVisibleFields", func(t *testing.T) {
		form := NewElmForm("/api/update", []Component{
			NewHidden("id", 123),
			NewInput("name", "名称").Required(),
			NewTextarea("description", "描述"),
			NewHidden("updated_at", 1234567890),
		}, nil)

		rules := form.FormRule()
		assert.Len(t, rules, 4)
		assert.Equal(t, "hidden", rules[0]["type"])
		assert.Equal(t, "input", rules[1]["type"])
		assert.Equal(t, "input", rules[2]["type"]) // Textarea is el-input
		assert.Equal(t, "hidden", rules[3]["type"])
	})
}

// TestHiddenWithValidation 测试验证功能
func TestHiddenWithValidation(t *testing.T) {
	t.Run("HiddenWithRequired", func(t *testing.T) {
		// Although unusual, hidden fields can have validation
		hidden := NewHidden("required_token", "").Required()

		data := hidden.GetData()
		assert.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("HiddenWithCustomValidation", func(t *testing.T) {
		hidden := NewHidden("token", "abc").
			Validate(PatternRule{
				Pattern: "^[a-z]+$",
				Message: "Token must be lowercase letters",
			})

		data := hidden.GetData()
		assert.Len(t, data.Validate, 1)
	})
}

// TestHiddenComparison 测试不同值类型
func TestHiddenComparison(t *testing.T) {
	t.Run("StringVsNumber", func(t *testing.T) {
		hiddenStr := NewHidden("field1", "123")
		hiddenNum := NewHidden("field2", 123)

		dataStr := hiddenStr.GetData()
		dataNum := hiddenNum.GetData()

		assert.IsType(t, "", dataStr.Value)
		assert.IsType(t, 0, dataNum.Value)
		assert.NotEqual(t, dataStr.Value, dataNum.Value)
	})

	t.Run("BooleanValues", func(t *testing.T) {
		hiddenTrue := NewHidden("bool_true", true)
		hiddenFalse := NewHidden("bool_false", false)

		dataTrue := hiddenTrue.GetData()
		dataFalse := hiddenFalse.GetData()

		assert.Equal(t, true, dataTrue.Value)
		assert.Equal(t, false, dataFalse.Value)
	})
}

// TestHiddenInterfaceMethods 测试接口方法
func TestHiddenInterfaceMethods(t *testing.T) {
	t.Run("GetField", func(t *testing.T) {
		hidden := NewHidden("test_field", "value")
		assert.Equal(t, "test_field", hidden.GetField())
	})

	t.Run("GetType", func(t *testing.T) {
		hidden := NewHidden("test_field", "value")
		assert.Equal(t, "hidden", hidden.GetType())
	})

	t.Run("AsComponent", func(t *testing.T) {
		hidden := NewHidden("test_field", "value")
		var component Component = hidden
		assert.NotNil(t, component)
		assert.Equal(t, "test_field", component.GetField())
		assert.Equal(t, "hidden", component.GetType())
	})
}

// BenchmarkHiddenCreation 性能测试
func BenchmarkHiddenCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewHidden("test", "value")
	}
}

// BenchmarkHiddenWithValue 性能测试
func BenchmarkHiddenWithValue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewHidden("token", "abc123xyz")
	}
}

// BenchmarkHiddenBuild 性能测试
func BenchmarkHiddenBuild(b *testing.B) {
	hidden := NewHidden("id", 123)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hidden.Build()
	}
}
