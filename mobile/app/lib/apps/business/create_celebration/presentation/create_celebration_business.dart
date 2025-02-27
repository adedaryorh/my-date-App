import 'package:flutter/material.dart';

class CreateCelebrationBusiness extends StatefulWidget {
  const CreateCelebrationBusiness({super.key});

  @override
  State<CreateCelebrationBusiness> createState() =>
      _CreateCelebrationBusinessState();
}

class _CreateCelebrationBusinessState extends State<CreateCelebrationBusiness> {
  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Text('Create Celebration'),
      ),
    );
  }
}
