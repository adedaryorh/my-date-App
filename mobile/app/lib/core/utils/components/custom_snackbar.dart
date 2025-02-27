import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class CustomSnackBar extends StatelessWidget {
  final String? title;
  final String message;
  final Color background;
  final void Function() onDismiss;

  const CustomSnackBar({
    required this.message,
    required this.background,
    required this.onDismiss,
    super.key,
    this.title,
  });

  @override
  Widget build(BuildContext context) {
    final size = MediaQuery.of(context).size;
    return Material(
      color: Colors.transparent,
      child: GestureDetector(
        onTapDown: (_) => onDismiss(),
        child: Container(
          padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 25),
          width: size.width,
          decoration: BoxDecoration(
            borderRadius:
                const BorderRadius.vertical(bottom: Radius.circular(10)),
            color: background,
          ),
          child: SafeArea(
            bottom: false,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                if (title != null) ...[
                  Text(
                    title!,
                    style: context.textTheme.bodyMedium?.copyWith(
                      color: Colors.white,
                      fontWeight: FontWeight.w600,
                    ),
                  ),
                  const SizedBox(height: 10),
                ],
                Text(
                  message,
                  style: context.textTheme.bodyMedium
                      ?.copyWith(color: Colors.white),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
