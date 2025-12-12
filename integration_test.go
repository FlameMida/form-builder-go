package formbuilder

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// integration_test.go 综合集成测试
// 测试多个组件和功能的集成工作

// TestCompleteUserRegistrationForm 测试完整的用户注册表单
func TestCompleteUserRegistrationForm(t *testing.T) {
	t.Run("CreateCompleteForm", func(t *testing.T) {
		form := NewElmForm("/api/user/register", []Component{
			NewInput("username", "用户名").
				Placeholder("请输入用户名").
				Clearable(true).
				MaxLength(50).
				Required().
				Validate(
					LengthRule{Min: 6, Max: 20, Message: "用户名长度6-20字符"},
					PatternRule{Pattern: "^[a-zA-Z0-9_]+$", Message: "只能包含字母数字下划线"},
				),

			Password("password", "密码").
				Placeholder("请输入密码").
				MinLength(6).
				ShowPassword(true).
				Required().
				Validate(
					LengthRule{Min: 8, Max: 32, Message: "密码长度8-32字符"},
				),

			Email("email", "邮箱").
				Placeholder("请输入邮箱").
				Required(),

			NewSelect("role", "角色").
				SetOptions([]Option{
					{Value: "user", Label: "普通用户"},
					{Value: "vip", Label: "VIP用户"},
					{Value: "admin", Label: "管理员"},
				}).
				Placeholder("请选择角色").
				Required(),

			NewRadio("gender", "性别", "male").
				SetOptions([]Option{
					{Value: "male", Label: "男"},
					{Value: "female", Label: "女"},
				}),

			NewSwitch("agree", "同意条款").
				ActiveText("同意").
				InactiveText("不同意").
				Required(),
		}, nil)

		// 验证表单创建成功
		assert.NotNil(t, form)
		assert.Len(t, form.rules, 6)

		// 验证JSON序列化
		jsonStr, err := form.ParseFormRule()
		require.NoError(t, err)
		assert.NotEmpty(t, jsonStr)

		var rules []map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &rules)
		require.NoError(t, err)
		assert.Len(t, rules, 6)

		// 验证每个组件的类型
		assert.Equal(t, "input", rules[0]["type"])
		assert.Equal(t, "input", rules[1]["type"])
		assert.Equal(t, "input", rules[2]["type"])
		assert.Equal(t, "select", rules[3]["type"])
		assert.Equal(t, "radio", rules[4]["type"])
		assert.Equal(t, "switch", rules[5]["type"])
	})
}

// TestFormWithControlIntegration 测试带条件显示的表单集成
func TestFormWithControlIntegration(t *testing.T) {
	t.Run("EmployeeTypeForm", func(t *testing.T) {
		form := NewElmForm("/api/employee/create", []Component{
			NewInput("name", "姓名").Required(),

			NewRadio("type", "员工类型", "trial").
				SetOptions([]Option{
					{Value: "trial", Label: "试用期"},
					{Value: "regular", Label: "正式员工"},
					{Value: "contract", Label: "合同工"},
				}).
				Control([]ControlRule{
					{
						Value: "trial",
						Rule: []Component{
							NewInputNumber("trial_days", "试用期天数").
								Min(1).
								Max(180).
								Required(),
						},
					},
					{
						Value: "regular",
						Rule: []Component{
							NewDatePicker("entry_date", "入职日期").Required(),
							NewInputNumber("salary", "薪资").Min(0).Required(),
						},
					},
					{
						Value: "contract",
						Rule: []Component{
							NewDatePicker("contract_start", "合同开始日期").Required(),
							NewDatePicker("contract_end", "合同结束日期").Required(),
						},
					},
				}),

			NewSelect("department", "部门").
				SetOptions([]Option{
					{Value: "tech", Label: "技术部"},
					{Value: "sales", Label: "销售部"},
					{Value: "hr", Label: "人力资源"},
				}).
				Required(),
		}, nil)

		// 验证表单结构
		rules := form.FormRule()
		require.Len(t, rules, 3)

		// 验证Control规则
		radioRule := rules[1]
		control, ok := radioRule["control"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, control, 3)

		// 验证试用期分支
		trialControl := control[0]
		assert.Equal(t, "trial", trialControl["value"])
		trialRules := trialControl["rule"].([]map[string]interface{})
		assert.Len(t, trialRules, 1)
		assert.Equal(t, "trial_days", trialRules[0]["field"])

		// 验证正式员工分支
		regularControl := control[1]
		assert.Equal(t, "regular", regularControl["value"])
		regularRules := regularControl["rule"].([]map[string]interface{})
		assert.Len(t, regularRules, 2)

		// 验证JSON序列化
		jsonStr, err := form.ParseFormRule()
		require.NoError(t, err)
		assert.True(t, strings.Contains(jsonStr, "control"))
	})
}

// TestFormDataApplicationIntegration 测试表单数据应用集成
func TestFormDataApplicationIntegration(t *testing.T) {
	t.Run("ApplyDataToMultipleComponents", func(t *testing.T) {
		form := NewElmForm("/api/update", []Component{
			NewInput("username", "用户名"),
			Email("email", "邮箱"),
			NewSelect("role", "角色").SetOptions([]Option{
				{Value: "admin", Label: "管理员"},
				{Value: "user", Label: "用户"},
			}),
			NewRadio("status", "状态").SetOptions([]Option{
				{Value: "1", Label: "启用"},
				{Value: "0", Label: "禁用"},
			}),
			NewSwitch("is_active", "是否激活"),
			NewInputNumber("age", "年龄"),
		}, nil)

		// 应用数据
		form.FormData(map[string]interface{}{
			"username":  "john_doe",
			"email":     "john@example.com",
			"role":      "admin",
			"status":    "1",
			"is_active": true,
			"age":       25,
		})

		rules := form.FormRule()

		// 验证每个字段的值都被正确应用
		assert.Equal(t, "john_doe", rules[0]["value"])
		assert.Equal(t, "john@example.com", rules[1]["value"])
		assert.Equal(t, "admin", rules[2]["value"])
		assert.Equal(t, "1", rules[3]["value"])
		assert.Equal(t, true, rules[4]["value"])
		assert.Equal(t, 25, rules[5]["value"])
	})

	t.Run("PartialDataApplication", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("field1", "字段1"),
			NewInput("field2", "字段2"),
			NewInput("field3", "字段3"),
		}, nil)

		// 只设置部分字段
		form.FormData(map[string]interface{}{
			"field1": "value1",
			"field3": "value3",
		})

		rules := form.FormRule()
		assert.Equal(t, "value1", rules[0]["value"])
		assert.Nil(t, rules[1]["value"]) // field2未设置
		assert.Equal(t, "value3", rules[2]["value"])
	})
}

// TestMultiUIFrameworkIntegration 测试多UI框架集成
func TestMultiUIFrameworkIntegration(t *testing.T) {
	components := []Component{
		NewInput("test", "测试").Required(),
	}

	t.Run("ElementUI", func(t *testing.T) {
		form := NewElmForm("/submit", components, nil)
		rules := form.FormRule()
		assert.Equal(t, "input", rules[0]["type"])
	})

	t.Run("IviewV3", func(t *testing.T) {
		form := NewIviewForm("/submit", []Component{
			NewIviewInput("test", "测试").Required(),
		}, nil)
		rules := form.FormRule()
		assert.Equal(t, "input", rules[0]["type"])
	})

	t.Run("IviewV4", func(t *testing.T) {
		form := NewIview4Form("/submit", []Component{
			NewIviewInput("test", "测试").Required(),
		}, nil)
		assert.NotNil(t, form)
	})
}

// TestComplexValidationIntegration 测试复杂验证规则集成
func TestComplexValidationIntegration(t *testing.T) {
	t.Run("MultipleValidationRules", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewInput("username", "用户名").
				Required().
				Validate(
					LengthRule{Min: 6, Max: 20, Message: "长度6-20"},
					PatternRule{Pattern: "^[a-zA-Z0-9]+$", Message: "只能字母数字"},
				),

			Email("email", "邮箱").
				Required(),

			NewInput("phone", "手机号").
				Validate(
					PatternRule{Pattern: "^1[3-9]\\d{9}$", Message: "手机号格式错误"},
				),

			NewInputNumber("age", "年龄").
				Validate(
					RangeRule{Min: 18, Max: 100, Message: "年龄18-100"},
				),
		}, nil)

		rules := form.FormRule()

		// 验证username的多个验证规则
		usernameValidates := rules[0]["validate"].([]map[string]interface{})
		require.Len(t, usernameValidates, 3) // Required + 2个自定义

		// 验证email自动添加的验证
		emailValidates := rules[1]["validate"].([]map[string]interface{})
		require.NotEmpty(t, emailValidates)
		hasEmailRule := false
		for _, v := range emailValidates {
			if v["type"] == "email" {
				hasEmailRule = true
				break
			}
		}
		assert.True(t, hasEmailRule)
	})
}

// TestNestedControlIntegration 测试嵌套Control集成
func TestNestedControlIntegration(t *testing.T) {
	t.Run("ThreeLevelNesting", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewRadio("level1", "一级分类", "a").
				SetOptions([]Option{
					{Value: "a", Label: "分类A"},
					{Value: "b", Label: "分类B"},
				}).
				Control([]ControlRule{
					{
						Value: "a",
						Rule: []Component{
							NewRadio("level2", "二级分类", "x").
								SetOptions([]Option{
									{Value: "x", Label: "X类"},
									{Value: "y", Label: "Y类"},
								}).
								Control([]ControlRule{
									{
										Value: "x",
										Rule: []Component{
											NewInput("level3", "详细信息").Required(),
										},
									},
								}),
						},
					},
				}),
		}, nil)

		jsonStr, err := form.ParseFormRule()
		require.NoError(t, err)

		var rules []map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &rules)
		require.NoError(t, err)

		// 验证嵌套结构
		control1 := rules[0]["control"].([]interface{})
		require.Len(t, control1, 1)

		rule1 := control1[0].(map[string]interface{})
		level2Rules := rule1["rule"].([]interface{})
		level2 := level2Rules[0].(map[string]interface{})

		control2 := level2["control"].([]interface{})
		require.Len(t, control2, 1)
	})
}

// TestFormConfigIntegration 测试表单配置集成
func TestFormConfigIntegration(t *testing.T) {
	t.Run("CustomConfig", func(t *testing.T) {
		config := NewElmConfig()
		config.SubmitBtn(true, "提交注册")
		config.ResetBtn(true, "重置表单")

		form := NewElmForm("/submit", []Component{
			NewInput("test", "测试"),
		}, config)

		assert.NotNil(t, form)
		assert.Equal(t, config, form.config)

		configMap := config.ToMap()
		submitBtn := configMap["submitBtn"].(map[string]interface{})
		assert.Equal(t, true, submitBtn["show"])
		assert.Equal(t, "提交注册", submitBtn["innerText"])
	})
}

// TestLargeFormPerformance 测试大型表单性能
func TestLargeFormPerformance(t *testing.T) {
	t.Run("FormWith50Components", func(t *testing.T) {
		components := make([]Component, 50)
		for i := 0; i < 50; i++ {
			components[i] = NewInput(
				"field"+string(rune(i)),
				"字段"+string(rune(i)),
			).Required()
		}

		form := NewElmForm("/submit", components, nil)

		// 验证表单创建成功
		assert.NotNil(t, form)
		assert.Len(t, form.rules, 50)

		// 验证JSON序列化不出错
		jsonStr, err := form.ParseFormRule()
		require.NoError(t, err)
		assert.NotEmpty(t, jsonStr)
	})
}

// TestRealWorldScenarios 测试真实场景
func TestRealWorldScenarios(t *testing.T) {
	t.Run("ArticlePublishForm", func(t *testing.T) {
		form := NewElmForm("/api/article/publish", []Component{
			NewInput("title", "标题").
				Placeholder("请输入文章标题").
				MaxLength(100).
				Required(),

			NewInput("author", "作者").Required(),

			NewSelect("category", "分类").
				SetOptions([]Option{
					{Value: "tech", Label: "技术"},
					{Value: "life", Label: "生活"},
					{Value: "travel", Label: "旅游"},
				}).
				Required(),

			NewRadio("publish_type", "发布方式", "immediate").
				SetOptions([]Option{
					{Value: "immediate", Label: "立即发布"},
					{Value: "scheduled", Label: "定时发布"},
				}).
				Control([]ControlRule{
					{
						Value: "scheduled",
						Rule: []Component{
							NewDatePicker("publish_time", "发布时间").
								DateType("datetime").
								Required(),
						},
					},
				}),

			NewSwitch("allow_comment", "允许评论").
				ActiveText("允许").
				InactiveText("禁止"),
		}, nil)

		// 应用数据
		form.FormData(map[string]interface{}{
			"title":         "Go语言测试最佳实践",
			"author":        "张三",
			"category":      "tech",
			"publish_type":  "scheduled",
			"allow_comment": true,
			"publish_time":  "2024-12-31 10:00:00",
		})

		rules := form.FormRule()
		assert.Len(t, rules, 5)

		// 验证数据应用
		assert.Equal(t, "Go语言测试最佳实践", rules[0]["value"])
		assert.Equal(t, "tech", rules[2]["value"])

		// 验证JSON生成
		jsonStr, err := form.ParseFormRule()
		require.NoError(t, err)
		assert.True(t, strings.Contains(jsonStr, "title"))
		assert.True(t, strings.Contains(jsonStr, "control"))
	})
}

// BenchmarkCompleteFormCreation 性能测试：完整表单创建
func BenchmarkCompleteFormCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewElmForm("/submit", []Component{
			NewInput("username", "用户名").Required(),
			Password("password", "密码").Required(),
			Email("email", "邮箱").Required(),
			NewSelect("role", "角色").SetOptions([]Option{
				{Value: "admin", Label: "管理员"},
				{Value: "user", Label: "用户"},
			}),
			NewRadio("status", "状态").SetOptions([]Option{
				{Value: "1", Label: "启用"},
				{Value: "0", Label: "禁用"},
			}),
		}, nil)
	}
}

// BenchmarkFormWithControl 性能测试：带Control的表单
func BenchmarkFormWithControl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		form := NewElmForm("/submit", []Component{
			NewRadio("type", "类型", "1").
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
				}),
		}, nil)
		_ = form.FormRule()
	}
}
