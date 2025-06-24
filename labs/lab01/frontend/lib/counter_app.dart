import 'package:flutter/material.dart';

class CounterApp extends StatefulWidget {
  const CounterApp({Key? key}) : super(key: key);

  @override
  State<CounterApp> createState() => _CounterAppState();
}

class _CounterAppState extends State<CounterApp> {
  int _counter = 0;

  void _increment() {
    // TODO: Implement increment
    setState(() { // triggers a UI update when the state of a widget changes
      _counter++;
    });
  }

  void _decrement() {
    // TODO: Implement decrement
    setState(() {
      _counter--;
    });
  }

  void _reset() {
    // TODO: Implement reset
    setState(() {
      _counter = 0;
    });
  }

  @override
  Widget build(BuildContext context) {
    // TODO: Implement counter UI
    return Column( // using Column widget for placing children widgets vertically from top to bottom
      mainAxisSize: MainAxisSize.min, // sets the space set by the widget along the axis
      children: [
        Text(
          '$_counter',
          style: TextStyle(fontSize: 40),
        ),
        Row(
          mainAxisAlignment: MainAxisAlignment.center, // sets the vertical alignment
          children: [
            IconButton(
              onPressed: _decrement,
              icon: Icon(Icons.remove),
            ),
            IconButton(
              onPressed: _reset,
              icon: Icon(Icons.refresh),
            ),
            IconButton(
              onPressed: _increment,
              icon: Icon(Icons.add),
            ),
          ],
        ),
      ],
    );
  }
}
