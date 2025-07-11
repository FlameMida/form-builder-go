package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequiredRule(t *testing.T) {
	t.Run("创建必填规则", func(t *testing.T) {
		rule := NewRequiredRule("此字段为必填项")
		
		assert.Equal(t, "required", rule.Type())
		assert.Equal(t, "此字段为必填项", rule.Message())
	})
	
	t.Run("默认消息", func(t *testing.T) {
		rule := NewRequiredRule("")
		assert.Equal(t, "此字段为必填项", rule.Message())
	})
	
	t.Run("验证空值", func(t *testing.T) {
		rule := NewRequiredRule("必填")
		
		// 测试nil
		err := rule.Validate(nil)
		assert.Error(t, err)
		assert.Equal(t, "必填", err.Error())
		
		// 测试空字符串
		err = rule.Validate("")
		assert.Error(t, err)
		
		// 测试空白字符串
		err = rule.Validate("   ")
		assert.Error(t, err)
		
		// 测试空数组
		err = rule.Validate([]interface{}{})
		assert.Error(t, err)
	})
	
	t.Run("验证有效值", func(t *testing.T) {
		rule := NewRequiredRule("必填")
		
		// 测试非空字符串
		err := rule.Validate("valid")
		assert.NoError(t, err)
		
		// 测试数字
		err = rule.Validate(123)
		assert.NoError(t, err)
		
		// 测试布尔值
		err = rule.Validate(true)
		assert.NoError(t, err)
		
		// 测试非空数组
		err = rule.Validate([]interface{}{"item"})
		assert.NoError(t, err)
	})
}

func TestLengthRule(t *testing.T) {
	t.Run("创建长度规则", func(t *testing.T) {
		rule := NewLengthRule(2, 10, "长度必须在2到10之间")
		
		assert.Equal(t, "length", rule.Type())
		assert.Equal(t, 2, rule.Min)
		assert.Equal(t, 10, rule.Max)
		assert.Equal(t, "长度必须在2到10之间", rule.Message())
	})
	
	t.Run("默认消息", func(t *testing.T) {
		rule := NewLengthRule(5, 15, "")
		assert.Contains(t, rule.Message(), "5")
		assert.Contains(t, rule.Message(), "15")
	})
	
	t.Run("验证字符串长度", func(t *testing.T) {
		rule := NewLengthRule(3, 8, "长度错误")
		
		// 测试nil值（应该通过）
		err := rule.Validate(nil)
		assert.NoError(t, err)
		
		// 测试太短
		err = rule.Validate("ab")
		assert.Error(t, err)
		
		// 测试太长
		err = rule.Validate("abcdefghijk")
		assert.Error(t, err)
		
		// 测试有效长度
		err = rule.Validate("abc")
		assert.NoError(t, err)
		
		err = rule.Validate("abcdefgh")
		assert.NoError(t, err)
		
		// 测试中文字符
		err = rule.Validate("你好世界")
		assert.NoError(t, err)
	})
	
	t.Run("非字符串类型", func(t *testing.T) {
		rule := NewLengthRule(1, 10, "长度错误")
		
		err := rule.Validate(123)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "字符串类型")
	})
}

func TestEmailRule(t *testing.T) {
	t.Run("创建邮箱规则", func(t *testing.T) {
		rule := NewEmailRule("邮箱格式错误")
		
		assert.Equal(t, "email", rule.Type())
		assert.Equal(t, "邮箱格式错误", rule.Message())
	})
	
	t.Run("默认消息", func(t *testing.T) {
		rule := NewEmailRule("")
		assert.Contains(t, rule.Message(), "邮箱")
	})
	
	t.Run("验证邮箱格式", func(t *testing.T) {
		rule := NewEmailRule("邮箱格式错误")
		
		// 测试nil值（应该通过）
		err := rule.Validate(nil)
		assert.NoError(t, err)
		
		// 测试空字符串（应该通过）
		err = rule.Validate("")
		assert.NoError(t, err)
		
		// 测试有效邮箱
		validEmails := []string{
			"test@example.com",
			"user.name@domain.co.uk",
			"test+tag@example.org",
			"123@456.com",
		}
		
		for _, email := range validEmails {
			err = rule.Validate(email)
			assert.NoError(t, err, "应该通过验证: %s", email)
		}
		
		// 测试无效邮箱
		invalidEmails := []string{
			"invalid",
			"@example.com",
			"test@",
			"test.example.com",
			"test@example",
		}
		
		for _, email := range invalidEmails {
			err = rule.Validate(email)
			assert.Error(t, err, "应该验证失败: %s", email)
		}
	})
	
	t.Run("非字符串类型", func(t *testing.T) {
		rule := NewEmailRule("邮箱格式错误")
		
		err := rule.Validate(123)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "字符串类型")
	})
}

func TestNumberRule(t *testing.T) {
	t.Run("创建数字规则", func(t *testing.T) {
		min := 10.0
		max := 100.0
		rule := NewNumberRule(&min, &max, "数值范围错误")
		
		assert.Equal(t, "number", rule.Type())
		assert.Equal(t, &min, rule.Min)
		assert.Equal(t, &max, rule.Max)
		assert.Equal(t, "数值范围错误", rule.Message())
	})
	
	t.Run("只有最小值", func(t *testing.T) {
		min := 5.0
		rule := NewNumberRule(&min, nil, "")
		assert.Contains(t, rule.Message(), "5")
		assert.Contains(t, rule.Message(), "大于等于")
	})
	
	t.Run("只有最大值", func(t *testing.T) {
		max := 100.0
		rule := NewNumberRule(nil, &max, "")
		assert.Contains(t, rule.Message(), "100")
		assert.Contains(t, rule.Message(), "小于等于")
	})
	
	t.Run("验证数字范围", func(t *testing.T) {
		min := 10.0
		max := 100.0
		rule := NewNumberRule(&min, &max, "范围错误")
		
		// 测试nil值（应该通过）
		err := rule.Validate(nil)
		assert.NoError(t, err)
		
		// 测试太小
		err = rule.Validate(5.0)
		assert.Error(t, err)
		
		err = rule.Validate(5)
		assert.Error(t, err)
		
		err = rule.Validate("5")
		assert.Error(t, err)
		
		// 测试太大
		err = rule.Validate(150.0)
		assert.Error(t, err)
		
		err = rule.Validate(150)
		assert.Error(t, err)
		
		err = rule.Validate("150")
		assert.Error(t, err)
		
		// 测试有效范围
		err = rule.Validate(50.0)
		assert.NoError(t, err)
		
		err = rule.Validate(50)
		assert.NoError(t, err)
		
		err = rule.Validate("50")
		assert.NoError(t, err)
		
		// 测试边界值
		err = rule.Validate(10.0)
		assert.NoError(t, err)
		
		err = rule.Validate(100.0)
		assert.NoError(t, err)
	})
	
	t.Run("无效数字字符串", func(t *testing.T) {
		rule := NewNumberRule(nil, nil, "")
		
		err := rule.Validate("abc")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "有效的数值")
	})
	
	t.Run("空字符串", func(t *testing.T) {
		rule := NewNumberRule(nil, nil, "")
		
		err := rule.Validate("")
		assert.NoError(t, err)
	})
}

func TestPatternRule(t *testing.T) {
	t.Run("创建模式规则", func(t *testing.T) {
		rule, err := NewPatternRule(`^\d{3}-\d{4}$`, "格式错误")
		require.NoError(t, err)
		
		assert.Equal(t, "pattern", rule.Type())
		assert.Equal(t, "格式错误", rule.Message())
	})
	
	t.Run("无效正则表达式", func(t *testing.T) {
		_, err := NewPatternRule(`[`, "")
		assert.Error(t, err)
	})
	
	t.Run("验证模式匹配", func(t *testing.T) {
		// 电话号码格式
		rule, err := NewPatternRule(`^\d{3}-\d{4}$`, "电话号码格式错误")
		require.NoError(t, err)
		
		// 测试nil值（应该通过）
		err = rule.Validate(nil)
		assert.NoError(t, err)
		
		// 测试空字符串（应该通过）
		err = rule.Validate("")
		assert.NoError(t, err)
		
		// 测试匹配的格式
		err = rule.Validate("123-4567")
		assert.NoError(t, err)
		
		// 测试不匹配的格式
		err = rule.Validate("1234567")
		assert.Error(t, err)
		
		err = rule.Validate("123-45678")
		assert.Error(t, err)
		
		err = rule.Validate("abc-defg")
		assert.Error(t, err)
	})
	
	t.Run("非字符串类型", func(t *testing.T) {
		rule, err := NewPatternRule(`\d+`, "")
		require.NoError(t, err)
		
		err = rule.Validate(123)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "字符串类型")
	})
}

func TestValidatorChain(t *testing.T) {
	t.Run("创建验证链", func(t *testing.T) {
		chain := NewValidatorChain()
		assert.NotNil(t, chain)
		assert.Len(t, chain.GetRules(), 0)
	})
	
	t.Run("添加规则", func(t *testing.T) {
		chain := NewValidatorChain()
		
		chain.Required("必填").
			Length(5, 20, "长度错误").
			Email("邮箱格式错误")
		
		rules := chain.GetRules()
		assert.Len(t, rules, 3)
		assert.Equal(t, "required", rules[0].Type())
		assert.Equal(t, "length", rules[1].Type())
		assert.Equal(t, "email", rules[2].Type())
	})
	
	t.Run("验证链执行", func(t *testing.T) {
		chain := NewValidatorChain()
		chain.Required("必填").
			Length(5, 20, "长度必须在5-20之间")
		
		// 测试通过所有规则
		err := chain.Validate("valid@email.com")
		assert.NoError(t, err)
		
		// 测试第一个规则失败
		err = chain.Validate("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "必填")
		
		// 测试第二个规则失败
		err = chain.Validate("abc")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "长度")
	})
	
	t.Run("数字范围验证链", func(t *testing.T) {
		min := 0.0
		max := 100.0
		chain := NewValidatorChain()
		chain.Required("分数必填").
			Number(&min, &max, "分数必须在0-100之间")
		
		// 测试有效分数
		err := chain.Validate("85.5")
		assert.NoError(t, err)
		
		err = chain.Validate(95)
		assert.NoError(t, err)
		
		// 测试无效分数
		err = chain.Validate("150")
		assert.Error(t, err)
		
		err = chain.Validate("-10")
		assert.Error(t, err)
	})
	
	t.Run("模式验证链", func(t *testing.T) {
		chain := NewValidatorChain()
		chain.Required("手机号必填").
			Pattern(`^1[3-9]\d{9}$`, "手机号格式错误")
		
		// 测试有效手机号
		err := chain.Validate("13812345678")
		assert.NoError(t, err)
		
		// 测试无效手机号
		err = chain.Validate("12812345678")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "格式错误")
	})
}