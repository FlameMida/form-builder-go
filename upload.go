package formbuilder

// upload.go 实现Upload文件上传组件

// Upload 文件上传组件
type Upload struct {
	Builder[*Upload]
}

// NewUpload 创建文件上传组件
func NewUpload(field, title string, value ...interface{}) *Upload {
	upload := &Upload{}
	upload.data = &ComponentData{
		Field:    field,
		Title:    title,
		RuleType: "upload",
		Props:    make(map[string]interface{}),
	}
	if len(value) > 0 {
		upload.data.Value = value[0]
	}
	upload.inst = upload
	return upload
}

// Action 设置上传地址
func (u *Upload) Action(url string) *Upload {
	u.data.Props["action"] = url
	return u
}

// Headers 设置请求头
func (u *Upload) Headers(headers map[string]string) *Upload {
	u.data.Props["headers"] = headers
	return u
}

// Data 设置上传时附带的额外参数
func (u *Upload) Data(data map[string]interface{}) *Upload {
	u.data.Props["data"] = data
	return u
}

// Name 设置上传的文件字段名
func (u *Upload) Name(name string) *Upload {
	u.data.Props["name"] = name
	return u
}

// WithCredentials 设置是否支持发送cookie凭证
func (u *Upload) WithCredentials(enable bool) *Upload {
	u.data.Props["with-credentials"] = enable
	return u
}

// Multiple 设置是否支持多选文件
func (u *Upload) Multiple(enable bool) *Upload {
	u.data.Props["multiple"] = enable
	return u
}

// Accept 设置接受的文件类型
func (u *Upload) Accept(accept string) *Upload {
	u.data.Props["accept"] = accept
	return u
}

// Limit 设置最大上传文件数
func (u *Upload) Limit(limit int) *Upload {
	u.data.Props["limit"] = limit
	return u
}

// Drag 设置是否启用拖拽上传
func (u *Upload) Drag(enable bool) *Upload {
	u.data.Props["drag"] = enable
	return u
}

// ListType 设置文件列表类型（text/picture/picture-card）
func (u *Upload) ListType(listType string) *Upload {
	u.data.Props["list-type"] = listType
	return u
}

// Disabled 设置是否禁用
func (u *Upload) Disabled(disabled bool) *Upload {
	u.data.Props["disabled"] = disabled
	return u
}

// GetField 实现Component接口
func (u *Upload) GetField() string {
	return u.data.Field
}

// GetType 实现Component接口
func (u *Upload) GetType() string {
	return u.data.RuleType
}

// Build 实现Component接口
func (u *Upload) Build() map[string]interface{} {
	return buildComponent(u.data)
}
