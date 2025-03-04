import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class InterestOptions extends StatelessWidget {
  const InterestOptions({
    required this.title,
    required this.bgColor,
    required this.textColor,
    required this.onTapped,
    super.key,
  });

  final String title;
  final Color bgColor;
  final Color textColor;
  final VoidCallback onTapped;

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: onTapped,
      child: Container(
        padding: const EdgeInsets.symmetric(vertical: 8, horizontal: 20),
        decoration: BoxDecoration(
          color: bgColor,
          borderRadius: BorderRadius.circular(24),
          border: Border.all(
            color: context.colorScheme.primary,
            width: 0.6,
          ),
        ),
        child: Text(
          title,
          style: context.textTheme.bodyLarge?.copyWith(color: textColor),
        ),
      ),
    );
  }
}
