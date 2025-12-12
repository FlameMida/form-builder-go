# FormBuilder Go API 参考文档

完整的API参考，包含所有组件、方法、验证规则和配置选项。

## 目录

- [核心类型](#核心类型)
- [表单组件](#表单组件)
- [验证规则](#验证规则)
- [工厂方法](#工厂方法)
- [表单配置](#表单配置)
- [表单方法](#表单方法)

---

## 核心类型

### Component

所有表单组件都实现此接口：

```go
type Component interface {
    GetField() string                      // 获取字段名
    GetType() string                       // 获取组件类型
    Build() map[string]interface{}         // 构建JSON规则
}
```

### ComponentData

组件数据结构：

```go
type ComponentData struct {
    Field    string                        // 字段名（必需，唯一）
    Title    string                        // 字段标题
    RuleType string                        // 组件类型（如 "input"）
    Value    interface{}                   // 默认值
    Props    map[string]interface{}        // 组件属性
    Validate []ValidateRule                // 验证规则列表
    Control  []ControlRule                 // 条件显示规则
    Children []Component                   // 子组件
    Emit     map[string]interface{}        // 事件配置
}
```

### Builder[T]

泛型Builder提供通用的链式方法：

```go
type Builder[T any] struct {
    data *ComponentData
    inst T
}

// 通用方法
func (b *Builder[T]) Value(val interface{}) T
func (b *Builder[T]) Title(title string) T
func (b *Builder[T]) Field(field string) T
func (b *Builder[T]) Required() T
func (b *Builder[T]) Validate(rules ...ValidateRule) T
func (b *Builder[T]) Control(rules []ControlRule) T
func (b *Builder[T]) AppendControl(rule ControlRule) T
func (b *Builder[T]) Props(key string, val interface{}) T
func (b *Builder[T]) SetProps(props map[string]interface{}) T
func (b *Builder[T]) Emit(event string, handler interface{}) T
func (b *Builder[T]) Children(children []Component) T
func (b *Builder[T]) AppendChild(child Component) T
func (b *Builder[T]) AppendRule(name string, value interface{}) T
```

**说明**：
- `Props()`: 设置单个属性
- `SetProps()`: 批量设置属性
- `AppendRule()`: 添加自定义规则字段，允许添加标准API之外的配置

---

## 表单组件

### Input - 输入框

**类型**: `input` (Element UI) / `input` (iView)

**构造函数**:
```go
func (ElmFactory) Input(field, title string, args ...interface{}) *Input
func (IviewFactory) Input(field, title string, args ...interface{}) *Input
```

**特有方法**:
```go
func (i *Input) Placeholder(text string) *Input          // 占位符
func (i *Input) Clearable(enable bool) *Input            // 可清空
func (i *Input) ShowPassword(enable bool) *Input         // 显示密码（密码输入框）
func (i *Input) Disabled(disabled bool) *Input           // 禁用
func (i *Input) Readonly(readonly bool) *Input           // 只读
func (i *Input) MaxLength(length int) *Input             // 最大长度
func (i *Input) MinLength(length int) *Input             // 最小长度
func (i *Input) ShowWordLimit(show bool) *Input          // 显示字数统计
func (i *Input) PrefixIcon(icon string) *Input           // 前缀图标
func (i *Input) SuffixIcon(icon string) *Input           // 后缀图标
func (i *Input) Rows(rows int) *Input                    // 行数（多行文本）
func (i *Input) Autosize(config interface{}) *Input      // 自适应高度
```

**示例**:
```go
fb.Elm.Input("username", "用户名").
    Placeholder("请输入用户名").
    Clearable(true).
    MaxLength(50).
    Required()
```

---

### Password - 密码输入框

继承Input的所有方法，默认type为password。

**构造函数**:
```go
func (ElmFactory) Password(field, title string, args ...interface{}) *Input
```

**示例**:
```go
fb.Elm.Password("password", "密码").
    Placeholder("请输入密码").
    ShowPassword(true).
    Required()
```

---

### Email - 邮箱输入框

继承Input的所有方法，自动添加邮箱验证规则。

**构造函数**:
```go
func (ElmFactory) Email(field, title string, args ...interface{}) *Input
```

---

### URL - URL输入框

继承Input的所有方法，自动添加URL验证规则。

**构造函数**:
```go
func (ElmFactory) URL(field, title string, args ...interface{}) *Input
```

---

### Textarea - 多行文本框

**类型**: `input` (type=textarea)

**构造函数**:
```go
func (ElmFactory) Textarea(field, title string, args ...interface{}) *Input
```

**特有方法**: 同Input，默认type为textarea

**示例**:
```go
fb.Elm.Textarea("description", "描述").
    Rows(4).
    Placeholder("请输入描述").
    MaxLength(500).
    ShowWordLimit(true)
```

---

### Select - 下拉选择框

**类型**: `select` / `select`

**构造函数**:
```go
func (ElmFactory) Select(field, title string, args ...interface{}) *Select
```

**特有方法**:
```go
func (s *Select) Placeholder(text string) *Select        // 占位符
func (s *Select) Multiple(enable bool) *Select           // 多选
func (s *Select) Clearable(enable bool) *Select          // 可清空
func (s *Select) Filterable(enable bool) *Select         // 可搜索
func (s *Select) Disabled(disabled bool) *Select         // 禁用
func (s *Select) Size(size string) *Select               // 尺寸 (large/medium/small/mini)
func (s *Select) SetOptions(options []Option) *Select    // 设置选项
```

**Option类型**:
```go
type Option struct {
    Value    interface{}  // 选项值
    Label    string       // 选项标签
    Disabled bool         // 是否禁用
    Children []Option     // 子选项（级联用）
}
```

**示例**:
```go
fb.Elm.Select("role", "角色").
    SetOptions([]fb.Option{
        {Value: "admin", Label: "管理员"},
        {Value: "user", Label: "普通用户"},
    }).
    Placeholder("请选择角色").
    Clearable(true).
    Required()
```

---

### Radio - 单选框

**类型**: `radio`

**构造函数**:
```go
func (ElmFactory) Radio(field, title string, defaultValue interface{}, args ...interface{}) *Radio
```

**特有方法**:
```go
func (r *Radio) SetOptions(options []Option) *Radio      // 设置选项
func (r *Radio) Size(size string) *Radio                 // 尺寸
func (r *Radio) Disabled(disabled bool) *Radio           // 禁用
func (r *Radio) TextColor(color string) *Radio           // 文字颜色
func (r *Radio) Fill(color string) *Radio                // 填充颜色
```

**示例**:
```go
fb.Elm.Radio("status", "状态", "1").
    SetOptions([]fb.Option{
        {Value: "1", Label: "启用"},
        {Value: "0", Label: "禁用"},
    })
```

---

### Checkbox - 复选框

**类型**: `checkbox`

**构造函数**:
```go
func (ElmFactory) Checkbox(field, title string, args ...interface{}) *Checkbox
```

**特有方法**:
```go
func (c *Checkbox) SetOptions(options []Option) *Checkbox  // 设置选项
func (c *Checkbox) Min(min int) *Checkbox                  // 最小选择数
func (c *Checkbox) Max(max int) *Checkbox                  // 最大选择数
func (c *Checkbox) Size(size string) *Checkbox             // 尺寸
func (c *Checkbox) Disabled(disabled bool) *Checkbox       // 禁用
```

**示例**:
```go
fb.Elm.Checkbox("permissions", "权限").
    SetOptions([]fb.Option{
        {Value: "read", Label: "读取"},
        {Value: "write", Label: "写入"},
        {Value: "delete", Label: "删除"},
    }).
    Min(1).
    Max(3)
```

---

### Number / InputNumber - 数字输入框

**类型**: `inputNumber` / `inputNumber`

**构造函数**:
```go
func (ElmFactory) Number(field, title string, args ...interface{}) *InputNumber
func (ElmFactory) InputNumber(field, title string, args ...interface{}) *InputNumber
```

**特有方法**:
```go
func (n *InputNumber) Min(min float64) *InputNumber          // 最小值
func (n *InputNumber) Max(max float64) *InputNumber          // 最大值
func (n *InputNumber) Step(step float64) *InputNumber        // 步长
func (n *InputNumber) Precision(precision int) *InputNumber  // 精度
func (n *InputNumber) Disabled(disabled bool) *InputNumber   // 禁用
func (n *InputNumber) Controls(show bool) *InputNumber       // 显示控制按钮
func (n *InputNumber) ControlsPosition(pos string) *InputNumber  // 按钮位置
func (n *InputNumber) Placeholder(text string) *InputNumber  // 占位符
```

**示例**:
```go
fb.Elm.Number("age", "年龄").
    Min(0).
    Max(150).
    Step(1).
    Placeholder("请输入年龄")
```

---

### DatePicker - 日期选择器

**类型**: `datePicker` / `datePicker`

**构造函数**:
```go
func (ElmFactory) DatePicker(field, title string, args ...interface{}) *DatePicker
```

**特有方法**:
```go
func (d *DatePicker) DateType(dateType string) *DatePicker  // 日期类型: date/datetime/daterange等
func (d *DatePicker) Placeholder(text string) *DatePicker   // 占位符
func (d *DatePicker) StartPlaceholder(text string) *DatePicker   // 范围开始占位符
func (d *DatePicker) EndPlaceholder(text string) *DatePicker     // 范围结束占位符
func (d *DatePicker) Format(format string) *DatePicker      // 显示格式
func (d *DatePicker) ValueFormat(format string) *DatePicker // 值格式
func (d *DatePicker) Clearable(enable bool) *DatePicker     // 可清空
func (d *DatePicker) Disabled(disabled bool) *DatePicker    // 禁用
func (d *DatePicker) Editable(enable bool) *DatePicker      // 可输入
func (d *DatePicker) RangeSeparator(sep string) *DatePicker // 范围分隔符
```

**日期类型**:
- `date` - 日期
- `datetime` - 日期时间
- `daterange` - 日期范围
- `datetimerange` - 日期时间范围
- `year` - 年
- `month` - 月
- `week` - 周

**示例**:
```go
fb.Elm.DatePicker("birthday", "生日").
    DateType("date").
    Placeholder("请选择日期").
    Format("yyyy-MM-dd").
    Clearable(true)
```

---

### TimePicker - 时间选择器

**类型**: `timePicker` / `timePicker`

**构造函数**:
```go
func (ElmFactory) TimePicker(field, title string, args ...interface{}) *TimePicker
```

**特有方法**:
```go
func (t *TimePicker) IsRange(enable bool) *TimePicker       // 是否范围选择
func (t *TimePicker) Placeholder(text string) *TimePicker   // 占位符
func (t *TimePicker) StartPlaceholder(text string) *TimePicker
func (t *TimePicker) EndPlaceholder(text string) *TimePicker
func (t *TimePicker) Format(format string) *TimePicker      // 显示格式
func (t *TimePicker) ValueFormat(format string) *TimePicker // 值格式
func (t *TimePicker) Clearable(enable bool) *TimePicker
func (t *TimePicker) Disabled(disabled bool) *TimePicker
func (t *TimePicker) Editable(enable bool) *TimePicker
func (t *TimePicker) RangeSeparator(sep string) *TimePicker
```

**示例**:
```go
fb.Elm.TimePicker("work_time", "工作时间").
    Placeholder("选择时间").
    Format("HH:mm:ss")
```

---

### Slider - 滑块

**类型**: `slider` / `slider`

**构造函数**:
```go
func (ElmFactory) Slider(field, title string, args ...interface{}) *Slider
```

**特有方法**:
```go
func (s *Slider) Min(min float64) *Slider                   // 最小值
func (s *Slider) Max(max float64) *Slider                   // 最大值
func (s *Slider) Step(step float64) *Slider                 // 步长
func (s *Slider) ShowInput(show bool) *Slider               // 显示输入框
func (s *Slider) ShowStops(show bool) *Slider               // 显示间断点
func (s *Slider) ShowTooltip(show bool) *Slider             // 显示提示
func (s *Slider) Range(enable bool) *Slider                 // 范围选择
func (s *Slider) Disabled(disabled bool) *Slider            // 禁用
func (s *Slider) Marks(marks map[int]string) *Slider        // 标记
```

**示例**:
```go
fb.Elm.Slider("score", "评分").
    Min(0).
    Max(100).
    Step(10).
    ShowInput(true).
    ShowStops(true)
```

---

### Switch - 开关

**类型**: `switch` / `switch`

**构造函数**:
```go
func (ElmFactory) Switch(field, title string, args ...interface{}) *Switch
```

**特有方法**:
```go
func (s *Switch) ActiveValue(val interface{}) *Switch       // 打开时的值
func (s *Switch) InactiveValue(val interface{}) *Switch     // 关闭时的值
func (s *Switch) ActiveText(text string) *Switch            // 打开时的文字
func (s *Switch) InactiveText(text string) *Switch          // 关闭时的文字
func (s *Switch) ActiveColor(color string) *Switch          // 打开时的颜色
func (s *Switch) InactiveColor(color string) *Switch        // 关闭时的颜色
func (s *Switch) Disabled(disabled bool) *Switch            // 禁用
```

**示例**:
```go
fb.Elm.Switch("is_active", "是否启用").
    ActiveValue(true).
    InactiveValue(false).
    ActiveText("启用").
    InactiveText("禁用").
    ActiveColor("#13ce66").
    InactiveColor("#ff4949")
```

---

### Upload - 文件上传

**类型**: `upload` / `upload`

**构造函数**:
```go
func (ElmFactory) Upload(field, title string, args ...interface{}) *Upload
```

**特有方法**:
```go
func (u *Upload) Action(url string) *Upload                 // 上传地址
func (u *Upload) Headers(headers map[string]string) *Upload // 请求头
func (u *Upload) Multiple(enable bool) *Upload              // 多文件上传
func (u *Upload) Data(data map[string]interface{}) *Upload  // 额外参数
func (u *Upload) Name(name string) *Upload                  // 文件字段名
func (u *Upload) WithCredentials(enable bool) *Upload       // 支持Cookie
func (u *Upload) Accept(accept string) *Upload              // 接受的文件类型
func (u *Upload) ListType(listType string) *Upload          // 列表类型: text/picture/picture-card
func (u *Upload) AutoUpload(enable bool) *Upload            // 自动上传
func (u *Upload) Disabled(disabled bool) *Upload            // 禁用
func (u *Upload) Limit(limit int) *Upload                   // 最大上传数
```

**示例**:
```go
fb.Elm.Upload("avatar", "头像").
    Action("/api/upload").
    Accept("image/*").
    ListType("picture-card").
    Limit(1)
```

---

### Cascader - 级联选择器

**类型**: `cascader` / `cascader`

**构造函数**:
```go
func (ElmFactory) Cascader(field, title string, args ...interface{}) *Cascader
```

**特有方法**:
```go
func (c *Cascader) SetOptions(options []Option) *Cascader   // 设置选项（支持Children）
func (c *Cascader) Placeholder(text string) *Cascader       // 占位符
func (c *Cascader) Clearable(enable bool) *Cascader         // 可清空
func (c *Cascader) Filterable(enable bool) *Cascader        // 可搜索
func (c *Cascader) Disabled(disabled bool) *Cascader        // 禁用
func (c *Cascader) ShowAllLevels(show bool) *Cascader       // 显示完整路径
func (c *Cascader) ExpandTrigger(trigger string) *Cascader  // 展开方式: click/hover
```

**示例**:
```go
fb.Elm.Cascader("region", "地区").
    SetOptions([]fb.Option{
        {
            Value: "beijing",
            Label: "北京",
            Children: []fb.Option{
                {Value: "chaoyang", Label: "朝阳区"},
                {Value: "haidian", Label: "海淀区"},
            },
        },
    }).
    Placeholder("请选择地区").
    Clearable(true)
```

---

### Tree - 树形控件

**类型**: `tree` / `tree`

**构造函数**:
```go
func (ElmFactory) Tree(field, title string, args ...interface{}) *Tree
```

**特有方法**:
```go
func (t *Tree) SetData(data []Option) *Tree                 // 设置树形数据
func (t *Tree) ShowCheckbox(show bool) *Tree                // 显示复选框
func (t *Tree) DefaultExpandAll(expand bool) *Tree          // 默认展开所有
func (t *Tree) ExpandOnClickNode(enable bool) *Tree         // 点击节点展开
func (t *Tree) CheckOnClickNode(enable bool) *Tree          // 点击节点选中
func (t *Tree) NodeKey(key string) *Tree                    // 节点唯一标识
func (t *Tree) Props(props map[string]interface{}) *Tree    // 配置选项
```

**示例**:
```go
fb.Elm.Tree("departments", "部门").
    SetData([]fb.Option{
        {
            Value: "1",
            Label: "技术部",
            Children: []fb.Option{
                {Value: "1-1", Label: "前端组"},
                {Value: "1-2", Label: "后端组"},
            },
        },
    }).
    ShowCheckbox(true).
    DefaultExpandAll(true)
```

---

### Rate - 评分

**类型**: `rate` / `rate`

**构造函数**:
```go
func (ElmFactory) Rate(field, title string, args ...interface{}) *Rate
```

**特有方法**:
```go
func (r *Rate) Max(max int) *Rate                           // 最大分值
func (r *Rate) AllowHalf(enable bool) *Rate                 // 允许半选
func (r *Rate) ShowText(show bool) *Rate                    // 显示辅助文字
func (r *Rate) ShowScore(show bool) *Rate                   // 显示分数
func (r *Rate) Disabled(disabled bool) *Rate                // 禁用
func (r *Rate) Texts(texts []string) *Rate                  // 辅助文字数组
func (r *Rate) Colors(colors []string) *Rate                // 图标颜色
```

**示例**:
```go
fb.Elm.Rate("rating", "评分").
    Max(5).
    AllowHalf(true).
    ShowText(true).
    Texts([]string{"极差", "失望", "一般", "满意", "惊喜"})
```

---

### ColorPicker - 颜色选择器

**类型**: `colorPicker` / `colorPicker`

**构造函数**:
```go
func (ElmFactory) ColorPicker(field, title string, args ...interface{}) *ColorPicker
```

**特有方法**:
```go
func (c *ColorPicker) ShowAlpha(show bool) *ColorPicker     // 支持透明度
func (c *ColorPicker) ColorFormat(format string) *ColorPicker  // 颜色格式: hex/rgb/hsl/hsv
func (c *ColorPicker) Disabled(disabled bool) *ColorPicker  // 禁用
func (c *ColorPicker) Size(size string) *ColorPicker        // 尺寸
func (c *ColorPicker) Predefine(colors []string) *ColorPicker  // 预定义颜色
```

**示例**:
```go
fb.Elm.ColorPicker("theme_color", "主题颜色").
    ShowAlpha(true).
    ColorFormat("hex").
    Predefine([]string{"#ff4500", "#ff8c00", "#ffd700"})
```

---

### Hidden - 隐藏字段

**类型**: `hidden`

**构造函数**:
```go
func (ElmFactory) Hidden(field string, defaultValue interface{}) *Hidden
```

**示例**:
```go
fb.Elm.Hidden("id", 123)
```

---

## 验证规则

所有验证规则都实现 `ValidateRule` 接口：

```go
type ValidateRule interface {
    ToMap() map[string]interface{}
}
```

### RequiredRule - 必填验证

```go
func NewRequired(message ...string) RequiredRule

type RequiredRule struct {
    Message string  // 错误信息
    Trigger string  // 触发方式: blur/change
}
```

**示例**:
```go
fb.NewRequired("用户名不能为空")
fb.NewRequired()  // 使用默认消息
```

---

### PatternRule - 正则验证

```go
func NewPattern(pattern, message string) PatternRule

type PatternRule struct {
    Pattern string  // 正则表达式
    Message string  // 错误信息
    Trigger string
}
```

**示例**:
```go
fb.NewPattern("^[a-zA-Z0-9]+$", "只能包含字母和数字")
fb.NewPattern("^1[3-9]\\d{9}$", "请输入正确的手机号")
```

---

### LengthRule - 长度验证

```go
func NewLength(min, max int, message ...string) LengthRule
func NewMinLength(min int, message ...string) LengthRule
func NewMaxLength(max int, message ...string) LengthRule

type LengthRule struct {
    Min     int
    Max     int
    Message string
    Trigger string
}
```

**示例**:
```go
fb.NewLength(6, 20, "长度必须在6-20个字符之间")
fb.NewMinLength(6, "最少6个字符")
fb.NewMaxLength(20, "最多20个字符")
```

---

### RangeRule - 数值范围验证

```go
func NewRange(min, max float64, message ...string) RangeRule

type RangeRule struct {
    Min     float64
    Max     float64
    Message string
    Trigger string
}
```

**示例**:
```go
fb.NewRange(18, 100, "年龄必须在18-100之间")
```

---

### EmailRule - 邮箱验证

```go
func NewEmail(message ...string) EmailRule

type EmailRule struct {
    Message string
    Trigger string
}
```

**示例**:
```go
fb.NewEmail("请输入正确的邮箱地址")
fb.NewEmail()  // 使用默认消息
```

---

### URLRule - URL验证

```go
func NewURL(message ...string) URLRule

type URLRule struct {
    Message string
    Trigger string
}
```

**示例**:
```go
fb.NewURL("请输入正确的URL")
```

---

### DateRule - 日期验证

```go
func NewDate(message ...string) DateRule

type DateRule struct {
    Message string
    Trigger string
}
```

---

### EnumRule - 枚举验证

```go
func NewEnum(enum []interface{}, message ...string) EnumRule

type EnumRule struct {
    Enum    []interface{}
    Message string
    Trigger string
}
```

**示例**:
```go
fb.NewEnum([]interface{}{"admin", "user", "guest"}, "角色必须是admin、user或guest")
```

---

### CustomRule - 自定义验证

```go
type CustomRule struct {
    Validator string  // JavaScript验证函数
    Message   string
    Trigger   string
}
```

**示例**:
```go
fb.CustomRule{
    Validator: "function(rule, value, callback) { if(value !== this.form.password) { callback(new Error('两次密码不一致')); } else { callback(); } }",
    Message:   "两次密码不一致",
    Trigger:   "blur",
}
```

---

## 工厂方法

### ElmFactory - Element UI 工厂

全局单例: `fb.Elm`

**表单创建**:
```go
func (ElmFactory) CreateForm(action string, rules ...interface{}) *Form
```

**配置创建**:
```go
func (ElmFactory) Config() *Config
```

**所有组件方法**: 参见[表单组件](#表单组件)章节

---

### IviewFactory - iView v3 工厂

全局单例: `fb.Iview`

方法同 `ElmFactory`，生成的组件类型值与 Element UI 相同（如 `input`、`datePicker` 等）。

**注意**：Element UI 和 iView 使用相同的 type 值，UI 框架由全局配置决定，而非 type 字段。

---

### Iview4Factory - iView v4 工厂

全局单例: `fb.Iview4`

方法同 `ElmFactory`，用于 iView v4 (View Design)。

---

## 表单配置

### Config

```go
type Config struct {
    // 提交按钮
    SubmitBtnShow bool
    SubmitBtnText string
    SubmitBtnProps map[string]interface{}

    // 重置按钮
    ResetBtnShow bool
    ResetBtnText string
    ResetBtnProps map[string]interface{}

    // 表单样式
    FormStyle map[string]interface{}

    // 全局配置
    GlobalConfig map[string]map[string]interface{}
}
```

**方法**:

```go
func (c *Config) SubmitBtn(show bool, text ...string) *Config
func (c *Config) ResetBtn(show bool, text ...string) *Config
func (c *Config) FormStyle(style map[string]interface{}) *Config
func (c *Config) Global(key string, config map[string]interface{}) *Config
func (c *Config) Build() map[string]interface{}
```

**示例**:
```go
config := fb.Elm.Config()
config.SubmitBtn(true, "提交表单")
config.ResetBtn(true, "重置")
config.FormStyle(map[string]interface{}{
    "labelWidth": "100px",
})
config.Global("upload", map[string]interface{}{
    "action": "/api/upload",
})
```

---

## 表单方法

### Form

**构造函数**:
```go
func NewElmForm(action string, rules []Component, config ...*Config) *Form
func NewIviewForm(action string, rules []Component, config ...*Config) *Form
func NewIview4Form(action string, rules []Component, config ...*Config) *Form
```

**表单配置**:
```go
func (f *Form) SetAction(action string) *Form              // 设置提交地址
func (f *Form) SetMethod(method string) *Form              // 设置提交方法
func (f *Form) SetTitle(title string) *Form                // 设置表单标题
func (f *Form) SetRule(rules []Component) *Form            // 设置表单规则
func (f *Form) AppendRule(rules ...Component) *Form        // 追加规则
func (f *Form) PrependRule(rules ...Component) *Form       // 前置规则
func (f *Form) SetConfig(config *Config) *Form             // 设置配置
```

**数据管理**:
```go
func (f *Form) FormData(data map[string]interface{}) *Form  // 预填充数据
func (f *Form) SetValue(field string, value interface{}) *Form  // 设置单个字段值
func (f *Form) RemoveField(field string) *Form              // 移除字段
```

**输出方法**:
```go
func (f *Form) FormRule() []map[string]interface{}         // 获取规则数组
func (f *Form) ParseFormRule() (string, error)             // 获取规则JSON
func (f *Form) FormConfig() map[string]interface{}         // 获取配置对象
func (f *Form) ParseFormConfig() (string, error)           // 获取配置JSON
func (f *Form) FormScript() string                         // 获取Vue初始化脚本
func (f *Form) View() (string, error)                      // 生成完整HTML页面
func (f *Form) Template(templateContent string) (string, error)  // 使用自定义模板
```

**内部方法**:
```go
func (f *Form) GetUI() UIBootstrap                         // 获取UI实例
func (f *Form) GetTitle() string                           // 获取标题
```

---

## 条件显示（Control）

### ControlRule

```go
type ControlRule struct {
    Value interface{}      // 触发值
    Rule  []Component      // 要显示的组件
}
```

**使用示例**:
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
                fb.Elm.Input("phone", "电话").Required(),
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

**嵌套Control**:
```go
fb.Elm.Switch("need_invoice", "是否需要发票").
    Control([]fb.ControlRule{
        {
            Value: true,
            Rule: []fb.Component{
                fb.Elm.Radio("invoice_type", "发票类型", "personal").
                    SetOptions([]fb.Option{
                        {Value: "personal", Label: "个人"},
                        {Value: "company", Label: "公司"},
                    }).
                    Control([]fb.ControlRule{
                        {
                            Value: "company",
                            Rule: []fb.Component{
                                fb.Elm.Input("company_name", "公司名称").Required(),
                                fb.Elm.Input("tax_number", "税号").Required(),
                            },
                        },
                    }),
            },
        },
    })
```

---

## 完整示例

```go
package main

import (
    "fmt"
    fb "github.com/maverick/form-builder-go/formbuilder"
)

func main() {
    // 创建配置
    config := fb.Elm.Config()
    config.SubmitBtn(true, "提交")
    config.ResetBtn(true, "重置")

    // 创建表单
    form := fb.Elm.CreateForm("/api/user/create", []fb.Component{
        // 基础输入
        fb.Elm.Input("username", "用户名").
            Placeholder("请输入用户名").
            Clearable(true).
            Validate(
                fb.NewRequired("用户名不能为空"),
                fb.NewLength(6, 20, "长度必须在6-20个字符之间"),
            ),

        // 邮箱（自动验证）
        fb.Email("email", "邮箱").
            Placeholder("请输入邮箱").
            Required(),

        // 下拉选择
        fb.Elm.Select("role", "角色").
            SetOptions([]fb.Option{
                {Value: "admin", Label: "管理员"},
                {Value: "user", Label: "普通用户"},
            }).
            Clearable(true).
            Required(),

        // 单选 + 条件显示
        fb.Elm.Radio("user_type", "用户类型", "1").
            SetOptions([]fb.Option{
                {Value: "1", Label: "试用期"},
                {Value: "2", Label: "正式"},
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
            }),

        // 开关
        fb.Elm.Switch("is_active", "是否启用").
            ActiveText("启用").
            InactiveText("禁用"),
    }, config)

    // 预填充数据
    form.FormData(map[string]interface{}{
        "user_type": "1",
        "is_active": true,
    })

    // 输出JSON（用于API）
    jsonRule, _ := form.ParseFormRule()
    fmt.Println(jsonRule)

    // 输出HTML（用于直接渲染）
    html, _ := form.View()
    fmt.Println(html)
}
```

---

## 类型定义总结

```go
// 核心接口
type Component interface
type ValidateRule interface
type UIBootstrap interface

// 核心结构
type ComponentData struct
type Builder[T any] struct
type Form struct
type Config struct
type Option struct
type ControlRule struct

// 组件类型
type Input struct { Builder[*Input] }
type Select struct { Builder[*Select] }
type Radio struct { Builder[*Radio] }
type Checkbox struct { Builder[*Checkbox] }
type InputNumber struct { Builder[*InputNumber] }
type DatePicker struct { Builder[*DatePicker] }
type TimePicker struct { Builder[*TimePicker] }
type Slider struct { Builder[*Slider] }
type Switch struct { Builder[*Switch] }
type Upload struct { Builder[*Upload] }
type Cascader struct { Builder[*Cascader] }
type Tree struct { Builder[*Tree] }
type Rate struct { Builder[*Rate] }
type ColorPicker struct { Builder[*ColorPicker] }
type Hidden struct { Builder[*Hidden] }

// 验证规则
type RequiredRule struct
type PatternRule struct
type LengthRule struct
type RangeRule struct
type EmailRule struct
type URLRule struct
type DateRule struct
type EnumRule struct
type CustomRule struct

// 工厂
type ElmFactory struct
type IviewFactory struct
type Iview4Factory struct

// UI Bootstrap
type ElmBootstrap struct
type IviewBootstrap struct
```

---

## 附录

### 常用UI属性

**尺寸 (size)**:
- `large` - 大
- `medium` - 中等（默认）
- `small` - 小
- `mini` - 迷你

**触发方式 (trigger)**:
- `blur` - 失去焦点时
- `change` - 值改变时

**日期格式 (format)**:
- `yyyy-MM-dd` - 2024-01-01
- `yyyy-MM-dd HH:mm:ss` - 2024-01-01 12:00:00
- `yyyy/MM/dd` - 2024/01/01

**颜色格式 (colorFormat)**:
- `hex` - #FFFFFF
- `rgb` - rgb(255, 255, 255)
- `hsl` - hsl(0, 0%, 100%)
- `hsv` - hsv(0, 0%, 100%)

---

**FormBuilder Go** - 使用Go 1.18+泛型构建的强大表单生成器
