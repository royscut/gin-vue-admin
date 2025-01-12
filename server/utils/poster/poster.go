package poster

import (
	"fmt"

	"github.com/hitailang/poster/core"
	"github.com/hitailang/poster/handler"
	"github.com/rs/xid"
)

func CreateMemberCardPoster(qr_value string, cardinfo_txt string) (filepath string, err error) {
	nullHandler := &handler.NullHandler{}
	ctx := &handler.Context{
		//图片都绘在这个PNG载体上
		PngCarrier: core.NewPNG(0, 0, 500, 1000),
	}
	// 绘制背景图
	backgroundHandler := &handler.BackgroundHandler{
		X:    0,                                      // 图片x坐标
		Y:    0,                                      // 图片y坐标
		Path: "./resource/assets/member_card_bg.png", //图片路径
	}
	// // 绘制圆形图像
	// imageCircleLocalHandler := &handler.ImageCircleLocalHandler{
	// 	X:    30, // 图片x坐标
	// 	Y:    50, // 图片y坐标
	// 	Path: "./resource/assets/reward.png",
	// }
	// 绘制本地图像
	// imageLocalHandler := &handler.ImageLocalHandler{
	// 	X:    30,                    // 图片x坐标
	// 	Y:    400,                   // 图片y坐标
	// 	Path: "./assets/reward.png", //图片路径
	// }

	// 绘制二维码
	qrCodeHandler := &handler.QRCodeHandler{
		X:   100,      // 二维码x坐标
		Y:   100,      // 二维码y坐标
		URL: qr_value, // 二维码跳转URL地址
	}
	// 绘制文字
	textHandler1 := &handler.TextHandler{
		Next:     handler.Next{},
		X:        20,  // 文字x坐标
		Y:        400, // 文字y坐标
		Size:     20,  // 文字大小
		R:        255, // 文字颜色RGB值
		G:        241,
		B:        250,
		Text:     cardinfo_txt,                 // 文字内容
		FontPath: "./resource/assets/msyh.ttf", // 字体文件
	}
	// // 绘制文字
	// textHandler2 := &handler.TextHandler{
	// 	Next:     handler.Next{},
	// 	X:        180,
	// 	Y:        150,
	// 	Size:     22,
	// 	R:        255,
	// 	G:        241,
	// 	B:        250,
	// 	Text:     "请随意赞赏~~",
	// 	FontPath: "./assets/msyh.ttf",
	// }
	// 结束绘制，把前面的内容合并成一张图片,输出到build目录
	out_file := "./resource/member_card_" + xid.New().String() + ".png"
	endHandler := &handler.EndHandler{
		Output: out_file,
	}

	// 链式调用绘制过程
	nullHandler.
		SetNext(backgroundHandler).
		//SetNext(imageCircleLocalHandler).
		SetNext(textHandler1).
		//SetNext(textHandler2).
		//SetNext(imageLocalHandler).
		SetNext(qrCodeHandler).
		SetNext(endHandler)

	// 开始执行业务
	if err = nullHandler.Run(ctx); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return "", err
	}
	// 成功
	fmt.Println("Success")
	return out_file, nil
}

func CreateMemberPoster(qr_value string, cardinfo_txt string) (filepath string, err error) {
	nullHandler := &handler.NullHandler{}
	ctx := &handler.Context{
		//图片都绘在这个PNG载体上
		PngCarrier: core.NewPNG(0, 0, 500, 1000),
	}
	// 绘制背景图
	backgroundHandler := &handler.BackgroundHandler{
		X:    0,                                 // 图片x坐标
		Y:    0,                                 // 图片y坐标
		Path: "./resource/assets/member_bg.png", //图片路径
	}
	// // 绘制圆形图像
	// imageCircleLocalHandler := &handler.ImageCircleLocalHandler{
	// 	X:    30, // 图片x坐标
	// 	Y:    50, // 图片y坐标
	// 	Path: "./resource/assets/reward.png",
	// }
	// 绘制本地图像
	// imageLocalHandler := &handler.ImageLocalHandler{
	// 	X:    30,                    // 图片x坐标
	// 	Y:    400,                   // 图片y坐标
	// 	Path: "./assets/reward.png", //图片路径
	// }

	// 绘制二维码
	qrCodeHandler := &handler.QRCodeHandler{
		X:   100,      // 二维码x坐标
		Y:   100,      // 二维码y坐标
		URL: qr_value, // 二维码跳转URL地址
	}
	// 绘制文字
	textHandler1 := &handler.TextHandler{
		Next:     handler.Next{},
		X:        20,  // 文字x坐标
		Y:        400, // 文字y坐标
		Size:     20,  // 文字大小
		R:        255, // 文字颜色RGB值
		G:        241,
		B:        250,
		Text:     cardinfo_txt,                 // 文字内容
		FontPath: "./resource/assets/msyh.ttf", // 字体文件
	}
	// // 绘制文字
	// textHandler2 := &handler.TextHandler{
	// 	Next:     handler.Next{},
	// 	X:        180,
	// 	Y:        150,
	// 	Size:     22,
	// 	R:        255,
	// 	G:        241,
	// 	B:        250,
	// 	Text:     "请随意赞赏~~",
	// 	FontPath: "./assets/msyh.ttf",
	// }
	// 结束绘制，把前面的内容合并成一张图片,输出到build目录
	out_file := "./resource/member_" + xid.New().String() + ".png"
	endHandler := &handler.EndHandler{
		Output: out_file,
	}

	// 链式调用绘制过程
	nullHandler.
		SetNext(backgroundHandler).
		//SetNext(imageCircleLocalHandler).
		SetNext(textHandler1).
		//SetNext(textHandler2).
		//SetNext(imageLocalHandler).
		SetNext(qrCodeHandler).
		SetNext(endHandler)

	// 开始执行业务
	if err := nullHandler.Run(ctx); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return "", err
	}
	// 成功
	fmt.Println("Success")
	return out_file, nil

}

func CreateMemberGiftCardPoster(qr_value string, cardinfo_txt string) (filepath string, err error) {
	nullHandler := &handler.NullHandler{}
	ctx := &handler.Context{
		//图片都绘在这个PNG载体上
		PngCarrier: core.NewPNG(0, 0, 500, 1000),
	}
	// 绘制背景图
	backgroundHandler := &handler.BackgroundHandler{
		X:    0,                                           // 图片x坐标
		Y:    0,                                           // 图片y坐标
		Path: "./resource/assets/member_gift_card_bg.png", //图片路径
	}
	// // 绘制圆形图像
	// imageCircleLocalHandler := &handler.ImageCircleLocalHandler{
	// 	X:    30, // 图片x坐标
	// 	Y:    50, // 图片y坐标
	// 	Path: "./resource/assets/reward.png",
	// }
	// 绘制本地图像
	// imageLocalHandler := &handler.ImageLocalHandler{
	// 	X:    30,                    // 图片x坐标
	// 	Y:    400,                   // 图片y坐标
	// 	Path: "./assets/reward.png", //图片路径
	// }

	// 绘制二维码
	qrCodeHandler := &handler.QRCodeHandler{
		X:   100,      // 二维码x坐标
		Y:   100,      // 二维码y坐标
		URL: qr_value, // 二维码跳转URL地址
	}
	// 绘制文字
	textHandler1 := &handler.TextHandler{
		Next:     handler.Next{},
		X:        20,  // 文字x坐标
		Y:        400, // 文字y坐标
		Size:     20,  // 文字大小
		R:        255, // 文字颜色RGB值
		G:        241,
		B:        250,
		Text:     cardinfo_txt,                 // 文字内容
		FontPath: "./resource/assets/msyh.ttf", // 字体文件
	}
	// // 绘制文字
	// textHandler2 := &handler.TextHandler{
	// 	Next:     handler.Next{},
	// 	X:        180,
	// 	Y:        150,
	// 	Size:     22,
	// 	R:        255,
	// 	G:        241,
	// 	B:        250,
	// 	Text:     "请随意赞赏~~",
	// 	FontPath: "./assets/msyh.ttf",
	// }
	// 结束绘制，把前面的内容合并成一张图片,输出到build目录
	out_file := "./resource/member_gift_card_" + xid.New().String() + ".png"
	endHandler := &handler.EndHandler{
		Output: out_file,
	}

	// 链式调用绘制过程
	nullHandler.
		SetNext(backgroundHandler).
		//SetNext(imageCircleLocalHandler).
		SetNext(textHandler1).
		//SetNext(textHandler2).
		//SetNext(imageLocalHandler).
		SetNext(qrCodeHandler).
		SetNext(endHandler)

	// 开始执行业务
	if err := nullHandler.Run(ctx); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return "", err
	}
	// 成功
	fmt.Println("Success")
	return out_file, nil

}
