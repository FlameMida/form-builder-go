package formbuilder

import (
	"encoding/json"
	"testing"
)

// TestFrameBasic 测试基础Frame组件
func TestFrameBasic(t *testing.T) {
	frame := Elm.Frame("file", "选择文件", "/api/select")

	rule := frame.Build()

	if rule["type"] != "frame" {
		t.Errorf("Expected type 'frame', got '%v'", rule["type"])
	}
	if rule["field"] != "file" {
		t.Errorf("Expected field 'file', got '%v'", rule["field"])
	}
	if rule["title"] != "选择文件" {
		t.Errorf("Expected title '选择文件', got '%v'", rule["title"])
	}

	props := rule["props"].(map[string]interface{})
	if props["src"] != "/api/select" {
		t.Errorf("Expected src '/api/select', got '%v'", props["src"])
	}
	if props["type"] != "input" {
		t.Errorf("Expected type 'input', got '%v'", props["type"])
	}
}

// TestFrameImage 测试单图片框架
func TestFrameImage(t *testing.T) {
	frame := Elm.FrameImage("avatar", "选择头像", "/api/select-image", "default.jpg")

	rule := frame.Build()

	if rule["type"] != "frame" {
		t.Errorf("Expected type 'frame', got '%v'", rule["type"])
	}
	if rule["field"] != "avatar" {
		t.Errorf("Expected field 'avatar', got '%v'", rule["field"])
	}
	if rule["value"] != "default.jpg" {
		t.Errorf("Expected value 'default.jpg', got '%v'", rule["value"])
	}

	props := rule["props"].(map[string]interface{})
	if props["type"] != "image" {
		t.Errorf("Expected type 'image', got '%v'", props["type"])
	}
	if props["maxLength"] != 1 {
		t.Errorf("Expected maxLength 1, got '%v'", props["maxLength"])
	}
}

// TestFrameImages 测试多图片框架
func TestFrameImages(t *testing.T) {
	defaultImages := []interface{}{"img1.jpg", "img2.jpg"}
	frame := Elm.FrameImages("images", "选择多张图片", "/api/select-images", defaultImages)

	rule := frame.Build()

	props := rule["props"].(map[string]interface{})
	if props["type"] != "image" {
		t.Errorf("Expected type 'image', got '%v'", props["type"])
	}
	if props["maxLength"] != 0 {
		t.Errorf("Expected maxLength 0 (unlimited), got '%v'", props["maxLength"])
	}
}

// TestFrameFiles 测试多文件框架
func TestFrameFiles(t *testing.T) {
	frame := Elm.FrameFiles("files", "选择多个文件", "/api/select-files")

	rule := frame.Build()

	props := rule["props"].(map[string]interface{})
	if props["type"] != "file" {
		t.Errorf("Expected type 'file', got '%v'", props["type"])
	}
	if props["maxLength"] != 0 {
		t.Errorf("Expected maxLength 0 (unlimited), got '%v'", props["maxLength"])
	}
}

// TestFrameChainedMethods 测试Frame的链式方法
func TestFrameChainedMethods(t *testing.T) {
	frame := Elm.FrameImage("banner", "选择Banner图", "/api/select").
		Height("600px").
		Width("90vw").
		Icon("el-icon-picture").
		FrameTitle("选择图片").
		Spin(true).
		HandleIcon(true).
		AllowRemove(true).
		Disabled(false)

	rule := frame.Build()
	props := rule["props"].(map[string]interface{})

	if props["height"] != "600px" {
		t.Errorf("Expected height '600px', got '%v'", props["height"])
	}
	if props["width"] != "90vw" {
		t.Errorf("Expected width '90vw', got '%v'", props["width"])
	}
	if props["icon"] != "el-icon-picture" {
		t.Errorf("Expected icon 'el-icon-picture', got '%v'", props["icon"])
	}
	if props["frameTitle"] != "选择图片" {
		t.Errorf("Expected frameTitle '选择图片', got '%v'", props["frameTitle"])
	}
}

// TestFrameJSON 测试Frame的JSON输出
func TestFrameJSON(t *testing.T) {
	frame := Elm.FrameImage("photo", "照片", "/select").
		Height("500px").
		Required()

	rule := frame.Build()
	jsonData, err := json.Marshal(rule)
	if err != nil {
		t.Fatalf("Failed to marshal frame to JSON: %v", err)
	}

	// 验证JSON可以正常序列化
	var decoded map[string]interface{}
	err = json.Unmarshal(jsonData, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal frame JSON: %v", err)
	}

	if decoded["type"] != "frame" {
		t.Errorf("Expected type 'frame' in JSON, got '%v'", decoded["type"])
	}
}

// TestUploadFile 测试单文件上传
func TestUploadFile(t *testing.T) {
	upload := Elm.UploadFile("document", "上传文档", "/api/upload", "doc.pdf")

	rule := upload.Build()

	if rule["type"] != "upload" {
		t.Errorf("Expected type 'upload', got '%v'", rule["type"])
	}
	if rule["value"] != "doc.pdf" {
		t.Errorf("Expected value 'doc.pdf', got '%v'", rule["value"])
	}

	props := rule["props"].(map[string]interface{})
	if props["action"] != "/api/upload" {
		t.Errorf("Expected action '/api/upload', got '%v'", props["action"])
	}
	if props["limit"] != 1 {
		t.Errorf("Expected limit 1, got '%v'", props["limit"])
	}
	if props["uploadType"] != "file" {
		t.Errorf("Expected uploadType 'file', got '%v'", props["uploadType"])
	}
}

// TestUploadImage 测试单图片上传
func TestUploadImage(t *testing.T) {
	upload := Elm.UploadImage("avatar", "上传头像", "/api/upload")

	rule := upload.Build()
	props := rule["props"].(map[string]interface{})

	if props["uploadType"] != "image" {
		t.Errorf("Expected uploadType 'image', got '%v'", props["uploadType"])
	}
	if props["accept"] != "image/*" {
		t.Errorf("Expected accept 'image/*', got '%v'", props["accept"])
	}
	if props["limit"] != 1 {
		t.Errorf("Expected limit 1, got '%v'", props["limit"])
	}
}

// TestUploadImages 测试多图片上传
func TestUploadImages(t *testing.T) {
	defaultImages := []interface{}{"img1.jpg", "img2.jpg"}
	upload := Elm.UploadImages("gallery", "上传相册", "/api/upload", defaultImages)

	rule := upload.Build()
	props := rule["props"].(map[string]interface{})

	if props["uploadType"] != "image" {
		t.Errorf("Expected uploadType 'image', got '%v'", props["uploadType"])
	}
	if props["accept"] != "image/*" {
		t.Errorf("Expected accept 'image/*', got '%v'", props["accept"])
	}
	// 多图上传不应该有limit限制
	if limit, exists := props["limit"]; exists && limit != float64(0) {
		t.Errorf("Expected no limit or limit 0 for multiple images, got '%v'", limit)
	}
}

// TestUploadFiles 测试多文件上传
func TestUploadFiles(t *testing.T) {
	upload := Elm.UploadFiles("attachments", "上传附件", "/api/upload")

	rule := upload.Build()
	props := rule["props"].(map[string]interface{})

	if props["uploadType"] != "file" {
		t.Errorf("Expected uploadType 'file', got '%v'", props["uploadType"])
	}
}

// TestCol 测试Col布局方法（整数）
func TestColInt(t *testing.T) {
	input := Elm.Input("username", "用户名").Col(12)

	rule := input.Build()

	col := rule["col"].(map[string]interface{})
	if col["span"] != 12 {
		t.Errorf("Expected col span 12, got '%v'", col["span"])
	}
}

// TestColMap 测试Col布局方法（map）
func TestColMap(t *testing.T) {
	input := Elm.Input("email", "邮箱").Col(map[string]interface{}{
		"span":   8,
		"offset": 2,
	})

	rule := input.Build()

	col := rule["col"].(map[string]interface{})
	if col["span"] != 8 {
		t.Errorf("Expected col span 8, got '%v'", col["span"])
	}
	if col["offset"] != 2 {
		t.Errorf("Expected col offset 2, got '%v'", col["offset"])
	}
}

// TestInputNumberPrecision 测试Precision方法
func TestInputNumberPrecision(t *testing.T) {
	num := Elm.Number("price", "价格").
		Min(0).
		Max(9999.99).
		Precision(2)

	rule := num.Build()
	props := rule["props"].(map[string]interface{})

	if props["precision"] != 2 {
		t.Errorf("Expected precision 2, got '%v'", props["precision"])
	}
	if props["min"] != float64(0) {
		t.Errorf("Expected min 0, got '%v'", props["min"])
	}
	if props["max"] != 9999.99 {
		t.Errorf("Expected max 9999.99, got '%v'", props["max"])
	}
}

// TestIviewFrameImage 测试iView版本的Frame
func TestIviewFrameImage(t *testing.T) {
	frame := Iview.FrameImage("photo", "照片", "/select")

	rule := frame.Build()

	if rule["type"] != "frame" {
		t.Errorf("Expected type 'frame', got '%v'", rule["type"])
	}

	props := rule["props"].(map[string]interface{})
	if props["type"] != "image" {
		t.Errorf("Expected type 'image', got '%v'", props["type"])
	}
}

// TestIviewUploadImage 测试iView版本的Upload
func TestIviewUploadImage(t *testing.T) {
	upload := Iview.UploadImage("avatar", "头像", "/upload")

	rule := upload.Build()

	if rule["type"] != "upload" {
		t.Errorf("Expected type 'upload', got '%v'", rule["type"])
	}

	props := rule["props"].(map[string]interface{})
	if props["uploadType"] != "image" {
		t.Errorf("Expected uploadType 'image', got '%v'", props["uploadType"])
	}
}
