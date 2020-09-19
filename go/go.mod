module hellogoflutter/go

go 1.13

require (
	// 非真实存在的github仓库，在下方被改写
	github.com/bettersun/go-flutter-plugin/hello v0.0.0

	github.com/go-flutter-desktop/go-flutter v0.42.0
	github.com/pkg/errors v0.9.1
)

// 使用本地目录改写上方的github仓库
replace github.com/bettersun/go-flutter-plugin/hello => ../go_plugin/hello
