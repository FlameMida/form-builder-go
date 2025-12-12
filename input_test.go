package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// input_test.go Input组件完整测试
// 这是组件测试的标准模板，其他组件可以参考此文件

// TestInputCreation 测试Input组件创建
func TestInputCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		input := NewInput("username", "用户名")

		assert.Equal(t, "username", input.GetField())
		assert.Equal(t, "input", input.GetType())

		data := input.GetData()
		assert.Equal(t, "username", data.Field)
		assert.Equal(t, "用户名", data.Title)
		assert.Equal(t, "input", data.RuleType)
		assert.NotNil(t, data.Props)
		assert.Equal(t, "text", data.Props["type"]) // 默认type
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		input := NewInput("test", "测试", "default")

		data := input.GetData()
		assert.Equal(t, "default", data.Value)
	})

	t.Run("CreationWithoutValue", func(t *testing.T) {
		input := NewInput("test", "测试")

		data := input.GetData()
		assert.Nil(t, data.Value)
	})

	t.Run("IviewCreation", func(t *testing.T) {
		input := NewIviewInput("test", "测试")

		assert.Equal(t, "input", input.GetType())
	})
}

// TestInputType 测试Type方法
func TestInputType(t *testing.T) {
	tests := []struct {
		name         string
		inputType    string
		expectedType string
	}{
		{"TextType", "text", "text"},
		{"PasswordType", "password", "password"},
		{"EmailType", "email", "email"},
		{"NumberType", "number", "number"},
		{"URLType", "url", "url"},
		{"TelType", "tel", "tel"},
		{"SearchType", "search", "search"},
		{"TextareaType", "textarea", "textarea"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := NewInput("test", "测试").Type(tt.inputType)

			data := input.GetData()
			assert.Equal(t, tt.expectedType, data.Props["type"])
		})
	}
}

// TestInputProperties 测试所有Props属性方法（表驱动测试）
func TestInputProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Input) *Input
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "Placeholder",
			setup:         func(i *Input) *Input { return i.Placeholder("请输入") },
			propKey:       "placeholder",
			expectedValue: "请输入",
		},
		{
			name:          "Clearable",
			setup:         func(i *Input) *Input { return i.Clearable(true) },
			propKey:       "clearable",
			expectedValue: true,
		},
		{
			name:          "ShowPassword",
			setup:         func(i *Input) *Input { return i.ShowPassword(true) },
			propKey:       "show-password",
			expectedValue: true,
		},
		{
			name:          "Disabled",
			setup:         func(i *Input) *Input { return i.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
		{
			name:          "Readonly",
			setup:         func(i *Input) *Input { return i.Readonly(true) },
			propKey:       "readonly",
			expectedValue: true,
		},
		{
			name:          "MaxLength",
			setup:         func(i *Input) *Input { return i.MaxLength(100) },
			propKey:       "maxlength",
			expectedValue: 100,
		},
		{
			name:          "MinLength",
			setup:         func(i *Input) *Input { return i.MinLength(6) },
			propKey:       "minlength",
			expectedValue: 6,
		},
		{
			name:          "ShowWordLimit",
			setup:         func(i *Input) *Input { return i.ShowWordLimit(true) },
			propKey:       "show-word-limit",
			expectedValue: true,
		},
		{
			name:          "PrefixIcon",
			setup:         func(i *Input) *Input { return i.PrefixIcon("el-icon-user") },
			propKey:       "prefix-icon",
			expectedValue: "el-icon-user",
		},
		{
			name:          "SuffixIcon",
			setup:         func(i *Input) *Input { return i.SuffixIcon("el-icon-search") },
			propKey:       "suffix-icon",
			expectedValue: "el-icon-search",
		},
		{
			name:          "SizeLarge",
			setup:         func(i *Input) *Input { return i.Size("large") },
			propKey:       "size",
			expectedValue: "large",
		},
		{
			name:          "SizeSmall",
			setup:         func(i *Input) *Input { return i.Size("small") },
			propKey:       "size",
			expectedValue: "small",
		},
		{
			name:          "Autocomplete",
			setup:         func(i *Input) *Input { return i.Autocomplete("off") },
			propKey:       "autocomplete",
			expectedValue: "off",
		},
		{
			name:          "Autofocus",
			setup:         func(i *Input) *Input { return i.Autofocus(true) },
			propKey:       "autofocus",
			expectedValue: true,
		},
		{
			name:          "Rows",
			setup:         func(i *Input) *Input { return i.Rows(5) },
			propKey:       "rows",
			expectedValue: 5,
		},
		{
			name:          "ValidateEvent",
			setup:         func(i *Input) *Input { return i.ValidateEvent(false) },
			propKey:       "validate-event",
			expectedValue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := NewInput("test", "测试")
			input = tt.setup(input)

			data := input.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestInputAutosize 测试Autosize方法（特殊情况，接受interface{}）
func TestInputAutosize(t *testing.T) {
	t.Run("BooleanAutosize", func(t *testing.T) {
		input := NewInput("test", "测试").Autosize(true)

		data := input.GetData()
		assert.Equal(t, true, data.Props["autosize"])
	})

	t.Run("ObjectAutosize", func(t *testing.T) {
		autosizeConfig := map[string]interface{}{
			"minRows": 2,
			"maxRows": 6,
		}

		input := NewInput("test", "测试").Autosize(autosizeConfig)

		data := input.GetData()
		assert.Equal(t, autosizeConfig, data.Props["autosize"])
	})
}

// TestInputChaining 测试链式调用
func TestInputChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		input := NewInput("username", "用户名").
			Placeholder("请输入用户名").
			Clearable(true).
			MaxLength(50).
			MinLength(6).
			ShowWordLimit(true).
			PrefixIcon("el-icon-user").
			Size("large").
			Required().
			Value("default")

		// 验证所有设置都生效
		data := input.GetData()
		assert.Equal(t, "请输入用户名", data.Props["placeholder"])
		assert.Equal(t, true, data.Props["clearable"])
		assert.Equal(t, 50, data.Props["maxlength"])
		assert.Equal(t, 6, data.Props["minlength"])
		assert.Equal(t, true, data.Props["show-word-limit"])
		assert.Equal(t, "el-icon-user", data.Props["prefix-icon"])
		assert.Equal(t, "large", data.Props["size"])
		assert.Equal(t, "default", data.Value)
		assert.NotEmpty(t, data.Validate)
	})

	t.Run("ChainPreservesType", func(t *testing.T) {
		// 测试链式调用后仍然返回*Input类型
		input := NewInput("test", "测试").
			Required().          // Builder方法
			Value("test").       // Builder方法
			Placeholder("test"). // Input方法
			Clearable(true).     // Input方法
			ShowPassword(true)   // Input方法

		assert.NotNil(t, input)
		// 如果类型不对，下面的调用会编译失败
		input.MaxLength(100)
	})
}

// TestInputBuild 测试Build方法
func TestInputBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		input := NewInput("username", "用户名").
			Placeholder("请输入").
			Value("test")

		result := input.Build()

		assert.Equal(t, "input", result["type"])
		assert.Equal(t, "username", result["field"])
		assert.Equal(t, "用户名", result["title"])
		assert.Equal(t, "test", result["value"])

		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "请输入", props["placeholder"])
		assert.Equal(t, "text", props["type"])
	})

	t.Run("BuildWithValidation", func(t *testing.T) {
		input := NewInput("test", "测试").
			Required().
			Validate(LengthRule{Min: 6, Max: 20, Message: "长度6-20"})

		result := input.Build()

		validates, ok := result["validate"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, validates, 2)
	})

	t.Run("BuildIview", func(t *testing.T) {
		input := NewIviewInput("test", "测试")
		result := input.Build()

		assert.Equal(t, "input", result["type"])
	})
}

// TestInputValidation 测试验证规则
func TestInputValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		input := NewInput("test", "测试").Required()

		data := input.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("MultipleValidationRules", func(t *testing.T) {
		input := NewInput("test", "测试").
			Required().
			Validate(
				LengthRule{Min: 6, Max: 20},
				PatternRule{Pattern: "^[a-zA-Z0-9]+$"},
			)

		data := input.GetData()
		assert.Len(t, data.Validate, 3)
	})
}

// TestInputConvenienceFunctions 测试便捷构造函数
func TestInputConvenienceFunctions(t *testing.T) {
	t.Run("Password", func(t *testing.T) {
		input := Password("pwd", "密码")

		data := input.GetData()
		assert.Equal(t, "password", data.Props["type"])
		assert.Equal(t, "pwd", data.Field)
	})

	t.Run("PasswordWithValue", func(t *testing.T) {
		input := Password("pwd", "密码", "secret")

		data := input.GetData()
		assert.Equal(t, "secret", data.Value)
	})

	t.Run("Email", func(t *testing.T) {
		input := Email("email", "邮箱")

		data := input.GetData()
		assert.Equal(t, "email", data.Props["type"])
		assert.NotEmpty(t, data.Validate) // 自动添加邮箱验证
	})

	t.Run("URL", func(t *testing.T) {
		input := URL("website", "网站")

		data := input.GetData()
		assert.Equal(t, "url", data.Props["type"])
		assert.NotEmpty(t, data.Validate) // 自动添加URL验证
	})

	t.Run("Tel", func(t *testing.T) {
		input := Tel("phone", "电话")

		data := input.GetData()
		assert.Equal(t, "tel", data.Props["type"])
	})

	t.Run("Search", func(t *testing.T) {
		input := Search("q", "搜索")

		data := input.GetData()
		assert.Equal(t, "search", data.Props["type"])
	})
}

// TestInputEdgeCases 测试边缘情况
func TestInputEdgeCases(t *testing.T) {
	t.Run("EmptyField", func(t *testing.T) {
		input := NewInput("", "")
		assert.Equal(t, "", input.GetField())
	})

	t.Run("ZeroMaxLength", func(t *testing.T) {
		input := NewInput("test", "测试").MaxLength(0)

		data := input.GetData()
		assert.Equal(t, 0, data.Props["maxlength"])
	})

	t.Run("NegativeMaxLength", func(t *testing.T) {
		input := NewInput("test", "测试").MaxLength(-1)

		data := input.GetData()
		assert.Equal(t, -1, data.Props["maxlength"]) // 允许负数，由前端验证
	})

	t.Run("VeryLongPlaceholder", func(t *testing.T) {
		longText := "这是一个非常非常非常非常非常非常长的占位符文本"
		input := NewInput("test", "测试").Placeholder(longText)

		data := input.GetData()
		assert.Equal(t, longText, data.Props["placeholder"])
	})

	t.Run("SpecialCharactersInValue", func(t *testing.T) {
		specialValue := `<script>alert("xss")</script>`
		input := NewInput("test", "测试").Value(specialValue)

		data := input.GetData()
		assert.Equal(t, specialValue, data.Value)
	})

	t.Run("NilValue", func(t *testing.T) {
		input := NewInput("test", "测试").Value(nil)

		data := input.GetData()
		assert.Nil(t, data.Value)
	})

	t.Run("MultipleTypeCalls", func(t *testing.T) {
		input := NewInput("test", "测试").
			Type("password").
			Type("email")

		data := input.GetData()
		assert.Equal(t, "email", data.Props["type"]) // 最后一个生效
	})
}

// TestInputWithControl 测试带Control的Input
func TestInputWithControl(t *testing.T) {
	t.Run("InputWithControl", func(t *testing.T) {
		input := NewInput("type", "类型").
			Control([]ControlRule{
				{
					Value: "show",
					Rule: []Component{
						NewInput("extra", "额外字段"),
					},
				},
			})

		data := input.GetData()
		require.Len(t, data.Control, 1)
		assert.Equal(t, "show", data.Control[0].Value)
	})
}

// BenchmarkInputCreation 性能测试：Input创建
func BenchmarkInputCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewInput("test", "测试")
	}
}

// BenchmarkInputChaining 性能测试：Input链式调用
func BenchmarkInputChaining(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewInput("username", "用户名").
			Placeholder("请输入用户名").
			Clearable(true).
			MaxLength(50).
			Required().
			Value("default")
	}
}

// BenchmarkInputBuild 性能测试：Input Build
func BenchmarkInputBuild(b *testing.B) {
	input := NewInput("username", "用户名").
		Placeholder("请输入").
		Clearable(true).
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = input.Build()
	}
}
