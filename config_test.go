package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// config_test.go Config配置测试

// TestConfigCreation 测试Config创建
func TestConfigCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		config := NewConfig()

		assert.NotNil(t, config)
		// Config is created with default values
		result := config.ToMap()
		assert.NotNil(t, result)
	})
}

// TestConfigSubmitBtn 测试提交按钮配置
func TestConfigSubmitBtn(t *testing.T) {
	t.Run("SetSubmitBtnProps", func(t *testing.T) {
		config := NewConfig()
		config.SetSubmitBtnProps(map[string]interface{}{
			"type": "primary",
			"size": "large",
		})

		result := config.ToMap()
		submitBtn, ok := result["submitBtn"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "primary", submitBtn["type"])
		assert.Equal(t, "large", submitBtn["size"])
	})

	t.Run("SubmitBtnWithText", func(t *testing.T) {
		config := NewConfig()
		config.SubmitBtn(true, "保存")

		result := config.ToMap()
		submitBtn, ok := result["submitBtn"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, true, submitBtn["show"])
		assert.Equal(t, "保存", submitBtn["innerText"])
	})

	t.Run("HideSubmitBtn", func(t *testing.T) {
		config := NewConfig()
		config.SubmitBtn(false)

		result := config.ToMap()
		submitBtn, ok := result["submitBtn"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, false, submitBtn["show"])
	})
}

// TestConfigResetBtn 测试重置按钮配置
func TestConfigResetBtn(t *testing.T) {
	t.Run("SetResetBtnProps", func(t *testing.T) {
		config := NewConfig()
		config.SetResetBtnProps(map[string]interface{}{
			"type": "default",
			"size": "medium",
		})

		result := config.ToMap()
		resetBtn, ok := result["resetBtn"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "default", resetBtn["type"])
		assert.Equal(t, "medium", resetBtn["size"])
	})

	t.Run("ResetBtnWithText", func(t *testing.T) {
		config := NewConfig()
		config.ResetBtn(true, "清空")

		result := config.ToMap()
		resetBtn, ok := result["resetBtn"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, true, resetBtn["show"])
		assert.Equal(t, "清空", resetBtn["innerText"])
	})

	t.Run("HideResetBtn", func(t *testing.T) {
		config := NewConfig()
		config.ResetBtn(false)

		result := config.ToMap()
		resetBtn, ok := result["resetBtn"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, false, resetBtn["show"])
	})
}

// TestConfigFormStyle 测试表单样式配置
func TestConfigFormStyle(t *testing.T) {
	t.Run("FormStyle", func(t *testing.T) {
		config := NewConfig()
		config.FormStyle(map[string]interface{}{
			"width":      "600px",
			"margin":     "0 auto",
			"background": "#fff",
		})

		result := config.ToMap()
		style, ok := result["formStyle"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "600px", style["width"])
		assert.Equal(t, "0 auto", style["margin"])
		assert.Equal(t, "#fff", style["background"])
	})
}

// TestConfigRow 测试Row配置
func TestConfigRow(t *testing.T) {
	t.Run("Row", func(t *testing.T) {
		config := NewConfig()
		config.Row(map[string]interface{}{
			"gutter": 20,
			"type":   "flex",
		})

		result := config.ToMap()
		row, ok := result["row"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, 20, row["gutter"])
		assert.Equal(t, "flex", row["type"])
	})
}

// TestConfigInfo 测试Info配置
func TestConfigInfo(t *testing.T) {
	t.Run("Info", func(t *testing.T) {
		config := NewConfig()
		config.Info("warning", "提示信息", true)

		result := config.ToMap()
		info, ok := result["info"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "warning", info["type"])
		assert.Equal(t, "提示信息", info["title"])
		assert.Equal(t, true, info["show"])
	})

	t.Run("InfoHidden", func(t *testing.T) {
		config := NewConfig()
		config.Info("info", "Hidden Info", false)

		result := config.ToMap()
		info, ok := result["info"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, false, info["show"])
	})
}

// TestConfigGlobal 测试全局配置
func TestConfigGlobal(t *testing.T) {
	t.Run("Global", func(t *testing.T) {
		config := NewConfig()
		result := config.Global("upload.action", "/api/upload")

		assert.Equal(t, config, result)

		configMap := config.ToMap()
		global, ok := configMap["global"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "/api/upload", global["upload.action"])
	})

	t.Run("MultipleGlobalSettings", func(t *testing.T) {
		config := NewConfig()
		config.Global("form.labelWidth", "120px")
		config.Global("form.size", "medium")

		result := config.ToMap()
		global := result["global"].(map[string]interface{})
		assert.Equal(t, "120px", global["form.labelWidth"])
		assert.Equal(t, "medium", global["form.size"])
	})

	t.Run("SetGlobal", func(t *testing.T) {
		config := NewConfig()
		config.SetGlobal(map[string]interface{}{
			"upload.action":   "/api/file/upload",
			"form.labelWidth": "150px",
		})

		result := config.ToMap()
		global := result["global"].(map[string]interface{})

		assert.Equal(t, "/api/file/upload", global["upload.action"])
		assert.Equal(t, "150px", global["form.labelWidth"])
	})
}

// TestConfigToMap 测试ToMap方法
func TestConfigToMap(t *testing.T) {
	t.Run("BasicToMap", func(t *testing.T) {
		config := NewConfig()
		result := config.ToMap()

		assert.NotNil(t, result)
		// NewConfig creates default submitBtn and resetBtn
		assert.NotNil(t, result["submitBtn"])
		assert.NotNil(t, result["resetBtn"])
	})

	t.Run("ToMapWithAllSettings", func(t *testing.T) {
		config := NewConfig()
		config.SubmitBtn(true, "保存")
		config.ResetBtn(true, "重置")
		config.FormStyle(map[string]interface{}{
			"width": "100%",
		})

		result := config.ToMap()

		submitBtn := result["submitBtn"].(map[string]interface{})
		assert.Equal(t, true, submitBtn["show"])
		assert.Equal(t, "保存", submitBtn["innerText"])

		resetBtn := result["resetBtn"].(map[string]interface{})
		assert.Equal(t, true, resetBtn["show"])
		assert.Equal(t, "重置", resetBtn["innerText"])

		style := result["formStyle"].(map[string]interface{})
		assert.Equal(t, "100%", style["width"])
	})
}

// TestConfigChaining 测试链式调用
func TestConfigChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		config := NewConfig().
			SubmitBtn(true, "提交").
			ResetBtn(true, "取消").
			Global("form.labelWidth", "100px")

		result := config.ToMap()

		submitBtn := result["submitBtn"].(map[string]interface{})
		assert.Equal(t, "提交", submitBtn["innerText"])

		resetBtn := result["resetBtn"].(map[string]interface{})
		assert.Equal(t, "取消", resetBtn["innerText"])

		global := result["global"].(map[string]interface{})
		assert.Equal(t, "100px", global["form.labelWidth"])
	})
}

// TestConfigInForm 测试Config在Form中的使用
func TestConfigInForm(t *testing.T) {
	t.Run("FormWithConfig", func(t *testing.T) {
		config := NewConfig().
			SubmitBtn(true, "保存数据").
			ResetBtn(false)

		form := NewElmForm("/api/submit", []Component{
			NewInput("name", "名称").Required(),
		}, config)

		assert.NotNil(t, form)
		assert.NotNil(t, form.config)

		configMap := form.config.ToMap()
		submitBtn := configMap["submitBtn"].(map[string]interface{})
		assert.Equal(t, "保存数据", submitBtn["innerText"])

		resetBtn := configMap["resetBtn"].(map[string]interface{})
		assert.Equal(t, false, resetBtn["show"])
	})
}
