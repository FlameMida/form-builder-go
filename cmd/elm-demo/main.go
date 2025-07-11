package main

import (
	"encoding/json"
	"fmt"
	"github.com/FlameMida/form-builder-go/components"
	"github.com/FlameMida/form-builder-go/contracts"
	"github.com/FlameMida/form-builder-go/factory"
	"log"
)

func main() {
	fmt.Println("=== Go FormBuilder Elm Demo ===")

	elm := factory.Elm{}

	// Elm::input('svip_name', '会员名：')->required()
	svipName := elm.Input("svip_name", "会员名：").Required().Placeholder("请输入会员名：")

	// Elm::radio('svip_type', '会员类别：', '2')
	// ->setOptions([...])
	// ->control([...])
	// ->appendRule('suffix', [...])
	svipType := elm.Radio("svip_type", "会员类别：", "2") // 直接在构造函数中设置默认值
	// 使用新的AddOptions方法传入结构体切片
	svipType.AddOptions([]contracts.Option{
		{Value: "1", Label: "试用期"},
		{Value: "2", Label: "有限期"},
		{Value: "3", Label: "永久期"},
	})

	svipNumberFor1 := elm.InputNumber("svip_number", "有效期（天）：").Required().Placeholder("请输入有效期（天）：").Min(0)

	svipNumberFor2 := elm.InputNumber("svip_number", "有效期（天）：").Required().Placeholder("请输入有效期（天）：").Min(0)

	svipNumber1Disabled := elm.Input("svip_number1", "有效期（天）：", "永久期").Placeholder("请输入有效期").Disabled(true)

	svipNumberHidden := elm.Input("svip_number", "有效期（天）：", "永久期").Placeholder("请输入有效期").Hidden(true) // 设置为隐藏

	// 添加控制规则 - 使用新的AddControls方法传入结构体切片
	svipType.AddControls([]components.ControlRule{
		{Value: "1", Rule: []contracts.Component{svipNumberFor1}},
		{Value: "2", Rule: []contracts.Component{svipNumberFor2}},
		{Value: "3", Rule: []contracts.Component{svipNumber1Disabled, svipNumberHidden}},
	})

	// 添加后缀元素
	suffixElement := components.SuffixElement{
		Type: "div",
		Style: map[string]interface{}{
			"color": "#999999",
		},
		DomProps: map[string]interface{}{
			"innerHTML": "试用期每个用户只能购买一次，购买过付费会员之后将不在展示，不可购买",
		},
	}
	svipType.AppendRule("suffix", suffixElement)

	// Elm::number('cost_price', '原价：')->required()
	costPrice := elm.InputNumber("cost_price", "原价：").Required().Placeholder("请输入原价：")

	// Elm::number('price', '优惠价：')->required()
	price := elm.InputNumber("price", "优惠价：").Required().Placeholder("请输入优惠价：")

	// Elm::number('sort', '排序：')
	sort := elm.InputNumber("sort", "排序：").Placeholder("请输入排序：")

	// Elm::switches('status', '是否显示：')->activeValue(1)->inactiveValue(0)->inactiveText('关')->activeText('开')
	status := elm.Switch("status", "是否显示：", "0").ActiveValue(1).InactiveValue(0).InactiveText("关").ActiveText("开")

	// === 创建表单 ===
	api := "/save"
	form, err := elm.CreateForm(api, []interface{}{
		svipName,
		svipType,
		costPrice,
		price,
		sort,
		status,
	}, map[string]interface{}{})
	if err != nil {
		log.Fatalf("创建表单失败: %v", err)
	}

	// 设置表单属性
	form.SetTitle("demo")
	form.SetMethod("POST")

	rule := form.FormRule()
	action := form.GetAction()
	method := form.GetMethod()
	title := form.GetTitle()
	view, err := form.View()
	if err != nil {
		log.Fatalf("生成视图失败: %v", err)
	}

	result := map[string]interface{}{
		"rule":   rule,
		"action": action,
		"method": method,
		"title":  title,
		"view":   view,
		"api":    api,
	}

	// === 输出完整的表单数据 ===
	fmt.Println("\n=== Complete Form Data  ===")
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("JSON序列化失败: %v", err)
	}

	fmt.Println(string(output))
}
