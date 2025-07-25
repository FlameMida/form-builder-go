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
	MenuName := "login_logo"
	ul := elm.UploadFile(MenuName, "config.Info", "/adminapi/file/upload/1?type=1").
		Name("file").Data(map[string]interface{}{
		"menu_name": MenuName,
	}).Headers(map[string]string{
		"Authori-zation": "token",
	}).AppendRule("suffix", map[string]interface{}{
		"type":  "div",
		"class": "tips-info",
		"domProps": map[string]any{
			"innerHTML": "config.Desc",
		},
	})
	cascader := elm.Cascader("menu_list", "父级id").Options([]components.CascaderOption{
		{
			Label:    "1",
			Value:    1,
			Pid:      0,
			Children: []components.CascaderOption{},
			Disabled: false,
		},
	}).Filterable(true)
	ints := []int{1}
	sel := elm.Select("roles", "管理员角色", ints).Options([]contracts.Option{
		{Value: 1, Label: "开启", Disabled: false},
		{Value: 0, Label: "关闭", Disabled: false},
	}).Multiple(true).Required()

	// Elm::input('svip_name', '会员名：')->required()
	svipName := elm.Input("svip_name", "会员名：").Required().Placeholder("请输入会员名：")

	// Elm::radio('svip_type', '会员类别：', '2')
	// ->setOptions([...])
	// ->control([...])
	// ->appendRule('suffix', [...])
	svipType := elm.Radio("svip_type", "会员类别：", "2").Col(13) // 直接在构造函数中设置默认值
	// 使用新的AddOptions方法传入结构体切片
	svipType.AddOptions([]contracts.Option{
		{Value: "1", Label: "试用期"},
		{Value: "2", Label: "有限期"},
		{Value: "3", Label: "永久期"},
	})

	svipNumber1Disabled := elm.Input("svip_number1", "有效期（天）：", "永久期").Placeholder("请输入有效期").Disabled(true)

	svipNumberHidden := elm.Input("svip_number", "有效期（天）：", "永久期").Placeholder("请输入有效期").Hidden(true) // 设置为隐藏

	// 添加控制规则 - 使用新的AddControls方法传入结构体切片
	svipType.AddControls([]components.ControlRule{
		{Value: "3", Rule: []contracts.Component{svipNumber1Disabled, svipNumberHidden}},
	})

	// Elm::number('sort', '排序：')
	sort := elm.InputNumber("sort", "排序：").Placeholder("请输入排序：")

	// Elm::switches('status', '是否显示：')->activeValue(1)->inactiveValue(0)->inactiveText('关')->activeText('开')
	status := elm.Switch("status", "是否显示：", "0").ActiveValue(1).InactiveValue(0).InactiveText("关").ActiveText("开")

	// === 创建表单 ===
	api := "/save"
	form, err := elm.CreateForm(api, []interface{}{
		elm.FrameImage("album", "相册", "/upload.php?type=image", "").Height("500px").Col(12),
		ul,
		cascader,
		sel,
		svipName,
		svipType,
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

	result := map[string]interface{}{
		"rule":   rule,
		"action": action,
		"method": method,
		"title":  title,
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
