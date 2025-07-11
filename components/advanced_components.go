// Package components provides complex and utility components
package components

import (
	"fmt"

	"github.com/FlameMida/form-builder-go/contracts"
	"github.com/FlameMida/form-builder-go/rules"
)

// TreeNode represents a tree node structure
type TreeNode struct {
	Label    string      `json:"label"`
	Value    interface{} `json:"value"`
	Children []TreeNode  `json:"children,omitempty"`
	Disabled bool        `json:"disabled,omitempty"`
}

// CascaderOption represents a cascader option
type CascaderOption struct {
	Label    string           `json:"label"`
	Value    interface{}      `json:"value"`
	Children []CascaderOption `json:"children,omitempty"`
	Disabled bool             `json:"disabled,omitempty"`
}

// Cascader component implementation
type Cascader struct {
	*BaseComponent
	options []CascaderOption
}

// NewCascader creates a new cascader component
func NewCascader(field, title string) *Cascader {
	return &Cascader{
		BaseComponent: NewBaseComponent(field, title),
		options:       make([]CascaderOption, 0),
	}
}

// Options sets cascader options
func (c *Cascader) Options(options []CascaderOption) *Cascader {
	c.options = options
	return c
}

// AddOption adds a single cascader option
func (c *Cascader) AddOption(option CascaderOption) *Cascader {
	c.options = append(c.options, option)
	return c
}

// Props sets cascader props configuration
func (c *Cascader) Props(props map[string]interface{}) *Cascader {
	c.SetProp("props", props)
	return c
}

// Size sets the cascader size
func (c *Cascader) Size(size string) *Cascader {
	c.SetProp("size", size)
	return c
}

// Clearable makes the cascader clearable
func (c *Cascader) Clearable(clearable bool) *Cascader {
	c.SetProp("clearable", clearable)
	return c
}

// ShowAllLevels shows all levels in the input
func (c *Cascader) ShowAllLevels(show bool) *Cascader {
	c.SetProp("show-all-levels", show)
	return c
}

// CollapseTags collapses tags in multiple mode
func (c *Cascader) CollapseTags(collapse bool) *Cascader {
	c.SetProp("collapse-tags", collapse)
	return c
}

// Separator sets the separator for display
func (c *Cascader) Separator(separator string) *Cascader {
	c.SetProp("separator", separator)
	return c
}

// Filterable enables filtering
func (c *Cascader) Filterable(filterable bool) *Cascader {
	c.SetProp("filterable", filterable)
	return c
}

// FilterMethod sets custom filter method
func (c *Cascader) FilterMethod(method string) *Cascader {
	c.SetProp("filter-method", method)
	return c
}

// DebounceDelay sets debounce delay for filtering
func (c *Cascader) DebounceDelay(delay int) *Cascader {
	c.SetProp("debounce", delay)
	return c
}

// BeforeFilter sets before filter handler
func (c *Cascader) BeforeFilter(handler string) *Cascader {
	c.SetProp("before-filter", handler)
	return c
}

// PopperClass sets popper class
func (c *Cascader) PopperClass(class string) *Cascader {
	c.SetProp("popper-class", class)
	return c
}

// Required makes the cascader required
func (c *Cascader) Required() contracts.FormComponent {
	c.AddValidateRule(rules.NewRequiredRule(fmt.Sprintf("%s 是必填项", c.title)))
	return c
}

// Placeholder sets placeholder
func (c *Cascader) Placeholder(text string) contracts.FormComponent {
	c.SetProp("placeholder", text)
	return c
}

// Disabled sets the disabled state
func (c *Cascader) Disabled(disabled bool) contracts.FormComponent {
	c.SetProp("disabled", disabled)
	return c
}

// Build returns the cascader component as a map
func (c *Cascader) Build() map[string]interface{} {
	result := c.BaseComponent.Build()
	result["type"] = "el-cascader"

	// Add options to props
	if len(c.options) > 0 {
		result["options"] = c.options
	}

	return result
}

// Tree component implementation
type Tree struct {
	*BaseComponent
	data []TreeNode
}

// NewTree creates a new tree component
func NewTree(field, title string) *Tree {
	return &Tree{
		BaseComponent: NewBaseComponent(field, title),
		data:          make([]TreeNode, 0),
	}
}

// Data sets tree data
func (t *Tree) Data(data []TreeNode) *Tree {
	t.data = data
	return t
}

// AddNode adds a single tree node
func (t *Tree) AddNode(node TreeNode) *Tree {
	t.data = append(t.data, node)
	return t
}

// EmptyText sets empty text
func (t *Tree) EmptyText(text string) *Tree {
	t.SetProp("empty-text", text)
	return t
}

// NodeKey sets node key property
func (t *Tree) NodeKey(key string) *Tree {
	t.SetProp("node-key", key)
	return t
}

// Props sets tree props configuration
func (t *Tree) Props(props map[string]interface{}) *Tree {
	t.SetProp("props", props)
	return t
}

// RenderAfterExpand enables render after expand
func (t *Tree) RenderAfterExpand(render bool) *Tree {
	t.SetProp("render-after-expand", render)
	return t
}

// Load sets lazy load function
func (t *Tree) Load(loadFunc string) *Tree {
	t.SetProp("load", loadFunc)
	return t
}

// RenderContent sets render content function
func (t *Tree) RenderContent(renderFunc string) *Tree {
	t.SetProp("render-content", renderFunc)
	return t
}

// HighlightCurrent highlights current node
func (t *Tree) HighlightCurrent(highlight bool) *Tree {
	t.SetProp("highlight-current", highlight)
	return t
}

// DefaultExpandAll expands all nodes by default
func (t *Tree) DefaultExpandAll(expand bool) *Tree {
	t.SetProp("default-expand-all", expand)
	return t
}

// ExpandOnClickNode expands on click node
func (t *Tree) ExpandOnClickNode(expand bool) *Tree {
	t.SetProp("expand-on-click-node", expand)
	return t
}

// CheckOnClickNode checks on click node
func (t *Tree) CheckOnClickNode(check bool) *Tree {
	t.SetProp("check-on-click-node", check)
	return t
}

// AutoExpandParent auto expands parent
func (t *Tree) AutoExpandParent(expand bool) *Tree {
	t.SetProp("auto-expand-parent", expand)
	return t
}

// DefaultExpandedKeys sets default expanded keys
func (t *Tree) DefaultExpandedKeys(keys []interface{}) *Tree {
	t.SetProp("default-expanded-keys", keys)
	return t
}

// ShowCheckbox shows checkbox
func (t *Tree) ShowCheckbox(show bool) *Tree {
	t.SetProp("show-checkbox", show)
	return t
}

// CheckStrictly enables strict checking
func (t *Tree) CheckStrictly(strict bool) *Tree {
	t.SetProp("check-strictly", strict)
	return t
}

// DefaultCheckedKeys sets default checked keys
func (t *Tree) DefaultCheckedKeys(keys []interface{}) *Tree {
	t.SetProp("default-checked-keys", keys)
	return t
}

// CurrentNodeKey sets current node key
func (t *Tree) CurrentNodeKey(key interface{}) *Tree {
	t.SetProp("current-node-key", key)
	return t
}

// FilterNodeMethod sets filter node method
func (t *Tree) FilterNodeMethod(method string) *Tree {
	t.SetProp("filter-node-method", method)
	return t
}

// Accordion enables accordion mode
func (t *Tree) Accordion(accordion bool) *Tree {
	t.SetProp("accordion", accordion)
	return t
}

// Indent sets indent size
func (t *Tree) Indent(indent int) *Tree {
	t.SetProp("indent", indent)
	return t
}

// IconClass sets icon class
func (t *Tree) IconClass(class string) *Tree {
	t.SetProp("icon-class", class)
	return t
}

// Lazy enables lazy loading
func (t *Tree) Lazy(lazy bool) *Tree {
	t.SetProp("lazy", lazy)
	return t
}

// Draggable enables dragging
func (t *Tree) Draggable(draggable bool) *Tree {
	t.SetProp("draggable", draggable)
	return t
}

// AllowDrag sets allow drag function
func (t *Tree) AllowDrag(allowFunc string) *Tree {
	t.SetProp("allow-drag", allowFunc)
	return t
}

// AllowDrop sets allow drop function
func (t *Tree) AllowDrop(allowFunc string) *Tree {
	t.SetProp("allow-drop", allowFunc)
	return t
}

// Required makes the tree required
func (t *Tree) Required() contracts.FormComponent {
	t.AddValidateRule(rules.NewRequiredRule(fmt.Sprintf("%s 是必填项", t.title)))
	return t
}

// Placeholder sets placeholder (not applicable for tree, but required by interface)
func (t *Tree) Placeholder(text string) contracts.FormComponent {
	return t
}

// Disabled sets the disabled state
func (t *Tree) Disabled(disabled bool) contracts.FormComponent {
	t.SetProp("disabled", disabled)
	return t
}

// Build returns the tree component as a map
func (t *Tree) Build() map[string]interface{} {
	result := t.BaseComponent.Build()
	result["type"] = "el-tree"

	// Add data to props
	if len(t.data) > 0 {
		result["data"] = t.data
	}

	return result
}

// Button component implementation
type Button struct {
	*BaseComponent
}

// NewButton creates a new button component
func NewButton(field, title string) *Button {
	return &Button{
		BaseComponent: NewBaseComponent(field, title),
	}
}

// Size sets the button size
func (b *Button) Size(size string) *Button {
	b.SetProp("size", size)
	return b
}

// Type sets the button type
func (b *Button) Type(buttonType string) *Button {
	b.SetProp("type", buttonType)
	return b
}

// Plain sets plain style
func (b *Button) Plain(plain bool) *Button {
	b.SetProp("plain", plain)
	return b
}

// Round sets round style
func (b *Button) Round(round bool) *Button {
	b.SetProp("round", round)
	return b
}

// Circle sets circle style
func (b *Button) Circle(circle bool) *Button {
	b.SetProp("circle", circle)
	return b
}

// Loading sets loading state
func (b *Button) Loading(loading bool) *Button {
	b.SetProp("loading", loading)
	return b
}

// Icon sets button icon
func (b *Button) Icon(icon string) *Button {
	b.SetProp("icon", icon)
	return b
}

// Autofocus sets autofocus
func (b *Button) Autofocus(autofocus bool) *Button {
	b.SetProp("autofocus", autofocus)
	return b
}

// NativeType sets native type
func (b *Button) NativeType(nativeType string) *Button {
	b.SetProp("native-type", nativeType)
	return b
}

// Required makes the button required (not typically applicable)
func (b *Button) Required() contracts.FormComponent {
	return b
}

// Placeholder sets placeholder (not applicable for button, but required by interface)
func (b *Button) Placeholder(text string) contracts.FormComponent {
	return b
}

// Disabled sets the disabled state
func (b *Button) Disabled(disabled bool) contracts.FormComponent {
	b.SetProp("disabled", disabled)
	return b
}

// Build returns the button component as a map
func (b *Button) Build() map[string]interface{} {
	result := b.BaseComponent.Build()
	result["type"] = "el-button"
	return result
}

// Hidden component implementation
type Hidden struct {
	*BaseComponent
}

// NewHidden creates a new hidden component
func NewHidden(field string) *Hidden {
	return &Hidden{
		BaseComponent: NewBaseComponent(field, ""),
	}
}

// Required makes the hidden field required (typically not needed)
func (h *Hidden) Required() contracts.FormComponent {
	return h
}

// Placeholder sets placeholder (not applicable for hidden, but required by interface)
func (h *Hidden) Placeholder(text string) contracts.FormComponent {
	return h
}

// Disabled sets the disabled state (not applicable for hidden)
func (h *Hidden) Disabled(disabled bool) contracts.FormComponent {
	return h
}

// Build returns the hidden component as a map
func (h *Hidden) Build() map[string]interface{} {
	result := h.BaseComponent.Build()
	result["type"] = "hidden"

	// Hidden fields typically don't show title
	delete(result, "title")

	return result
}
