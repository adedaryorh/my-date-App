import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

class ProfileAvatar extends StatelessWidget {
  const ProfileAvatar({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Align(
      child: SizedBox(
        height: 100,
        child: Stack(
          children: [
            Container(
              height: 100,
              width: 100,
              decoration: const BoxDecoration(
                shape: BoxShape.circle,
                image: DecorationImage(
                  image: AssetImage(
                    AppAssets.defaultAvatarGirl,
                  ),
                  fit: BoxFit.cover,
                ),
              ),
            ),
            Positioned(
              bottom: 0,
              right: 0,
              child: SizedBox(
                height: 32,
                width: 32,
                child: FloatingActionButton.small(
                  heroTag: null,
                  onPressed: () {},
                  backgroundColor: const Color(0xffF5F5FF),
                  shape: const CircleBorder(),
                  child: SvgPicture.asset(
                    AppAssets.camera,
                    width: 18,
                    colorFilter: ColorFilter.mode(
                      context.colorScheme.primary,
                      BlendMode.srcIn,
                    ),
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
