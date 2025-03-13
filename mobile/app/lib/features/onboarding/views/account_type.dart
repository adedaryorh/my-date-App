import 'package:celebut/core/core.dart';
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class AccountTypeView extends StatefulWidget {
  const AccountTypeView({super.key});

  @override
  State<AccountTypeView> createState() => _AccountTypeViewState();
}

class _AccountTypeViewState extends State<AccountTypeView> {
  String? selectedOption;
  List<String> options = ['Personal', 'Business'];

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
                'Select Account Type \nTo Create',
                style: context.textTheme.headlineLarge,
              ),
              const Space(50),
              for (int index = 0; index < options.length; index++)
                CustomRadioListTile(
                  option: options[index],
                  selectedOption: selectedOption,
                  onSelected: (String value) {
                    selectedOption = value;
                    setState(() {});
                  },
                ),
              const Space(50),
              MainButton(
                loading: false,
                text: 'Continue',
                pressed: selectedOption != null
                    ? () {
                        if (selectedOption == options[0]) {
                          context.push('/personalSignUp');
                        } else if (selectedOption == options[1]) {
                          context.push('/businessSignUp');
                        }
                      }
                    : null,
              ),
              const Space(50),
              Align(
                child: TextButton(
                  onPressed: () => context.pushNamed(AppRoute.signIn.name),
                  child: RichText(
                    textAlign: TextAlign.center,
                    text: TextSpan(
                      style: context.textTheme.bodyMedium,
                      children: [
                        const TextSpan(
                          text: 'Already have an account? ',
                          style: TextStyle(fontWeight: FontWeight.w600),
                        ),
                        TextSpan(
                          text: 'Log in',
                          style: TextStyle(
                            color: context.colorScheme.primary,
                            fontWeight: FontWeight.w600,
                          ),
                          recognizer: TapGestureRecognizer()..onTap = () {},
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
