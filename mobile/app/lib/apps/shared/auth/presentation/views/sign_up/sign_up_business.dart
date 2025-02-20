import 'package:flutter/material.dart';

class SignUpBusiness extends StatefulWidget {
  const SignUpBusiness({super.key});

  @override
  State<SignUpBusiness> createState() => _SignUpBusinessState();
}

class _SignUpBusinessState extends State<SignUpBusiness> {
  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Text('Business'),
      ),
    );
  }
}
