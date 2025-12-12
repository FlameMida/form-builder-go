# FormBuilder Go

一个强大的Go表单生成器库，使用泛型Builder模式实现优雅的链式调用API。完全兼容PHP版本FormBuilder的功能和JSON输出格式。

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/maverick/form-builder-go)
[![License](https://img.shields.io/badge/license-MIT-blue)](LICENSE)

## ✨ 特性

- 🎯 **泛型Builder模式** - 利用Go 1.18+泛型特性，实现类型安全的链式调用
- 🧩 **15+表单组件** - Input、Select、Radio、Checkbox、DatePicker等常用组件
- ✅ **完整验证系统** - 内置10+种验证规则，支持自定义验证
- 🔄 **条件显示** - Control系统支持动态表单，根据字段值显示/隐藏组件
- 📦 **多UI框架** - 支持Element UI、Element Plus、iView v3/v4
- 🔗 **100%兼容PHP版本** - JSON输出格式完全兼容，可无缝迁移
- 📝 **表单数据管理** - 支持预填充数据，自动应用到组件
- 🎨 **模板系统** - 内置模板或自定义模板生成HTML页面

## 📦 安装

```bash
go get github.com/maverick/form-builder-go/formbuilder
```

**要求**: Go 1.18+ (支持泛型)

## 🚀 快速开始

### 基础示例

```go
package main

import (
    "fmt"
    fb "github.com/maverick/form-builder-go/formbuilder"
)

func main() {
    // 创建表单
    form := fb.Elm.CreateForm("/api/login", []fb.Component{
        fb.Elm.Input("username", "用户名").
            Placeholder("请输入用户名").
            Required(),

        fb.Elm.Password("password", "密码").
            Placeholder("请输入密码").
            Required(),
    })

    // 输出JSON规则
    jsonRule, _ := form.ParseFormRule()
    fmt.Println(jsonRule)

    // 输出HTML页面
    html, _ := form.View()
    fmt.Println(html)
}
```

### 完整示例

```go
form := fb.Elm.CreateForm("/api/user/create", []fb.Component{
    // 输入框 + 验证
    fb.Elm.Input("username", "用户名").
        Placeholder("请输入用户名").
        Clearable(true).
        MaxLength(50).
        Validate(
            fb.NewRequired("用户名不能为空"),
            fb.NewLength(6, 20, "长度必须在6-20个字符之间"),
        ),

    // 下拉选择
    fb.Elm.Select("role", "角色").
        SetOptions([]fb.Option{
            {Value: "admin", Label: "管理员"},
            {Value: "user", Label: "普通用户"},
        }).
        Required(),

    // 单选框 + 条件显示
    fb.Elm.Radio("user_type", "用户类型", "1").
        SetOptions([]fb.Option{
            {Value: "1", Label: "试用期"},
            {Value: "2", Label: "正式"},
        }).
        Control([]fb.ControlRule{
            {
                Value: "1",
                Rule: []fb.Component{
                    fb.Elm.Number("trial_days", "试用天数").Required(),
                },
            },
        }),

    // 开关
    fb.Elm.Switch("is_active", "是否启用").
        ActiveText("启用").
        InactiveText("禁用"),
})

// 预填充数据
form.FormData(map[string]interface{}{
    "user_type": "1",
    "is_active": true,
})
```

## 📚 核心概念

### 组件系统

所有组件都支持链式调用，每个方法返回组件自身：

```go
input := fb.Elm.Input("email", "邮箱").
    Placeholder("请输入邮箱").     // Input特有方法
    Clearable(true).              // Input特有方法
    Required().                   // Builder通用方法
    Value("default@example.com")  // Builder通用方法
```

### 验证规则

内置多种验证规则：

```go
fb.Elm.Input("username", "用户名").
    Validate(
        fb.NewRequired("必填"),
        fb.NewLength(6, 20, "长度6-20"),
        fb.NewPattern("^[a-zA-Z0-9]+$", "只能字母数字"),
    )

// 便捷方法
fb.Email("email", "邮箱")  // 自动添加邮箱验证
fb.URL("website", "网站")  // 自动添加URL验证
```

支持的验证规则：
- `RequiredRule` - 必填
- `PatternRule` - 正则表达式
- `LengthRule` - 长度（Min/Max）
- `RangeRule` - 数值范围
- `EmailRule` - 邮箱格式
- `URLRule` - URL格式
- `DateRule` - 日期格式
- `EnumRule` - 枚举值
- `CustomRule` - 自定义验证

### 条件显示（Control）

根据字段值动态显示/隐藏组件：

```go
fb.Elm.Radio("delivery", "配送方式", "express").
    SetOptions([]fb.Option{
        {Value: "express", Label: "快递"},
        {Value: "pickup", Label: "自提"},
    }).
    Control([]fb.ControlRule{
        {
            Value: "express",
            Rule: []fb.Component{
                fb.Elm.Input("address", "地址").Required(),
            },
        },
        {
            Value: "pickup",
            Rule: []fb.Component{
                fb.Elm.Select("store", "门店").Required(),
            },
        },
    })
```

### 表单数据

预填充表单数据：

```go
form.FormData(map[string]interface{}{
    "username": "john_doe",
    "email":    "john@example.com",
    "role":     "admin",
})

// 或单个设置
form.SetValue("username", "john_doe")
```

## 🧩 支持的组件

| 组件 | Type值 | 说明 |
|------|--------|------|
| Input | `input` | 输入框 |
| Select | `select` | 下拉选择 |
| Radio | `radio` | 单选框 |
| Checkbox | `checkbox` | 复选框 |
| InputNumber | `inputNumber` | 数字输入 |
| DatePicker | `datePicker` | 日期选择 |
| TimePicker | `timePicker` | 时间选择 |
| Slider | `slider` | 滑块 |
| Switch | `switch` | 开关 |
| Upload | `upload` | 文件上传 |
| Cascader | `cascader` | 级联选择 |
| Tree | `tree` | 树形控件 |
| Rate | `rate` | 评分 |
| ColorPicker | `colorPicker` | 颜色选择器 |
| Hidden | `hidden` | 隐藏字段 |

**注意**：Element UI 和 iView 使用相同的 type 值，框架的选择由全局配置决定，而非 type 字段。

## 🎨 UI框架

### Element UI (Vue 2)

```go
form := fb.Elm.CreateForm("/submit", rules)
```

### Element Plus (Vue 3)

```go
// 使用相同的API，只需更换前端库
form := fb.Elm.CreateForm("/submit", rules)
```

### iView v3

```go
form := fb.Iview.CreateForm("/submit", rules)
```

### iView v4 (View Design)

```go
form := fb.Iview4.CreateForm("/submit", rules)
```

### 自定义CDN

```go
form := fb.Elm.CreateForm("/submit", rules)
bootstrap := form.GetUI().(*fb.ElmBootstrap)
bootstrap.SetScripts([]string{
    "https://your-cdn.com/vue.js",
    "https://your-cdn.com/element-ui.js",
    // ...
})
```

## 📤 输出格式

### JSON规则（用于API）

```go
// 获取规则数组
rules := form.FormRule()

// 获取JSON字符串
jsonStr, _ := form.ParseFormRule()
```

输出示例：
```json
[
  {
    "type": "input",
    "field": "username",
    "title": "用户名",
    "props": {
      "placeholder": "请输入用户名",
      "clearable": true
    },
    "validate": [
      {"required": true, "message": "用户名不能为空"}
    ]
  }
]
```

### HTML页面

```go
html, _ := form.View()
```

生成包含Vue实例的完整HTML页面。

### 自定义模板

```go
template := `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    {{range .Styles}}<link href="{{.}}" rel="stylesheet">{{end}}
</head>
<body>
    <div id="app">
        <form-create v-model="fApi" :rule="rule" :option="option"></form-create>
    </div>
    {{range .Scripts}}<script src="{{.}}"></script>{{end}}
    <script>{{.FormScript}}</script>
</body>
</html>
`

html, _ := form.Template(template)
```

## 🔧 配置

```go
config := fb.Elm.Config()

// 提交按钮
config.SubmitBtn(true, "提交表单")

// 重置按钮
config.ResetBtn(true, "重置")

// 表单样式
config.FormStyle(map[string]interface{}{
    "labelWidth": "100px",
})

// 全局配置
config.Global("upload", map[string]interface{}{
    "action": "/upload",
})

form := fb.Elm.CreateForm("/submit", rules, config)
```

## 🧪 测试

运行测试：

```bash
go test -v
```

运行性能测试：

```bash
go test -bench=. -benchmem
```

测试覆盖率：

```bash
go test -cover
```

## 📖 示例

查看 `examples/` 目录获取更多示例：

- `basic.go` - 基础用法
- `validation.go` - 验证规则示例
- `control.go` - 条件显示示例
- `demo.go` - 完整功能演示

运行示例：

```bash
cd examples
go run demo.go
```

## 🔄 与PHP版本对比

| 特性 | PHP | Go |
|------|-----|-----|
| 链式调用 | ✅ Trait混入 | ✅ 泛型Builder |
| 组件数量 | 15+ | 15+ |
| 验证规则 | ✅ | ✅ |
| 条件显示 | ✅ | ✅ |
| JSON输出 | ✅ | ✅ 100%兼容 |
| UI框架 | Element UI, iView | Element UI, iView |
| 类型安全 | ❌ 动态类型 | ✅ 编译时检查 |
| 性能 | - | 更快 |

## 📝 迁移指南

从PHP版本迁移：

```php
// PHP
use FormBuilder\Factory\Elm;

$form = Elm::createForm('/submit', [
    Elm::input('username', '用户名')->required(),
]);
```

```go
// Go
form := fb.Elm.CreateForm("/submit", []fb.Component{
    fb.Elm.Input("username", "用户名").Required(),
})
```

主要区别：
1. 方法名：PHP使用小驼峰，Go使用大驼峰（公开方法）
2. 数组：PHP使用 `[]`，Go使用 `[]fb.Component{}`
3. 工厂：PHP静态方法，Go全局单例 `fb.Elm`

## 🤝 贡献

欢迎贡献！请查看 [CONTRIBUTING.md](CONTRIBUTING.md)

## 📄 许可证

MIT License - 详见 [LICENSE](LICENSE)

## 🔗 相关链接

- [PHP版本 FormBuilder](https://github.com/xaboy/form-builder)
- [@form-create 文档](https://www.form-create.com/)
- [Element UI](https://element.eleme.io/)
- [Element Plus](https://element-plus.org/)
- [iView](http://iview.talkingdata.com/)
