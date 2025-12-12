package formbuilder

// textarea.go 实现Textarea多行文本框组件
// 注：Element UI中textarea是input组件的type="textarea"
// 但为了API的清晰性，我们单独提供Textarea构造函数

// Textarea 创建多行文本框
func Textarea(field, title string, value ...interface{}) *Input {
	input := NewInput(field, title, value...)
	input.Type("textarea")
	return input
}

// NewTextarea 创建多行文本框（与Textarea相同）
func NewTextarea(field, title string, value ...interface{}) *Input {
	return Textarea(field, title, value...)
}
