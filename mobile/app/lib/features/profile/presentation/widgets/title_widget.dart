import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class TitleWidget extends StatelessWidget {
  const TitleWidget({
    required this.constraints,
    super.key,
  });

  final BoxConstraints constraints;

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.only(left: 20, right: 20, top: 50),
      width: double.maxFinite,
      height: constraints.maxHeight / 3.5,
      decoration: BoxDecoration(
        color: context.colorScheme.primary,
        borderRadius: const BorderRadius.only(
          bottomLeft: Radius.circular(35),
          bottomRight: Radius.circular(35),
        ),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          InkWell(
            onTap: () => context.pop(),
            child: const Icon(
              Icons.close,
              size: 23,
              color: Colors.white,
            ),
          ),
          const Space(30),
          Align(
            child: Text(
              'Premium',
              style: context.textTheme.titleMedium
                  ?.copyWith(color: Colors.white, fontSize: 32),
            ),
          ),
          const Space(10),
          Align(
            child: Text(
              'Subscribe to our premium package and '
              '\nhave access to more features',
              style:
                  context.textTheme.bodyMedium?.copyWith(color: Colors.white),
              textAlign: TextAlign.center,
            ),
          ),
        ],
      ),
    );
  }
}
