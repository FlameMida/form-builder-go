package components

import (
	"testing"

	"github.com/FlameMida/form-builder-go/contracts"
	formbuildererrors "github.com/FlameMida/form-builder-go/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidationErrorHandling(t *testing.T) {
	t.Run("组件验证错误", func(t *testing.T) {
		input := NewInput("test", "测试字段")
		input.Required()

		// 空值应该验证失败
		err := input.DoValidate()
		assert.Error(t, err)

		// 检查是否是正确的验证错误类型
		var fbErr *formbuildererrors.FormBuilderError
		assert.ErrorAs(t, err, &fbErr)
		assert.Equal(t, formbuildererrors.ErrTypeValidation, fbErr.Type)
		assert.Equal(t, "test", fbErr.Field)
		assert.Contains(t, err.Error(), "test")
	})

	t.Run("组件验证成功", func(t *testing.T) {
		input := NewInput("test", "测试字段")
		input.Required()
		input.SetValue("有效值")

		// 有值应该验证成功
		err := input.DoValidate()
		assert.NoError(t, err)
	})

	t.Run("多组件验证", func(t *testing.T) {
		input1 := NewInput("field1", "字段1")
		input1.Required()
		input1.SetValue("值1")

		input2 := NewInput("field2", "字段2")
		input2.Required()
		// input2 没有设置值，应该失败

		components := []contracts.Component{input1, input2}

		err := ValidateComponents(components)
		assert.Error(t, err)

		var fbErr *formbuildererrors.FormBuilderError
		assert.ErrorAs(t, err, &fbErr)
		assert.Equal(t, "field2", fbErr.Field)
	})
}
