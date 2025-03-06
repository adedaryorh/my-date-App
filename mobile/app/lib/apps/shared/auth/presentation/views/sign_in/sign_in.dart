import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class SignInView extends StatefulWidget {
  const SignInView({super.key});

  @override
  State<SignInView> createState() => _SignInViewState();
}

class _SignInViewState extends State<SignInView> {
  final TextEditingController emailCtr = TextEditingController();
  final TextEditingController passWordCtr = TextEditingController();
  final _formKey = GlobalKey<FormState>();
  bool obscure = true;

  @override
  void dispose() {
    emailCtr.dispose();
    passWordCtr.dispose();
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
                'Sign in to your \nAccount',
                style: context.textTheme.headlineLarge,
              ),
              const Space(20),
              const Text('Enter your email and password to log in '),
              const Space(20),
              Form(
                key: _formKey,
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'Email',
                      style: context.textTheme.bodyMedium
                          ?.copyWith(fontWeight: FontWeight.w600),
                    ),
                    const Space(5),
                    TextFormInput(
                      autoValidateMode: AutovalidateMode.onUserInteraction,
                      controller: emailCtr,
                      validator: validateEmailAddress,
                      labelText: 'Enter email',
                    ),
                    const Space(10),
                    Text(
                      'Password',
                      style: context.textTheme.bodyMedium
                          ?.copyWith(fontWeight: FontWeight.w600),
                    ),
                    const Space(5),
                    TextFormInput(
                      obscureText: obscure,
                      autoValidateMode: AutovalidateMode.onUserInteraction,
                      controller: passWordCtr,
                      maxLines: 1,
                      validator: validatePassword,
                      labelText: '********',
                      suffixIcon: IconButton(
                        onPressed: () {
                          setState(() {
                            obscure = !obscure;
                          });
                        },
                        icon: obscure
                            ? const Icon(
                                Icons.visibility_off,
                              )
                            : const Icon(Icons.visibility),
                        iconSize: 19,
                      ),
                    ),
                    Align(
                      alignment: Alignment.centerRight,
                      child: TextButton(
                        onPressed: () {
                          context.pushNamed(AppRoute.forgetPassword.name);
                        },
                        child: Text(
                          'Forgot Password?',
                          style:
                              Theme.of(context).textTheme.bodySmall?.copyWith(
                                    fontWeight: FontWeight.w700,
                                    color: context.colorScheme.primary,
                                  ),
                        ),
                      ),
                    ),
                    const Space(20),
                    MainButton(
                      loading: false,
                      text: 'Log In',
                      pressed: () {
                        //context.goNamed(AppRoute.bCreateCelebration.name);
                        context.goNamed(AppRoute.pCreateCelebration.name);
                      },
                    ),
                  ],
                ),
              ),
              const Space(20),
              TextButton(
                onPressed: () => context.pushNamed(AppRoute.accountType.name),
                child: Center(
                  child: RichText(
                    selectionColor: context.colorScheme.primary,
                    text: TextSpan(
                      text: "Don't have an account?",
                      style: context.textTheme.bodySmall,
                      children: <TextSpan>[
                        TextSpan(
                          text: ' Sign Up',
                          style: context.textTheme.bodySmall?.copyWith(
                            fontWeight: FontWeight.w600,
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
