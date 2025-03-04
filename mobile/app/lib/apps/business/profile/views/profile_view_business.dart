import 'package:celebut/apps/personal/profile/presentation/widgets/main_profile_widget.dart';
import 'package:celebut/apps/personal/profile/presentation/widgets/profile_item.dart';
import 'package:celebut/apps/personal/profile/presentation/widgets/top_background_widget.dart';
import 'package:celebut/core/utils/constants/app_assets.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

class ProfileViewBusiness extends StatefulWidget {
  const ProfileViewBusiness({super.key});

  @override
  State<ProfileViewBusiness> createState() => _ProfileViewBusinessState();
}

class _ProfileViewBusinessState extends State<ProfileViewBusiness> {
  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (BuildContext context, BoxConstraints constraints) {
        return Scaffold(
          body: Stack(
            children: [
              Column(
                children: [
                  TopBackgroundWidget(
                    constraints: constraints,
                  ),
                  SizedBox(
                    height: constraints.maxHeight / 7.5,
                  ),
                  ListView.builder(
                    shrinkWrap: true,
                    physics: const NeverScrollableScrollPhysics(),
                    itemCount: profileBusinessItemComponents.length,
                    itemBuilder: (context, index) {
                      final item = profileBusinessItemComponents[index];
                      return ProfileItem(
                        model: item,
                      );
                    },
                  ),
                ],
              ),
              MainProfileWidget(
                constraints: constraints,
              ),
              Positioned(
                top: constraints.maxHeight / 7.5,
                left: (constraints.maxWidth - 80) / 2,
                child: Container(
                  height: 80,
                  width: 80,
                  decoration: const BoxDecoration(
                    color: Colors.grey,
                    borderRadius: BorderRadius.all(
                      Radius.circular(10),
                    ),
                    image: DecorationImage(
                      fit: BoxFit.cover,
                      image: AssetImage(
                        AppAssets.defaultAvatarGirl,
                      ),
                    ),
                  ),
                ),
              ),
              Positioned(
                top: constraints.maxHeight / 7.8,
                left: (constraints.maxWidth - 100) / 2,
                child: SvgPicture.asset(
                  AppAssets.starCheck,
                  width: 20,
                  height: 20,
                ),
              ),
            ],
          ),
        );
      },
    );
  }
}
