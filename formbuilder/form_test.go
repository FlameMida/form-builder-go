package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/FlameMida/form-builder-go/components"
	"github.com/FlameMida/form-builder-go/ui/elm"
)

func TestFormBuilder(t *testing.T) {
	t.Run("创建表单", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		form, err := NewForm(bootstrap, "/test", []interface{}{}, map[string]interface{}{})

		require.NoError(t, err)
		assert.NotNil(t, form)
		assert.Equal(t, "/test", form.GetAction())
		assert.Equal(t, "POST", form.GetMethod())
		assert.Equal(t, "", form.GetTitle()) // 默认title为空
	})

	t.Run("设置表单属性", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		form, err := NewForm(bootstrap, "", []interface{}{}, map[string]interface{}{})
		require.NoError(t, err)

		form.SetAction("/save").
			SetMethod("PUT").
			SetTitle("测试表单")

		assert.Equal(t, "/save", form.GetAction())
		assert.Equal(t, "PUT", form.GetMethod())
		assert.Equal(t, "测试表单", form.GetTitle())
	})

	t.Run("添加组件", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		form, err := NewForm(bootstrap, "", []interface{}{}, map[string]interface{}{})
		require.NoError(t, err)

		input := components.NewInput("name", "姓名")
		textarea := components.NewTextarea("desc", "描述")

		form.Append(input).Append(textarea)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "name", rules[0]["field"])
		assert.Equal(t, "desc", rules[1]["field"])
	})

	t.Run("前置添加组件", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		input1 := components.NewInput("field1", "字段1")
		form, err := NewForm(bootstrap, "", []interface{}{input1}, map[string]interface{}{})
		require.NoError(t, err)

		input2 := components.NewInput("field2", "字段2")
		form.Prepend(input2)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "field2", rules[0]["field"])
		assert.Equal(t, "field1", rules[1]["field"])
	})

	t.Run("设置表单数据", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		input := components.NewInput("username", "用户名")
		form, err := NewForm(bootstrap, "", []interface{}{input}, map[string]interface{}{})
		require.NoError(t, err)

		form.FormData(map[string]interface{}{
			"username": "testuser",
			"other":    "value",
		})

		rules := form.FormRule()
		assert.Equal(t, "testuser", rules[0]["value"])
	})

	t.Run("设置单个字段值", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		input := components.NewInput("email", "邮箱")
		form, err := NewForm(bootstrap, "", []interface{}{input}, map[string]interface{}{})
		require.NoError(t, err)

		form.SetValue("email", "test@example.com")

		rules := form.FormRule()
		assert.Equal(t, "test@example.com", rules[0]["value"])
	})

	t.Run("字段唯一性检查", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		input1 := components.NewInput("name", "姓名1")
		input2 := components.NewInput("name", "姓名2") // 重复字段名

		_, err := NewForm(bootstrap, "", []interface{}{input1, input2}, map[string]interface{}{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "不能重复")
	})

	t.Run("生成JSON规则", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		input := components.NewInput("test", "测试")
		input.Placeholder("请输入测试值").Required()

		form, err := NewForm(bootstrap, "", []interface{}{input}, map[string]interface{}{})
		require.NoError(t, err)

		jsonBytes, err := form.ParseFormRule()
		require.NoError(t, err)

		jsonStr := string(jsonBytes)
		assert.Contains(t, jsonStr, "test")
		assert.Contains(t, jsonStr, "测试")
		assert.Contains(t, jsonStr, "请输入测试值")
	})

	t.Run("生成JSON配置", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		config := map[string]interface{}{
			"submitBtn": true,
			"resetBtn":  false,
		}

		form, err := NewForm(bootstrap, "", []interface{}{}, config)
		require.NoError(t, err)

		jsonBytes, err := form.ParseFormConfig()
		require.NoError(t, err)

		jsonStr := string(jsonBytes)
		assert.Contains(t, jsonStr, "submitBtn")
		assert.Contains(t, jsonStr, "resetBtn")
	})

	t.Run("设置HTTP头", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		form, err := NewForm(bootstrap, "", []interface{}{}, map[string]interface{}{})
		require.NoError(t, err)

		form.SetHeader("Authorization", "Bearer token").
			SetHeader("Content-SetPropType", "application/json")

		headersJSON, err := form.ParseHeaders()
		require.NoError(t, err)

		headersStr := string(headersJSON)
		assert.Contains(t, headersStr, "Authorization")
		assert.Contains(t, headersStr, "Bearer token")
		assert.Contains(t, headersStr, "Content-SetPropType")
	})

	t.Run("设置多个HTTP头", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		form, err := NewForm(bootstrap, "", []interface{}{}, map[string]interface{}{})
		require.NoError(t, err)

		headers := map[string]string{
			"X-Custom-Header": "custom-value",
			"X-Api-Key":       "api-key-value",
		}
		form.SetHeaders(headers)

		headersJSON, err := form.ParseHeaders()
		require.NoError(t, err)

		headersStr := string(headersJSON)
		assert.Contains(t, headersStr, "X-Custom-Header")
		assert.Contains(t, headersStr, "X-Api-Key")
	})

	t.Run("设置依赖脚本", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		form, err := NewForm(bootstrap, "", []interface{}{}, map[string]interface{}{})
		require.NoError(t, err)

		scripts := []string{
			"<script src=\"custom1.js\"></script>",
			"<script src=\"custom2.js\"></script>",
		}
		form.SetDependScript(scripts)

		scriptStr := form.ParseDependScript()
		assert.Contains(t, scriptStr, "custom1.js")
		assert.Contains(t, scriptStr, "custom2.js")
	})

	t.Run("渲染表单视图", func(t *testing.T) {
		bootstrap := elm.NewBootstrap()
		input := components.NewInput("name", "姓名")

		form, err := NewForm(bootstrap, "/save", []interface{}{input}, map[string]interface{}{
			"submitBtn": true,
		})
		require.NoError(t, err)

		form.SetTitle("用户注册表单")

		html, err := form.View()
		require.NoError(t, err)

		assert.Contains(t, html, "用户注册表单")
		assert.Contains(t, html, "/save")
		assert.Contains(t, html, "POST")
		assert.Contains(t, html, "form-create")
	})
}
