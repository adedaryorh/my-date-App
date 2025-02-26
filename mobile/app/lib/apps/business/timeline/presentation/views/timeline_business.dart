import 'package:flutter/material.dart';

class TimelineBusiness extends StatefulWidget {
  const TimelineBusiness({super.key});

  @override
  State<TimelineBusiness> createState() => _TimelineBusinessState();
}

class _TimelineBusinessState extends State<TimelineBusiness> {
  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Text('Timeline'),
      ),
    );
  }
}
