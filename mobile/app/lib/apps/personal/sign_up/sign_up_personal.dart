import 'package:flutter/material.dart';

class SignUpPersonal extends StatefulWidget {
  const SignUpPersonal({super.key});

  @override
  State<SignUpPersonal> createState() => _SignUpPersonalState();
}

class _SignUpPersonalState extends State<SignUpPersonal> {
  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Text('Personal'),
      ),
    );
  }
}
