import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class VerificationSuccessView extends StatefulWidget {
  const VerificationSuccessView({super.key});

  @override
  State<VerificationSuccessView> createState() =>
      _VerificationSuccessViewState();
}

class _VerificationSuccessViewState extends State<VerificationSuccessView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          children: [
            SizedBox(
              height: context.screenSize.height * 0.2,
            ),
            Image.asset(
              AppAssets.successCheck,
              height: 208,
              width: 208,
            ),
            const Space(50),
            Text(
              'Account Created',
              style: context.textTheme.headlineLarge,
            ),
            const Space(10),
            const Text('Verification has been done Successfully'),
            const Spacer(),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 20),
              child:
                  MainButton(loading: false, text: 'Continue', pressed: () {}),
            ),
            const Space(50),
          ],
        ),
      ),
    );
  }
}
