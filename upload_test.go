package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// upload_test.go Upload组件完整测试

// TestUploadCreation 测试Upload组件创建
func TestUploadCreation(t *testing.T) {
	t.Run("BasicCreation", func(t *testing.T) {
		upload := NewUpload("avatar", "头像")

		assert.Equal(t, "avatar", upload.GetField())
		assert.Equal(t, "upload", upload.GetType())

		data := upload.GetData()
		assert.Equal(t, "avatar", data.Field)
		assert.Equal(t, "头像", data.Title)
		assert.NotNil(t, data.Props)
	})

	t.Run("CreationWithValue", func(t *testing.T) {
		fileList := []interface{}{
			map[string]interface{}{"name": "file1.jpg", "url": "http://example.com/file1.jpg"},
		}
		upload := NewUpload("files", "文件", fileList)

		data := upload.GetData()
		assert.NotNil(t, data.Value)
	})
}

// TestUploadProperties 测试所有属性方法
func TestUploadProperties(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Upload) *Upload
		propKey       string
		expectedValue interface{}
	}{
		{
			name:          "Action",
			setup:         func(u *Upload) *Upload { return u.Action("/api/upload") },
			propKey:       "action",
			expectedValue: "/api/upload",
		},
		{
			name:          "Name",
			setup:         func(u *Upload) *Upload { return u.Name("file") },
			propKey:       "name",
			expectedValue: "file",
		},
		{
			name:          "WithCredentials",
			setup:         func(u *Upload) *Upload { return u.WithCredentials(true) },
			propKey:       "with-credentials",
			expectedValue: true,
		},
		{
			name:          "Multiple",
			setup:         func(u *Upload) *Upload { return u.Multiple(true) },
			propKey:       "multiple",
			expectedValue: true,
		},
		{
			name:          "Accept",
			setup:         func(u *Upload) *Upload { return u.Accept("image/*") },
			propKey:       "accept",
			expectedValue: "image/*",
		},
		{
			name:          "Limit",
			setup:         func(u *Upload) *Upload { return u.Limit(5) },
			propKey:       "limit",
			expectedValue: 5,
		},
		{
			name:          "Drag",
			setup:         func(u *Upload) *Upload { return u.Drag(true) },
			propKey:       "drag",
			expectedValue: true,
		},
		{
			name:          "ListType",
			setup:         func(u *Upload) *Upload { return u.ListType("picture-card") },
			propKey:       "list-type",
			expectedValue: "picture-card",
		},
		{
			name:          "Disabled",
			setup:         func(u *Upload) *Upload { return u.Disabled(true) },
			propKey:       "disabled",
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upload := NewUpload("test", "测试")
			upload = tt.setup(upload)

			data := upload.GetData()
			assert.Equal(t, tt.expectedValue, data.Props[tt.propKey])
		})
	}
}

// TestUploadChaining 测试链式调用
func TestUploadChaining(t *testing.T) {
	t.Run("CompleteChain", func(t *testing.T) {
		headers := map[string]string{
			"Authorization": "Bearer token123",
			"X-Custom":      "value",
		}

		extraData := map[string]interface{}{
			"userId": 123,
			"type":   "avatar",
		}

		upload := NewUpload("avatar", "头像上传").
			Action("/api/upload").
			Headers(headers).
			Data(extraData).
			Name("file").
			WithCredentials(true).
			Multiple(false).
			Accept("image/png,image/jpeg").
			Limit(1).
			ListType("picture-card").
			Required()

		data := upload.GetData()
		assert.Equal(t, "/api/upload", data.Props["action"])
		assert.Equal(t, headers, data.Props["headers"])
		assert.Equal(t, extraData, data.Props["data"])
		assert.Equal(t, "file", data.Props["name"])
		assert.Equal(t, true, data.Props["with-credentials"])
		assert.Equal(t, false, data.Props["multiple"])
		assert.Equal(t, "image/png,image/jpeg", data.Props["accept"])
		assert.Equal(t, 1, data.Props["limit"])
		assert.Equal(t, "picture-card", data.Props["list-type"])
		assert.NotEmpty(t, data.Validate)
	})

	t.Run("DragUploadChain", func(t *testing.T) {
		upload := NewUpload("files", "文件上传").
			Action("/api/upload").
			Drag(true).
			Multiple(true).
			Accept("*").
			Limit(10)

		data := upload.GetData()
		assert.Equal(t, true, data.Props["drag"])
		assert.Equal(t, true, data.Props["multiple"])
		assert.Equal(t, "*", data.Props["accept"])
		assert.Equal(t, 10, data.Props["limit"])
	})
}

// TestUploadBuild 测试Build方法
func TestUploadBuild(t *testing.T) {
	t.Run("BasicBuild", func(t *testing.T) {
		upload := NewUpload("file", "文件")

		result := upload.Build()

		assert.Equal(t, "upload", result["type"])
		assert.Equal(t, "file", result["field"])
		assert.Equal(t, "文件", result["title"])
	})

	t.Run("BuildWithAction", func(t *testing.T) {
		upload := NewUpload("file", "文件").
			Action("/api/upload")

		result := upload.Build()
		props, ok := result["props"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "/api/upload", props["action"])
	})

	t.Run("BuildWithHeaders", func(t *testing.T) {
		headers := map[string]string{
			"Authorization": "Bearer token",
		}
		upload := NewUpload("file", "文件").
			Headers(headers)

		result := upload.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, headers, props["headers"])
	})

	t.Run("BuildWithMultiple", func(t *testing.T) {
		upload := NewUpload("files", "多文件").
			Multiple(true).
			Limit(5)

		result := upload.Build()
		props := result["props"].(map[string]interface{})
		assert.Equal(t, true, props["multiple"])
		assert.Equal(t, 5, props["limit"])
	})

	t.Run("BuildWithListType", func(t *testing.T) {
		listTypes := []string{"text", "picture", "picture-card"}
		for _, listType := range listTypes {
			upload := NewUpload("file", "文件").
				ListType(listType)

			result := upload.Build()
			props := result["props"].(map[string]interface{})
			assert.Equal(t, listType, props["list-type"])
		}
	})
}

// TestUploadEdgeCases 测试边缘情况
func TestUploadEdgeCases(t *testing.T) {
	t.Run("AcceptAllFiles", func(t *testing.T) {
		upload := NewUpload("any_file", "任意文件").
			Accept("*")

		data := upload.GetData()
		assert.Equal(t, "*", data.Props["accept"])
	})

	t.Run("AcceptSpecificTypes", func(t *testing.T) {
		upload := NewUpload("images", "图片").
			Accept("image/png,image/jpeg,image/gif")

		data := upload.GetData()
		assert.Equal(t, "image/png,image/jpeg,image/gif", data.Props["accept"])
	})

	t.Run("ZeroLimit", func(t *testing.T) {
		upload := NewUpload("files", "文件").
			Limit(0)

		data := upload.GetData()
		assert.Equal(t, 0, data.Props["limit"])
	})

	t.Run("LargeLimit", func(t *testing.T) {
		upload := NewUpload("files", "文件").
			Limit(100)

		data := upload.GetData()
		assert.Equal(t, 100, data.Props["limit"])
	})

	t.Run("DisabledWithMultiple", func(t *testing.T) {
		upload := NewUpload("files", "文件").
			Multiple(true).
			Disabled(true)

		data := upload.GetData()
		assert.Equal(t, true, data.Props["multiple"])
		assert.Equal(t, true, data.Props["disabled"])
	})

	t.Run("DragWithoutMultiple", func(t *testing.T) {
		upload := NewUpload("file", "文件").
			Drag(true)

		data := upload.GetData()
		assert.Equal(t, true, data.Props["drag"])
		_, hasMultiple := data.Props["multiple"]
		assert.False(t, hasMultiple)
	})

	t.Run("EmptyHeaders", func(t *testing.T) {
		upload := NewUpload("file", "文件").
			Headers(map[string]string{})

		data := upload.GetData()
		headers, ok := data.Props["headers"].(map[string]string)
		require.True(t, ok)
		assert.Len(t, headers, 0)
	})

	t.Run("EmptyData", func(t *testing.T) {
		upload := NewUpload("file", "文件").
			Data(map[string]interface{}{})

		data := upload.GetData()
		extraData, ok := data.Props["data"].(map[string]interface{})
		require.True(t, ok)
		assert.Len(t, extraData, 0)
	})
}

// TestUploadWithValidation 测试验证功能
func TestUploadWithValidation(t *testing.T) {
	t.Run("RequiredValidation", func(t *testing.T) {
		upload := NewUpload("file", "文件").Required()

		data := upload.GetData()
		require.Len(t, data.Validate, 1)

		rule := data.Validate[0].ToMap()
		assert.Equal(t, true, rule["required"])
	})

	t.Run("CustomValidation", func(t *testing.T) {
		upload := NewUpload("file", "文件").
			Validate(CustomRule{
				Validator: "function(rule, value, callback) { if(value.length === 0) callback(new Error('请上传文件')); else callback(); }",
				Message:   "请上传文件",
			})

		data := upload.GetData()
		require.Len(t, data.Validate, 1)
	})
}

// TestUploadInForm 测试在表单中的使用
func TestUploadInForm(t *testing.T) {
	t.Run("FormWithUpload", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			NewUpload("avatar", "头像").
				Action("/api/upload/avatar").
				Accept("image/*").
				ListType("picture-card").
				Limit(1).
				Required(),
			NewUpload("files", "附件").
				Action("/api/upload/files").
				Multiple(true).
				Drag(true).
				Limit(5),
		}, nil)

		assert.NotNil(t, form)
		assert.Len(t, form.rules, 2)

		rules := form.FormRule()
		assert.Len(t, rules, 2)
		assert.Equal(t, "upload", rules[0]["type"])
		assert.Equal(t, "upload", rules[1]["type"])
	})

	t.Run("UploadWithControl", func(t *testing.T) {
		upload := NewUpload("file_type", "文件类型", "image").
			Control([]ControlRule{
				{
					Value: "image",
					Rule: []Component{
						NewUpload("image_file", "图片文件").
							Accept("image/*").
							Required(),
					},
				},
				{
					Value: "document",
					Rule: []Component{
						NewUpload("doc_file", "文档文件").
							Accept(".pdf,.doc,.docx").
							Required(),
					},
				},
			})

		data := upload.GetData()
		assert.Len(t, data.Control, 2)
		assert.Len(t, data.Control[0].Rule, 1)
		assert.Len(t, data.Control[1].Rule, 1)
	})
}

// TestUploadWithHeaders 测试Headers详细场景
func TestUploadWithHeaders(t *testing.T) {
	t.Run("AuthorizationHeader", func(t *testing.T) {
		upload := NewUpload("file", "文件").
			Headers(map[string]string{
				"Authorization": "Bearer token123",
			})

		data := upload.GetData()
		headers := data.Props["headers"].(map[string]string)
		assert.Equal(t, "Bearer token123", headers["Authorization"])
	})

	t.Run("MultipleHeaders", func(t *testing.T) {
		upload := NewUpload("file", "文件").
			Headers(map[string]string{
				"Authorization": "Bearer token",
				"X-API-Key":     "key123",
				"X-Request-ID":  "req-001",
			})

		data := upload.GetData()
		headers := data.Props["headers"].(map[string]string)
		assert.Len(t, headers, 3)
		assert.Equal(t, "Bearer token", headers["Authorization"])
		assert.Equal(t, "key123", headers["X-API-Key"])
		assert.Equal(t, "req-001", headers["X-Request-ID"])
	})
}

// TestUploadWithData 测试Data详细场景
func TestUploadWithData(t *testing.T) {
	t.Run("SimpleData", func(t *testing.T) {
		upload := NewUpload("file", "文件").
			Data(map[string]interface{}{
				"userId": 123,
			})

		data := upload.GetData()
		extraData := data.Props["data"].(map[string]interface{})
		assert.Equal(t, 123, extraData["userId"])
	})

	t.Run("ComplexData", func(t *testing.T) {
		upload := NewUpload("file", "文件").
			Data(map[string]interface{}{
				"userId":   123,
				"category": "avatar",
				"metadata": map[string]interface{}{
					"width":  200,
					"height": 200,
				},
			})

		data := upload.GetData()
		extraData := data.Props["data"].(map[string]interface{})
		assert.Equal(t, 123, extraData["userId"])
		assert.Equal(t, "avatar", extraData["category"])

		metadata, ok := extraData["metadata"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, 200, metadata["width"])
		assert.Equal(t, 200, metadata["height"])
	})
}

// BenchmarkUploadCreation 性能测试
func BenchmarkUploadCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewUpload("test", "测试")
	}
}

// BenchmarkUploadWithProperties 性能测试
func BenchmarkUploadWithProperties(b *testing.B) {
	headers := map[string]string{
		"Authorization": "Bearer token",
	}
	extraData := map[string]interface{}{
		"userId": 123,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewUpload("file", "文件").
			Action("/api/upload").
			Headers(headers).
			Data(extraData).
			Multiple(true).
			Drag(true).
			Accept("image/*").
			Limit(5).
			Required()
	}
}

// BenchmarkUploadBuild 性能测试
func BenchmarkUploadBuild(b *testing.B) {
	headers := map[string]string{
		"Authorization": "Bearer token",
	}

	upload := NewUpload("file", "文件").
		Action("/api/upload").
		Headers(headers).
		Multiple(true).
		Drag(true).
		Accept("image/*").
		Limit(5).
		Required()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = upload.Build()
	}
}
