package formbuilder

import (
	"bytes"
	"html/template"
)

// template.go 实现模板和视图生成
// 对应PHP的Template/form.php和formScript()方法

// FormScript 生成表单初始化JavaScript脚本
// 对应PHP的formScript()方法
func (f *Form) FormScript() string {
	ruleJSON, _ := f.ParseFormRule()
	configJSON, _ := f.ParseFormConfig()

	script := `
new Vue({
    el: '#app',
    data: {
        fApi: null,
        rule: ` + ruleJSON + `,
        option: ` + configJSON + `
    },
    mounted() {
        console.log('Form created:', this.fApi);
    }
});
`
	return script
}

// View 生成完整的HTML页面
// 对应PHP的view()方法
func (f *Form) View() (string, error) {
	tmpl := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    {{range .Styles}}
    <link rel="stylesheet" href="{{.}}">
    {{end}}
    {{range .Scripts}}
    <script src="{{.}}"></script>
    {{end}}
</head>
<body>
    <div id="app">
        <form-create v-model="fApi" :rule="rule" :option="option"></form-create>
    </div>
    <script>
        {{.FormScript}}
    </script>
</body>
</html>`

	data := struct {
		Title      string
		Styles     []string
		Scripts    []string
		FormScript string
	}{
		Title:      f.getTitle(),
		Styles:     f.ui.GetStyles(),
		Scripts:    f.ui.GetScripts(),
		FormScript: f.FormScript(),
	}

	t, err := template.New("form").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// getTitle 获取表单标题
func (f *Form) getTitle() string {
	if f.title != "" {
		return f.title
	}
	return "Form"
}

// Template 使用自定义模板生成HTML
// templateContent: 模板内容字符串
func (f *Form) Template(templateContent string) (string, error) {
	data := struct {
		Title      string
		Styles     []string
		Scripts    []string
		FormScript string
		FormRule   string
		FormConfig string
		Action     string
		Method     string
	}{
		Title:      f.getTitle(),
		Styles:     f.ui.GetStyles(),
		Scripts:    f.ui.GetScripts(),
		FormScript: f.FormScript(),
		FormRule:   func() string { s, _ := f.ParseFormRule(); return s }(),
		FormConfig: func() string { s, _ := f.ParseFormConfig(); return s }(),
		Action:     f.action,
		Method:     f.method,
	}

	t, err := template.New("custom").Parse(templateContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// SetTemplate 设置自定义模板并返回渲染结果
// 这是Template方法的便捷封装
func (f *Form) SetTemplate(templateContent string) (string, error) {
	return f.Template(templateContent)
}

// GetUI 获取UI实例（允许自定义CDN资源）
func (f *Form) GetUI() Bootstrap {
	return f.ui
}

// SetUI 设置UI实例
func (f *Form) SetUI(ui Bootstrap) *Form {
	f.ui = ui
	f.ui.Init(f)
	return f
}
