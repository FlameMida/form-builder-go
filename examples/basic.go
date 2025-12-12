package main

import (
	"fmt"
	"log"

	fb "github.com/maverick/form-builder-go/formbuilder"
)

// basic.go - 基础用法示例

func main() {
	// 创建一个简单的登录表单
	form := fb.Elm.CreateForm("/api/login", []fb.Component{
		fb.Elm.Input("username", "用户名").
			Placeholder("请输入用户名").
			Required(),

		fb.Elm.Password("password", "密码").
			Placeholder("请输入密码").
			Required(),

		fb.Elm.Checkbox("remember", "记住我").
			SetOptions([]fb.Option{
				{Value: "1", Label: "记住登录状态"},
			}),
	})

	// 获取JSON规则
	jsonRule, err := form.ParseFormRule()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("表单规则（JSON）:")
	fmt.Println(jsonRule)

	// 获取HTML视图
	html, err := form.View()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nHTML页面:")
	fmt.Println(html)
}
