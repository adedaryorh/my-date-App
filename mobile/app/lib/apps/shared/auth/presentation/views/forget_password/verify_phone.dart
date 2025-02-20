import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:pinput/pinput.dart';

class VerifyPhoneView extends StatefulWidget {
  const VerifyPhoneView({super.key});

  @override
  State<VerifyPhoneView> createState() => _VerifyPhoneViewState();
}

class _VerifyPhoneViewState extends State<VerifyPhoneView> {
  final TextEditingController code = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Space(20),
              Image.asset(
                AppAssets.celebutLogo,
                width: 105,
                height: 20,
              ),
              const Space(20),
              Text(
                'Verification',
                style: context.textTheme.headlineLarge,
              ),
              const Space(20),
              const Text(
                'We have sent the verification code to your mobile number',
              ),
              const Space(40),
              Center(
                child: Pinput(
                  controller: code,
                  defaultPinTheme: buildPinTheme(context),
                  focusedPinTheme:
                      buildPinTheme(context, const Color(0xFF14202D)),
                  submittedPinTheme: buildPinTheme(context),
                  errorPinTheme:
                      buildPinTheme(context, const Color(0xffb42c3a)),
                ),
              ),
              const Space(20),
              MainButton(
                loading: false,
                text: 'Verify',
                pressed: () {
                  context.pushNamed(AppRoute.verificationSuccess.name);
                },
              ),
              const Space(30),
              TextButton(
                onPressed: () => AppRoute.signIn.name,
                child: Center(
                  child: RichText(
                    selectionColor: context.colorScheme.primary,
                    text: TextSpan(
                      text: 'Resend ',
                      style: context.textTheme.bodySmall,
                      children: <TextSpan>[
                        TextSpan(
                          text: 'code',
                          style: context.textTheme.bodySmall?.copyWith(
                            color: context.colorScheme.primary,
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
