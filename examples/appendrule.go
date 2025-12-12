// Package main 演示 AppendRule 功能的使用
// AppendRule 允许添加标准API之外的自定义配置
package main

import (
	"encoding/json"
	"fmt"

	fb "github.com/maverick/form-builder-go/formbuilder"
)

func main() {
	fmt.Println("=== AppendRule 功能示例 ===\n")

	// 示例1：添加后缀说明（对应PHP示例）
	example1()

	// 示例2：添加前缀和后缀
	example2()

	// 示例3：添加复杂的自定义配置
	example3()

	// 示例4：在表单中使用 AppendRule
	example4()
}

// 示例1：单选框添加后缀说明（对应PHP示例）
func example1() {
	fmt.Println("【示例1】单选框添加后缀说明（对应PHP示例）")
	fmt.Println("PHP 代码：")
	fmt.Println(`Elm::radio('svip_type', '会员类别：', '2')
    ->appendRule('suffix', [
        'type' => 'div',
        'style' => ['color' => '#999999'],
        'domProps' => [
            'innerHTML' =>'试用期每个用户只能购买一次',
        ]
    ])`)
	fmt.Println()
	fmt.Println("Go 代码：")

	radio := fb.NewRadio("svip_type", "会员类别：", "2").
		SetOptions([]fb.Option{
			{Value: "1", Label: "试用期"},
			{Value: "2", Label: "有限期"},
			{Value: "3", Label: "永久期"},
		}).
		AppendRule("suffix", map[string]interface{}{
			"type": "div",
			"style": map[string]interface{}{
				"color": "#999999",
			},
			"domProps": map[string]interface{}{
				"innerHTML": "试用期每个用户只能购买一次",
			},
		})

	printJSON(radio.Build())
	fmt.Println()
}

// 示例2：输入框添加前缀和后缀
func example2() {
	fmt.Println("【示例2】输入框添加前缀和后缀")

	input := fb.NewInput("price", "价格").
		Placeholder("请输入价格").
		Required().
		AppendRule("prefix", "¥").
		AppendRule("suffix", "元")

	printJSON(input.Build())
	fmt.Println()
}

// 示例3：添加复杂的自定义配置
func example3() {
	fmt.Println("【示例3】添加复杂的自定义配置")

	input := fb.NewInput("amount", "金额").
		Placeholder("请输入金额").
		Required().
		AppendRule("prefix", map[string]interface{}{
			"type": "i",
			"props": map[string]interface{}{
				"class": "el-icon-money",
			},
		}).
		AppendRule("suffix", map[string]interface{}{
			"type": "span",
			"props": map[string]interface{}{
				"class": "amount-suffix",
			},
			"domProps": map[string]interface{}{
				"innerHTML": "人民币",
			},
		}).
		AppendRule("customData", map[string]interface{}{
			"min":  0,
			"max":  99999,
			"step": 0.01,
		})

	printJSON(input.Build())
	fmt.Println()
}

// 示例4：在表单中使用 AppendRule
func example4() {
	fmt.Println("【示例4】在表单中使用 AppendRule")

	form := fb.Elm.CreateForm("/api/vip/save", []fb.Component{
		fb.NewInput("svip_name", "会员名称").
			Placeholder("请输入会员名称").
			Required(),

		fb.NewRadio("svip_type", "会员类别", "2").
			SetOptions([]fb.Option{
				{Value: "1", Label: "试用期"},
				{Value: "2", Label: "有限期"},
				{Value: "3", Label: "永久期"},
			}).
			AppendRule("suffix", map[string]interface{}{
				"type": "div",
				"style": map[string]interface{}{
					"color":      "#999999",
					"fontSize":   "12px",
					"marginTop":  "5px",
					"lineHeight": "1.5",
				},
				"domProps": map[string]interface{}{
					"innerHTML": "试用期每个用户只能购买一次，购买过付费会员之后将不再展示",
				},
			}),

		fb.NewInputNumber("cost_price", "原价").
			Min(0).
			Precision(2).
			Required().
			AppendRule("prefix", "¥"),

		fb.NewInputNumber("price", "优惠价").
			Min(0).
			Precision(2).
			Required().
			AppendRule("prefix", "¥"),

		fb.NewInputNumber("sort", "排序").
			Value(0),

		fb.NewSwitch("status", "是否显示").
			ActiveValue(1).
			InactiveValue(0).
			ActiveText("开").
			InactiveText("关"),
	})

	form.SetTitle("会员配置")

	// 输出表单规则
	rules := form.FormRule()
	jsonBytes, _ := json.MarshalIndent(rules, "", "  ")
	fmt.Println(string(jsonBytes))
	fmt.Println()
}

// printJSON 辅助函数：格式化输出 JSON
func printJSON(data map[string]interface{}) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("JSON 序列化错误: %v\n", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
