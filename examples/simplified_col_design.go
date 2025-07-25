package main

import (
	"fmt"
	"github.com/FlameMida/form-builder-go/components"
	"github.com/FlameMida/form-builder-go/contracts"
)

func main() {
	fmt.Println("=== Col() 方法简化设计示例 ===\n")

	// 1. 流畅的链式调用 - 直接使用 Col() 方法
	fmt.Println("1. 流畅的链式调用：")
	radio := components.NewRadio("gender", "性别").
		Col(12).
		AddOption("male", "男性").
		AddOption("female", "女性").
		Required().
		Size("large")

	fmt.Printf("Radio组件类型: %T\n", radio)
	fmt.Printf("Col配置: %v\n\n", radio.Build()["col"])

	// 2. 接口兼容性 - 组件仍然实现 Component 接口
	fmt.Println("2. 接口兼容性：")
	componentList := []contracts.Component{
		components.NewInput("name", "姓名"),
		components.NewRadio("gender", "性别"),
		components.NewSelect("city", "城市"),
	}

	for i, component := range componentList {
		fmt.Printf("Component %d: %s - %s\n", i+1, component.Field(), component.Title())
	}
	fmt.Println()

	// 3. 具体组件的 Col() 方法返回具体类型
	fmt.Println("3. 具体类型的Col()方法：")
	input := components.NewInput("email", "邮箱").Col(8)          // 返回 *Input
	textarea := components.NewTextarea("content", "内容").Col(12) // 返回 *Textarea
	switchComp := components.NewSwitch("enabled", "启用").Col(6)  // 返回 *Switch

	fmt.Printf("Input类型: %T\n", input)
	fmt.Printf("Textarea类型: %T\n", textarea)
	fmt.Printf("Switch类型: %T\n", switchComp)
	fmt.Println()

	// 4. FormBuilder 兼容性
	fmt.Println("4. FormBuilder兼容性：")
	formComponents := []contracts.Component{
		components.NewInput("username", "用户名").Col(12), // 具体类型 -> 接口类型
		components.NewRadio("gender", "性别").Col(6),     // 具体类型 -> 接口类型
		components.NewSelect("city", "城市").Col(8),      // 具体类型 -> 接口类型
	}

	fmt.Printf("FormBuilder可以接受 %d 个组件\n", len(formComponents))
	for _, comp := range formComponents {
		result := comp.Build()
		fmt.Printf("  - %s: col配置 = %v\n", result["field"], result["col"])
	}

	fmt.Println("\n=== 设计优势 ===")
	fmt.Println("✓ 接口简化：Component接口不再包含Col()方法")
	fmt.Println("✓ 类型安全：每个组件的Col()方法返回具体类型")
	fmt.Println("✓ 链式调用：完美支持流畅的链式调用语法")
	fmt.Println("✓ 接口兼容：组件仍然可以作为Component接口使用")
	fmt.Println("✓ 代码简洁：不需要ColChain()这样的双重方法")
}
