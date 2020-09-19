package main

import (
	// go.mod中引入的路径
	"github.com/bettersun/go-flutter-plugin/hello"
	"github.com/go-flutter-desktop/go-flutter"
)

var options = []flutter.Option{
	// 设置窗口宽高
	flutter.WindowInitialDimensions(800, 600),

	// 添加插件
	flutter.AddPlugin(hello.HelloPlugin{}),
}
