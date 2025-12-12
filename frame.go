package formbuilder

// frame.go 实现Frame框架组件
// 对应PHP版本的 Frame 组件，用于在弹出框中选择内容

// Frame 框架组件
// 用于在iframe弹出框中选择内容（图片、文件、输入等）
type Frame struct {
	Builder[*Frame]
}

// Frame类型常量
const (
	FrameTypeInput = "input"
	FrameTypeFile  = "file"
	FrameTypeImage = "image"
)

// NewFrame 创建一个新的Frame组件
func NewFrame(field, title, src string, value ...interface{}) *Frame {
	frame := &Frame{}
	frame.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "frame",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		frame.data.Value = value[0]
	}
	frame.inst = frame
	// 设置默认值
	frame.Src(src).Type(FrameTypeInput).MaxLength(0)
	return frame
}

// NewIviewFrame 创建iView版本的Frame组件
func NewIviewFrame(field, title, src string, value ...interface{}) *Frame {
	frame := NewFrame(field, title, src, value...)
	frame.data.RuleType = "frame"
	return frame
}

// Type 设置frame类型
// 可选值: input, file, image
func (f *Frame) Type(frameType string) *Frame {
	f.data.Props["type"] = frameType
	return f
}

// Src 设置iframe地址
func (f *Frame) Src(src string) *Frame {
	f.data.Props["src"] = src
	return f
}

// MaxLength 设置value的最大数量
// 默认为0表示无限制，设置为1表示单选
func (f *Frame) MaxLength(length int) *Frame {
	f.data.Props["maxLength"] = length
	return f
}

// Icon 设置打开弹出框的按钮图标
func (f *Frame) Icon(icon string) *Frame {
	f.data.Props["icon"] = icon
	return f
}

// Height 设置弹出框高度
// 例如: "500px", "80vh"
func (f *Frame) Height(height string) *Frame {
	f.data.Props["height"] = height
	return f
}

// Width 设置弹出框宽度
// 例如: "800px", "90vw"
func (f *Frame) Width(width string) *Frame {
	f.data.Props["width"] = width
	return f
}

// Spin 设置是否显示加载动画
// 默认为true
func (f *Frame) Spin(enable bool) *Frame {
	f.data.Props["spin"] = enable
	return f
}

// FrameTitle 设置弹出框标题
func (f *Frame) FrameTitle(title string) *Frame {
	f.data.Props["frameTitle"] = title
	return f
}

// Modal 设置弹出框props
func (f *Frame) Modal(modalProps map[string]interface{}) *Frame {
	f.data.Props["modal"] = modalProps
	return f
}

// HandleIcon 设置操作按钮的图标
// 设置为false将不显示，设置为true为默认的预览图标
// 类型为file时默认为false，image类型默认为true
func (f *Frame) HandleIcon(enable bool) *Frame {
	f.data.Props["handleIcon"] = enable
	return f
}

// AllowRemove 设置是否可删除
// 设置为false时不显示删除按钮
func (f *Frame) AllowRemove(enable bool) *Frame {
	f.data.Props["allowRemove"] = enable
	return f
}

// Disabled 设置是否禁用
func (f *Frame) Disabled(disabled bool) *Frame {
	f.data.Props["disabled"] = disabled
	return f
}

// GetField 实现Component接口
func (f *Frame) GetField() string {
	return f.data.Field
}

// GetType 实现Component接口
func (f *Frame) GetType() string {
	return f.data.RuleType
}

// Build 实现Component接口
func (f *Frame) Build() map[string]interface{} {
	return buildComponent(f.data)
}
