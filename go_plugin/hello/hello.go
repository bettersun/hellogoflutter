package hello

import (
	"encoding/json"

	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// go-flutter插件需要声明包名和函数名
// dart代码中调用时需要指定相应的包名和函数名
const (
	channelName = "bettersun.go-flutter.plugin.hello"
	hello       = "hello"
	message     = "message"
)

// 声明插件结构体
type HelloPlugin struct{}

// 指定为go-flutter插件
var _ flutter.Plugin = &HelloPlugin{}

// 初始化插件
func (HelloPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc(hello, helloFunc)
	channel.HandleFunc(message, messageFunc)
	return nil
}

// 具体的逻辑处理函数，无参数传递
func helloFunc(arguments interface{}) (reply interface{}, err error) {

	return "Hello, go-flutter", nil
}

// 具体的逻辑处理函数，有参数传递
func messageFunc(arguments interface{}) (reply interface{}, err error) {

	var param string
	if arguments == nil {
		param = ""
	}

	switch arguments.(type) {
	case string:
		param = arguments.(string)
	default:
		newValue, _ := json.Marshal(arguments)
		param = string(newValue)
	}

	return "Welcome to go-flutter, " + param, nil
}
