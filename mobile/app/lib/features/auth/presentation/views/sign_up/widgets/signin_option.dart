import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class SignInOption extends StatelessWidget {
  const SignInOption({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return TextButton(
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
    );
  }
}
