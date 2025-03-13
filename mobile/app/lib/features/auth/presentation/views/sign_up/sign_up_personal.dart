import 'package:celebut/core/core.dart';
import 'package:celebut/features/auth/presentation/views/sign_up/widgets/signin_option.dart';
import 'package:celebut/features/auth/presentation/views/sign_up/widgets/user_agreement.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:intl/intl.dart';

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
  final TextEditingController phoneNumberCtr = TextEditingController();
  final TextEditingController date = TextEditingController();

  final _formKey = GlobalKey<FormState>();
  bool obscure = true;
  bool obscure1 = true;

  PhoneNumber number = PhoneNumber(isoCode: 'NG');
  String dialCode = '';

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
      body: GestureDetector(
        onTap: () {
          FocusScope.of(context).unfocus();
        },
        behavior: HitTestBehavior.opaque,
        child: Form(
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
                'Phone Number',
                style: context.textTheme.bodyMedium
                    ?.copyWith(fontWeight: FontWeight.w600),
              ),
              const SizedBox(
                height: 5,
              ),
              InternationalPhoneNumberInput(
                onInputChanged: (PhoneNumber number) {
                  dialCode = '${number.dialCode}';
                },
                autoValidateMode: AutovalidateMode.onUserInteraction,
                onInputValidated: (bool value) {},
                initialValue: number,
                textFieldController: phoneNumberCtr,
                keyboardType: const TextInputType.numberWithOptions(
                  signed: true,
                  decimal: true,
                ),
                validator: validatePhoneNumber,
                onPressed: () {},
              ),
              const SizedBox(
                height: 10,
              ),
              Text(
                'Date of Birth',
                style: context.textTheme.bodyMedium
                    ?.copyWith(fontWeight: FontWeight.w600),
              ),
              const SizedBox(
                height: 5,
              ),
              TextFormInput(
                readOnly: true,
                controller: date,
                validator: validateDate,
                labelText: 'DD/MM/YYYY',
                suffixIcon: const Icon(
                  Icons.calendar_month,
                  color: Colors.black,
                  size: 20,
                ),
                inputFormatters: [
                  FilteringTextInputFormatter.deny(RegExp('[ ]')),
                ],
                onTap: () async {
                  final value = await showPlatformDatePicker(
                    context: context,
                    initialDate: DateTime.now(),
                    firstDate:
                        DateTime.now().subtract(const Duration(days: 365 * 50)),
                    lastDate: DateTime.now(),
                  );
                  if (value == null) return;
                  final formatter = DateFormat('yyyy-MM-dd');
                  final formattedDate = formatter.format(value);
                  setState(() {
                    date.text = formattedDate;
                  });
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
      ),
    );
  }
}
