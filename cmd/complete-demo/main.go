package main

import (
	"encoding/json"
	"fmt"
	"github.com/FlameMida/form-builder-go/factory"
	"log"
	"strings"
)

func main() {
	fmt.Println("=== Go FormBuilder Complete Demo  ===")

	// === 创建Elm组件 ===
	elm := factory.Elm{}

	// 会员名输入框
	svipName := elm.Input("svip_name", "会员名：")
	svipName.Required()

	// 会员类别单选框
	svipType := elm.Radio("svip_type", "会员类别：")
	svipType.AddOption("1", "试用期")
	svipType.AddOption("2", "有限期")
	svipType.AddOption("3", "永久期")
	svipType.SetValue("2") // 默认值

	// 有效期数字输入框
	svipNumber := elm.InputNumber("svip_number", "有效期（天）：")
	svipNumber.Required()
	svipNumber.Min(0)

	// 原价
	costPrice := elm.InputNumber("cost_price", "原价：")
	costPrice.Required()

	// 优惠价
	price := elm.InputNumber("price", "优惠价：")
	price.Required()

	// 排序
	sort := elm.InputNumber("sort", "排序：")

	// 是否显示开关
	status := elm.Switch("status", "是否显示：")
	status.ActiveValue(1)
	status.InactiveValue(0)
	status.InactiveText("关")
	status.ActiveText("开")

	// 图片上传
	logo := elm.UploadImage("logo", "logo", "/upload.php")

	//相册
	album := elm.FrameImage("album", "相册", "/upload.php?type=image", "").Height("500px").Col(12)

	// === 创建表单 ===
	api := "/save"
	form, err := elm.CreateForm(api, []interface{}{
		svipName,
		svipType,
		svipNumber,
		costPrice,
		price,
		sort,
		status,
		logo,
		album,
	}, map[string]interface{}{})
	if err != nil {
		log.Fatalf("创建表单失败: %v", err)
	}

	// 设置表单属性
	form.SetTitle("demo")
	form.SetMethod("POST")

	// 设置一些默认数据
	form.FormData(map[string]interface{}{
		"svip_name":   "VIP会员",
		"svip_type":   "2",
		"svip_number": 365,
		"cost_price":  199.0,
		"price":       99.0,
		"sort":        1,
		"status":      1,
	})

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
	fmt.Println("\n=== Complete Form Data ===")
	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("JSON序列化失败: %v", err)
	}

	fmt.Println(string(output))

	// === 分别显示各个部分 (便于调试) ===
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("FORM SUMMARY")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("API: %s\n", api)
	fmt.Printf("Action: %s\n", action)
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Components Count: %d\n", len(rule))

	fmt.Println("\nComponent Types:")
	for i, r := range rule {
		field := r["field"]
		componentType := r["type"]
		titleField := r["title"]
		fmt.Printf("  %d. %s (%s) - %s\n", i+1, field, componentType, titleField)
	}
}
