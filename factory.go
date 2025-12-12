package formbuilder

// factory.go 实现工厂方法
// 对应PHP的Factory/Elm.php和Factory/Iview.php
// 提供便捷的组件创建API

// ElmFactory Element UI工厂
type ElmFactory struct{}

// 全局单例
var Elm ElmFactory

// Input 创建输入框
func (ElmFactory) Input(field, title string, args ...interface{}) *Input {
	var value interface{}
	inputType := "text"

	if len(args) > 0 {
		value = args[0]
	}
	if len(args) > 1 {
		if t, ok := args[1].(string); ok {
			inputType = t
		}
	}

	input := NewInput(field, title)
	if value != nil {
		input.Value(value)
	}
	input.Type(inputType)
	return input
}

// Password 创建密码输入框
func (ElmFactory) Password(field, title string, value ...interface{}) *Input {
	return Password(field, title, value...)
}

// Textarea 创建多行文本框
func (ElmFactory) Textarea(field, title string, value ...interface{}) *Input {
	return Textarea(field, title, value...)
}

// Select 创建下拉选择框
func (ElmFactory) Select(field, title string, value ...interface{}) *Select {
	return NewSelect(field, title, value...)
}

// Radio 创建单选框
func (ElmFactory) Radio(field, title string, value ...interface{}) *Radio {
	return NewRadio(field, title, value...)
}

// Checkbox 创建复选框
func (ElmFactory) Checkbox(field, title string, value ...interface{}) *Checkbox {
	return NewCheckbox(field, title, value...)
}

// Number 创建数字输入框
func (ElmFactory) Number(field, title string, value ...interface{}) *InputNumber {
	return NewInputNumber(field, title, value...)
}

// DatePicker 创建日期选择器
func (ElmFactory) DatePicker(field, title string, value ...interface{}) *DatePicker {
	return NewDatePicker(field, title, value...)
}

// TimePicker 创建时间选择器
func (ElmFactory) TimePicker(field, title string, value ...interface{}) *TimePicker {
	return NewTimePicker(field, title, value...)
}

// Slider 创建滑块
func (ElmFactory) Slider(field, title string, value ...interface{}) *Slider {
	return NewSlider(field, title, value...)
}

// Switch 创建开关
func (ElmFactory) Switch(field, title string, value ...interface{}) *Switch {
	return NewSwitch(field, title, value...)
}

// Upload 创建文件上传
func (ElmFactory) Upload(field, title string, value ...interface{}) *Upload {
	return NewUpload(field, title, value...)
}

// Cascader 创建级联选择器
func (ElmFactory) Cascader(field, title string, value ...interface{}) *Cascader {
	return NewCascader(field, title, value...)
}

// Tree 创建树形控件
func (ElmFactory) Tree(field, title string, value ...interface{}) *Tree {
	return NewTree(field, title, value...)
}

// Rate 创建评分组件
func (ElmFactory) Rate(field, title string, value ...interface{}) *Rate {
	return NewRate(field, title, value...)
}

// ColorPicker 创建颜色选择器
func (ElmFactory) ColorPicker(field, title string, value ...interface{}) *ColorPicker {
	return NewColorPicker(field, title, value...)
}

// Hidden 创建隐藏字段
func (ElmFactory) Hidden(field string, value ...interface{}) *Hidden {
	return NewHidden(field, value...)
}

// Frame 创建框架组件（通用）
func (ElmFactory) Frame(field, title, src string, value ...interface{}) *Frame {
	return NewFrame(field, title, src, value...)
}

// FrameImage 创建单图片框架组件
// value为string类型
func (ElmFactory) FrameImage(field, title, src string, value ...interface{}) *Frame {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewFrame(field, title, src, val).
		Type(FrameTypeImage).
		MaxLength(1)
}

// FrameImages 创建多图片框架组件
// value为Array类型
func (ElmFactory) FrameImages(field, title, src string, value ...interface{}) *Frame {
	var val interface{} = []interface{}{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewFrame(field, title, src, val).
		Type(FrameTypeImage).
		MaxLength(0)
}

// FrameFile 创建单文件框架组件
// value为string类型
func (ElmFactory) FrameFile(field, title, src string, value ...interface{}) *Frame {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewFrame(field, title, src, val).
		Type(FrameTypeFile).
		MaxLength(1)
}

// FrameFiles 创建多文件框架组件
// value为Array类型
func (ElmFactory) FrameFiles(field, title, src string, value ...interface{}) *Frame {
	var val interface{} = []interface{}{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewFrame(field, title, src, val).
		Type(FrameTypeFile).
		MaxLength(0)
}

// FrameInput 创建单输入框架组件
// value为string类型
func (ElmFactory) FrameInput(field, title, src string, value ...interface{}) *Frame {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewFrame(field, title, src, val).
		Type(FrameTypeInput).
		MaxLength(1)
}

// FrameInputs 创建多输入框架组件
// value为Array类型
func (ElmFactory) FrameInputs(field, title, src string, value ...interface{}) *Frame {
	var val interface{} = []interface{}{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewFrame(field, title, src, val).
		Type(FrameTypeInput).
		MaxLength(0)
}

// UploadFile 创建单文件上传组件
// value为string类型，自动设置limit为1
func (ElmFactory) UploadFile(field, title, action string, value ...interface{}) *Upload {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	}
	upload := NewUpload(field, title, val)
	upload.data.Props["uploadType"] = "file"
	return upload.Action(action).Limit(1)
}

// UploadFiles 创建多文件上传组件
// value为Array类型
func (ElmFactory) UploadFiles(field, title, action string, value ...interface{}) *Upload {
	var val interface{} = []interface{}{}
	if len(value) > 0 {
		val = value[0]
	}
	upload := NewUpload(field, title, val)
	upload.data.Props["uploadType"] = "file"
	return upload.Action(action)
}

// UploadImage 创建单图片上传组件
// value为string类型，自动设置accept为image/*，limit为1
func (ElmFactory) UploadImage(field, title, action string, value ...interface{}) *Upload {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	}
	upload := NewUpload(field, title, val)
	upload.data.Props["uploadType"] = "image"
	return upload.Action(action).Accept("image/*").Limit(1)
}

// UploadImages 创建多图片上传组件
// value为Array类型，自动设置accept为image/*
func (ElmFactory) UploadImages(field, title, action string, value ...interface{}) *Upload {
	var val interface{} = []interface{}{}
	if len(value) > 0 {
		val = value[0]
	}
	upload := NewUpload(field, title, val)
	upload.data.Props["uploadType"] = "image"
	return upload.Action(action).Accept("image/*")
}

// CreateForm 创建表单
func (ElmFactory) CreateForm(action string, args ...interface{}) *Form {
	var rules []Component
	var config *Config

	if len(args) > 0 {
		if r, ok := args[0].([]Component); ok {
			rules = r
		}
	}
	if len(args) > 1 {
		if c, ok := args[1].(*Config); ok {
			config = c
		}
	}

	return NewElmForm(action, rules, config)
}

// Config 创建配置
func (ElmFactory) Config() *Config {
	return NewElmConfig()
}

// Option 创建选项
func (ElmFactory) Option(value interface{}, label string, disabled ...bool) Option {
	opt := Option{
		Value: value,
		Label: label,
	}
	if len(disabled) > 0 {
		opt.Disabled = disabled[0]
	}
	return opt
}

// IviewFactory iView工厂
type IviewFactory struct {
	version int // 3 or 4
}

// 全局单例
var (
	Iview  = IviewFactory{version: 3}
	Iview4 = IviewFactory{version: 4}
)

// Input 创建输入框
func (f IviewFactory) Input(field, title string, args ...interface{}) *Input {
	var value interface{}
	inputType := "text"

	if len(args) > 0 {
		value = args[0]
	}
	if len(args) > 1 {
		if t, ok := args[1].(string); ok {
			inputType = t
		}
	}

	input := NewIviewInput(field, title)
	if value != nil {
		input.Value(value)
	}
	input.Type(inputType)
	return input
}

// Password 创建密码输入框
func (f IviewFactory) Password(field, title string, value ...interface{}) *Input {
	input := NewIviewInput(field, title, value...)
	input.Type("password")
	return input
}

// Textarea 创建多行文本框
func (f IviewFactory) Textarea(field, title string, value ...interface{}) *Input {
	input := NewIviewInput(field, title, value...)
	input.Type("textarea")
	return input
}

// Select 创建下拉选择框
func (f IviewFactory) Select(field, title string, value ...interface{}) *Select {
	return NewIviewSelect(field, title, value...)
}

// Radio 创建单选框
func (f IviewFactory) Radio(field, title string, value ...interface{}) *Radio {
	return NewIviewRadio(field, title, value...)
}

// Checkbox 创建复选框
func (f IviewFactory) Checkbox(field, title string, value ...interface{}) *Checkbox {
	return NewIviewCheckbox(field, title, value...)
}

// Number 创建数字输入框
func (f IviewFactory) Number(field, title string, value ...interface{}) *InputNumber {
	return NewIviewInputNumber(field, title, value...)
}

// DatePicker 创建日期选择器
func (f IviewFactory) DatePicker(field, title string, value ...interface{}) *DatePicker {
	dp := NewDatePicker(field, title, value...)
	dp.data.RuleType = "datePicker"
	return dp
}

// TimePicker 创建时间选择器
func (f IviewFactory) TimePicker(field, title string, value ...interface{}) *TimePicker {
	tp := NewTimePicker(field, title, value...)
	tp.data.RuleType = "timePicker"
	return tp
}

// Slider 创建滑块
func (f IviewFactory) Slider(field, title string, value ...interface{}) *Slider {
	slider := NewSlider(field, title, value...)
	slider.data.RuleType = "slider"
	return slider
}

// Switch 创建开关
func (f IviewFactory) Switch(field, title string, value ...interface{}) *Switch {
	sw := NewSwitch(field, title, value...)
	sw.data.RuleType = "switch"
	return sw
}

// Upload 创建文件上传
func (f IviewFactory) Upload(field, title string, value ...interface{}) *Upload {
	upload := NewUpload(field, title, value...)
	upload.data.RuleType = "upload"
	return upload
}

// Cascader 创建级联选择器
func (f IviewFactory) Cascader(field, title string, value ...interface{}) *Cascader {
	cascader := NewCascader(field, title, value...)
	cascader.data.RuleType = "cascader"
	return cascader
}

// Tree 创建树形控件
func (f IviewFactory) Tree(field, title string, value ...interface{}) *Tree {
	tree := NewTree(field, title, value...)
	tree.data.RuleType = "tree"
	return tree
}

// Rate 创建评分组件
func (f IviewFactory) Rate(field, title string, value ...interface{}) *Rate {
	rate := NewRate(field, title, value...)
	rate.data.RuleType = "rate"
	return rate
}

// ColorPicker 创建颜色选择器
func (f IviewFactory) ColorPicker(field, title string, value ...interface{}) *ColorPicker {
	cp := NewColorPicker(field, title, value...)
	cp.data.RuleType = "colorPicker"
	return cp
}

// Hidden 创建隐藏字段
func (f IviewFactory) Hidden(field string, value ...interface{}) *Hidden {
	return NewHidden(field, value...)
}

// Frame 创建框架组件（通用）
func (f IviewFactory) Frame(field, title, src string, value ...interface{}) *Frame {
	return NewIviewFrame(field, title, src, value...)
}

// FrameImage 创建单图片框架组件
// value为string类型
func (f IviewFactory) FrameImage(field, title, src string, value ...interface{}) *Frame {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewIviewFrame(field, title, src, val).
		Type(FrameTypeImage).
		MaxLength(1)
}

// FrameImages 创建多图片框架组件
// value为Array类型
func (f IviewFactory) FrameImages(field, title, src string, value ...interface{}) *Frame {
	var val interface{} = []interface{}{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewIviewFrame(field, title, src, val).
		Type(FrameTypeImage).
		MaxLength(0)
}

// FrameFile 创建单文件框架组件
// value为string类型
func (f IviewFactory) FrameFile(field, title, src string, value ...interface{}) *Frame {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewIviewFrame(field, title, src, val).
		Type(FrameTypeFile).
		MaxLength(1)
}

// FrameFiles 创建多文件框架组件
// value为Array类型
func (f IviewFactory) FrameFiles(field, title, src string, value ...interface{}) *Frame {
	var val interface{} = []interface{}{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewIviewFrame(field, title, src, val).
		Type(FrameTypeFile).
		MaxLength(0)
}

// FrameInput 创建单输入框架组件
// value为string类型
func (f IviewFactory) FrameInput(field, title, src string, value ...interface{}) *Frame {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewIviewFrame(field, title, src, val).
		Type(FrameTypeInput).
		MaxLength(1)
}

// FrameInputs 创建多输入框架组件
// value为Array类型
func (f IviewFactory) FrameInputs(field, title, src string, value ...interface{}) *Frame {
	var val interface{} = []interface{}{}
	if len(value) > 0 {
		val = value[0]
	}
	return NewIviewFrame(field, title, src, val).
		Type(FrameTypeInput).
		MaxLength(0)
}

// UploadFile 创建单文件上传组件
// value为string类型，自动设置limit为1
func (f IviewFactory) UploadFile(field, title, action string, value ...interface{}) *Upload {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	}
	upload := NewUpload(field, title, val)
	upload.data.RuleType = "upload"
	upload.data.Props["uploadType"] = "file"
	return upload.Action(action).Limit(1)
}

// UploadFiles 创建多文件上传组件
// value为Array类型
func (f IviewFactory) UploadFiles(field, title, action string, value ...interface{}) *Upload {
	var val interface{} = []interface{}{}
	if len(value) > 0 {
		val = value[0]
	}
	upload := NewUpload(field, title, val)
	upload.data.RuleType = "upload"
	upload.data.Props["uploadType"] = "file"
	return upload.Action(action)
}

// UploadImage 创建单图片上传组件
// value为string类型，自动设置accept为image/*，limit为1
func (f IviewFactory) UploadImage(field, title, action string, value ...interface{}) *Upload {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	}
	upload := NewUpload(field, title, val)
	upload.data.RuleType = "upload"
	upload.data.Props["uploadType"] = "image"
	return upload.Action(action).Accept("image/*").Limit(1)
}

// UploadImages 创建多图片上传组件
// value为Array类型，自动设置accept为image/*
func (f IviewFactory) UploadImages(field, title, action string, value ...interface{}) *Upload {
	var val interface{} = []interface{}{}
	if len(value) > 0 {
		val = value[0]
	}
	upload := NewUpload(field, title, val)
	upload.data.RuleType = "upload"
	upload.data.Props["uploadType"] = "image"
	return upload.Action(action).Accept("image/*")
}

// CreateForm 创建表单
func (f IviewFactory) CreateForm(action string, args ...interface{}) *Form {
	var rules []Component
	var config *Config

	if len(args) > 0 {
		if r, ok := args[0].([]Component); ok {
			rules = r
		}
	}
	if len(args) > 1 {
		if c, ok := args[1].(*Config); ok {
			config = c
		}
	}

	if f.version == 4 {
		return NewIview4Form(action, rules, config)
	}
	return NewIviewForm(action, rules, config)
}

// Config 创建配置
func (f IviewFactory) Config() *Config {
	return NewIviewConfig()
}

// Option 创建选项
func (f IviewFactory) Option(value interface{}, label string, disabled ...bool) Option {
	opt := Option{
		Value: value,
		Label: label,
	}
	if len(disabled) > 0 {
		opt.Disabled = disabled[0]
	}
	return opt
}
