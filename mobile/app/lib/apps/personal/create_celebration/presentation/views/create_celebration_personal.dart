import 'package:flutter/material.dart';

class CreateCelebrationPersonal extends StatefulWidget {
  const CreateCelebrationPersonal({super.key});

  @override
  State<CreateCelebrationPersonal> createState() =>
      _CreateCelebrationPersonalState();
}

class _CreateCelebrationPersonalState extends State<CreateCelebrationPersonal> {
  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Text('Create Celebration'),
      ),
    );
  }
}
