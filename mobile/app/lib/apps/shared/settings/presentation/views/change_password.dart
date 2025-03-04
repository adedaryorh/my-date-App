import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ChangePassword extends StatefulWidget {
  const ChangePassword({super.key});

  @override
  State<ChangePassword> createState() => _ChangePasswordState();
}

class _ChangePasswordState extends State<ChangePassword> {
  bool obscureOld = true;
  bool obscureNew = true;
  bool obscureConfirm = true;
  final TextEditingController newPasswordCtr = TextEditingController();
  final TextEditingController oldPassWordCtr = TextEditingController();
  final TextEditingController confirmPassWordCtr = TextEditingController();
  final TextEditingController codeCtr = TextEditingController();
  final _formKey = GlobalKey<FormState>();

  @override
  void dispose() {
    oldPassWordCtr.dispose();
    newPasswordCtr.dispose();
    confirmPassWordCtr.dispose();
    codeCtr.dispose();
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
              const SizedBox(
                height: 50,
              ),
              InkWell(
                onTap: () {
                  context.pop();
                },
                child: const SizedBox(
                  height: 20,
                  width: 20,
                  child: Icon(
                    Icons.arrow_back_ios,
                    size: 13,
                  ),
                ),
              ),
              const SizedBox(
                height: 20,
              ),
              Text(
                'Change Password',
                style: context.textTheme.headlineMedium,
              ),
              const Text(
                  'Feeling worried about your account been easily preyed on? '
                  'Then change that password now!'),
              Form(
                key: _formKey,
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    const Space(20),
                    Text(
                      'Previous Password',
                      style: context.textTheme.bodyMedium
                          ?.copyWith(fontWeight: FontWeight.w600),
                    ),
                    const Space(5),
                    TextFormInput(
                      obscureText: obscureOld,
                      autoValidateMode: AutovalidateMode.onUserInteraction,
                      controller: oldPassWordCtr,
                      maxLines: 1,
                      validator: validatePassword,
                      labelText: '********',
                      suffixIcon: IconButton(
                        onPressed: () {
                          setState(() {
                            obscureOld = !obscureOld;
                          });
                        },
                        icon: obscureOld
                            ? const Icon(
                                Icons.visibility_off,
                              )
                            : const Icon(Icons.visibility),
                        iconSize: 19,
                      ),
                    ),
                    const Space(10),
                    Text(
                      'New Password',
                      style: context.textTheme.bodyMedium
                          ?.copyWith(fontWeight: FontWeight.w600),
                    ),
                    const Space(5),
                    TextFormInput(
                      obscureText: obscureNew,
                      autoValidateMode: AutovalidateMode.onUserInteraction,
                      controller: newPasswordCtr,
                      maxLines: 1,
                      validator: validatePassword,
                      labelText: '********',
                      suffixIcon: IconButton(
                        onPressed: () {
                          setState(() {
                            obscureNew = !obscureNew;
                          });
                        },
                        icon: obscureNew
                            ? const Icon(
                                Icons.visibility_off,
                              )
                            : const Icon(Icons.visibility),
                        iconSize: 19,
                      ),
                    ),
                    const Space(10),
                    Text(
                      'Confirm Password',
                      style: context.textTheme.bodyMedium
                          ?.copyWith(fontWeight: FontWeight.w600),
                    ),
                    const Space(5),
                    TextFormInput(
                      obscureText: obscureConfirm,
                      autoValidateMode: AutovalidateMode.onUserInteraction,
                      controller: confirmPassWordCtr,
                      maxLines: 1,
                      validator: validatePassword,
                      labelText: '********',
                      suffixIcon: IconButton(
                        onPressed: () {
                          setState(() {
                            obscureConfirm = !obscureConfirm;
                          });
                        },
                        icon: obscureConfirm
                            ? const Icon(
                                Icons.visibility_off,
                              )
                            : const Icon(Icons.visibility),
                        iconSize: 19,
                      ),
                    ),
                    const Space(10),
                    Text(
                      'Two-Factor Authentication',
                      style: context.textTheme.bodyMedium
                          ?.copyWith(fontWeight: FontWeight.w600),
                    ),
                    const Space(5),
                    TextFormInput(
                      autoValidateMode: AutovalidateMode.onUserInteraction,
                      controller: codeCtr,
                      maxLines: 1,
                      validator: validatePassword,
                      labelText: '********',
                    ),
                  ],
                ),
              )
            ],
          ),
        ),
      ),
      bottomNavigationBar: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: MainButton(
              loading: false,
              text: 'Save',
              pressed: () {},
            ),
          ),
          const Space(30),
        ],
      ),
    );
  }
}
