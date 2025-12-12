package formbuilder

// tree.go 实现Tree树形控件组件

// Tree 树形控件组件
type Tree struct {
	Builder[*Tree]
}

// NewTree 创建树形控件
func NewTree(field, title string, value ...interface{}) *Tree {
	tree := &Tree{}
	tree.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "tree",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		tree.data.Value = value[0]
	}
	tree.inst = tree
	return tree
}

// Data 设置树形数据
func (t *Tree) Data(data interface{}) *Tree {
	t.data.Props["data"] = data
	return t
}

// TreeProps 设置配置选项
func (t *Tree) TreeProps(props map[string]interface{}) *Tree {
	t.data.Props["props"] = props
	return t
}

// ShowCheckbox 设置是否显示复选框
func (t *Tree) ShowCheckbox(show bool) *Tree {
	t.data.Props["show-checkbox"] = show
	return t
}

// NodeKey 设置节点唯一标识
func (t *Tree) NodeKey(key string) *Tree {
	t.data.Props["node-key"] = key
	return t
}

// DefaultExpandAll 设置是否默认展开所有节点
func (t *Tree) DefaultExpandAll(expand bool) *Tree {
	t.data.Props["default-expand-all"] = expand
	return t
}

// ExpandOnClickNode 设置是否点击节点时展开
func (t *Tree) ExpandOnClickNode(expand bool) *Tree {
	t.data.Props["expand-on-click-node"] = expand
	return t
}

// CheckOnClickNode 设置是否点击节点时选中
func (t *Tree) CheckOnClickNode(check bool) *Tree {
	t.data.Props["check-on-click-node"] = check
	return t
}

// GetField 实现Component接口
func (t *Tree) GetField() string {
	return t.data.Field
}

// GetType 实现Component接口
func (t *Tree) GetType() string {
	return t.data.RuleType
}

// Build 实现Component接口
func (t *Tree) Build() map[string]interface{} {
	return buildComponent(t.data)
}
