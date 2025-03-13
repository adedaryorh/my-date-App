import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class PrivacyOptionItem extends StatelessWidget {
  const PrivacyOptionItem({
    required this.title,
    required this.tapped,
    super.key,
  });

  final String title;
  final VoidCallback tapped;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: tapped,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            title,
            style: context.textTheme.bodyLarge,
          ),
          const Space(20),
          Container(
            height: 1,
            width: double.maxFinite,
            color: const Color(0xffE6E6E6),
          ),
          const Space(10),
        ],
      ),
    );
  }
}
