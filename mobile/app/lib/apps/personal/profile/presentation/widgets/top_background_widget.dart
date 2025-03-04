import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:go_router/go_router.dart';

class TopBackgroundWidget extends StatelessWidget {
  const TopBackgroundWidget({
    required this.constraints,
    super.key,
  });

  final BoxConstraints constraints;

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.only(left: 50, right: 50, top: 70),
      width: double.maxFinite,
      height: constraints.maxHeight / 3.5,
      decoration: BoxDecoration(
        color: context.colorScheme.primary,
        borderRadius: const BorderRadius.all(
          Radius.circular(45),
        ),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          InkWell(
            onTap: () => context.pushNamed(AppRoute.pSettings.name),
            //context.pushNamed(AppRoute.bSettings.name),
            child: SvgPicture.asset(
              AppAssets.settings,
              width: 23,
              height: 23,
            ),
          ),
          SvgPicture.asset(
            AppAssets.profileNotification,
            width: 23,
            height: 23,
          ),
        ],
      ),
    );
  }
}
