// Package components provides date/time and other complex components
package components

import (
	"fmt"
	"github.com/FlameMida/form-builder-go/contracts"
)

// Upload component implementation
type Upload struct {
	*BaseComponent
}

// NewUpload creates a new upload component
func NewUpload(field, title string) *Upload {
	return &Upload{
		BaseComponent: NewBaseComponent(field, title),
	}
}

// Action sets the upload action URL
func (u *Upload) Action(action string) *Upload {
	u.SetProp("action", action)
	return u
}

// Headers sets upload headers
func (u *Upload) Headers(headers map[string]string) *Upload {
	u.SetProp("headers", headers)
	return u
}

// Multiple enables multiple file upload
func (u *Upload) Multiple(multiple bool) *Upload {
	u.SetProp("multiple", multiple)
	return u
}

// Data sets additional data
func (u *Upload) Data(data map[string]interface{}) *Upload {
	u.SetProp("data", data)
	return u
}

// Name sets the file field name
func (u *Upload) Name(name string) *Upload {
	u.SetProp("name", name)
	return u
}

// WithCredentials enables credentials
func (u *Upload) WithCredentials(with bool) *Upload {
	u.SetProp("with-credentials", with)
	return u
}

// ShowFileList shows file list
func (u *Upload) ShowFileList(show bool) *Upload {
	u.SetProp("show-file-list", show)
	return u
}

// Drag enables drag upload
func (u *Upload) Drag(drag bool) *Upload {
	u.SetProp("drag", drag)
	return u
}

// Accept sets accepted file types
func (u *Upload) Accept(accept string) *Upload {
	u.SetProp("accept", accept)
	return u
}

// OnPreview sets preview handler
func (u *Upload) OnPreview(handler string) *Upload {
	u.SetProp("on-preview", handler)
	return u
}

// OnRemove sets remove handler
func (u *Upload) OnRemove(handler string) *Upload {
	u.SetProp("on-remove", handler)
	return u
}

// OnSuccess sets success handler
func (u *Upload) OnSuccess(handler string) *Upload {
	u.SetProp("on-success", handler)
	return u
}

// OnError sets error handler
func (u *Upload) OnError(handler string) *Upload {
	u.SetProp("on-error", handler)
	return u
}

// OnProgress sets progress handler
func (u *Upload) OnProgress(handler string) *Upload {
	u.SetProp("on-progress", handler)
	return u
}

// OnChange sets change handler
func (u *Upload) OnChange(handler string) *Upload {
	u.SetProp("on-change", handler)
	return u
}

// BeforeUpload sets before upload handler
func (u *Upload) BeforeUpload(handler string) *Upload {
	u.SetProp("before-upload", handler)
	return u
}

// BeforeRemove sets before remove handler
func (u *Upload) BeforeRemove(handler string) *Upload {
	u.SetProp("before-remove", handler)
	return u
}

// ListType sets list type
func (u *Upload) ListType(listType string) *Upload {
	u.SetProp("list-type", listType)
	return u
}

// AutoUpload enables auto upload
func (u *Upload) AutoUpload(auto bool) *Upload {
	u.SetProp("auto-upload", auto)
	return u
}

// FileList sets file list
func (u *Upload) FileList(files []interface{}) *Upload {
	u.SetProp("file-list", files)
	return u
}

// HttpRequest sets custom HTTP request
func (u *Upload) HttpRequest(request string) *Upload {
	u.SetProp("http-request", request)
	return u
}

// Limit sets file limit
func (u *Upload) Limit(limit int) *Upload {
	u.SetProp("limit", limit)
	return u
}

// OnExceed sets exceed handler
func (u *Upload) OnExceed(handler string) *Upload {
	u.SetProp("on-exceed", handler)
	return u
}

// Icon sets the button icon
func (u *Upload) Icon(icon string) *Upload {
	u.SetProp("icon", icon)
	return u
}

// Width sets the modal width
func (u *Upload) Width(width string) *Upload {
	u.SetProp("width", width)
	return u
}

// Height sets the modal height
func (u *Upload) Height(height string) *Upload {
	u.SetProp("height", height)
	return u
}

// MaxSize sets the file size limit in KB.
func (u *Upload) MaxSize(size int) *Upload {
	u.SetProp("max-size", size)
	return u
}

// Format sets the supported file formats.
func (u *Upload) Format(formats []string) *Upload {
	u.SetProp("format", formats)
	return u
}

// Required makes the upload required
func (u *Upload) Required() contracts.FormComponent {
	// Create conditional required rule based on upload limit
	rule := NewConditionalRequiredRule(
		fmt.Sprintf("请上传%s", u.title),
		u,
		func(component interface{}) string {
			upload := component.(*Upload)
			if limit, exists := upload.props["limit"]; exists {
				if val, ok := limit.(int); ok && val == 1 {
					return "string" // Single file upload uses string validation
				}
			}
			return "array" // Multiple file upload uses array validation
		},
	)
	u.AddValidateRule(rule)
	return u
}

// Placeholder sets placeholder (not applicable for upload, but required by interface)
func (u *Upload) Placeholder(text string) contracts.FormComponent {
	return u
}

// Disabled sets the disabled state
func (u *Upload) Disabled(disabled bool) contracts.FormComponent {
	u.SetProp("disabled", disabled)
	return u
}

// Build returns the upload component as a map
func (u *Upload) Build() map[string]interface{} {
	result := u.BaseComponent.Build()
	result["type"] = "el-upload"

	// Override value based on upload limit
	if u.value == nil {
		if limit, exists := u.props["limit"]; exists {
			if val, ok := limit.(int); ok && val == 1 {
				result["value"] = "" // Empty string for single file upload
			} else {
				result["value"] = []interface{}{} // Empty array for multiple file upload
			}
		} else {
			result["value"] = []interface{}{} // Default to empty array
		}
	}

	return result
}
