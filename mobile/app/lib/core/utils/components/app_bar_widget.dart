import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

class TimelineAppBar extends StatelessWidget implements PreferredSizeWidget {
  const TimelineAppBar({
    required this.onTapAvatar,
    required this.onTapStore,
    required this.onTapSearch,
    super.key,
  });

  final VoidCallback onTapAvatar;
  final VoidCallback onTapStore;
  final VoidCallback onTapSearch;

  @override
  Widget build(BuildContext context) {
    return AppBar(
      backgroundColor: context.colorScheme.primary,
      centerTitle: false,
      title: const SizedBox.shrink(),
      actions: [
        Expanded(
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              InkWell(
                onTap: onTapAvatar,
                child: Stack(
                  children: [
                    Container(
                      margin: const EdgeInsets.only(left: 15),
                      height: 30,
                      width: 30,
                      decoration: const BoxDecoration(
                        color: Colors.grey,
                        borderRadius: BorderRadius.all(
                          Radius.circular(5),
                        ),
                        image: DecorationImage(
                          fit: BoxFit.cover,
                          image: AssetImage(
                            AppAssets.defaultAvatarGirl,
                          ),
                        ),
                      ),
                    ),
                    Positioned(
                      top: -20,
                      bottom: 0,
                      left: 10,
                      child: SvgPicture.asset(
                        AppAssets.starCheck,
                        width: 10,
                        height: 10,
                      ),
                    ),
                  ],
                ),
              ),
              InkWell(
                onTap: onTapStore,
                child: Image.asset(
                  AppAssets.store,
                  width: 38,
                  height: 38,
                ),
              ),
              InkWell(
                onTap: onTapSearch,
                child: Padding(
                  padding: const EdgeInsets.only(right: 15),
                  child: SvgPicture.asset(
                    AppAssets.search,
                    width: 20,
                    height: 20,
                  ),
                ),
              ),
            ],
          ),
        ),
      ],
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);
}
