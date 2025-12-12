package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// tree_test.go Tree组件完整测试

// TestTreeCreation 测试Tree组件创建
func TestTreeCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		tree := NewTree("categories", "分类")

		assert.Equal(t, "categories", tree.GetField())
		assert.Equal(t, "tree", tree.GetType())

		data := tree.GetData()
		assert.Equal(t, "categories", data.Field)
		assert.Equal(t, "分类", data.Title)
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		tree := NewTree("selected", "已选", []interface{}{1, 2, 3})

		data := tree.GetData()
		assert.NotNil(t, data.Value)
	})
}

// TestTreeProperties 测试所有属性方法
func TestTreeProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Tree) *Tree
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "ShowCheckbox",
			setup:         func(tr *Tree) *Tree { return tr.ShowCheckbox(true) },
			propKey:       "show-checkbox",
			expectedValue: true,
		},
		{
			name:          "NodeKey",
			setup:         func(tr *Tree) *Tree { return tr.NodeKey("id") },
			propKey:       "node-key",
			expectedValue: "id",
		},
		{
			name:          "DefaultExpandAll",
			setup:         func(tr *Tree) *Tree { return tr.DefaultExpandAll(true) },
			propKey:       "default-expand-all",
			expectedValue: true,
		},
		{
			name:          "ExpandOnClickNode",
			setup:         func(tr *Tree) *Tree { return tr.ExpandOnClickNode(false) },
			propKey:       "expand-on-click-node",
			expectedValue: false,
		},
		{
			name:          "CheckOnClickNode",
			setup:         func(tr *Tree) *Tree { return tr.CheckOnClickNode(true) },
			propKey:       "check-on-click-node",
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewTree("test", "测试")
			tree = tt.setup(tree)

			data := tree.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestTreeChaining 测试链式调用
func TestTreeChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		treeData := []map[string]interface{}{
			{
				"id":    1,
				"label": "一级 1",
				"children": []map[string]interface{}{
					{"id": 2, "label": "二级 1-1"},
					{"id": 3, "label": "二级 1-2"},
				},
			},
			{
				"id":    4,
				"label": "一级 2",
				"children": []map[string]interface{}{
					{"id": 5, "label": "二级 2-1"},
				},
			},
		}

		tree := NewTree("permissions", "权限").
			Data(treeData).
			ShowCheckbox(true).
			NodeKey("id").
			DefaultExpandAll(true).
			Required()

		data := tree.GetData()
		assert.Equal(t, true, data.Props["show-checkbox"])
		assert.Equal(t, "id", data.Props["node-key"])
		assert.Equal(t, true, data.Props["default-expand-all"])
		assert.NotEmpty(t, data.Validate)
		assert.NotNil(t, data.Props["data"])
	})

	t.Run("TreePropsChain", func(t *testing.T) {
		props := map[string]interface{}{
			"label":    "name",
			"children": "subNodes",
			"disabled": "isDisabled",
		}

		tree := NewTree("tree", "树形").
			TreeProps(props).
			ExpandOnClickNode(false).
			CheckOnClickNode(true)

		data := tree.GetData()
		treeProps, ok := data.Props["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "name", treeProps["label"])
		assert.Equal(t, "subNodes", treeProps["children"])
		assert.Equal(t, false, data.Props["expand-on-click-node"])
		assert.Equal(t, true, data.Props["check-on-click-node"])
	})
}

// TestTreeBuild 测试Build方法
func TestTreeBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		tree := NewTree("menu", "菜单")

		result := tree.Build()

		assert.Equal(t, "tree", result["type"])
		assert.Equal(t, "menu", result["field"])
		assert.Equal(t, "菜单", result["title"])
	})

	t.Run("BuildWithData", func(t *testing.T) {
		treeData := []map[string]interface{}{
			{"id": 1, "label": "节点1"},
			{"id": 2, "label": "节点2"},
		}

		tree := NewTree("nodes", "节点").
			Data(treeData)

		result := tree.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, treeData, props["data"])
	})

	t.Run("BuildWithCheckbox", func(t *testing.T) {
		tree := NewTree("selected", "已选").
			ShowCheckbox(true).
			NodeKey("id")

		result := tree.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, true, props["show-checkbox"])
		assert.Equal(t, "id", props["node-key"])
	})

	t.Run("BuildWithTreeProps", func(t *testing.T) {
		customProps := map[string]interface{}{
			"label": "title",
			"value": "key",
		}

		tree := NewTree("custom", "自定义").
			TreeProps(customProps)

		result := tree.Build()
		props := result["props"].(map[string]interface{})
		treeProps, ok := props["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "title", treeProps["label"])
		assert.Equal(t, "key", treeProps["value"])
	})
}

// TestTreeEdgeCases 测试边缘情况
func TestTreeEdgeCases(t *testing.T) {
	t.Run("EmptyTreeData", func(t *testing.T) {
		tree := NewTree("empty", "空树").
			Data([]map[string]interface{}{})

		result := tree.Build()
		props := result["props"].(map[string]interface{})
		data, ok := props["data"].([]map[string]interface{})
		require.True(t, ok)
		assert.Len(t, data, 0)
	})

	t.Run("DeepNesting", func(t *testing.T) {
		treeData := []map[string]interface{}{
			{
				"id":    1,
				"label": "Level 1",
				"children": []map[string]interface{}{
					{
						"id":    2,
						"label": "Level 2",
						"children": []map[string]interface{}{
							{
								"id":    3,
								"label": "Level 3",
								"children": []map[string]interface{}{
									{"id": 4, "label": "Level 4"},
								},
							},
						},
					},
				},
			},
		}

		tree := NewTree("deep", "深层嵌套").
			Data(treeData).
			DefaultExpandAll(true)

		data := tree.GetData()
		assert.NotNil(t, data.Props["data"])
		assert.Equal(t, true, data.Props["default-expand-all"])
	})

	t.Run("AllExpandCollapseCombinations", func(t *testing.T) {
		tree := NewTree("test", "测试").
			DefaultExpandAll(false).
			ExpandOnClickNode(true).
			CheckOnClickNode(false)

		data := tree.GetData()
		assert.Equal(t, false, data.Props["default-expand-all"])
		assert.Equal(t, true, data.Props["expand-on-click-node"])
		assert.Equal(t, false, data.Props["check-on-click-node"])
	})

	t.Run("CheckboxWithoutNodeKey", func(t *testing.T) {
		tree := NewTree("test", "测试").
			ShowCheckbox(true)

		data := tree.GetData()
		assert.Equal(t, true, data.Props["show-checkbox"])
		_, hasNodeKey := data.Props["node-key"]
		assert.False(t, hasNodeKey)
	})
}

// TestTreeWithValidation 测试验证功能
func TestTreeWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		tree := NewTree("permissions", "权限").Required()

		data := tree.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("CustomValidation", func(t *testing.T) {
		tree := NewTree("nodes", "节点").
			Validate(CustomRule{
				Validator: "function(rule, value, callback) { ... }",
				Message:   "请至少选择一个节点",
			})

		data := tree.GetData()
		require.Len(t, data.Validate, 1)
	})
}

// TestTreeInForm 测试在表单中的使用
func TestTreeInForm(t *testing.T) {
	t.Run("FormWithTree", func(t *testing.T) {
		menuData := []map[string]interface{}{
			{
				"id":    1,
				"label": "系统管理",
				"children": []map[string]interface{}{
					{"id": 2, "label": "用户管理"},
					{"id": 3, "label": "角色管理"},
				},
			},
		}

		form := NewElmForm("/submit", []Component{
			NewTree("permissions", "权限").
				Data(menuData).
				ShowCheckbox(true).
				NodeKey("id").
				Required(),
			NewTree("departments", "部门").
				DefaultExpandAll(true),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 2)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "tree", rules[0]["type"])
		assert.Equal(t, "tree", rules[1]["type"])
	})

	t.Run("TreeWithControl", func(t *testing.T) {
		tree := NewTree("enable_tree", "启用树形", []interface{}{}).
			Control([]ControlRule{
				{
					Value: []interface{}{1, 2},
					Rule: []Component{
						NewInput("tree_note", "备注").Required(),
					},
				},
			})

		data := tree.GetData()
		assert.Len(t, data.Control, 1)
		assert.Len(t, data.Control[0].Rule, 1)
	})
}

// BenchmarkTreeCreation 性能测试
func BenchmarkTreeCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewTree("test", "测试")
	}
}

// BenchmarkTreeWithProperties 性能测试
func BenchmarkTreeWithProperties(b *testing.B) {
	treeData := []map[string]interface{}{
		{"id": 1, "label": "节点1"},
		{"id": 2, "label": "节点2"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewTree("nodes", "节点").
			Data(treeData).
			ShowCheckbox(true).
			NodeKey("id").
			DefaultExpandAll(true).
			Required()
	}
}

// BenchmarkTreeBuild 性能测试
func BenchmarkTreeBuild(b *testing.B) {
	treeData := []map[string]interface{}{
		{
			"id":    1,
			"label": "一级",
			"children": []map[string]interface{}{
				{"id": 2, "label": "二级"},
			},
		},
	}

	tree := NewTree("menu", "菜单").
		Data(treeData).
		ShowCheckbox(true).
		NodeKey("id").
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = tree.Build()
	}
}
