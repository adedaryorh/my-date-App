import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ForgetPasswordView extends StatefulWidget {
  const ForgetPasswordView({super.key});

  @override
  State<ForgetPasswordView> createState() => _ForgetPasswordViewState();
}

class _ForgetPasswordViewState extends State<ForgetPasswordView> {
  final TextEditingController numberCtr = TextEditingController();
  @override
  void dispose() {
    numberCtr.dispose();
    super.dispose();
  }

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
                'Forgot Password',
                style: context.textTheme.headlineLarge,
              ),
              const Space(20),
              const Text(
                  'Enter your registered phone number to reset password'),
              const Space(40),
              TextFormInput(
                autoValidateMode: AutovalidateMode.onUserInteraction,
                controller: numberCtr,
                validator: validatePhoneNumber,
                labelText: 'Enter registered phone number',
              ),
              const Space(50),
              MainButton(
                loading: false,
                text: 'Recover password',
                pressed: () {
                  context.pushNamed(AppRoute.verifyPhone.name);
                },
              ),
              const Space(30),
              TextButton(
                onPressed: () => AppRoute.signIn.name,
                child: Center(
                  child: RichText(
                    selectionColor: context.colorScheme.primary,
                    text: TextSpan(
                      text: 'Remembered Credentials? ',
                      style: context.textTheme.bodySmall,
                      children: <TextSpan>[
                        TextSpan(
                          text: 'Login',
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
