import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class BottomWidgetComponents extends StatelessWidget {
  const BottomWidgetComponents({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Row(
            children: [
              Image.asset(
                AppAssets.postClap,
                height: 20,
                width: 20,
              ),
              const Space(5),
              Text(
                '247',
                style: context.textTheme.bodySmall
                    ?.copyWith(color: context.colorScheme.primary),
              ),
            ],
          ),
          Row(
            children: [
              Image.asset(
                AppAssets.postComment,
                height: 20,
                width: 20,
              ),
              const Space(5),
              Text(
                '47',
                style: context.textTheme.bodySmall?.copyWith(
                  color: const Color(0xff7A8FA6),
                ),
              ),
            ],
          ),
          GestureDetector(
            onTap: () {
              context.pushNamed(AppRoute.pSharePost.name);
            },
            child: Container(
              height: 20,
              width: 20,
              padding: const EdgeInsets.all(2),
              decoration: const BoxDecoration(
                color: Color(0xff7A8FA6),
                borderRadius: BorderRadius.all(
                  Radius.circular(5),
                ),
              ),
              child: Image.asset(
                AppAssets.postArrowUp,
              ),
            ),
          ),
        ],
      ),
    );
  }
}
