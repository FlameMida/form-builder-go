package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// cascader_test.go Cascader组件完整测试

// TestCascaderCreation 测试Cascader组件创建
func TestCascaderCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		cascader := NewCascader("region", "地区")

		assert.Equal(t, "region", cascader.GetField())
		assert.Equal(t, "cascader", cascader.GetType())

		data := cascader.GetData()
		assert.Equal(t, "region", data.Field)
		assert.Equal(t, "地区", data.Title)
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		cascader := NewCascader("city", "城市", []interface{}{"beijing", "chaoyang"})

		data := cascader.GetData()
		assert.NotNil(t, data.Value)
	})
}

// TestCascaderProperties 测试所有属性方法
func TestCascaderProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Cascader) *Cascader
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "Separator",
			setup:         func(c *Cascader) *Cascader { return c.Separator(" / ") },
			propKey:       "separator",
			expectedValue: " / ",
		},
		{
			name:          "Filterable",
			setup:         func(c *Cascader) *Cascader { return c.Filterable(true) },
			propKey:       "filterable",
			expectedValue: true,
		},
		{
			name:          "Clearable",
			setup:         func(c *Cascader) *Cascader { return c.Clearable(true) },
			propKey:       "clearable",
			expectedValue: true,
		},
		{
			name:          "ShowAllLevels",
			setup:         func(c *Cascader) *Cascader { return c.ShowAllLevels(false) },
			propKey:       "show-all-levels",
			expectedValue: false,
		},
		{
			name:          "Placeholder",
			setup:         func(c *Cascader) *Cascader { return c.Placeholder("请选择地区") },
			propKey:       "placeholder",
			expectedValue: "请选择地区",
		},
		{
			name:          "Disabled",
			setup:         func(c *Cascader) *Cascader { return c.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
		{
			name:          "Size",
			setup:         func(c *Cascader) *Cascader { return c.Size("large") },
			propKey:       "size",
			expectedValue: "large",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cascader := NewCascader("test", "测试")
			cascader = tt.setup(cascader)

			data := cascader.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestCascaderChaining 测试链式调用
func TestCascaderChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		options := []Option{
			{
				Value: "beijing",
				Label: "北京",
				Children: []Option{
					{Value: "chaoyang", Label: "朝阳区"},
					{Value: "haidian", Label: "海淀区"},
				},
			},
			{
				Value: "shanghai",
				Label: "上海",
				Children: []Option{
					{Value: "pudong", Label: "浦东新区"},
					{Value: "huangpu", Label: "黄浦区"},
				},
			},
		}

		cascader := NewCascader("region", "地区").
			SetOptions(options).
			Separator(" > ").
			Filterable(true).
			Clearable(true).
			Placeholder("请选择地区").
			Required()

		data := cascader.GetData()
		assert.Equal(t, " > ", data.Props["separator"])
		assert.Equal(t, true, data.Props["filterable"])
		assert.Equal(t, true, data.Props["clearable"])
		assert.Equal(t, "请选择地区", data.Props["placeholder"])
		assert.NotEmpty(t, data.Validate)
	})

	t.Run("CascaderPropsChain", func(t *testing.T) {
		props := map[string]interface{}{
			"value":    "id",
			"label":    "name",
			"children": "subItems",
		}

		cascader := NewCascader("custom", "自定义").
			CascaderProps(props).
			ShowAllLevels(false)

		data := cascader.GetData()
		propsValue, ok := data.Props["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "id", propsValue["value"])
		assert.Equal(t, "name", propsValue["label"])
		assert.Equal(t, "subItems", propsValue["children"])
		assert.Equal(t, false, data.Props["show-all-levels"])
	})
}

// TestCascaderBuild 测试Build方法
func TestCascaderBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		cascader := NewCascader("region", "地区")

		result := cascader.Build()

		assert.Equal(t, "cascader", result["type"])
		assert.Equal(t, "region", result["field"])
		assert.Equal(t, "地区", result["title"])
	})

	t.Run("BuildWithOptions", func(t *testing.T) {
		options := []Option{
			{Value: "1", Label: "选项1"},
			{Value: "2", Label: "选项2"},
		}

		cascader := NewCascader("test", "测试").
			SetOptions(options)

		result := cascader.Build()
		opts, ok := result["options"].([]map[string]interface{})
		require.True(t, ok)
		assert.Len(t, opts, 2)
		assert.Equal(t, "1", opts[0]["value"])
		assert.Equal(t, "选项1", opts[0]["label"])
	})

	t.Run("BuildWithNestedOptions", func(t *testing.T) {
		options := []Option{
			{
				Value: "parent",
				Label: "父级",
				Children: []Option{
					{Value: "child1", Label: "子级1"},
					{Value: "child2", Label: "子级2"},
				},
			},
		}

		cascader := NewCascader("nested", "嵌套").
			SetOptions(options)

		result := cascader.Build()
		opts := result["options"].([]map[string]interface{})
		assert.Len(t, opts, 1)
		assert.Equal(t, "parent", opts[0]["value"])

		children, ok := opts[0]["children"].([]map[string]interface{})
		require.True(t, ok)
		assert.Len(t, children, 2)
		assert.Equal(t, "child1", children[0]["value"])
		assert.Equal(t, "子级1", children[0]["label"])
	})

	t.Run("BuildWithCascaderProps", func(t *testing.T) {
		props := map[string]interface{}{
			"value": "code",
			"label": "title",
		}

		cascader := NewCascader("custom", "自定义").
			CascaderProps(props)

		result := cascader.Build()
		resultProps, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		cascaderProps, ok := resultProps["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "code", cascaderProps["value"])
		assert.Equal(t, "title", cascaderProps["label"])
	})
}

// TestCascaderEdgeCases 测试边缘情况
func TestCascaderEdgeCases(t *testing.T) {
	t.Run("EmptyOptions", func(t *testing.T) {
		cascader := NewCascader("test", "测试").
			SetOptions([]Option{})

		result := cascader.Build()
		_, hasOptions := result["options"]
		assert.False(t, hasOptions)
	})

	t.Run("ThreeLevelNesting", func(t *testing.T) {
		options := []Option{
			{
				Value: "china",
				Label: "中国",
				Children: []Option{
					{
						Value: "beijing",
						Label: "北京",
						Children: []Option{
							{Value: "chaoyang", Label: "朝阳区"},
							{Value: "haidian", Label: "海淀区"},
						},
					},
				},
			},
		}

		cascader := NewCascader("location", "位置").
			SetOptions(options)

		result := cascader.Build()
		opts := result["options"].([]map[string]interface{})
		level1 := opts[0]
		level2 := level1["children"].([]map[string]interface{})[0]
		level3, ok := level2["children"].([]map[string]interface{})
		require.True(t, ok)
		assert.Len(t, level3, 2)
	})

	t.Run("MixedSeparators", func(t *testing.T) {
		cascader := NewCascader("test", "测试").
			Separator(" / ").
			Separator(" > ")

		data := cascader.GetData()
		assert.Equal(t, " > ", data.Props["separator"])
	})

	t.Run("DisabledWithFilterable", func(t *testing.T) {
		cascader := NewCascader("test", "测试").
			Disabled(true).
			Filterable(true)

		data := cascader.GetData()
		assert.Equal(t, true, data.Props["disabled"])
		assert.Equal(t, true, data.Props["filterable"])
	})
}

// TestCascaderWithValidation 测试验证功能
func TestCascaderWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		cascader := NewCascader("region", "地区").Required()

		data := cascader.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("CustomValidation", func(t *testing.T) {
		cascader := NewCascader("city", "城市").
			Validate(PatternRule{
				Pattern: "^[a-z]+$",
				Message: "只能输入小写字母",
			})

		data := cascader.GetData()
		require.Len(t, data.Validate, 1)
	})
}

// TestCascaderInForm 测试在表单中的使用
func TestCascaderInForm(t *testing.T) {
	t.Run("FormWithCascader", func(t *testing.T) {
		options := []Option{
			{
				Value: "tech",
				Label: "技术",
				Children: []Option{
					{Value: "frontend", Label: "前端"},
					{Value: "backend", Label: "后端"},
				},
			},
		}

		form := NewElmForm("/submit", []Component{
			NewCascader("category", "分类").
				SetOptions(options).
				Required(),
			NewCascader("location", "位置").
				Placeholder("请选择位置"),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 2)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "cascader", rules[0]["type"])
		assert.Equal(t, "cascader", rules[1]["type"])
	})

	t.Run("CascaderWithControl", func(t *testing.T) {
		cascader := NewCascader("region_type", "地区类型", "domestic").
			SetOptions([]Option{
				{Value: "domestic", Label: "国内"},
				{Value: "international", Label: "国际"},
			}).
			Control([]ControlRule{
				{
					Value: "domestic",
					Rule: []Component{
						NewCascader("domestic_region", "国内地区").Required(),
					},
				},
				{
					Value: "international",
					Rule: []Component{
						NewCascader("country", "国家").Required(),
					},
				},
			})

		data := cascader.GetData()
		assert.Len(t, data.Control, 2)
		assert.Len(t, data.Control[0].Rule, 1)
	})
}

// BenchmarkCascaderCreation 性能测试
func BenchmarkCascaderCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewCascader("test", "测试")
	}
}

// BenchmarkCascaderWithProperties 性能测试
func BenchmarkCascaderWithProperties(b *testing.B) {
	options := []Option{
		{Value: "1", Label: "选项1"},
		{Value: "2", Label: "选项2"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewCascader("region", "地区").
			SetOptions(options).
			Separator(" > ").
			Filterable(true).
			Clearable(true).
			Required()
	}
}

// BenchmarkCascaderBuild 性能测试
func BenchmarkCascaderBuild(b *testing.B) {
	options := []Option{
		{
			Value: "parent",
			Label: "父级",
			Children: []Option{
				{Value: "child1", Label: "子级1"},
				{Value: "child2", Label: "子级2"},
			},
		},
	}

	cascader := NewCascader("nested", "嵌套").
		SetOptions(options).
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cascader.Build()
	}
}
