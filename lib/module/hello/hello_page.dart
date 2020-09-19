import 'package:flutter/material.dart';
import 'package:hellogoflutter/plugin/go/plugin.dart';

class HelloPage extends StatefulWidget {
  @override
  _HelloPageState createState() => _HelloPageState();
}

class _HelloPageState extends State<HelloPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Hello'),
      ),
      body: Center(
          child: Column(
        children: <Widget>[
          FutureBuilder<String>(
            future: HelloPlugin.hello(),
            builder: (c, snapshot) {
              if (!snapshot.hasData) {
                return Text('Hello插件执行出错');
              }
              return Text(snapshot.data);
            },
          )
        ],
      )),
    );
  }
}
