import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class StoryWidget extends StatelessWidget {
  const StoryWidget({
    required this.assetPath,
    super.key,
  });

  final String assetPath;

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Container(
          width: 143,
          height: 200,
          margin: const EdgeInsets.only(right: 20),
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(25),
            color: Colors.black38,
            image: DecorationImage(
              image: AssetImage(assetPath),
              fit: BoxFit.cover,
            ),
          ),
        ),
        Positioned(
          right: 0,
          left: 10,
          bottom: 30,
          child: Text(
            'Jane \nGeorge',
            style: context.textTheme.bodyMedium
                ?.copyWith(color: Colors.white, fontWeight: FontWeight.w700),
          ),
        ),
      ],
    );
  }
}
