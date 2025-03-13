import 'package:celebut/core/core.dart';
import 'package:celebut/features/profile/presentation/widgets/main_business_profile.dart';
import 'package:celebut/features/profile/presentation/widgets/main_profile_widget.dart';
import 'package:celebut/features/profile/presentation/widgets/profile_item.dart';
import 'package:celebut/features/profile/presentation/widgets/tab_view_widget.dart';
import 'package:celebut/features/profile/presentation/widgets/top_background_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:go_router/go_router.dart';

class ProfileViewPersonal extends StatefulWidget {
  const ProfileViewPersonal({super.key});

  @override
  State<ProfileViewPersonal> createState() => _ProfileViewPersonalState();
}

class _ProfileViewPersonalState extends State<ProfileViewPersonal> {
  bool showTabView = false;
  bool isBusiness = false;
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
                    tapped: () {
                      if (showTabView) {
                        setState(() {
                          showTabView = false;
                        });
                      } else {
                        context.pop();
                      }
                    },
                  ),
                  if (isBusiness) ...[
                    SizedBox(
                      height: constraints.maxHeight / 5,
                    ),
                  ] else ...[
                    SizedBox(
                      height: constraints.maxHeight / 7.5,
                    ),
                  ],
                  if (showTabView) ...[
                    const TabViewWidget(),
                  ] else ...[
                    ListView.builder(
                      shrinkWrap: true,
                      physics: const NeverScrollableScrollPhysics(),
                      itemCount: profileItemComponents.length,
                      itemBuilder: (context, index) {
                        final item = profileItemComponents[index];
                        return ProfileItem(
                          model: item,
                          tapped: () {
                            switch (item.title) {
                              case 'Friends':
                                context.pushNamed(AppRoute.pFriendsList.name);
                              case 'Business':
                                context.pushNamed(AppRoute.pBusinessList.name);
                              case 'Contents':
                                setState(() {
                                  showTabView = true;
                                });
                              case 'Celebrated':
                                setState(() {
                                  showTabView = true;
                                });
                            }
                          },
                        );
                      },
                    ),
                  ],
                ],
              ),
              if (isBusiness) ...[
                MainBusinessProfile(constraints: constraints),
              ] else ...[
                MainProfileWidget(
                  constraints: constraints,
                ),
              ],
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
