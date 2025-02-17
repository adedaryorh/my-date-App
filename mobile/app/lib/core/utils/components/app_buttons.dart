import 'package:celebut/core/utils/components/loading_indicator.dart';
import 'package:celebut/core/utils/extension/context_extension.dart';
import 'package:flutter/material.dart';

class MainButton extends StatelessWidget {
  const MainButton(
      {required this.loading,
      required this.text,
      required this.pressed,
      super.key});

  final bool loading;
  final String text;
  final VoidCallback? pressed;

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: pressed,
      child: loading
          ? const LoadingIndicator(
              color: Colors.white,
            )
          : Text(
              text,
              style: Theme.of(context)
                  .textTheme
                  .bodySmall
                  ?.copyWith(color: Colors.white, fontWeight: FontWeight.w600),
            ),
    );
  }
}

class BorderButton extends StatelessWidget {
  const BorderButton(
      {required this.text, required this.pressed, super.key, this.loading});

  final String text;
  final VoidCallback pressed;
  final bool? loading;

  @override
  Widget build(BuildContext context) {
    return OutlinedButton(
      onPressed: pressed,
      child: loading ?? false
          ? LoadingIndicator(
              color: context.colorScheme.primary,
            )
          : Text(
              text,
              style: Theme.of(context).textTheme.bodySmall?.copyWith(
                  color: Theme.of(context).primaryColor,
                  fontWeight: FontWeight.w600),
            ),
    );
  }
}
