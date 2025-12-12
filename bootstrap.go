package formbuilder

// bootstrap.go 实现UI引导类
// 负责初始化不同UI框架的CDN资源

// ElmBootstrap Element UI引导类
type ElmBootstrap struct {
	scripts []string // JavaScript脚本列表
	styles  []string // CSS样式列表
}

// NewElmBootstrap 创建Element UI引导实例
func NewElmBootstrap() *ElmBootstrap {
	return &ElmBootstrap{
		scripts: []string{
			"https://unpkg.com/vue@2/dist/vue.js",
			"https://unpkg.com/element-ui/lib/index.js",
			"https://unpkg.com/@form-create/element-ui/dist/form-create.min.js",
		},
		styles: []string{
			"https://unpkg.com/element-ui/lib/theme-chalk/index.css",
		},
	}
}

// Init 初始化表单
func (b *ElmBootstrap) Init(form *Form) {
	form.dependScript = append(form.dependScript, b.scripts...)
}

// GetScripts 获取脚本列表
func (b *ElmBootstrap) GetScripts() []string {
	return b.scripts
}

// GetStyles 获取样式列表
func (b *ElmBootstrap) GetStyles() []string {
	return b.styles
}

// SetScripts 自定义脚本（用于切换CDN或使用本地资源）
func (b *ElmBootstrap) SetScripts(scripts []string) {
	b.scripts = scripts
}

// SetStyles 自定义样式
func (b *ElmBootstrap) SetStyles(styles []string) {
	b.styles = styles
}

// IviewBootstrap iView引导类
type IviewBootstrap struct {
	scripts []string
	styles  []string
	version int // 3 or 4
}

// NewIviewBootstrap 创建iView引导实例
// version: 3表示iView v3, 4表示iView v4 (view-design)
func NewIviewBootstrap(version int) *IviewBootstrap {
	bootstrap := &IviewBootstrap{
		version: version,
	}

	if version == 4 {
		// iView v4 (view-design)
		bootstrap.scripts = []string{
			"https://unpkg.com/vue@2/dist/vue.js",
			"https://unpkg.com/view-design/dist/iview.js",
			"https://unpkg.com/@form-create/iview4/dist/form-create.min.js",
		}
		bootstrap.styles = []string{
			"https://unpkg.com/view-design/dist/styles/iview.css",
		}
	} else {
		// iView v3
		bootstrap.scripts = []string{
			"https://unpkg.com/vue@2/dist/vue.js",
			"https://unpkg.com/iview@3/dist/iview.js",
			"https://unpkg.com/@form-create/iview/dist/form-create.min.js",
		}
		bootstrap.styles = []string{
			"https://unpkg.com/iview@3/dist/styles/iview.css",
		}
	}

	return bootstrap
}

// Init 初始化表单
func (b *IviewBootstrap) Init(form *Form) {
	form.dependScript = append(form.dependScript, b.scripts...)
}

// GetScripts 获取脚本列表
func (b *IviewBootstrap) GetScripts() []string {
	return b.scripts
}

// GetStyles 获取样式列表
func (b *IviewBootstrap) GetStyles() []string {
	return b.styles
}

// SetScripts 自定义脚本
func (b *IviewBootstrap) SetScripts(scripts []string) {
	b.scripts = scripts
}

// SetStyles 自定义样式
func (b *IviewBootstrap) SetStyles(styles []string) {
	b.styles = styles
}

// GetVersion 获取版本号
func (b *IviewBootstrap) GetVersion() int {
	return b.version
}
