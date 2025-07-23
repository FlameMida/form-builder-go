// 简化版本的demo，直接使用基础功能
package main

import (
	"fmt"
	"log"

	"github.com/FlameMida/form-builder-go/components"
	"github.com/FlameMida/form-builder-go/formbuilder"
	"github.com/FlameMida/form-builder-go/ui/elm"
)

func main() {
	fmt.Println("=== Go FormBuilder 基础Demo ===")

	// 创建ElementUI引导程序
	bootstrap := elm.NewBootstrap()

	// 创建表单组件
	input := components.NewInput("goods_name", "商品名称")
	input.Placeholder("请输入商品名称")
	input.Required()

	textarea := components.NewTextarea("goods_info", "商品简介")
	textarea.Placeholder("请输入商品简介")
	textarea.Rows(4)
	textarea.AppendRule("suffix", map[string]interface{}{
		"type": "div",
		"style": map[string]interface{}{
			"color": "#999999",
		},
		"domProps": map[string]interface{}{
			"innerHTML": "这是一个后缀",
		},
	})

	switchComp := components.NewSwitch("is_open", "是否开启")
	switchComp.ActiveText("开启")
	switchComp.InactiveText("关闭")
	switchComp.AppendRule("suffix", map[string]interface{}{
		"type": "div",
		"style": map[string]interface{}{
			"color": "#999999",
		},
		"domProps": map[string]interface{}{
			"innerHTML": "这是一个后缀",
		},
	})
	switchComp.AppendRule("suffix", map[string]interface{}{
		"type": "div",
		"style": map[string]interface{}{
			"color": "#999999",
		},
		"domProps": map[string]interface{}{
			"innerHTML": "这是一个后缀",
		},
	})

	// 创建组件数组
	componentsList := []interface{}{input, textarea, switchComp}

	// 创建表单配置
	config := map[string]interface{}{
		"submitBtn": true,
		"resetBtn":  false,
		"form": map[string]interface{}{
			"labelPosition": "right",
			"labelWidth":    "120px",
		},
	}

	// 创建表单
	form, err := formbuilder.NewForm(bootstrap, "/save", componentsList, config)
	if err != nil {
		log.Fatalf("创建表单失败: %v", err)
	}

	// 设置表单数据
	form.FormData(map[string]interface{}{
		"goods_name": "测试商品",
		"goods_info": "这是一个测试商品",
		"is_open":    true,
	})

	// 输出表单规则JSON
	ruleJSON, err := form.ParseFormRule()
	if err != nil {
		log.Fatalf("解析表单规则失败: %v", err)
	}

	// 输出表单配置JSON
	configJSON, err := form.ParseFormConfig()
	if err != nil {
		log.Fatalf("解析表单配置失败: %v", err)
	}

	fmt.Println("\n表单规则JSON:")
	fmt.Println(string(ruleJSON))

	fmt.Println("\n表单配置JSON:")
	fmt.Println(string(configJSON))

	// 验证组件
	fmt.Println("\n=== 组件验证测试 ===")

	// 测试必填验证
	input.SetValue("")
	if err := input.DoValidate(); err != nil {
		fmt.Printf("商品名称验证失败: %v\n", err)
	}

	input.SetValue("有效商品名称")
	if err := input.DoValidate(); err != nil {
		fmt.Printf("商品名称验证失败: %v\n", err)
	} else {
		fmt.Println("商品名称验证通过")
	}

	fmt.Println("\n=== Demo完成 ===")
}
