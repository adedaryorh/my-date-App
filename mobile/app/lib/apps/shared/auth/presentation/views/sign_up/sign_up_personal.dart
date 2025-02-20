import 'package:celebut/core/core.dart';
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class SignUpPersonal extends StatefulWidget {
  const SignUpPersonal({super.key});

  @override
  State<SignUpPersonal> createState() => _SignUpPersonalState();
}

class _SignUpPersonalState extends State<SignUpPersonal> {
  final TextEditingController nameCtr = TextEditingController();
  final TextEditingController usernameCtr = TextEditingController();
  final TextEditingController emailCtr = TextEditingController();
  final TextEditingController passWordCtr = TextEditingController();
  final TextEditingController secondPassWordCtr = TextEditingController();
  final _formKey = GlobalKey<FormState>();
  bool obscure = true;
  bool obscure1 = true;
  bool agreed = false;

  @override
  void dispose() {
    emailCtr.dispose();
    passWordCtr.dispose();
    nameCtr.dispose();
    usernameCtr.dispose();
    secondPassWordCtr.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Form(
        key: _formKey,
        child: ListView(
          padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 20),
          children: [
            const SizedBox(
              height: 60,
            ),
            Align(
              alignment: Alignment.centerLeft,
              child: Image.asset(
                AppAssets.celebutLogo,
                width: 105,
                height: 20,
              ),
            ),
            const SizedBox(
              height: 20,
            ),
            Text(
              'Sign up',
              style: context.textTheme.headlineLarge,
            ),
            const SizedBox(
              height: 20,
            ),
            const Text('Create an account to continue!'),
            const SizedBox(
              height: 20,
            ),
            Text(
              'Full Name',
              style: context.textTheme.bodyMedium
                  ?.copyWith(fontWeight: FontWeight.w600),
            ),
            const SizedBox(
              height: 5,
            ),
            TextFormInput(
              autoValidateMode: AutovalidateMode.onUserInteraction,
              controller: nameCtr,
              validator: validateName,
              labelText: 'Enter full name',
            ),
            const SizedBox(
              height: 10,
            ),
            Text(
              'Username',
              style: context.textTheme.bodyMedium
                  ?.copyWith(fontWeight: FontWeight.w600),
            ),
            const SizedBox(
              height: 5,
            ),
            TextFormInput(
              autoValidateMode: AutovalidateMode.onUserInteraction,
              controller: usernameCtr,
              validator: validateUsername,
              labelText: 'Enter username',
            ),
            const SizedBox(
              height: 10,
            ),
            Text(
              'Email',
              style: context.textTheme.bodyMedium
                  ?.copyWith(fontWeight: FontWeight.w600),
            ),
            const SizedBox(
              height: 5,
            ),
            TextFormInput(
              autoValidateMode: AutovalidateMode.onUserInteraction,
              controller: emailCtr,
              validator: validateEmailAddress,
              labelText: 'Enter email',
            ),
            const SizedBox(
              height: 10,
            ),
            Text(
              'Password',
              style: context.textTheme.bodyMedium
                  ?.copyWith(fontWeight: FontWeight.w600),
            ),
            const SizedBox(
              height: 5,
            ),
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
            const SizedBox(
              height: 10,
            ),
            Text(
              'Confirm Password',
              style: context.textTheme.bodyMedium
                  ?.copyWith(fontWeight: FontWeight.w600),
            ),
            const SizedBox(
              height: 5,
            ),
            TextFormInput(
              obscureText: obscure1,
              autoValidateMode: AutovalidateMode.onUserInteraction,
              controller: secondPassWordCtr,
              maxLines: 1,
              validator: validatePassword,
              labelText: '********',
              suffixIcon: IconButton(
                onPressed: () {
                  setState(() {
                    obscure1 = !obscure1;
                  });
                },
                icon: obscure1
                    ? const Icon(
                        Icons.visibility_off,
                      )
                    : const Icon(Icons.visibility),
                iconSize: 19,
              ),
            ),
            const SizedBox(
              height: 10,
            ),
            Row(
              children: [
                Transform.scale(
                  scale: 0.7,
                  child: Checkbox(
                    materialTapTargetSize: MaterialTapTargetSize.shrinkWrap,
                    value: agreed,
                    onChanged: (val) {
                      if (val == null) return;
                      setState(() {
                        agreed = val;
                      });
                    },
                  ),
                ),
                Expanded(
                  child: RichText(
                    selectionColor: context.colorScheme.primary,
                    text: TextSpan(
                      text: 'I agree to Celebut',
                      style: context.textTheme.bodySmall,
                      children: <TextSpan>[
                        TextSpan(
                          recognizer: TapGestureRecognizer()
                            ..onTap = () {
                              //print('Tap Here onTap');
                            },
                          text: '“Terms and conditions',
                          style: context.textTheme.bodySmall?.copyWith(
                            color: context.colorScheme.primary,
                          ),
                        ),
                        TextSpan(
                            text: ' and ', style: context.textTheme.bodySmall),
                        TextSpan(
                          recognizer: TapGestureRecognizer()
                            ..onTap = () {
                              //print('Tap Here onTap');
                            },
                          text: 'Privacy policy”',
                          style:
                              Theme.of(context).textTheme.bodySmall?.copyWith(
                                    color: context.colorScheme.primary,
                                  ),
                        ),
                      ],
                    ),
                  ),
                ),
              ],
            ),
            const SizedBox(
              height: 20,
            ),
            MainButton(
              loading: false,
              text: 'Register',
              pressed: () {},
            ),
            const SizedBox(
              height: 10,
            ),
            TextButton(
              onPressed: () => context.pushNamed(AppRoute.signIn.name),
              child: Center(
                child: RichText(
                  selectionColor: context.colorScheme.primary,
                  text: TextSpan(
                    text: "Don't have an account?",
                    style: context.textTheme.bodySmall,
                    children: <TextSpan>[
                      TextSpan(
                        text: ' Sign In',
                        style: Theme.of(context).textTheme.bodySmall?.copyWith(
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
    );
  }
}
