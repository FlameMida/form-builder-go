package components

import (
	"testing"

	"github.com/FlameMida/form-builder-go/rules"
	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	t.Run("创建Input组件", func(t *testing.T) {
		input := NewInput("test_field", "测试字段")

		assert.Equal(t, "test_field", input.Field())
		assert.Equal(t, "测试字段", input.Title())
		assert.Nil(t, input.GetValue())
	})

	t.Run("设置Input属性", func(t *testing.T) {
		input := NewInput("email", "邮箱")
		input.Type("email")
		input.Placeholder("请输入邮箱")
		input.Required()
		input.Maxlength(100)

		// 构建组件
		built := input.Build()

		assert.Equal(t, "email", built["field"])
		assert.Equal(t, "邮箱", built["title"])
		assert.Equal(t, "el-input", built["type"])

		props := built["props"].(map[string]interface{})
		assert.Equal(t, "email", props["type"])
		assert.Equal(t, "请输入邮箱", props["placeholder"])
		assert.Equal(t, 100, props["maxlength"])

		// 检查验证规则
		validateRules := built["validate"].([]map[string]interface{})
		assert.Len(t, validateRules, 1)
		assert.Equal(t, true, validateRules[0]["required"])
	})

	t.Run("Input值设置和验证", func(t *testing.T) {
		input := NewInput("username", "用户名").Required()

		// 测试空值验证
		input.SetValue("")
		err := input.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "必填项")

		// 测试有效值验证
		input.SetValue("validuser")
		err = input.Validate()
		assert.NoError(t, err)
		assert.Equal(t, "validuser", input.GetValue())
	})
}

func TestTextarea(t *testing.T) {
	t.Run("创建Textarea组件", func(t *testing.T) {
		textarea := NewTextarea("content", "内容")

		assert.Equal(t, "content", textarea.Field())
		assert.Equal(t, "内容", textarea.Title())

		built := textarea.Build()
		assert.Equal(t, "el-input", built["type"])

		props := built["props"].(map[string]interface{})
		assert.Equal(t, "textarea", props["type"])
	})

	t.Run("设置Textarea属性", func(t *testing.T) {
		textarea := NewTextarea("description", "描述")
		textarea.Rows(5).
			Placeholder("请输入描述")

		built := textarea.Build()
		props := built["props"].(map[string]interface{})

		assert.Equal(t, 5, props["rows"])
		assert.Equal(t, "请输入描述", props["placeholder"])
	})
}

func TestSwitch(t *testing.T) {
	t.Run("创建Switch组件", func(t *testing.T) {
		switchComp := NewSwitch("enabled", "启用状态")

		assert.Equal(t, "enabled", switchComp.Field())
		assert.Equal(t, "启用状态", switchComp.Title())

		built := switchComp.Build()
		assert.Equal(t, "el-switch", built["type"])
	})

	t.Run("设置Switch属性", func(t *testing.T) {
		switchComp := NewSwitch("status", "状态")
		switchComp.ActiveText("开启").
			InactiveText("关闭").
			ActiveValue(1).
			InactiveValue(0).
			Required()

		built := switchComp.Build()
		props := built["props"].(map[string]interface{})

		assert.Equal(t, "开启", props["active-text"])
		assert.Equal(t, "关闭", props["inactive-text"])
		assert.Equal(t, 1, props["active-value"])
		assert.Equal(t, 0, props["inactive-value"])

		// 检查验证规则
		validateRules := built["validate"].([]map[string]interface{})
		assert.Len(t, validateRules, 1)
		assert.Equal(t, true, validateRules[0]["required"])
	})

	t.Run("Switch验证", func(t *testing.T) {
		switchComp := NewSwitch("agree", "同意条款").Required()

		// 测试空值验证
		switchComp.SetValue(nil)
		err := switchComp.Validate()
		assert.Error(t, err)

		// 测试有效值验证
		switchComp.SetValue(true)
		err = switchComp.Validate()
		assert.NoError(t, err)
	})
}

func TestBaseComponent(t *testing.T) {
	t.Run("基础组件功能", func(t *testing.T) {
		base := NewBaseComponent("test", "测试")

		assert.Equal(t, "test", base.Field())
		assert.Equal(t, "测试", base.Title())
		assert.Nil(t, base.GetValue())

		// 设置属性
		base.SetProp("placeholder", "测试占位符")
		assert.Equal(t, "测试占位符", base.GetProp("placeholder"))

		// 设置值
		base.SetValue("test_value")
		assert.Equal(t, "test_value", base.GetValue())

		// 构建组件
		built := base.Build()
		assert.Equal(t, "test", built["field"])
		assert.Equal(t, "测试", built["title"])
		assert.Equal(t, "test_value", built["value"])
		assert.Equal(t, "测试占位符", built["props"].(map[string]interface{})["placeholder"])
	})

	t.Run("验证规则管理", func(t *testing.T) {
		base := NewBaseComponent("test", "测试")

		// 添加验证规则
		rule := rules.NewRequiredRule("必填")
		base.AddValidateRule(rule)

		rules := base.GetValidateRules()
		assert.Len(t, rules, 1)
		assert.Equal(t, rule, rules[0])

		// 验证功能
		base.SetValue("")
		err := base.Validate()
		assert.Error(t, err)

		base.SetValue("有效值")
		err = base.Validate()
		assert.NoError(t, err)
	})
}

func TestRequiredRule(t *testing.T) {
	t.Run("必填规则验证", func(t *testing.T) {
		rule := rules.NewRequiredRule("此字段必填")

		assert.Equal(t, "required", rule.Type())
		assert.Equal(t, "此字段必填", rule.Message())

		// 测试空值
		err := rule.Validate(nil)
		assert.Error(t, err)
		assert.Equal(t, "此字段必填", err.Error())

		err = rule.Validate("")
		assert.Error(t, err)

		err = rule.Validate("   ")
		assert.Error(t, err)

		// 测试有效值
		err = rule.Validate("valid")
		assert.NoError(t, err)

		// 测试ToMap
		ruleMap := rule.ToMap()
		assert.Equal(t, true, ruleMap["required"])
		assert.Equal(t, "此字段必填", ruleMap["message"])
		assert.Equal(t, "required", ruleMap["type"])
	})
}
