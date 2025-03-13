import 'package:celebut/core/core.dart';
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';

class UserAgreement extends StatefulWidget {
  const UserAgreement({
    super.key,
  });

  @override
  State<UserAgreement> createState() => _UserAgreementState();
}

class _UserAgreementState extends State<UserAgreement> {
  bool agreed = false;
  @override
  Widget build(BuildContext context) {
    return Row(
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
                TextSpan(text: ' and ', style: context.textTheme.bodySmall),
                TextSpan(
                  recognizer: TapGestureRecognizer()
                    ..onTap = () {
                      //print('Tap Here onTap');
                    },
                  text: 'Privacy policy”',
                  style: Theme.of(context).textTheme.bodySmall?.copyWith(
                        color: context.colorScheme.primary,
                      ),
                ),
              ],
            ),
          ),
        ),
      ],
    );
  }
}
