import 'dart:async';

import 'package:flutter/services.dart';

class HelloPlugin {
  // go-flutter插件中的包名，两者必须一致
  static const _channel = const MethodChannel("bettersun.go-flutter.plugin.hello");

  // go-flutter插件中的函数名
  static Future<String> hello() async => _channel.invokeMethod("hello");
}
