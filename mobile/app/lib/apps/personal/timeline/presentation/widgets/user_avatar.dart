import 'package:celebut/core/utils/constants/app_assets.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

class UserAvatar extends StatelessWidget {
  const UserAvatar({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Container(
          margin: const EdgeInsets.only(left: 10, top: 10),
          height: 60,
          width: 60,
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
          top: -45,
          bottom: 0,
          left: 0,
          child: SvgPicture.asset(
            AppAssets.starCheck,
            width: 20,
            height: 20,
          ),
        ),
      ],
    );
  }
}
