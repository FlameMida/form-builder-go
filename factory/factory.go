// Package factory provides factory methods for creating form components
package factory

import (
	"github.com/FlameMida/form-builder-go/components"
	"github.com/FlameMida/form-builder-go/contracts"
	"github.com/FlameMida/form-builder-go/formbuilder"
	"github.com/FlameMida/form-builder-go/ui/elm"
	"github.com/FlameMida/form-builder-go/ui/iview"
)

// Elm provides ElementUI component factory methods
type Elm struct{}

// Input creates an ElementUI input component
func (Elm) Input(field, title string, value ...string) *components.Input {
	input := components.NewInput(field, title)
	if len(value) > 0 && value[0] != "" {
		input.SetValue(value[0])
	}
	return input
}

// Textarea creates an ElementUI textarea component
func (Elm) Textarea(field, title string, value ...string) *components.Textarea {
	textarea := components.NewTextarea(field, title)
	if len(value) > 0 && value[0] != "" {
		textarea.SetValue(value[0])
	}
	return textarea
}

// Switch creates an ElementUI switch component
func (Elm) Switch(field, title string, value ...string) *components.Switch {
	switch_ := components.NewSwitch(field, title)
	if len(value) > 0 && value[0] != "" {
		switch_.SetValue(value[0])
	}
	return switch_
}

// Select creates an ElementUI select component
func (Elm) Select(field, title string, value ...string) *components.Select {
	select_ := components.NewSelect(field, title)
	if len(value) > 0 && value[0] != "" {
		select_.SetValue(value[0])
	}
	return select_
}

// Radio creates an ElementUI radio component
func (Elm) Radio(field, title string, value ...string) *components.Radio {
	radio := components.NewRadio(field, title)
	if len(value) > 0 && value[0] != "" {
		radio.SetValue(value[0])
	}
	return radio
}

// Checkbox creates an ElementUI checkbox component
func (Elm) Checkbox(field, title string, value ...string) *components.Checkbox {
	checkbox := components.NewCheckbox(field, title)
	if len(value) > 0 {
		checkbox.SetValue(value[0])
	}
	return checkbox
}

// InputNumber creates an ElementUI input number component
func (Elm) InputNumber(field, title string, value ...string) *components.InputNumber {
	inputNumber := components.NewInputNumber(field, title)
	if len(value) > 0 {
		inputNumber.SetValue(value[0])
	}
	return inputNumber
}

// Slider creates an ElementUI slider component
func (Elm) Slider(field, title string, value ...string) *components.Slider {
	slider := components.NewSlider(field, title)
	if len(value) > 0 {
		slider.SetValue(value[0])
	}
	return slider
}

// Rate creates an ElementUI rate component
func (Elm) Rate(field, title string, value ...string) *components.Rate {
	rate := components.NewRate(field, title)
	if len(value) > 0 {
		rate.SetValue(value[0])
	}
	return rate
}

// DatePicker creates an ElementUI date picker component
func (Elm) DatePicker(field, title string, value ...string) *components.DatePicker {
	datePicker := components.NewDatePicker(field, title)
	if len(value) > 0 {
		datePicker.SetValue(value[0])
	}
	return datePicker
}

// TimePicker creates an ElementUI time picker component
func (Elm) TimePicker(field, title string, value ...string) *components.TimePicker {
	timePicker := components.NewTimePicker(field, title)
	if len(value) > 0 {
		timePicker.SetValue(value[0])
	}
	return timePicker
}

// ColorPicker creates an ElementUI color picker component
func (Elm) ColorPicker(field, title string, value ...string) *components.ColorPicker {
	colorPicker := components.NewColorPicker(field, title)
	if len(value) > 0 {
		colorPicker.SetValue(value[0])
	}
	return colorPicker
}

// Upload creates an ElementUI upload component
func (Elm) Upload(field, title string, value ...string) *components.Upload {
	upload := components.NewUpload(field, title)
	if len(value) > 0 {
		upload.SetValue(value[0])
	}
	return upload
}

// Cascader creates an ElementUI cascader component
func (Elm) Cascader(field, title string, value ...string) *components.Cascader {
	cascader := components.NewCascader(field, title)
	if len(value) > 0 {
		cascader.SetValue(value[0])
	}
	return cascader
}

// Tree creates an ElementUI tree component
func (Elm) Tree(field, title string, value ...string) *components.Tree {
	tree := components.NewTree(field, title)
	if len(value) > 0 {
		tree.SetValue(value[0])
	}
	return tree
}

// Button creates an ElementUI button component
func (Elm) Button(field, title string, value ...string) *components.Button {
	button := components.NewButton(field, title)
	if len(value) > 0 {
		button.SetValue(value[0])
	}
	return button
}

// Hidden creates a hidden field component
func (Elm) Hidden(field string, value ...string) *components.Hidden {
	hidden := components.NewHidden(field)
	if len(value) > 0 {
		hidden.SetValue(value[0])
	}
	return hidden
}

// Option creates an option for select components
func (Elm) Option(value interface{}, label string) contracts.Option {
	return contracts.Option{
		Value: value,
		Label: label,
	}
}

// TreeNode creates a tree node
func (Elm) TreeNode(label string, value interface{}) components.TreeNode {
	return components.TreeNode{
		Label: label,
		Value: value,
	}
}

// CascaderOption creates a cascader option
func (Elm) CascaderOption(label string, value interface{}) components.CascaderOption {
	return components.CascaderOption{
		Label: label,
		Value: value,
	}
}

// Config creates ElementUI form configuration
func (Elm) Config(config map[string]interface{}) contracts.ConfigInterface {
	return elm.NewConfig(config)
}

// CreateForm creates a new ElementUI form
func (Elm) CreateForm(action string, rules []interface{}, config map[string]interface{}) (contracts.Form, error) {
	bootstrap := elm.NewBootstrap()
	return formbuilder.NewForm(bootstrap, action, rules, config)
}

// Iview provides IView component factory methods
type Iview struct{}

// Input creates an IView input component
func (Iview) Input(field, title string, value ...string) *components.Input {
	input := components.NewInput(field, title)
	if len(value) > 0 {
		input.SetValue(value[0])
	}
	// IView specific modifications could go here
	return input
}

// Textarea creates an IView textarea component
func (Iview) Textarea(field, title string, value ...string) *components.Textarea {
	textarea := components.NewTextarea(field, title)
	if len(value) > 0 {
		textarea.SetValue(value[0])
	}
	// IView specific modifications could go here
	return textarea
}

// Switch creates an IView switch component
func (Iview) Switch(field, title string, value ...string) *components.Switch {
	switchComp := components.NewSwitch(field, title)
	if len(value) > 0 {
		switchComp.SetValue(value[0])
	}
	// IView specific modifications could go here
	return switchComp
}

// Select creates an IView select component
func (Iview) Select(field, title string, value ...string) *components.Select {
	select_ := components.NewSelect(field, title)
	if len(value) > 0 {
		select_.SetValue(value[0])
	}
	return select_
}

// Radio creates an IView radio component
func (Iview) Radio(field, title string, value ...string) *components.Radio {
	radio := components.NewRadio(field, title)
	if len(value) > 0 {
		radio.SetValue(value[0])
	}
	return radio
}

// Checkbox creates an IView checkbox component
func (Iview) Checkbox(field, title string, value ...string) *components.Checkbox {
	checkbox := components.NewCheckbox(field, title)
	if len(value) > 0 {
		checkbox.SetValue(value[0])
	}
	return checkbox
}

// InputNumber creates an IView input number component
func (Iview) InputNumber(field, title string, value ...string) *components.InputNumber {
	inputNumber := components.NewInputNumber(field, title)
	if len(value) > 0 {
		inputNumber.SetValue(value[0])
	}
	return inputNumber
}

// Slider creates an IView slider component
func (Iview) Slider(field, title string, value ...string) *components.Slider {
	slider := components.NewSlider(field, title)
	if len(value) > 0 {
		slider.SetValue(value[0])
	}
	return slider
}

// Rate creates an IView rate component
func (Iview) Rate(field, title string, value ...string) *components.Rate {
	rate := components.NewRate(field, title)
	if len(value) > 0 {
		rate.SetValue(value[0])
	}
	return rate
}

// DatePicker creates an IView date picker component
func (Iview) DatePicker(field, title string, value ...string) *components.DatePicker {
	datePicker := components.NewDatePicker(field, title)
	if len(value) > 0 {
		datePicker.SetValue(value[0])
	}
	return datePicker
}

// TimePicker creates an IView time picker component
func (Iview) TimePicker(field, title string, value ...string) *components.TimePicker {
	timePicker := components.NewTimePicker(field, title)
	if len(value) > 0 {
		timePicker.SetValue(value[0])
	}
	return timePicker
}

// ColorPicker creates an IView color picker component
func (Iview) ColorPicker(field, title string, value ...string) *components.ColorPicker {
	colorPicker := components.NewColorPicker(field, title)
	if len(value) > 0 {
		colorPicker.SetValue(value[0])
	}
	return colorPicker
}

// Upload creates an IView upload component
func (Iview) Upload(field, title string, value ...string) *components.Upload {
	upload := components.NewUpload(field, title)
	if len(value) > 0 {
		upload.SetValue(value[0])
	}
	return upload
}

// Cascader creates an IView cascader component
func (Iview) Cascader(field, title string, value ...string) *components.Cascader {
	cascader := components.NewCascader(field, title)
	if len(value) > 0 {
		cascader.SetValue(value[0])
	}
	return cascader
}

// Tree creates an IView tree component
func (Iview) Tree(field, title string, value ...string) *components.Tree {
	tree := components.NewTree(field, title)
	if len(value) > 0 {
		tree.SetValue(value[0])
	}
	return tree
}

// Button creates an IView button component
func (Iview) Button(field, title string, value ...string) *components.Button {
	button := components.NewButton(field, title)
	if len(value) > 0 {
		button.SetValue(value[0])
	}
	return button
}

// Hidden creates a hidden field component
func (Iview) Hidden(field string, value ...string) *components.Hidden {
	hidden := components.NewHidden(field)
	if len(value) > 0 {
		hidden.SetValue(value[0])
	}
	return hidden
}

// Option creates an option for select components
func (Iview) Option(value interface{}, label string) contracts.Option {
	return contracts.Option{
		Value: value,
		Label: label,
	}
}

// TreeNode creates a tree node
func (Iview) TreeNode(label string, value interface{}) components.TreeNode {
	return components.TreeNode{
		Label: label,
		Value: value,
	}
}

// CascaderOption creates a cascader option
func (Iview) CascaderOption(label string, value interface{}) components.CascaderOption {
	return components.CascaderOption{
		Label: label,
		Value: value,
	}
}

// Config creates IView form configuration
func (Iview) Config(config map[string]interface{}) contracts.ConfigInterface {
	return iview.NewConfig(config)
}

// CreateForm creates a new IView form
func (Iview) CreateForm(action string, rules []interface{}, config map[string]interface{}) (contracts.Form, error) {
	bootstrap := iview.NewBootstrap()
	return formbuilder.NewForm(bootstrap, action, rules, config)
}
