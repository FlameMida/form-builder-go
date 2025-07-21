// Package components provides form component implementations
package components

import (
	"fmt"
	"github.com/FlameMida/form-builder-go/contracts"
)

const (
	FrameTypeImage = "image"
	FrameTypeFile  = "file"
	FrameTypeInput = "input"
)

// Frame component implementation
type Frame struct {
	*BaseComponent
}

// NewFrame creates a new frame component
func NewFrame(field, title string) *Frame {
	frame := &Frame{
		BaseComponent: NewBaseComponent(field, title),
	}
	frame.SetProp("type", FrameTypeInput)
	frame.SetProp("maxLength", 0)
	frame.Placeholder(fmt.Sprintf("请选择%s", title))
	frame.SetType("frame")
	return frame
}

// SetPropType sets the frame type
func (f *Frame) SetPropType(frameType string) *Frame {
	f.SetProp("type", frameType)
	return f
}

// Src sets the iframe src
func (f *Frame) Src(src string) *Frame {
	f.SetProp("src", src)
	return f
}

// MaxLength sets the max length of value
func (f *Frame) MaxLength(length int) *Frame {
	f.SetProp("maxLength", length)
	return f
}

// Icon sets the button icon
func (f *Frame) Icon(icon string) *Frame {
	f.SetProp("icon", icon)
	return f
}

// Spin sets the loading spin status
func (f *Frame) Spin(spin bool) *Frame {
	f.SetProp("spin", spin)
	return f
}

// FrameTitle sets the modal title
func (f *Frame) FrameTitle(title string) *Frame {
	f.SetProp("title", title)
	return f
}

// Modal sets the modal props
func (f *Frame) Modal(modalProps map[string]interface{}) *Frame {
	f.SetProp("modal", modalProps)
	return f
}

// HandleIcon sets the handle icon visibility
func (f *Frame) HandleIcon(show bool) *Frame {
	f.SetProp("handleIcon", show)
	return f
}

// AllowRemove sets the remove button visibility
func (f *Frame) AllowRemove(allow bool) *Frame {
	f.SetProp("allowRemove", allow)
	return f
}

// Width sets the modal width
func (f *Frame) Width(width string) *Frame {
	f.SetProp("width", width)
	return f
}

// Height sets the modal height
func (f *Frame) Height(height string) *Frame {
	f.SetProp("height", height)
	return f
}

// Required makes the frame required
func (f *Frame) Required() contracts.FormComponent {
	maxLength, _ := f.GetProp("maxLength").(int)
	var rule contracts.ValidateRule
	if maxLength == 1 {
		rule = NewStringRequiredRule(fmt.Sprintf("请选择%s", f.title))
	} else {
		rule = NewArrayRequiredRule(fmt.Sprintf("请选择%s", f.title))
	}
	f.AddValidateRule(rule)
	return f
}

// Placeholder sets placeholder
func (f *Frame) Placeholder(text string) contracts.FormComponent {
	f.SetProp("placeholder", text)
	return f
}

// Disabled sets the disabled state
func (f *Frame) Disabled(disabled bool) contracts.FormComponent {
	f.SetProp("disabled", disabled)
	return f
}

// Build returns the frame component as a map
func (f *Frame) Build() map[string]interface{} {
	result := f.BaseComponent.Build()
	result["type"] = "frame"
	return result
}
