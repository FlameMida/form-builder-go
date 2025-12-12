package formbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// factory_test.go Factory方法完整测试

// TestElmFactoryMethods 测试Element UI工厂方法
func TestElmFactoryMethods(t *testing.T) {
	t.Run("ElmPassword", func(t *testing.T) {
		password := Elm.Password("pwd", "密码", "secret")

		assert.NotNil(t, password)
		assert.Equal(t, "pwd", password.GetField())
		assert.Equal(t, "input", password.GetType())
		data := password.GetData()
		assert.Equal(t, "password", data.Props["type"])
		assert.Equal(t, "secret", data.Value)
	})

	t.Run("ElmTextarea", func(t *testing.T) {
		textarea := Elm.Textarea("content", "内容", "sample text")

		assert.NotNil(t, textarea)
		assert.Equal(t, "content", textarea.GetField())
		data := textarea.GetData()
		assert.Equal(t, "textarea", data.Props["type"])
		assert.Equal(t, "sample text", data.Value)
	})

	t.Run("ElmRadio", func(t *testing.T) {
		options := []Option{
			{Value: "1", Label: "选项1"},
			{Value: "2", Label: "选项2"},
		}
		radio := Elm.Radio("gender", "性别", options)

		assert.NotNil(t, radio)
		assert.Equal(t, "gender", radio.GetField())
		assert.Equal(t, "radio", radio.GetType())
	})

	t.Run("ElmCheckbox", func(t *testing.T) {
		options := []Option{
			{Value: "1", Label: "选项1"},
			{Value: "2", Label: "选项2"},
		}
		checkbox := Elm.Checkbox("hobbies", "爱好", options)

		assert.NotNil(t, checkbox)
		assert.Equal(t, "hobbies", checkbox.GetField())
		assert.Equal(t, "checkbox", checkbox.GetType())
	})

	t.Run("ElmNumber", func(t *testing.T) {
		number := Elm.Number("age", "年龄", 18)

		assert.NotNil(t, number)
		assert.Equal(t, "age", number.GetField())
		assert.Equal(t, "inputNumber", number.GetType())
		data := number.GetData()
		assert.Equal(t, 18, data.Value)
	})

	t.Run("ElmDatePicker", func(t *testing.T) {
		datepicker := Elm.DatePicker("birthday", "生日")

		assert.NotNil(t, datepicker)
		assert.Equal(t, "birthday", datepicker.GetField())
		assert.Equal(t, "datePicker", datepicker.GetType())
	})

	t.Run("ElmTimePicker", func(t *testing.T) {
		timepicker := Elm.TimePicker("time", "时间")

		assert.NotNil(t, timepicker)
		assert.Equal(t, "time", timepicker.GetField())
		assert.Equal(t, "timePicker", timepicker.GetType())
	})

	t.Run("ElmSlider", func(t *testing.T) {
		slider := Elm.Slider("volume", "音量", 50)

		assert.NotNil(t, slider)
		assert.Equal(t, "volume", slider.GetField())
		assert.Equal(t, "slider", slider.GetType())
	})

	t.Run("ElmSwitch", func(t *testing.T) {
		sw := Elm.Switch("enabled", "启用", true)

		assert.NotNil(t, sw)
		assert.Equal(t, "enabled", sw.GetField())
		assert.Equal(t, "switch", sw.GetType())
	})

	t.Run("ElmUpload", func(t *testing.T) {
		upload := Elm.Upload("file", "文件")

		assert.NotNil(t, upload)
		assert.Equal(t, "file", upload.GetField())
		assert.Equal(t, "upload", upload.GetType())
	})

	t.Run("ElmCascader", func(t *testing.T) {
		cascader := Elm.Cascader("region", "地区")

		assert.NotNil(t, cascader)
		assert.Equal(t, "region", cascader.GetField())
		assert.Equal(t, "cascader", cascader.GetType())
	})

	t.Run("ElmTree", func(t *testing.T) {
		tree := Elm.Tree("permissions", "权限")

		assert.NotNil(t, tree)
		assert.Equal(t, "permissions", tree.GetField())
		assert.Equal(t, "tree", tree.GetType())
	})

	t.Run("ElmRate", func(t *testing.T) {
		rate := Elm.Rate("rating", "评分", 5)

		assert.NotNil(t, rate)
		assert.Equal(t, "rating", rate.GetField())
		assert.Equal(t, "rate", rate.GetType())
	})

	t.Run("ElmColorPicker", func(t *testing.T) {
		colorpicker := Elm.ColorPicker("color", "颜色", "#FF0000")

		assert.NotNil(t, colorpicker)
		assert.Equal(t, "color", colorpicker.GetField())
		assert.Equal(t, "colorPicker", colorpicker.GetType())
	})

	t.Run("ElmHidden", func(t *testing.T) {
		hidden := Elm.Hidden("token", "abc123")

		assert.NotNil(t, hidden)
		assert.Equal(t, "token", hidden.GetField())
		assert.Equal(t, "hidden", hidden.GetType())
	})

	t.Run("ElmConfig", func(t *testing.T) {
		config := Elm.Config()
		assert.NotNil(t, config)
	})

	t.Run("ElmOption", func(t *testing.T) {
		option := Elm.Option("1", "选项1")
		assert.Equal(t, "1", option.Value)
		assert.Equal(t, "选项1", option.Label)
	})
}

// TestIviewFactoryMethods 测试iView工厂方法
func TestIviewFactoryMethods(t *testing.T) {
	t.Run("IviewPassword", func(t *testing.T) {
		password := Iview.Password("pwd", "密码", "secret")

		assert.NotNil(t, password)
		assert.Equal(t, "pwd", password.GetField())
		assert.Equal(t, "input", password.GetType())
		data := password.GetData()
		assert.Equal(t, "password", data.Props["type"])
		assert.Equal(t, "secret", data.Value)
	})

	t.Run("IviewTextarea", func(t *testing.T) {
		textarea := Iview.Textarea("content", "内容", "text")

		assert.NotNil(t, textarea)
		assert.Equal(t, "content", textarea.GetField())
		data := textarea.GetData()
		assert.Equal(t, "textarea", data.Props["type"])
		assert.Equal(t, "text", data.Value)
	})

	t.Run("IviewSelect", func(t *testing.T) {
		options := []Option{{Value: "1", Label: "选项1"}}
		sel := Iview.Select("status", "状态", options)

		assert.NotNil(t, sel)
		assert.Equal(t, "status", sel.GetField())
		assert.Equal(t, "select", sel.GetType())
	})

	t.Run("IviewRadio", func(t *testing.T) {
		options := []Option{{Value: "1", Label: "选项1"}}
		radio := Iview.Radio("type", "类型", options)

		assert.NotNil(t, radio)
		assert.Equal(t, "type", radio.GetField())
		assert.Equal(t, "radio", radio.GetType())
	})

	t.Run("IviewCheckbox", func(t *testing.T) {
		options := []Option{{Value: "1", Label: "选项1"}}
		checkbox := Iview.Checkbox("tags", "标签", options)

		assert.NotNil(t, checkbox)
		assert.Equal(t, "tags", checkbox.GetField())
		assert.Equal(t, "checkbox", checkbox.GetType())
	})

	t.Run("IviewNumber", func(t *testing.T) {
		number := Iview.Number("count", "数量", 10)

		assert.NotNil(t, number)
		assert.Equal(t, "count", number.GetField())
		assert.Equal(t, "inputNumber", number.GetType())
	})

	t.Run("IviewDatePicker", func(t *testing.T) {
		datepicker := Iview.DatePicker("date", "日期")

		assert.NotNil(t, datepicker)
		assert.Equal(t, "date", datepicker.GetField())
		assert.Equal(t, "datePicker", datepicker.GetType())
	})

	t.Run("IviewTimePicker", func(t *testing.T) {
		timepicker := Iview.TimePicker("time", "时间")

		assert.NotNil(t, timepicker)
		assert.Equal(t, "time", timepicker.GetField())
		assert.Equal(t, "timePicker", timepicker.GetType())
	})

	t.Run("IviewSlider", func(t *testing.T) {
		slider := Iview.Slider("progress", "进度", 75)

		assert.NotNil(t, slider)
		assert.Equal(t, "progress", slider.GetField())
		assert.Equal(t, "slider", slider.GetType())
	})

	t.Run("IviewSwitch", func(t *testing.T) {
		sw := Iview.Switch("active", "激活", false)

		assert.NotNil(t, sw)
		assert.Equal(t, "active", sw.GetField())
		assert.Equal(t, "switch", sw.GetType())
	})

	t.Run("IviewUpload", func(t *testing.T) {
		upload := Iview.Upload("avatar", "头像")

		assert.NotNil(t, upload)
		assert.Equal(t, "avatar", upload.GetField())
		assert.Equal(t, "upload", upload.GetType())
	})

	t.Run("IviewCascader", func(t *testing.T) {
		cascader := Iview.Cascader("area", "地区")

		assert.NotNil(t, cascader)
		assert.Equal(t, "area", cascader.GetField())
		assert.Equal(t, "cascader", cascader.GetType())
	})

	t.Run("IviewTree", func(t *testing.T) {
		tree := Iview.Tree("menu", "菜单")

		assert.NotNil(t, tree)
		assert.Equal(t, "menu", tree.GetField())
		assert.Equal(t, "tree", tree.GetType())
	})

	t.Run("IviewRate", func(t *testing.T) {
		rate := Iview.Rate("score", "评分", 4)

		assert.NotNil(t, rate)
		assert.Equal(t, "score", rate.GetField())
		assert.Equal(t, "rate", rate.GetType())
	})

	t.Run("IviewColorPicker", func(t *testing.T) {
		colorpicker := Iview.ColorPicker("theme", "主题色", "#409EFF")

		assert.NotNil(t, colorpicker)
		assert.Equal(t, "theme", colorpicker.GetField())
		assert.Equal(t, "colorPicker", colorpicker.GetType())
	})

	t.Run("IviewHidden", func(t *testing.T) {
		hidden := Iview.Hidden("id", 123)

		assert.NotNil(t, hidden)
		assert.Equal(t, "id", hidden.GetField())
		assert.Equal(t, "hidden", hidden.GetType())
	})

	t.Run("IviewConfig", func(t *testing.T) {
		config := Iview.Config()
		assert.NotNil(t, config)
	})

	t.Run("IviewOption", func(t *testing.T) {
		option := Iview.Option("a", "选项A")
		assert.Equal(t, "a", option.Value)
		assert.Equal(t, "选项A", option.Label)
	})
}

// TestFactoryWithComponents 测试工厂方法创建的组件能正常使用
func TestFactoryWithComponents(t *testing.T) {
	t.Run("ElmFactoryInForm", func(t *testing.T) {
		form := NewElmForm("/submit", []Component{
			Elm.Password("password", "密码").Required(),
			Elm.Number("age", "年龄", 18),
			Elm.Slider("score", "评分", 50),
		}, nil)

		rules := form.FormRule()
		require.Len(t, rules, 3)
		assert.Equal(t, "input", rules[0]["type"])
		assert.Equal(t, "password", rules[0]["props"].(map[string]interface{})["type"])
		assert.Equal(t, "inputNumber", rules[1]["type"])
		assert.Equal(t, "slider", rules[2]["type"])
	})

	t.Run("IviewFactoryInForm", func(t *testing.T) {
		form := NewIviewForm("/save", []Component{
			Iview.Password("pwd", "密码").Required(),
			Iview.Number("count", "数量", 1),
		}, nil)

		rules := form.FormRule()
		require.Len(t, rules, 2)
		assert.Equal(t, "input", rules[0]["type"])
		assert.Equal(t, "inputNumber", rules[1]["type"])
	})
}
