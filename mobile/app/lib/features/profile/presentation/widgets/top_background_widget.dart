import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:go_router/go_router.dart';

class TopBackgroundWidget extends StatelessWidget {
  const TopBackgroundWidget({
    required this.constraints,
    required this.tapped,
    super.key,
  });

  final BoxConstraints constraints;
  final VoidCallback tapped;

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.only(left: 30, right: 30, top: 70),
      width: double.maxFinite,
      height: constraints.maxHeight / 3.5,
      decoration: BoxDecoration(
        color: context.colorScheme.primary,
        borderRadius: const BorderRadius.only(
          bottomLeft: Radius.circular(35),
          bottomRight: Radius.circular(35),
        ),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          InkWell(
            onTap: tapped,
            child: const Icon(
              Icons.arrow_back,
              size: 23,
            ),
          ),
          PopupMenuButton<int>(
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(10),
            ),
            elevation: 10,
            itemBuilder: (context) => [
              PopupMenuItem<int>(
                value: 0,
                child: Row(
                  children: [
                    SvgPicture.asset(
                      AppAssets.shareIcon,
                      width: 15,
                      height: 15,
                    ),
                    const SizedBox(width: 8),
                    const Text('Share Profile'),
                  ],
                ),
              ),
              PopupMenuItem<int>(
                value: 1,
                child: Row(
                  children: [
                    SvgPicture.asset(
                      AppAssets.settings2,
                      width: 15,
                      height: 15,
                    ),
                    const SizedBox(width: 8),
                    const Text('Settings'),
                  ],
                ),
              ),
            ],
            onSelected: (value) {
              if (value == 0) {
              } else {
                context.pushNamed(AppRoute.pSettings.name);
                //context.pushNamed(AppRoute.bSettings.name),
              }
            },
            child: const Icon(
              Icons.more_vert,
              size: 23,
            ),
          ),
        ],
      ),
    );
  }
}
