package main

import (
	"fmt"
	"log"

	fb "github.com/maverick/form-builder-go/formbuilder"
)

// control.go - 条件显示（Control）示例

func main() {
	form := fb.Elm.CreateForm("/api/order", []fb.Component{
		// 配送方式选择
		fb.Elm.Radio("delivery_type", "配送方式", "express").
			SetOptions([]fb.Option{
				{Value: "express", Label: "快递配送"},
				{Value: "pickup", Label: "到店自提"},
			}).
			Control([]fb.ControlRule{
				// 选择快递配送时显示
				{
					Value: "express",
					Rule: []fb.Component{
						fb.Elm.Input("address", "收货地址").
							Placeholder("请输入详细地址").
							Required(),

						fb.Elm.Input("receiver", "收货人").
							Placeholder("请输入收货人姓名").
							Required(),

						fb.Elm.Input("phone", "联系电话").
							Placeholder("请输入联系电话").
							Required(),
					},
				},
				// 选择到店自提时显示
				{
					Value: "pickup",
					Rule: []fb.Component{
						fb.Elm.Select("store", "自提门店").
							SetOptions([]fb.Option{
								{Value: "store1", Label: "北京朝阳店"},
								{Value: "store2", Label: "北京海淀店"},
								{Value: "store3", Label: "上海浦东店"},
							}).
							Required(),

						fb.Elm.DatePicker("pickup_date", "自提日期").
							DateType("date").
							Placeholder("请选择自提日期").
							Required(),
					},
				},
			}),

		// 支付方式选择
		fb.Elm.Radio("payment_method", "支付方式", "online").
			SetOptions([]fb.Option{
				{Value: "online", Label: "在线支付"},
				{Value: "cod", Label: "货到付款"},
			}).
			Control([]fb.ControlRule{
				{
					Value: "online",
					Rule: []fb.Component{
						fb.Elm.Radio("payment_channel", "支付渠道", "alipay").
							SetOptions([]fb.Option{
								{Value: "alipay", Label: "支付宝"},
								{Value: "wechat", Label: "微信支付"},
								{Value: "credit", Label: "信用卡"},
							}),
					},
				},
			}),

		// 是否需要发票
		fb.Elm.Switch("need_invoice", "是否需要发票").
			ActiveValue(true).
			InactiveValue(false).
			Control([]fb.ControlRule{
				{
					Value: true,
					Rule: []fb.Component{
						fb.Elm.Radio("invoice_type", "发票类型", "personal").
							SetOptions([]fb.Option{
								{Value: "personal", Label: "个人"},
								{Value: "company", Label: "公司"},
							}).
							Control([]fb.ControlRule{
								{
									Value: "company",
									Rule: []fb.Component{
										fb.Elm.Input("company_name", "公司名称").
											Required(),
										fb.Elm.Input("tax_number", "税号").
											Required(),
									},
								},
							}),
					},
				},
			}),
	})

	jsonRule, err := form.ParseFormRule()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("带条件显示的复杂表单:")
	fmt.Println(jsonRule)
}
