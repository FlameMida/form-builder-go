package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// bootstrap_test.go Bootstrap引导类测试

// TestElmBootstrap 测试Element UI Bootstrap
func TestElmBootstrap(t *testing.T) {
	t.Run("Creation", func(t *testing.T) {
		bootstrap := NewElmBootstrap()

		assert.NotNil(t, bootstrap)
		assert.Len(t, bootstrap.GetScripts(), 3)
		assert.Len(t, bootstrap.GetStyles(), 1)
	})

	t.Run("GetScripts", func(t *testing.T) {
		bootstrap := NewElmBootstrap()
		scripts := bootstrap.GetScripts()

		require.Len(t, scripts, 3)
		assert.Contains(t, scripts[0], "vue")
		assert.Contains(t, scripts[1], "element-ui")
		assert.Contains(t, scripts[2], "form-create")
	})

	t.Run("GetStyles", func(t *testing.T) {
		bootstrap := NewElmBootstrap()
		styles := bootstrap.GetStyles()

		require.Len(t, styles, 1)
		assert.Contains(t, styles[0], "element-ui")
		assert.Contains(t, styles[0], "theme-chalk")
	})

	t.Run("SetScripts", func(t *testing.T) {
		bootstrap := NewElmBootstrap()
		customScripts := []string{
			"https://custom-cdn.com/vue.js",
			"https://custom-cdn.com/element-ui.js",
		}

		bootstrap.SetScripts(customScripts)

		scripts := bootstrap.GetScripts()
		assert.Equal(t, customScripts, scripts)
		assert.Len(t, scripts, 2)
	})

	t.Run("SetStyles", func(t *testing.T) {
		bootstrap := NewElmBootstrap()
		customStyles := []string{
			"https://custom-cdn.com/element-ui.css",
			"https://custom-cdn.com/custom.css",
		}

		bootstrap.SetStyles(customStyles)

		styles := bootstrap.GetStyles()
		assert.Equal(t, customStyles, styles)
		assert.Len(t, styles, 2)
	})

	t.Run("Init", func(t *testing.T) {
		bootstrap := NewElmBootstrap()
		form := NewElmForm("/api/submit", []Component{
			NewInput("name", "名称"),
		}, nil)

		// Form is already initialized with bootstrap scripts by NewElmForm
		// Init adds 3 scripts during form creation
		assert.Len(t, form.dependScript, 3)
		assert.Contains(t, form.dependScript[0], "vue")

		// Calling Init again appends the scripts again (duplicates)
		bootstrap.Init(form)
		assert.Len(t, form.dependScript, 6)
	})
}

// TestIviewBootstrap 测试iView Bootstrap
func TestIviewBootstrap(t *testing.T) {
	t.Run("CreationV3", func(t *testing.T) {
		bootstrap := NewIviewBootstrap(3)

		assert.NotNil(t, bootstrap)
		assert.Equal(t, 3, bootstrap.version)
		assert.NotEmpty(t, bootstrap.GetScripts())
		assert.NotEmpty(t, bootstrap.GetStyles())
	})

	t.Run("CreationV4", func(t *testing.T) {
		bootstrap := NewIviewBootstrap(4)

		assert.NotNil(t, bootstrap)
		assert.Equal(t, 4, bootstrap.version)
		assert.NotEmpty(t, bootstrap.GetScripts())
		assert.NotEmpty(t, bootstrap.GetStyles())
	})

	t.Run("GetScriptsV3", func(t *testing.T) {
		bootstrap := NewIviewBootstrap(3)
		scripts := bootstrap.GetScripts()

		require.NotEmpty(t, scripts)
		assert.Contains(t, scripts[0], "vue")
		// iView v3 scripts should be present
		foundIview := false
		for _, script := range scripts {
			if containsAny(script, []string{"iview", "view-design"}) {
				foundIview = true
				break
			}
		}
		assert.True(t, foundIview)
	})

	t.Run("GetScriptsV4", func(t *testing.T) {
		bootstrap := NewIviewBootstrap(4)
		scripts := bootstrap.GetScripts()

		require.NotEmpty(t, scripts)
		assert.Contains(t, scripts[0], "vue")
		// iView v4 (view-design) scripts should be present
		foundViewDesign := false
		for _, script := range scripts {
			if containsAny(script, []string{"view-design", "iview"}) {
				foundViewDesign = true
				break
			}
		}
		assert.True(t, foundViewDesign)
	})

	t.Run("GetStyles", func(t *testing.T) {
		bootstrap := NewIviewBootstrap(3)
		styles := bootstrap.GetStyles()

		require.NotEmpty(t, styles)
		foundIviewStyle := false
		for _, style := range styles {
			if containsAny(style, []string{"iview", "view-design"}) {
				foundIviewStyle = true
				break
			}
		}
		assert.True(t, foundIviewStyle)
	})

	t.Run("SetScripts", func(t *testing.T) {
		bootstrap := NewIviewBootstrap(3)
		customScripts := []string{
			"https://custom-cdn.com/vue.js",
			"https://custom-cdn.com/iview.js",
		}

		bootstrap.SetScripts(customScripts)

		scripts := bootstrap.GetScripts()
		assert.Equal(t, customScripts, scripts)
	})

	t.Run("SetStyles", func(t *testing.T) {
		bootstrap := NewIviewBootstrap(3)
		customStyles := []string{
			"https://custom-cdn.com/iview.css",
		}

		bootstrap.SetStyles(customStyles)

		styles := bootstrap.GetStyles()
		assert.Equal(t, customStyles, styles)
	})

	t.Run("GetVersion", func(t *testing.T) {
		bootstrapV3 := NewIviewBootstrap(3)
		assert.Equal(t, 3, bootstrapV3.GetVersion())

		bootstrapV4 := NewIviewBootstrap(4)
		assert.Equal(t, 4, bootstrapV4.GetVersion())
	})

	t.Run("Init", func(t *testing.T) {
		bootstrap := NewIviewBootstrap(3)
		form := NewIviewForm("/api/save", []Component{
			NewIviewInput("title", "标题"),
		}, nil)

		// Form is already initialized with bootstrap scripts by NewIviewForm
		assert.Greater(t, len(form.dependScript), 0)

		// Calling Init again appends more scripts
		initialLen := len(form.dependScript)
		bootstrap.Init(form)
		assert.Equal(t, initialLen+3, len(form.dependScript))
	})
}

// TestBootstrapCustomization 测试Bootstrap自定义配置
func TestBootstrapCustomization(t *testing.T) {
	t.Run("ElmBootstrapWithLocalResources", func(t *testing.T) {
		bootstrap := NewElmBootstrap()

		// Use local resources instead of CDN
		bootstrap.SetScripts([]string{
			"/static/js/vue.js",
			"/static/js/element-ui.js",
			"/static/js/form-create.js",
		})
		bootstrap.SetStyles([]string{
			"/static/css/element-ui.css",
		})

		scripts := bootstrap.GetScripts()
		assert.Len(t, scripts, 3)
		assert.Contains(t, scripts[0], "/static/js/vue.js")

		styles := bootstrap.GetStyles()
		assert.Len(t, styles, 1)
		assert.Contains(t, styles[0], "/static/css/element-ui.css")
	})

	t.Run("IviewBootstrapWithLocalResources", func(t *testing.T) {
		bootstrap := NewIviewBootstrap(4)

		bootstrap.SetScripts([]string{
			"/assets/vue.min.js",
			"/assets/view-design.min.js",
		})
		bootstrap.SetStyles([]string{
			"/assets/view-design.css",
		})

		scripts := bootstrap.GetScripts()
		assert.Len(t, scripts, 2)
		assert.Contains(t, scripts[0], "/assets/vue.min.js")

		styles := bootstrap.GetStyles()
		assert.Len(t, styles, 1)
		assert.Contains(t, styles[0], "/assets/view-design.css")
	})
}

// TestBootstrapWithForm 测试Bootstrap与Form集成
func TestBootstrapWithForm(t *testing.T) {
	t.Run("ElmFormWithBootstrap", func(t *testing.T) {
		bootstrap := NewElmBootstrap()
		form := NewElmForm("/submit", []Component{
			NewInput("username", "���户名").Required(),
			NewInput("email", "邮箱").Required(),
		}, nil)

		bootstrap.Init(form)

		// Verify scripts were added to form (form already has 3, now should have 6)
		assert.Len(t, form.dependScript, 6)

		// Form should still work normally
		rules := form.FormRule()
		assert.Len(t, rules, 2)
	})

	t.Run("IviewFormWithBootstrap", func(t *testing.T) {
		bootstrap := NewIviewBootstrap(3)
		form := NewIviewForm("/save", []Component{
			NewIviewInput("name", "名称").Required(),
		}, nil)

		bootstrap.Init(form)

		// Verify scripts were added
		assert.Greater(t, len(form.dependScript), 0)

		// Form should still work normally
		rules := form.FormRule()
		assert.Len(t, rules, 1)
	})
}

// Helper function for testing
func containsAny(s string, substrs []string) bool {
	for _, substr := range substrs {
		if contains(s, substr) {
			return true
		}
	}
	return false
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
