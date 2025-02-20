import 'package:celebut/apps/shared/auth/presentation/views/sign_up/widgets/industry_picker.dart';
import 'package:celebut/apps/shared/auth/presentation/views/sign_up/widgets/signin_option.dart';
import 'package:celebut/apps/shared/auth/presentation/views/sign_up/widgets/user_agreement.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

class SignUpBusiness extends StatefulWidget {
  const SignUpBusiness({super.key});

  @override
  State<SignUpBusiness> createState() => _SignUpBusinessState();
}

class _SignUpBusinessState extends State<SignUpBusiness> {
  final TextEditingController businessNameCtr = TextEditingController();
  final TextEditingController industryTypeCtr = TextEditingController();
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
    businessNameCtr.dispose();
    industryTypeCtr.dispose();
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
              'Business Name',
              style: context.textTheme.bodyMedium
                  ?.copyWith(fontWeight: FontWeight.w600),
            ),
            const SizedBox(
              height: 5,
            ),
            TextFormInput(
              autoValidateMode: AutovalidateMode.onUserInteraction,
              controller: businessNameCtr,
              validator: validateName,
              labelText: 'Enter business name',
            ),
            const SizedBox(
              height: 10,
            ),
            Text(
              'Industry Type',
              style: context.textTheme.bodyMedium
                  ?.copyWith(fontWeight: FontWeight.w600),
            ),
            const SizedBox(
              height: 5,
            ),
            TextFormInput(
              autoValidateMode: AutovalidateMode.onUserInteraction,
              controller: industryTypeCtr,
              validator: validateUsername,
              labelText: 'select industry type',
              readOnly: true,
              suffixIcon: IconButton(
                onPressed: () {},
                icon: SvgPicture.asset(
                  AppAssets.downArrow,
                  width: 15,
                  height: 15,
                ),
              ),
              onTap: () {
                showDialog<void>(
                  context: context,
                  builder: (BuildContext context) {
                    return IndustryPicker();
                  },
                );
              },
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
            const UserAgreement(),
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
            const SignInOption(),
          ],
        ),
      ),
    );
  }
}
