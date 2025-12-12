package main

import (
	"fmt"
	"log"

	fb "github.com/maverick/form-builder-go/formbuilder"
)

// demo.go - 完整的表单示例
// 对应PHP版本的demo/elm.php

func main() {
	// 创建表单规则
	rules := []fb.Component{
		// 基础输入框
		fb.Elm.Input("username", "用户名").
			Placeholder("请输入用户名").
			Clearable(true).
			Required(),

		// 密码输入框
		fb.Elm.Password("password", "密码").
			Placeholder("请输入密码").
			ShowPassword(true).
			Required(),

		// 邮箱输入框（带邮箱验证）
		fb.Email("email", "邮箱").
			Placeholder("请输入邮箱").
			Required(),

		// 单选框带条件显示
		fb.Elm.Radio("user_type", "用户类型", "1").
			SetOptions([]fb.Option{
				{Value: "1", Label: "试用期用户"},
				{Value: "2", Label: "正式用户"},
			}).
			Control([]fb.ControlRule{
				{
					Value: "1",
					Rule: []fb.Component{
						fb.Elm.Number("trial_days", "试用天数").
							Min(1).
							Max(30).
							Required(),
					},
				},
				{
					Value: "2",
					Rule: []fb.Component{
						fb.Elm.DatePicker("expire_date", "到期日期").
							DateType("date").
							Placeholder("请选择到期日期").
							Required(),
					},
				},
			}),

		// 下拉选择框
		fb.Elm.Select("role", "角色").
			SetOptions([]fb.Option{
				{Value: "admin", Label: "管理员"},
				{Value: "editor", Label: "编辑"},
				{Value: "user", Label: "普通用户"},
			}).
			Placeholder("请选择角色").
			Clearable(true).
			Required(),

		// 复选框
		fb.Elm.Checkbox("permissions", "权限").
			SetOptions([]fb.Option{
				{Value: "read", Label: "读取"},
				{Value: "write", Label: "写入"},
				{Value: "delete", Label: "删除"},
			}),

		// 数字输入框
		fb.Elm.Number("age", "年龄").
			Min(18).
			Max(100).
			Placeholder("请输入年龄"),

		// 滑块
		fb.Elm.Slider("score", "评分").
			Min(0).
			Max(100).
			ShowInput(true),

		// 开关
		fb.Elm.Switch("is_active", "是否启用").
			ActiveText("启用").
			InactiveText("禁用"),

		// 多行文本
		fb.Elm.Textarea("description", "描述").
			Rows(4).
			Placeholder("请输入描述信息").
			MaxLength(500).
			ShowWordLimit(true),

		// 评分
		fb.Elm.Rate("rating", "满意度").
			Max(5).
			ShowText(true),

		// 颜色选择器
		fb.Elm.ColorPicker("theme_color", "主题颜色").
			ShowAlpha(true),
	}

	// 创建配置
	config := fb.Elm.Config()
	config.SubmitBtn(true, "提交表单")
	config.ResetBtn(true, "重置")

	// 创建表单
	form := fb.Elm.CreateForm("/api/user/create", rules, config)

	// 设置表单标题
	form.SetTitle("用户注册表单")

	// 预填充数据（可选）
	form.FormData(map[string]interface{}{
		"user_type": "1",
		"role":      "user",
		"is_active": true,
		"score":     80,
	})

	// 1. 输出JSON规则（用于API）
	fmt.Println("========== JSON Rule ==========")
	ruleJSON, err := form.ParseFormRule()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ruleJSON)

	fmt.Println("\n========== JSON Config ==========")
	configJSON, err := form.ParseFormConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(configJSON)

	// 2. 输出完整HTML页面
	//fmt.Println("\n========== HTML View ==========")
	//html, err := form.View()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(html)
}
