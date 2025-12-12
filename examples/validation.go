package main

import (
	"fmt"
	"log"

	fb "github.com/maverick/form-builder-go/formbuilder"
)

// validation.go - 验证规则示例

func main() {
	form := fb.Elm.CreateForm("/api/register", []fb.Component{
		// 用户名：必填 + 长度验证
		fb.Elm.Input("username", "用户名").
			Placeholder("6-20个字符").
			Validate(
				fb.NewRequired("用户名不能为空"),
				fb.NewLength(6, 20, "用户名长度必须在6-20个字符之间"),
				fb.NewPattern("^[a-zA-Z0-9_]+$", "用户名只能包含字母、数字和下划线"),
			),

		// 邮箱：必填 + 邮箱格式验证
		fb.Email("email", "邮箱").
			Placeholder("请输入邮箱"),

		// 手机号：正则验证
		fb.Elm.Input("phone", "手机号").
			Placeholder("请输入手机号").
			Validate(
				fb.NewRequired("手机号不能为空"),
				fb.NewPattern("^1[3-9]\\d{9}$", "请输入正确的手机号"),
			),

		// 年龄：数值范围验证
		fb.Elm.Number("age", "年龄").
			Min(18).
			Max(100).
			Validate(
				fb.NewRequired("年龄不能为空"),
				fb.NewRange(18, 100, "年龄必须在18-100之间"),
			),

		// 密码：必填 + 长度验证
		fb.Elm.Password("password", "密码").
			Placeholder("8-20个字符").
			Validate(
				fb.NewRequired("密码不能为空"),
				fb.NewLength(8, 20, "密码长度必须在8-20个字符之间"),
			),

		// 确认密码：自定义验证（JavaScript）
		fb.Elm.Password("password_confirm", "确认密码").
			Placeholder("请再次输入密码").
			Validate(
				fb.NewRequired("请确认密码"),
				fb.CustomRule{
					Validator: "function(rule, value, callback) { if(value !== this.form.password) { callback(new Error('两次密码不一致')); } else { callback(); } }",
					Message:   "两次密码不一致",
				},
			),
	})

	jsonRule, err := form.ParseFormRule()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("带完整验证规则的表单:")
	fmt.Println(jsonRule)
}
